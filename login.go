package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIWrapper is main API handler for Zabbix API
type APIWrapper struct {
	Token   string
	Address string
}

type loginRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  someParams  `json:"params"`
	ID      uint        `json:"id"`
	Auth    interface{} `json:"auth"`
}
type loginResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      uint   `json:"id"`
}

type someParams map[string]string

// Login into Zabbix
func Login(login, password, address string) (w APIWrapper, err error) {
	req := loginRequest{}
	req.JSONRPC = "2.0"
	req.Method = "user.login"
	req.Params = make(map[string]string)
	req.Params["user"] = login
	req.Params["password"] = password
	req.ID = 1
	req.Auth = nil
	reqJSON, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("Error while marshalling request: %s", err.Error())
		return
	}
	r := bytes.NewReader(reqJSON)
	resp, err := http.Post(address, "application/json-rpc", r)
	if err != nil {
		err = fmt.Errorf("Error while making post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error while reading response body: %s", err.Error())
	}
	result := loginResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		err = fmt.Errorf("Error while unmarshalling response: %s", err.Error())
	}
	w.Token = result.Result
	return
}
