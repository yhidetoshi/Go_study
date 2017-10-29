package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

func main() {

	fmt.Println(os.Args)

	if len(os.Args) != 2{
		fmt.Println("引数エラー")
		os.Exit(1)
	}
	var _instances []string = []string{os.Args[1]}

	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
	instances := aws.StringSlice(_instances)

	params := &ec2.StartInstancesInput{
		InstanceIds: instances,
	}
	result, err := svc.StartInstances(params)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(result)

	for _, r := range result.StartingInstances {
		fmt.Printf("%s を起動しました.\n", *r.InstanceId)
	}
}
