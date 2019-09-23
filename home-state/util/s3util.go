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

type S3UtilTest struct {
    S3GetBuffer *aws.WriteAtBuffer
    S3GetError error
    S3GetMock bool
    S3PutError error
    S3PutMock bool
}
var S3UtilTestSetting = S3UtilTest{nil, nil, false, nil, false}

func S3Get (key string) (*aws.WriteAtBuffer, error) {
    if S3UtilTestSetting.S3GetMock == true {
        return S3UtilTestSetting.S3GetBuffer,S3UtilTestSetting.S3GetError
    }
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
    if S3UtilTestSetting.S3PutMock == true {
        return S3UtilTestSetting.S3PutError
    }
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

