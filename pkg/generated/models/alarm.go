package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	FQName        []string       `json:"fq_name,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules,omitempty"`
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
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		UUID:          "",
		UveKeys:       MakeUveKeysType(),
		FQName:        []string{},
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		ParentUUID:    "",
		AlarmRules:    MakeAlarmOrList(),
		AlarmSeverity: MakeAlarmSeverity(),
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
