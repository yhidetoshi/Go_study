package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})

	res, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	var tagName string

	for _, resInfo := range res.Reservations {
		for _, instanceInfo := range resInfo.Instances {

			for _, tagInfo := range instanceInfo.Tags {
				if *tagInfo.Key == "Name" {
					tagName = *tagInfo.Value
				}
			}
			fmt.Println("{")
			fmt.Println(
				" InstanceNmae:\t"+tagName+"\n",
				   "InstanceType:\t"+*instanceInfo.InstanceType+"\n",
				   "AvailabilityZone:\t"+*instanceInfo.Placement.AvailabilityZone+"\n",
				   "State:\t"+*instanceInfo.State.Name+"\n",
				   "PrivateIP:\t"+*instanceInfo.PrivateIpAddress,
			)
			fmt.Println("}\n")
		}
	}
}
