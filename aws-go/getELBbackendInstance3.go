package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func main(){
	fmt.Println(os.Args)

	if len(os.Args) != 2{
		fmt.Println("引数エラー")
		os.Exit(1)
	}

	elbName := os.Args[1]

	svc := elb.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	input := &elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(elbName),
	}

	resELBInfo, err := svc.DescribeInstanceHealth(input)
	if err != nil {
		os.Exit(1)
	}

	for i := range resELBInfo.InstanceStates {
		resultInstance      := *resELBInfo.InstanceStates[i].InstanceId
		resultInstanceState := *resELBInfo.InstanceStates[i].State
		fmt.Println(resultInstance+"\t"+resultInstanceState)
	}

}
