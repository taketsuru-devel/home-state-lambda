package main

import (
	"encoding/json"
	//"fmt"
    //"bytes"
    "net/http"
    "time"
    "strings"
    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/followedwind/home-state-lambda/home-state/util"
)


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    var requestParam util.Param
    json.Unmarshal([]byte(request.Body), &requestParam)
    t := time.Now()
    t_string := t.Format("2006-01-02 15:04:05")
    err_strs := []string{}
    command := requestParam.Command

    var res []byte
    var err error
    if command == "etherwake" || command == "aircon" || command == "picture"  {
        res, err = util.HomeAccess(command, requestParam)
    } else if command == "log" {
        logs := util.LogGet(5)
        res, err = json.Marshal(*logs)
    } else {
        res, err = util.S3Access(command, requestParam)
    }
    if err != nil {
        err_strs = append(err_strs, err.Error())
    }
	evt := util.LogType{0, t_string, command, strings.Join(requestParam.Params,","), strings.Join(err_strs,",")}
	err = util.LogPut(evt)
    if err != nil {
        err_strs = append(err_strs, err.Error())
    }

    if len(err_strs) > 0 {
        return makeResponse(http.StatusInternalServerError, strings.Join(err_strs,","), err)
    } else {
        return makeResponse(http.StatusOK, string(res), nil)
    }
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

func makeResponse(statusCode int, body string, err error) (events.APIGatewayProxyResponse,error) {
	return events.APIGatewayProxyResponse{
        Headers: map[string]string{
            "Access-Control-Allow-Methods": "OPTIONS,POST",
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
            "Content-Type":                 "application/x-www-form-urlencoded, application/json",
        },
		Body:       body,
		StatusCode: statusCode,
	}, err

}
