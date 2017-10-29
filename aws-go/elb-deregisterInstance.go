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

	if len(os.Args) != 3{
		fmt.Println("引数エラー")
		os.Exit(1)
	}

	elbName := os.Args[1]
	instanceID := os.Args[2]

	svc := elb.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	//func (c *ELB) DeregisterInstancesFromLoadBalancer(input *DeregisterInstancesFromLoadBalancerInput) (*DeregisterInstancesFromLoadBalancerOutput, error)


	input := &elb.DeregisterInstancesFromLoadBalancerInput{
		Instances: []*elb.Instance{
			{
				InstanceId: aws.String(instanceID),
			},
		},
		LoadBalancerName: aws.String(elbName),
	}

	resELBInfo, err := svc.DeregisterInstancesFromLoadBalancer(input)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(resELBInfo)


}
