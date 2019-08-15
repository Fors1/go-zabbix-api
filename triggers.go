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
	Flags              uint8  `json:"flags"`
	LastChange         uint64 `json:"lastchange"`
	Priority           uint8  `json:"priority"`
	State              uint8  `json:"state"`
	Status             uint8  `json:"status"`
	TemplateID         string `json:"templateid"`
	Type               uint8  `json:"type"`
	URL                string `json:"url"`
	Value              uint8  `json:"value"`
	RecoveryMode       uint8  `json:"recovery_mode"`
	RecoveryExpression string `json:"recovery_expression"`
	CorrelationMode    uint8  `json:"correlation_mode"`
	CorrelationTag     string `json:"correlation_tag"`
	ManualClose        uint8  `json:"manual_close"`
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

// GetTrigger retrievs trigger objects from Zabbix
func (w *APIWrapper) GetTrigger(triggerParams TriggerParams) ([]Trigger, error) {
	params := convertStructToMap(&triggerParams)
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
