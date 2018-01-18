package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
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
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		UveKeys:       MakeUveKeysType(),
		AlarmSeverity: MakeAlarmSeverity(),
		UUID:          "",
		ParentUUID:    "",
		Annotations:   MakeKeyValuePairs(),
		AlarmRules:    MakeAlarmOrList(),
		Perms2:        MakePermType2(),
		FQName:        []string{},
		DisplayName:   "",
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
