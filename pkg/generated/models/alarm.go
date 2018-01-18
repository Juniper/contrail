package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	AlarmRules    *AlarmOrList   `json:"alarm_rules,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
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
		AlarmRules:    MakeAlarmOrList(),
		DisplayName:   "",
		ParentUUID:    "",
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		UveKeys:       MakeUveKeysType(),
		AlarmSeverity: MakeAlarmSeverity(),
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		UUID:          "",
		FQName:        []string{},
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
