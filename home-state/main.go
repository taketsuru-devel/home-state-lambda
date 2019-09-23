package main

import (
	"encoding/json"
	//"fmt"
    //"bytes"
    //"net/url"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/followedwind/home-state-lambda/home-state/util"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    var requestParam util.Param
    json.Unmarshal([]byte(request.Body), &requestParam)

    command := requestParam.Command

    if command == "etherwake" {
        res, err := util.HomeAccess(command, requestParam)
		return events.APIGatewayProxyResponse{
			Body:       string(res),
			StatusCode: 200,
		}, err
    } else {
        res, err := util.S3Access(command, requestParam)
		return events.APIGatewayProxyResponse{
            Headers: map[string]string{
                "Access-Control-Allow-Methods": "OPTIONS,POST",
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
                "Content-Type":                 "application/x-www-form-urlencoded, application/json",
            },
            Body:       string(res),
			StatusCode: 200,
		}, err
    }

}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
