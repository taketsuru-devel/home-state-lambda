package util

import (
    "sort"
    "fmt"
	"time"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
)

type BillingType struct {
    Cost float64 `json:"cost"`
    Day int `json:"day"`
}

//func GetBilling() *[]*cloudwatch.Datapoint {
func GetBilling() *[]BillingType {
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

    datas := resp.Datapoints
    sort.Slice(datas, func(i int, j int) bool {
        return datas[i].Timestamp.Before(*(datas[j].Timestamp))
    })

    ret := make([]BillingType, len(datas))
    for i,v := range datas {
        ret[i].Cost = *(v.Maximum)
        ret[i].Day = v.Timestamp.Day()
    }
    return &ret
}
