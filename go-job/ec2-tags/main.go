package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Go v1 - next time will be using go sdk v2
func updateTagsForInstance(ec2Client *ec2.EC2, instanceID string) error {
    input := &ec2.CreateTagsInput{
        Resources: []*string{aws.String(instanceID)},
        Tags:      []*ec2.Tag{{Key: aws.String("Name"), Value: aws.String("qrsocketweb")}},
    }

    _, err := ec2Client.CreateTags(input)
    return err
}
func main() {
	// Initialize a session that the SDK will use to load credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err!= nil{
		fmt.Println("Not able to connect to aws account")
	}

	// Create an EC2 service client.
	client := ec2.New(sess)

	err = updateTagsForInstance(client, "i-0e36f91124875d0b9")

	// resp, err := client.DescribeInstances(&ec2.DescribeInstancesInput{
	// 	InstanceIds: []*string{aws.String("i-0e36f91124875d0b9")},})

	if err!=nil{
		fmt.Println("Error", err)
	}else{
		fmt.Println("Success")
	}

	// fmt.Println(resp)
}
