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
    jsonParam, err := json.Marshal(param)
    if err != nil {
        return []byte{}, err
    }
    req, err := http.NewRequest("POST", str, bytes.NewBuffer(jsonParam))
    if err != nil {
        return []byte{}, err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return []byte{}, err
    }
    defer resp.Body.Close()

    byteArray, err := ioutil.ReadAll(resp.Body)
    return byteArray, err
}
