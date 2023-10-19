package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
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
		DesiredCount: aws.Int32(0),
	}
	_, err := client.UpdateService(context.TODO(), input)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("success")
	}

	serviceResourceID := fmt.Sprintf("service/%s/%s",clusterName,serviceName)

	// Specify the new minimum and maximum capacity values
	newMinCapacity := int32(1)
	newMaxCapacity := int32(2)

	// Create a variable to hold the updated scalable target settings
	scalableTargetUpdate := &applicationautoscaling.RegisterScalableTargetInput{
		ServiceNamespace: types.ServiceNamespaceEcs,
		ResourceId:      aws.String(serviceResourceID),
		ScalableDimension: types.ScalableDimensionECSServiceDesiredCount,
		MinCapacity:     aws.Int32(newMinCapacity),
		MaxCapacity:     aws.Int32(newMaxCapacity),
	}

	// Update the scalable target settings
	_, err = autoScalingClient.RegisterScalableTarget(context.TODO(), scalableTargetUpdate)
	if err != nil {
		fmt.Println("Error updating scalable target:", err)
		os.Exit(1)
	}else{
		fmt.Println("Updated autoscaling policy")
	}
}
