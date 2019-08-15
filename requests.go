package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type genericRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      uint                   `json:"id"`
	Auth    string                 `json:"auth"`
}
type genericResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	ID      uint            `json:"id"`
	Error   zabbixError     `json:"error"`
}
type zabbixError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// func (r *genericResponse) UnmarshallJSON(data []byte) error {
//
// }

func requestConstruct(method string) genericRequest {
	return genericRequest{JSONRPC: "2.0", ID: 1, Method: method}
}

func (req *genericRequest) Send(w *APIWrapper) (resp genericResponse, err error) {
	req.Auth = w.Token
	reqJSON, err := json.Marshal(*req)
	if err != nil {
		err = fmt.Errorf("Error while marshalling request: %s", err.Error())
		return
	}
	r := bytes.NewReader(reqJSON)
	httpResp, err := http.Post(w.Address, "application/json-rpc", r)
	if err != nil {
		err = fmt.Errorf("Error while making post request: %s", err.Error())
		return
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		err = fmt.Errorf("Error while reading response body: %s", err.Error())
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("Error while unmarshalling response: %s", err.Error())
		return
	}
	return
}

// ConvertParamsToMap accepts struct or map and converts it into map[string]interface{} to put it into generic request
func ConvertParamsToMap(params interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	switch reflect.TypeOf(params).Kind() {
	case reflect.Struct:
		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			k := val.Type().Field(i).Name
			v := val.Field(i).Interface()
			result[k] = v
		}
	case reflect.Map:
		iter := reflect.ValueOf(params).MapRange()
		for iter.Next() {
			k := iter.Key().String()
			v := iter.Value().Interface()
			result[k] = v
		}
	default:
		return nil, fmt.Errorf("Error: params is not map or struct")
	}
	return result, nil
}
