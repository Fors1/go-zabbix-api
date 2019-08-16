package zabbix

import (
	"encoding/json"
	"fmt"
)

// Host is host object in Zabbix
type Host struct {
	HostID            string `json:"hostid"` //RO
	Host              string `json:"host"`
	Available         string `json:"available"` //RO
	Description       string `json:"description"`
	DisableUntil      string `json:"disable_until"`  //RO
	Error             string `json:"error"`          //RO
	ErrorsFrom        string `json:"errors_from"`    //RO
	Flags             string `json:"flags"`          //RO
	InventoryMode     string `json:"inventory_mode"` //WO
	IPMIAuthType      string `json:"ipmi_authtype"`
	IPMIAvailable     string `json:"ipmi_available"`     //RO
	IPMIDisableUntil  string `json:"ipmi_disable_until"` //RO
	IPMIError         string `json:"ipmi_error"`         //RO
	IPMIErrorsFrom    string `json:"ipmi_errors_from"`   //RO
	IPMIPassword      string `json:"ipmi_password"`
	IPMIPrivilege     string `json:"ipmi_privilege"`
	IPMIUsername      string `json:"ipmi_username"`
	JMXAvailable      string `json:"jmx_available"`      //RO
	JMXDisableUntil   string `json:"jmx_disable_until"`  //RO
	JMXError          string `json:"jmx_error"`          //RO
	JMXErrorsFrom     string `json:"jmx_errors_from"`    //RO
	MaintenanceFrom   string `json:"maintenance_from"`   //RO
	MaintenanceStatus string `json:"maintenance_status"` //RO
	MaintenanceType   string `json:"maintenance_type"`   //RO
	MaintenanceID     string `json:"maintenanceid"`      // RO
	Name              string `json:"name"`
	ProxyHostID       string `json:"proxy_hostid"`
	SNMPAvailable     string `json:"snmp_available"`     //RO
	SNMPDisableUntil  string `json:"snmp_disable_until"` //RO
	SNMPError         string `json:"snmp_error"`         //RO
	SNMPErrorsFrom    string `json:"snmp_errors_from"`   //RO
	Status            string `json:"status"`
	TLSConnect        string `json:"tls_connect"`
	TLSAccept         string `json:"tls_accept"`
	TLSIssuer         string `json:"tls_issuer"`
	TLSSubject        string `json:"tls_subject"`
	TLSPSKIdentity    string `json:"tls_psk_identity"`
	TLSPSK            string `json:"tls_psk"`
	AutoCompress      string `json:"auto_compress"`
}

// HostParams describe supported parameters for hosts
type HostParams struct {
	GroupIDs               []string               `json:"groupids,omitempty"`
	ApplicationIDs         []string               `json:"applicationids,omitempty"`
	DServiceIDs            []string               `json:"dserviceids,omitempty"`
	GraphIDs               []string               `json:"graphids,omitempty"`
	HostIDs                []string               `json:"hostids,omitempty"`
	HTTPTestIDs            []string               `json:"httptestids,omitempty"`
	InterfaceIDs           []string               `json:"interfaceids,omitempty"`
	ItemIDs                []string               `json:"itemids,omitempty"`
	MaintenanceIDs         []string               `json:"maintenanceids,omitempty"`
	MonitoredHosts         bool                   `json:"monitored_hosts,omitempty"`
	ProxyHosts             bool                   `json:"proxy_hosts,omitempty"`
	ProxyIDs               []string               `json:"proxyids,omitempty"`
	TemplatedHosts         bool                   `json:"templated_hosts,omitempty"`
	TemplateIDs            []string               `json:"templateids,omitempty"`
	TriggerIDs             []string               `json:"triggerids,omitempty"`
	WithItems              bool                   `json:"with_items,omitempty"`
	WithApplications       bool                   `json:"with_applications,omitempty"`
	WithGraphs             bool                   `json:"with_graphs,omitempty"`
	WithHTTPTests          bool                   `json:"with_httptests,omitempty"`
	WithMonitoredHTTPTests bool                   `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems     bool                   `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers  bool                   `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems   bool                   `json:"with_simple_graph_items,omitempty"`
	WithTriggers           bool                   `json:"with_triggers,omitempty"`
	WithInventory          bool                   `json:"withInventory,omitempty"`
	EvalType               uint8                  `json:"evaltype,omitempty"`
	Tags                   []string               `json:"tags,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	SelectGroups           map[string]interface{} `json:"selectGroups,omitempty"`
	SelectTags             map[string]interface{} `json:"selectTags,omitempty"`
	SelectApplications     map[string]interface{} `json:"selectApplications,omitempty"`
	SelectDiscoveries      map[string]interface{} `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule    map[string]interface{} `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs           map[string]interface{} `json:"selectGraphs,omitempty"`
	SelectHostDiscovery    map[string]interface{} `json:"selectHostDiscovery,omitempty"`
	SelectHTTPTests        map[string]interface{} `json:"selectHttpTests,omitempty"`
	SelectInterfaces       map[string]interface{} `json:"selectInterfaces,omitempty"`
	SelectInventory        map[string]interface{} `json:"selectInventory,omitempty"`
	SelectItems            map[string]interface{} `json:"selectItems,omitempty"`
	SelectMacros           map[string]interface{} `json:"selectMacros,omitempty"`
	SelectParentTemplates  map[string]interface{} `json:"selectParentTemplates,omitempty"`
	SelectScreens          map[string]interface{} `json:"selectScreens,omitempty"`
	SelectTriggers         []string               `json:"selectTriggers,omitempty"`
	LimitSelects           uint                   `json:"limitSelects,omitempty"`
	Search                 map[string]interface{} `json:"search,omitempty"`
	SearchInventory        map[string]interface{} `json:"searchInventory,omitempty"`
	Sortfield              []string               `json:"sortfield,omitempty"`
	CountOutput            bool                   `json:"countOutput,omitempty"`
	Editable               bool                   `json:"editable,omitempty"`
	ExcludeSearch          bool                   `json:"excludeSearch,omitempty"`
	Limit                  uint                   `json:"limit,omitempty"`
	Output                 map[string]interface{} `json:"output,omitempty"`
	PreserveKeys           bool                   `json:"preservekeys,omitempty"`
	SearchByAny            bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortOrder              []string               `json:"sortorder,omitempty"`
	StartSearch            bool                   `json:"startSearch,omitempty"`
}

// HostsGet returns host object from zabbix
func (w *APIWrapper) HostsGet(hostParams HostParams) ([]Host, error) {
	req := requestConstruct("host.get")
	b, _ := json.Marshal(hostParams)       //
	params := make(map[string]interface{}) // this is to convert HostParams to map[string]interface{}. ugly but working
	json.Unmarshal(b, &params)             //
	req.Params = params
	resp, err := req.Send(w)
	if err != nil {
		return []Host{}, fmt.Errorf("Error while sending request to zabbix - %s", err.Error())
	}
	if resp.Error.Code != 0 {
		return []Host{}, fmt.Errorf("Zabbix Server returned error: %d - %s; %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}
	host := []Host{}
	err = json.Unmarshal(resp.Result, &host)
	if err != nil {
		return []Host{}, fmt.Errorf("Error while unmarshalling response json - %s", err.Error())
	}
	return host, nil
}
