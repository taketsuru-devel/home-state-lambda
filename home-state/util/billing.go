package util

import (
    "sort"
    "fmt"
	"time"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
)

func GetBilling() *[]*cloudwatch.Datapoint {
    //請求はバージニア北部(us-east-1)からしか取れない
    client := cloudwatch.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

    now := time.Now()
    dateOfFirst := time.Date(now.Year(),now.Month(),1,0,0,0,0, time.UTC)

    req := &cloudwatch.GetMetricStatisticsInput{
        Namespace: aws.String("AWS/Billing"),
        MetricName: aws.String("EstimatedCharges"),
        Period: aws.Int64(86400), //一日刻みのデータ
        StartTime: aws.Time(dateOfFirst),
        EndTime: aws.Time(dateOfFirst.AddDate(0,0,now.Day())),
        Statistics: []*string{
            aws.String(cloudwatch.StatisticMaximum),
        },
        Dimensions: []*cloudwatch.Dimension{
            {
                Name: aws.String("Currency"),
                Value: aws.String("USD"),
            },
        },
        Unit: aws.String(cloudwatch.StandardUnitNone),
    }

    resp, err := client.GetMetricStatistics(req)
    if err != nil {
        fmt.Println(err)
    }

    ret := resp.Datapoints
    sort.Slice(ret, func(i int, j int) bool {
        return ret[i].Timestamp.Before(*(ret[j].Timestamp))
    })
    return &ret
}
