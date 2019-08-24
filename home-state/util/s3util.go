package util

import (
    "io"
	//"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const BUCKETNAME = "home-state"

func S3Get (key string) (*aws.WriteAtBuffer, error) {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
    buffer := aws.NewWriteAtBuffer([]byte{})
	// _ -> numBytes
	_, err := downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(BUCKETNAME),
			Key:    aws.String(key),
		})

    return buffer, err
}

func S3Put (key string, body io.Reader) error {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	// _ -> type UploadOutput
    _, err := uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(BUCKETNAME),
			Key:    aws.String(key),
            Body:   body,
		})

    return err
}

