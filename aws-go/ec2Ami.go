package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	config = aws.Config{Region: aws.String("ap-northeast-1")}
	svc    = ec2.New(session.New(&config))
	layout = "2006-01-02T15:04:05"
)

func main() {
	amilist := ListAMI()
	times := make([]time.Time, len(amilist))

	for i, _ := range amilist {
		times[i], _ = time.Parse(layout, strings.Replace(amilist[i][1], ".000Z", "", 1))
	}

	sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })

	for i := 0; i < len(times); i++ {
		for j := 0; j < len(times); j++ {

			if amilist[i][1] == timeToString(times[j])+".000Z" {
				fmt.Printf("%s %s\n", amilist[j][0], amilist[j][1])
			}
		}
	}
}

func timeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func ListAMI() [][]string {
	var owner, images []*string
	var _owner []string = []string{"self"}

	owner = aws.StringSlice(_owner)
	params := &ec2.DescribeImagesInput{
		ImageIds: images,
		Owners:   owner,
	}
	res, err := svc.DescribeImages(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	allAmiInfo := [][]string{}
	for _, resInfo := range res.Images {
		amiInfo := []string{
			*resInfo.ImageId,
			*resInfo.CreationDate,
		}
		allAmiInfo = append(allAmiInfo, amiInfo)
	}
	return allAmiInfo
}

func DeregisterAMI(ec2AMIid *string) {
	params := &ec2.DeregisterImageInput{
		ImageId: ec2AMIid,
	}
	_, err := svc.DeregisterImage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Success!!")
}

