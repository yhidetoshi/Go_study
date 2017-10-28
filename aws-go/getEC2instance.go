
package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil{
		panic(err)
	}

	fmt.Println("インスタンス合計数", len(resp.Reservations))

	for index, _ := range resp.Reservations {
		//fmt.Println(len(res.Instances))

		for _, inst := range resp.Reservations[index].Instances{
			fmt.Print("インスタンス情報\t" , *inst.InstanceId, "\t", *inst.ImageId, "\t", *inst.PrivateIpAddress, "\t", *inst.SubnetId, "\t", *inst.InstanceType, "\n")

		}
	}

}
