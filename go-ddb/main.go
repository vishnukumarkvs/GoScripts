package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	region := "us-east-2"
	tableName := "QrChatSummary3"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := dynamodb.New(sess)

	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(scanInput)
	if err != nil {
		fmt.Println("Error scanning table: ", err)
		return
	}

	fmt.Println("Scan complete")

	for _, item := range result.Items {
		deleteInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key:       item,
		}
		_, err = svc.DeleteItem(deleteInput)
		if err != nil {
			fmt.Println("Error deleteing item: ", err)
		} else {
			fmt.Println("Deleted Item: ", item)
		}
	}

	fmt.Println("Delete Complete")
}
