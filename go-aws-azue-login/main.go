package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	region := "eu-west-1"
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithSharedConfigProfile("otinp"),
	)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	client := sts.NewFromConfig(cfg)

	identity, err := client.GetCallerIdentity(
		context.TODO(),
		&sts.GetCallerIdentityInput{},
	)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf(
		"Account: %s\nUserID: %s\nARN: %s\n",
		aws.ToString(identity.Account),
		aws.ToString(identity.UserId),
		aws.ToString(identity.Arn),
	)
	// In this code, identity.Account, identity.UserId, and identity.Arn are all *string pointers returned by the AWS SDK for Go v2. By using aws.ToString, you ensure that even if any of these fields are nil, you won't encounter a runtime error when trying to print them as strings.
}
