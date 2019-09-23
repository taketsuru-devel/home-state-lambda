package util

import (
	"bytes"
	//"fmt"
    //"net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

//func HomeAccess (subdir string, param url.Values) ([]byte, error) {
func HomeAccess (subdir string, param Param) ([]byte, error) {

    buf, err := S3Get("config/endpoint.txt")
    str := string(buf.Bytes())
    jsonParam, _ := json.Marshal(param)
    req, _ := http.NewRequest("POST", str, bytes.NewBuffer(jsonParam))

    //以下適当
    //param挟む
    client := &http.Client{}
    resp, _ := client.Do(req)
    //fmt.Println(str)
    //fmt.Println(param["command"])
    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
    return byteArray, err
}
