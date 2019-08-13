package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type genericRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      uint                   `json:"id"`
	Auth    string                 `json:"auth"`
}
type genericResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      uint        `json:"id"`
}

func requestConstruct() genericRequest {
	return genericRequest{JSONRPC: "2.0", ID: 1}
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
	fmt.Println(string(body))
	return
}
