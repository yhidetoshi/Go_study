package main

import (
	"fmt"
	"os"
  "time"
  "sort"
  "strings"

  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
  config = aws.Config{Region: aws.String("ap-northeast-1")}
	svc = ec2.New(session.New(&config))
)

func main() {

  // ec2
  result := getInstanceInfo()
	for v, _ := range result {
		fmt.Println(result[v])
	}

  //ami
  amilist := ListAMI()
  const layout = "2006-01-02T15:04:05"
  times := make([]time.Time, len(amilist))

  //fmt.Println(amilist)


  for i, _ := range amilist {
    times[i], _ = time.Parse(layout, strings.Replace(amilist[i][1], ".000Z", "", 1))
  }


  //times := make([]time.Time, 3)
/*
  times[0], _ = time.Parse(DFmt, "2018-06-29T10:33:59.000Z")
  times[1], _ = time.Parse(DFmt, "2018-07-02T06:26:02.000Z")
  times[2], _ = time.Parse(DFmt, "2018-06-29T10:10:06.000Z")
*/
/*
  times[0], _ = time.Parse(DFmt, strings.Replace("2018-06-29T10:33:59.000Z", ".000Z", "", 1))
  times[1], _ = time.Parse(DFmt, strings.Replace("2018-07-02T06:26:02.000Z", ".000Z", "", 1))
  times[2], _ = time.Parse(DFmt, strings.Replace("2018-06-29T10:10:06.000Z", ".000Z", "", 1))
*/
  sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })

    for _, t := range times {
        fmt.Printf("%+v\n", t)
    }
}


func ListAMI() [][]string {
	var owner, images []*string
	var _owner []string = []string{"self"}
	// Convert []string to []*string
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
    //fmt.Println(allAmiInfo)
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


func getInstanceInfo() []string {
	var tagName string
	var instances []string

	params := &ec2.DescribeInstancesInput{}
	res, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, resInfo := range res.Reservations {
		for _, instanceInfo := range resInfo.Instances {
			for _, tagInfo := range instanceInfo.Tags {
				if *tagInfo.Key == "Name" {
					tagName = *tagInfo.Value
				}
			}
			instances = append(instances, tagName)
		}
	}
	return instances
}
