package util

import (
	//"time"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type LogType struct {
    Event int
	Created string
	Request string
	RequestParam string
	LogMessage string
}

func LogPut(evt LogType) error {
    ddb := dynamo.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	table := ddb.Table("home-state-log")

	return table.Put(evt).Run()
}
