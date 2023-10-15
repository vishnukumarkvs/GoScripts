package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	start := time.Now()
	region := "us-east-2"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	svc := s3.New(sess)

	bucket := aws.String(os.Getenv("S3_BUCKET"))

	listInput := &s3.ListObjectsInput{
		Bucket: bucket,
	}

	listResult, err:= svc.ListObjects(listInput)
	if err != nil {
		fmt.Println("Error listing objects:", err)
		return
	}

	for _, object := range listResult.Contents{
		deleteInput := &s3.DeleteObjectInput{
			Bucket: bucket,
			Key: object.Key,
		}

		_,err := svc.DeleteObject(deleteInput)
		if err != nil {
			fmt.Println("Error deleting object:", err)
		}else{
			fmt.Println("Deleted object:", *object.Key)
		}

	}


	fmt.Println(time.Since(start))
}