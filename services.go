package zabbix

// Service Zabbix object
type Service struct {
	ServiceID string
	Algorithm uint8
	Name      string
	ShowSLA   uint8
	SortOrder uint16
	GoodSLA   float32
	Status    uint8
	TriggerID string
}

// ServiceTime Zabbix object
type ServiceTime struct {
	TimeID    string
	ServiceID string
	TSFrom    uint16
	TSTo      uint16
	Type      uint8
	Note      string
}

// ServiceDependency Zabbix object
type ServiceDependency struct {
	LinkID        string
	ServiceDownID string
	ServiceUpID   string
	Soft          uint8
}

// ServiceAlarm Zabbix object
type ServiceAlarm struct {
	ServiceAlarmID string
	ServiceID      string
	Clock          uint32
	Value          uint8
}
