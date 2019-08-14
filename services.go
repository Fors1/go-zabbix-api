package zabbix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Service Zabbix object
type Service struct {
	ServiceID    string              `json:"serviceid"`
	Algorithm    uint8               `json:"algorithm"`
	Name         string              `json:"name"`
	ShowSLA      uint8               `json:"showsla"`
	SortOrder    uint16              `json:"sortorder"`
	GoodSLA      float32             `json:"goodsla"`
	Status       uint8               `json:"status"`
	TriggerID    string              `json:"triggerid"`
	Dependencies []ServiceDependency `json:"dependencies"`
}

// ServiceTime Zabbix object
type ServiceTime struct {
	TimeID    string `json:"timeid"`
	ServiceID string `json:"serviceid"`
	TSFrom    uint16 `json:"ts_from"`
	TSTo      uint16 `json:"ts_to"`
	Type      uint8  `json:"type"`
	Note      string `json:"note"`
}

// ServiceDependency Zabbix object
type ServiceDependency struct {
	LinkID        string `json:"linkid"`
	ServiceDownID string `json:"servicedownid"`
	ServiceUpID   string `json:"serviceupid"`
	Soft          uint8  `json:"soft"`
}

// ServiceAlarm Zabbix object
type ServiceAlarm struct {
	ServiceAlarmID string `json:"servicealarmid"`
	ServiceID      string `json:"serviceid"`
	Clock          uint32 `json:"clock"`
	Value          uint8  `json:"value"`
}

// GetServices gets zabbix IT Services
func (w *APIWrapper) GetServices(params map[string]interface{}) (services []Service, err error) {
	req := requestConstruct("service.get")
	req.Params = params
	resp, err := req.Send(w)
	if err != nil {
		return
	}
	if resp.Error.Code != 0 {
		return services, fmt.Errorf("Zabbix Server returned error: %d - %s; %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}
	err = json.Unmarshal([]byte(resp.Result), &services)
	if err != nil {
		return services, fmt.Errorf("Error while unmarshalling response json - %s", err.Error())
	}
	return
}

//func (s *Service) GetSLA(params map[string]interface{}) () {}

// CreateService creates service and returns ID of it
func (w *APIWrapper) CreateService(s Service) (int, error) {
	req := requestConstruct("service.create")
	params := make(map[string]interface{})
	params["name"] = s.Name
	params["algorithm"] = s.Algorithm
	params["showsla"] = s.ShowSLA
	params["goodsla"] = s.GoodSLA
	params["sortorder"] = s.SortOrder
	if s.TriggerID != "" {
		params["triggerid"] = s.TriggerID
	}
	req.Params = params
	resp, err := req.Send(w)
	if err != nil {
		return 0, err
	}
	if resp.Error.Code != 0 {
		return 0, fmt.Errorf("Zabbix Server returned error: %d - %s; %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}
	resultMap := make(map[string][]string)
	err = json.Unmarshal(resp.Result, &resultMap)
	if err != nil {
		return 0, fmt.Errorf("Error while unmarshalling response json - %s", err.Error())
	}
	if len(resultMap["serviceids"]) == 0 {
		return 0, fmt.Errorf("Error - no service IDs were returned")
	}
	serviceID, err := strconv.Atoi(resultMap["serviceids"][0])
	if err != nil {
		return 0, fmt.Errorf("Error while parsing created service ID returned by Zabbix - %s", err.Error())
	}
	return serviceID, nil
}

//AddDependency links one service to another one
func (w *APIWrapper) AddDependency(parentServiceID, childServiceID, soft int) error {
	req := requestConstruct("service.adddependencies")
	params := make(map[string]interface{})
	params["serviceid"] = strconv.Itoa(parentServiceID) // Convert to string because zabbix accepts this parameters as strings
	params["dependsOnServiceid"] = strconv.Itoa(childServiceID)
	params["soft"] = soft
	req.Params = params
	resp, err := req.Send(w)
	if err != nil {
		return err
	}
	if resp.Error.Code != 0 {
		return fmt.Errorf("Zabbix Server returned error: %d - %s; %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}
	return nil
}
