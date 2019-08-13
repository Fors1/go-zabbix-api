package zabbix

import (
	"encoding/json"
	"fmt"
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
	res, err := req.Send(w)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(fmt.Sprintf("%+v", res.Result)), &services)
	return
}
