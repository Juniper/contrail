package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	Perms2        *PermType2     `json:"perms2,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
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
		ParentUUID:    "",
		IDPerms:       MakeIdPermsType(),
		Annotations:   MakeKeyValuePairs(),
		UveKeys:       MakeUveKeysType(),
		AlarmSeverity: MakeAlarmSeverity(),
		Perms2:        MakePermType2(),
		UUID:          "",
		ParentType:    "",
		FQName:        []string{},
		DisplayName:   "",
		AlarmRules:    MakeAlarmOrList(),
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
