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

func deleteTableItems(svc *dynamodb.DynamoDB, tableName string, wg *sync.WaitGroup){
	defer wg.Done()

	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(scanInput)
	if err != nil{
		fmt.Println("Error scanning table: ", err)
		return
	}

	fmt.Printf("Scan complete for %v\n",tableName)

	for _, item := range result.Items {
		deleteInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key : item,
		}
		_, err = svc.DeleteItem(deleteInput)
		if err!= nil{
			fmt.Println("Error deleteing item: ",err)
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
	start:= time.Now()
	region := "us-east-2"
	tableNamesStr := os.Getenv("TableNames")
	tableNames := strings.Split(tableNamesStr,",")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil{
		fmt.Println("Error creating session:", err)
		return
	}

	svc := dynamodb.New(sess)

	wg := sync.WaitGroup{}

	for _,tableName := range tableNames{
		wg.Add(1)
		go deleteTableItems(svc,tableName,&wg)
	}

    wg.Wait()

	fmt.Println(time.Since(start))
}

// 2 sec without goroutines
// 1.16 with go routines