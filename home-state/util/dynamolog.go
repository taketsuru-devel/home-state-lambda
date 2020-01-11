package util

import (
	//"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type LogType struct {
	Event        int
	Created      string
	Request      string
	RequestParam string
	LogMessage   string
}

func LogPut(evt LogType) error {
	ddb := dynamo.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	table := ddb.Table("home-state-log")

	return table.Put(evt).Run()
}

func LogGet(limit int64) *[]LogType {
	ddb := dynamo.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	table := ddb.Table("home-state-log")

	var logs []LogType
	table.Get("Event", 0).Order(false).Limit(limit).All(&logs)
	return &logs
}
