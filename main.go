package main

import (
	//"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Get() ([]byte, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{Profile: "s3-lambda"}))
	downloader := s3manager.NewDownloader(sess)
	buffer := aws.NewWriteAtBuffer([]byte{})
	// _ -> numBytes
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String("home-state"),
			Key:    aws.String("config/ip.txt"),
		})
	//fmt.Println(string(buffer.Bytes()))
	return buffer.Bytes(), err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Get)
	//Get()
}
