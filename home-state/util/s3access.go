package util

import (
	//"bytes"
	//"fmt"
	//"net/url"
	"encoding/csv"
	"encoding/json"
	"strings"
)

//func HomeAccess (subdir string, param url.Values) ([]byte, error) {
func S3Access(subdir string, param Param) ([]byte, error) {

	//sample
	buf, err := S3Get("data/" + param.Params[0] + ".csv")
	if err != nil {
		return []byte{}, err
	}
	str := string(buf.Bytes())
	r := csv.NewReader(strings.NewReader(str))
	result, err := r.ReadAll()
	if err != nil {
		return []byte{}, err
	}
	json_result, err := json.Marshal(result)

	return json_result, err
}
