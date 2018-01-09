package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	AlarmSeverity AlarmSeverity  `json:"alarm_severity"`
	ParentUUID    string         `json:"parent_uuid"`
	ParentType    string         `json:"parent_type"`
	FQName        []string       `json:"fq_name"`
	Perms2        *PermType2     `json:"perms2"`
	UUID          string         `json:"uuid"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules"`
	UveKeys       *UveKeysType   `json:"uve_keys"`
	IDPerms       *IdPermsType   `json:"id_perms"`
	DisplayName   string         `json:"display_name"`
	Annotations   *KeyValuePairs `json:"annotations"`
}

// String returns json representation of the object
func (model *Alarm) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAlarm makes Alarm
func MakeAlarm() *Alarm {
	return &Alarm{
		//TODO(nati): Apply default
		UveKeys:       MakeUveKeysType(),
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		AlarmRules:    MakeAlarmOrList(),
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		Perms2:        MakePermType2(),
		UUID:          "",
		AlarmSeverity: MakeAlarmSeverity(),
	}
}

// InterfaceToAlarm makes Alarm from interface
func InterfaceToAlarm(iData interface{}) *Alarm {
	data := iData.(map[string]interface{})
	return &Alarm{
		AlarmSeverity: InterfaceToAlarmSeverity(data["alarm_severity"]),

		//{"description":"Severity level for the alarm.","type":"integer","minimum":0,"maximum":2}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		AlarmRules: InterfaceToAlarmOrList(data["alarm_rules"]),

		//{"description":"Rules based on the UVE attributes specified as OR-of-ANDs of AlarmExpression template. Example: \"alarm_rules\": {\"or_list\": [{\"and_list\": [{AlarmExpression1}, {AlarmExpression2}, ...]}, {\"and_list\": [{AlarmExpression3}, {AlarmExpression4}, ...]}]}","type":"object","properties":{"or_list":{"type":"array","item":{"type":"object","properties":{"and_list":{"type":"array","item":{"type":"object","properties":{"operand1":{"type":"string"},"operand2":{"type":"object","properties":{"json_value":{"type":"string"},"uve_attribute":{"type":"string"}}},"operation":{"type":"string","enum":["==","!=","\u003c","\u003c=","\u003e","\u003e=","in","not in","range","size==","size!="]},"variables":{"type":"array","item":{"type":"string"}}}}}}}}}}
		UveKeys: InterfaceToUveKeysType(data["uve_keys"]),

		//{"description":"List of UVE tables or UVE objects where this alarm config should be applied. For example, rules based on NodeStatus UVE can be applied to multiple object types or specific uve objects such as analytics-node, config-node, control-node:\u003chostname\u003e, etc.,","type":"object","properties":{"uve_key":{"type":"array","item":{"type":"string"}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToAlarmSlice makes a slice of Alarm from interface
func InterfaceToAlarmSlice(data interface{}) []*Alarm {
	list := data.([]interface{})
	result := MakeAlarmSlice()
	for _, item := range list {
		result = append(result, InterfaceToAlarm(item))
	}
	return result
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
