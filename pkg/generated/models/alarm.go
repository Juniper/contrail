package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	ParentUUID    string         `json:"parent_uuid,omitempty"`
	ParentType    string         `json:"parent_type,omitempty"`
	FQName        []string       `json:"fq_name,omitempty"`
	IDPerms       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName   string         `json:"display_name,omitempty"`
	UUID          string         `json:"uuid,omitempty"`
	UveKeys       *UveKeysType   `json:"uve_keys,omitempty"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity,omitempty"`
	Perms2        *PermType2     `json:"perms2,omitempty"`
	Annotations   *KeyValuePairs `json:"annotations,omitempty"`
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
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		UUID:          "",
		UveKeys:       MakeUveKeysType(),
		AlarmSeverity: MakeAlarmSeverity(),
		Perms2:        MakePermType2(),
		Annotations:   MakeKeyValuePairs(),
		AlarmRules:    MakeAlarmOrList(),
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
