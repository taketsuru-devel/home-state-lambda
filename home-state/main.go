package main

import (
	//"fmt"
    //"bytes"
    //"net/url"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/followedwind/home-state-lambda/home-state/util"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    command := "connection_check"
    //param := url.Values{}
    //param.Set("command", command)
    param := make(map[string]interface{})
    param["command"] = command
    res, err := util.HomeAccess(command, param)

	return events.APIGatewayProxyResponse{
		Body:       string(res),
		StatusCode: 200,
	}, err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
