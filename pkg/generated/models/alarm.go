package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
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
		AlarmSeverity: MakeAlarmSeverity(),
		Perms2:        MakePermType2(),
		ParentUUID:    "",
		IDPerms:       MakeIdPermsType(),
		AlarmRules:    MakeAlarmOrList(),
		UveKeys:       MakeUveKeysType(),
		UUID:          "",
		ParentType:    "",
		FQName:        []string{},
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
