package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	config = aws.Config{Region: aws.String("ap-northeast-1")}
	svcEc2 = ec2.New(session.New(&config))
	svcAsg = autoscaling.New(session.New(&config))
)

type Calc interface {
	GetAmiNums() int
	GetLCNums() int
}

type Nums struct {
	Calc
	amiNumbers int
	lcNumbers  int
}

func main() {
	var c Calc = &Nums{}

	fmt.Println(c.GetAmiNums())
	fmt.Println(c.GetLCNums())
}

// 登録されているAMI数を返す
func (n *Nums) GetAmiNums() int {
	var owner, images []*string
	var _owner []string = []string{"self"}

	owner = aws.StringSlice(_owner)
	params := &ec2.DescribeImagesInput{
		ImageIds: images,
		Owners:   owner,
	}
	res, err := svcEc2.DescribeImages(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	n.amiNumbers = len(res.Images)
	return n.amiNumbers
}

// 起動設定(LC)数を返す
func (n *Nums) GetLCNums() int {
	params := &autoscaling.DescribeLaunchConfigurationsInput{}
	res, err := svcAsg.DescribeLaunchConfigurations(params)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	n.lcNumbers = len(res.LaunchConfigurations)
	return n.lcNumbers
}
