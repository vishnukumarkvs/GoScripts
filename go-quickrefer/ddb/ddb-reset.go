package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
)

// config.go: contains KeySchema and the tableKeyMapping map
// only works for N and S types for partition and sort keys
// for example:
// tableKeyMapping["Table1"] = KeySchema{PartitionKey: "pk", PartitionType: "N", SortKey: "sk", SortKeyType: "N"}

func deleteTableItems(svc *dynamodb.DynamoDB, tableName string, wg *sync.WaitGroup) {
	defer wg.Done()

	keySchema, found := tableKeyMapping[tableName]
	if !found {
		fmt.Println("Table not found: ", tableName)
		return
	}

	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(scanInput)
	if err != nil {
		fmt.Println("Error scanning table: ", err)
		return
	}
	for _, item := range result.Items {
		partitionValue := item[keySchema.PartitionKey]
		sortValue := item[keySchema.SortKey]

		var partitionAttributeValue *dynamodb.AttributeValue
		var sortAttributeValue *dynamodb.AttributeValue

		if keySchema.PartitionType == "S" {
			partitionAttributeValue = &dynamodb.AttributeValue{S: partitionValue.S}
		} else if keySchema.PartitionType == "N" {
			partitionAttributeValue = &dynamodb.AttributeValue{N: partitionValue.N}
		}

		if keySchema.SortKeyType == "S" {
			sortAttributeValue = &dynamodb.AttributeValue{S: sortValue.S}
		} else if keySchema.SortKeyType == "N" {
			sortAttributeValue = &dynamodb.AttributeValue{N: sortValue.N}
		}

		deleteInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				keySchema.PartitionKey: partitionAttributeValue,
				keySchema.SortKey:      sortAttributeValue,
			},
		}

		_, err = svc.DeleteItem(deleteInput)
		if err != nil {
			fmt.Println("Error deleting item: ", tableName, err, item)
		}
	}

	fmt.Printf("Delete Complete for %v\n", tableName)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	start := time.Now()
	region := "us-east-2"
	tableNamesStr := os.Getenv("TableNames")
	tableNames := strings.Split(tableNamesStr, ",")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := dynamodb.New(sess)

	wg := sync.WaitGroup{}

	for _, tableName := range tableNames {
		wg.Add(1)
		go deleteTableItems(svc, tableName, &wg)
	}

	wg.Wait()

	fmt.Println(time.Since(start))
}

// 2 sec without goroutines
// 1.16 with go routines
