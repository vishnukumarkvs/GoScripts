package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Failed to load AWS configuration:", err)
		return
	}

	// Create an ECS service client.
	client := ecs.NewFromConfig(cfg)
	autoScalingClient := applicationautoscaling.NewFromConfig(cfg)


	upscaleEcsService(client, autoScalingClient, "test2", "test-service-02")

}

func upscaleEcsService(client *ecs.Client, autoScalingClient *applicationautoscaling.Client,  clusterName string, serviceName string) {
	input := &ecs.UpdateServiceInput{
		Service:      aws.String(serviceName),
		Cluster:      aws.String(clusterName),
		DesiredCount: aws.Int32(5),
	}
	_, err := client.UpdateService(context.TODO(), input)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("success")
	}
}
