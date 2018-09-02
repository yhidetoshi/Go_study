package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	KEEPNUM = 2
	SSMPARA = "source_ami_id"
)

var (
	config = aws.Config{Region: aws.String("ap-northeast-1")}
	svcEc2 = ec2.New(session.New(&config))
	svcSsm = ssm.New(session.New(&config))
	layout = "2006-01-02T15:04:05"
)

func main() {

	var sortedAmiList = []string{}
	ssmPara := SSMPARA

	sourceAmiId := GetSourceAmi(&ssmPara)
	//fmt.Println(sourceAmiId)

	amiList := ListAMI()
	times := make([]time.Time, len(amiList))

	// trim ".000Z"
	for i, _ := range amiList {
		times[i], _ = time.Parse(layout, strings.Replace(amiList[i][1], ".000Z", "", 1))
		amiList[i][1] = strings.Replace(amiList[i][1], ".000Z", "", 1)
	}

	// sort by createTime(ascending order)
	sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })

	for i, _ := range amiList {
		for j, _ := range amiList {
			if TimeToString(times[i]) == amiList[j][1] {
				sortedAmiList = append(sortedAmiList, amiList[j][0])
				//fmt.Printf("%s %s\n", amiList[j][0], amiList[j][1])
			}
		}
	}

	// deregister ami(older)
	for i := 0; i < (len(sortedAmiList) - KEEPNUM); i++ {
		// excluded source ami
		if sourceAmiId != sortedAmiList[i] {
			DeregisterAMI(sortedAmiList[i])
		}
	}
	fmt.Println("Deleted \nFin")
}

// convert time to string
func TimeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

// get ssm parameter
func GetSourceAmi(sourceAmi *string) string {
	var _sourceAmi string
	params := &ssm.GetParameterInput{
		Name: sourceAmi,
	}
	res, err := svcSsm.GetParameter(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	_sourceAmi = *res.Parameter.Value

	return _sourceAmi

}

// get ami list
func ListAMI() [][]string {
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

// remove ami
func DeregisterAMI(ec2AMIid string) {
	fmt.Printf("%s \n", ec2AMIid)

	// using pointer for params
	var _ec2AMIid *string
	_ec2AMIid = &ec2AMIid

	params := &ec2.DeregisterImageInput{
		ImageId: _ec2AMIid,
	}
	_, err := svcEc2.DeregisterImage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

