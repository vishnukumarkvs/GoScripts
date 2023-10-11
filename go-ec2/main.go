package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    // Initialize a session that the SDK will use to load credentials
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )

    // Create an EC2 service client.
    svc := ec2.New(sess)

    // Call the DescribeVolumes method from the EC2 client
    result, err := svc.DescribeVolumes(nil)
    if err != nil {
        fmt.Println("Error", err)
        return
    }

    // Loop through the volumes and print details
    for _, volume := range result.Volumes {
        fmt.Println("Volume ID:", *volume.VolumeId)
        fmt.Println("Size:", *volume.Size, "GiB")
        fmt.Println("State:", *volume.State)
        fmt.Println("Type:", *volume.VolumeType)
        fmt.Println("---------------------------")
    }
}