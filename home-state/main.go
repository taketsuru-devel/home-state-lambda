package main

import (
	//"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/followedwind/home-state-lambda/home-state/util"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    buffer, err := util.S3Get("config/ip.txt")
	return events.APIGatewayProxyResponse{
		Body:       string(buffer.Bytes()),
		StatusCode: 200,
	}, err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
