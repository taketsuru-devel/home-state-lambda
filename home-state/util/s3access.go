package util

import (
	//"bytes"
	//"fmt"
    //"net/url"
    "encoding/json"
    "encoding/csv"
    "strings"
)

//func HomeAccess (subdir string, param url.Values) ([]byte, error) {
func S3Access (subdir string, param Param) ([]byte, error) {

    //sample
    buf, err := S3Get("data/"+param.Params[0]+".csv")
    str := string(buf.Bytes())
    r := csv.NewReader(strings.NewReader(str))
    result, _ := r.ReadAll()

    json, _ := json.Marshal(result)

    return json, err
}
