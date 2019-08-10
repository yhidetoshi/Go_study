package main

import (
	"fmt"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

var (
	projectName = "sample_pjt"
	config = aws.Config{Region: aws.String("ap-northeast-1")}
	svcCodeBuild = codebuild.New(session.New(&config))
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) {
	var ptrProjectName *string

	ptrProjectName = &projectName

	params := &codebuild.StartBuildInput{
		ProjectName: ptrProjectName,
	}

	req, resp := svcCodeBuild.StartBuildRequest(params)
	err := req.Send()
	if err == nil {
		fmt.Println(resp)
	}
}
