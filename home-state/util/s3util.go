package util

import (
	//"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3GetBuffer struct {
    Buffer *aws.WriteAtBuffer
}

// buffer = aws.NewWriteAtBufferを返しても関数内テンポラリなので参照先は消える
func (res *S3GetBuffer) Get (key string) error {
	sess := session.Must(session.NewSession())
	downloader := s3manager.NewDownloader(sess)
	res.Buffer = aws.NewWriteAtBuffer([]byte{})
	// _ -> numBytes
	_, err := downloader.Download(res.Buffer,
		&s3.GetObjectInput{
			Bucket: aws.String("home-state"),
			Key:    aws.String(key),
		})

    return err
}

