package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func main(){
	elbName := "LBNAME"

	svc := elb.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	input := &elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(elbName),
	}

	resELBInfo, err := svc.DescribeInstanceHealth(input)
	if err != nil {
		panic(err)
	}

	for i := range resELBInfo.InstanceStates {
		resultInstance      := *resELBInfo.InstanceStates[i].InstanceId
		resultInstanceState := *resELBInfo.InstanceStates[i].State
		fmt.Println(resultInstance+"\t"+resultInstanceState)
	}

}
