package util

import (
	//"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Get (key string) (*aws.WriteAtBuffer, error) {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
    buffer := aws.NewWriteAtBuffer([]byte{})
	// _ -> numBytes
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String("home-state"),
			Key:    aws.String(key),
		})

    return buffer, err
}

