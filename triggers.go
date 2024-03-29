package zabbix

import (
	"encoding/json"
	"fmt"
)

//Trigger is trigger object from Zabbix
type Trigger struct {
	TriggerID          string `json:"triggerid"`
	Description        string `json:"description"`
	Expression         string `json:"expression"`
	Comments           string `json:"comments"`
	Error              string `json:"error"`
	Flags              string `json:"flags"`
	LastChange         string `json:"lastchange"`
	Priority           string `json:"priority"`
	State              string `json:"state"`
	Status             string `json:"status"`
	TemplateID         string `json:"templateid"`
	Type               string `json:"type"`
	URL                string `json:"url"`
	Value              string `json:"value"`
	RecoveryMode       string `json:"recovery_mode"`
	RecoveryExpression string `json:"recovery_expression"`
	CorrelationMode    string `json:"correlation_mode"`
	CorrelationTag     string `json:"correlation_tag"`
	ManualClose        string `json:"manual_close"`
}

//TriggerTag is trigger tag object from Zabbix
type TriggerTag struct {
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Operator uint8  `json:"operator,omitempty"`
}

//TriggerParams describes parameters for trigger objects
type TriggerParams struct {
	TriggerIDs                  []string               `json:"triggerids,omitempty"`
	GroupIDs                    []string               `json:"groupids,omitempty"`
	TemplateIDs                 []string               `json:"templateids,omitempty"`
	HostIDs                     []string               `json:"hostids,omitempty"`
	ItemIDs                     []string               `json:"itemids,omitempty"`
	ApplicationIDs              []string               `json:"applicationids,omitempty"`
	Functions                   []string               `json:"functions,omitempty"`
	Group                       string                 `json:"group,omitempty"`
	Host                        string                 `json:"host,omitempty"`
	Inherited                   bool                   `json:"inherited,omitempty"`
	Templated                   bool                   `json:"templated,omitempty"`
	Dependent                   bool                   `json:"dependent,omitempty"`
	Monitored                   bool                   `json:"monitored,omitempty"`
	Active                      bool                   `json:"active,omitempty"`
	Maintenance                 bool                   `json:"maintenance,omitempty"`
	WithUnacknowledgedEvents    bool                   `json:"withUnacknowledgedEvents,omitempty"`
	WithAcknowledgedEvents      bool                   `json:"withAcknowledgedEvents,omitempty"`
	WithLastEventUnacknowledged bool                   `json:"withLastEventUnacknowledged,omitempty"`
	SkipDependent               bool                   `json:"skipDependent,omitempty"`
	LastChangeSince             uint64                 `json:"lastChangeSince,omitempty"`
	LastChangeTill              uint64                 `json:"lastChangeTill,omitempty"`
	OnlyTrue                    bool                   `json:"only_true,omitempty"`
	MinSeverity                 uint8                  `json:"min_severity,omitempty"`
	EvalType                    uint8                  `json:"evaltype,omitempty"`
	Tags                        []TriggerTag           `json:"tags,omitempty"`
	ExpandComment               bool                   `json:"expandComment,omitempty"`
	ExpandDescription           bool                   `json:"expandDescription,omitempty"`
	ExpandExpression            bool                   `json:"expandExpression,omitempty"`
	SelectGroups                []string               `json:"selectGroups,omitempty"`
	SelectHosts                 []string               `json:"selectHosts,omitempty"`
	SelectItems                 []string               `json:"selectItems,omitempty"`
	SelectFunctions             []string               `json:"selectFunctions,omitempty"`
	SelectDependencies          []string               `json:"selectDependencies,omitempty"`
	SelectDiscoveryRule         []string               `json:"selectDiscoveryRule,omitempty"`
	SelectLastEvent             []string               `json:"selectLastEvent,omitempty"`
	SelectTags                  []string               `json:"selectTags,omitempty"`
	SelectTriggerDiscovery      []string               `json:"selectTriggerDiscovery,omitempty"`
	Filter                      map[string]interface{} `json:"filter,omitempty"`
	LimitSelects                uint                   `json:"limitSelects,omitempty"`
	SortField                   []string               `json:"sortfield,omitempty"`
	CountOutput                 bool                   `json:"countOutput,omitempty"`
	Editable                    bool                   `json:"editable,omitempty"`
	ExcludeSearch               bool                   `json:"excludeSearch,omitempty"`
	Limit                       uint                   `json:"limit,omitempty"`
	Output                      []string               `json:"output,omitempty"`
	PreserveKeys                bool                   `json:"preservekeys,omitempty"`
	Search                      map[string]interface{} `json:"search,omitempty"`
	SearchByAny                 bool                   `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled      bool                   `json:"searchWildcardsEnabled,omitempty"`
	SortOrder                   []string               `json:"sortorder,omitempty"`
	StartSearch                 bool                   `json:"startSearch,omitempty"`
}

// TriggersGet retrievs trigger objects from Zabbix
func (w *APIWrapper) TriggersGet(triggerParams TriggerParams) ([]Trigger, error) {
	b, _ := json.Marshal(triggerParams)    //
	params := make(map[string]interface{}) // this is to convert HostParams to map[string]interface{}. ugly but working
	json.Unmarshal(b, &params)             //
	req := requestConstruct("trigger.get")
	req.Params = params
	resp, err := req.Send(w)
	if err != nil {
		return []Trigger{}, fmt.Errorf("Error while sending request to zabbix - %s", err.Error())
	}
	if resp.Error.Code != 0 {
		return []Trigger{}, fmt.Errorf("Zabbix Server returned error: %d - %s; %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}
	t := []Trigger{}
	err = json.Unmarshal(resp.Result, &t)
	if err != nil {
		return []Trigger{}, fmt.Errorf("Error while unmarshalling response json - %s", err.Error())
	}
	return t, nil
}
