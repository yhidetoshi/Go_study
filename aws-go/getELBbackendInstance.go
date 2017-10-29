package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func main(){
	elbName := "ELBName"
	var inserviceInstances []string

	svc := elb.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	params := &elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(elbName),
	}

	resELBInfo, err := svc.DescribeInstanceHealth(params)
	if err != nil {
		panic(err)
	}

	for i := range resELBInfo.InstanceStates {
		result := *resELBInfo.InstanceStates[i].InstanceId + ":" + *resELBInfo.InstanceStates[i].State
		inserviceInstances = append(inserviceInstances, result)
	}

	fmt.Println(inserviceInstances)
}
