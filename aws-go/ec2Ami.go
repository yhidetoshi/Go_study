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
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	KEEPNUM       = 0
	SSMSEARCHWORD = "_source_ami"
)

var (
	config = aws.Config{Region: aws.String("ap-northeast-1")}
	svcEc2 = ec2.New(session.New(&config))
	svcSsm = ssm.New(session.New(&config))
	layout = "2006-01-02T15:04:05"
)

type Params struct {
	Name, Value []*string
}

func main() {

	var sortedAmiList = []string{}
	var excludedSourceAmi = []string{}
	var excludedSourceAmiId = []string{}
	var checkFlag = true

	// 除外されるssmパラメータのkeyを取得
	excludedSourceAmi = FetchSsmParams()

	// 除外されるssmパラメータの値(ami-id)を取得
	for i, _ := range excludedSourceAmi {
		excludedSourceAmiId = append(excludedSourceAmiId, GetSourceAmi(&excludedSourceAmi[i]))
	}

	amiList := GetListAMI()
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

	// 不要なAMI削除 | checkFlag is true -> delete
	for i := 0; i < (len(sortedAmiList) - KEEPNUM); i++ {
		for j, _ := range excludedSourceAmiId {
			if sortedAmiList[i] == excludedSourceAmiId[j] {
				checkFlag = false
			}
		}
		if checkFlag == true {
			//DeregisterAMI(sortedAmiList[i])
			fmt.Println(sortedAmiList[i])
		}
		checkFlag = true
	}
	fmt.Println("Deleted \nFin")
}

// convert time to string
func TimeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

// ssmパラメータで指定したkeyでvalueを返す
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

// ssmパラメータから `_source_ami` に該当するkeyを返す
func FetchSsmParams() []string {

	allSourceAmi := []string{}
	//var result []string
	params := &ssm.DescribeParametersInput{}
	res, err := svcSsm.DescribeParameters(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, value := range res.Parameters {
		if strings.Contains(*value.Name, SSMSEARCHWORD) {
			allSourceAmi = append(allSourceAmi, *value.Name)
		}
	}
	return allSourceAmi
}

// 登録されているAMI一覧を返す
func GetListAMI() [][]string {
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

// AMI-IDを指定してイメージを削除(登録解除)する
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
