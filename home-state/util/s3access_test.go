package util

import (
	"errors"
	"testing"
	"github.com/aws/aws-sdk-go/aws"
)

func TestS3Access(t *testing.T) {
    //csvreader,ReadAll()やjson.Marshal()でエラーを起こすパスが思いつかない
    //そこでエラーを起こすなら事前の読み込みとかでエラーのはず
	t.Run("ItemExist", func(t *testing.T) {
        S3UtilTestSetting.S3GetMock = true
        S3UtilTestSetting.S3GetBuffer = aws.NewWriteAtBuffer([]byte{'t','e','s','t'})
        var p = Param{"test",[]string{""}}
        buf, err := S3Access("tekito", p)
        if err != nil {
			t.Fatalf("Error failed to existing item: %v", err)
        }
        if string(buf) == "" {
			t.Fatal("Something shuld be exist")
        }
    })
	t.Run("ItemNotExist", func(t *testing.T) {
        S3UtilTestSetting.S3GetBuffer = aws.NewWriteAtBuffer([]byte{})
        S3UtilTestSetting.S3GetError = errors.New("error")
        var p = Param{"test",[]string{""}}
        buf, err := S3Access("tekito", p)
        if string(buf) != "" {
			t.Fatal("should be empty")
        }
        if err == nil {
			t.Fatal("should be error")
        }
    })
}
