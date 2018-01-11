package models

// Alarm

import "encoding/json"

// Alarm
type Alarm struct {
	UUID          string         `json:"uuid"`
	ParentType    string         `json:"parent_type"`
	DisplayName   string         `json:"display_name"`
	IDPerms       *IdPermsType   `json:"id_perms"`
	Annotations   *KeyValuePairs `json:"annotations"`
	AlarmRules    *AlarmOrList   `json:"alarm_rules"`
	UveKeys       *UveKeysType   `json:"uve_keys"`
	AlarmSeverity AlarmSeverity  `json:"alarm_severity"`
	Perms2        *PermType2     `json:"perms2"`
	ParentUUID    string         `json:"parent_uuid"`
	FQName        []string       `json:"fq_name"`
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
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		Annotations:   MakeKeyValuePairs(),
		AlarmRules:    MakeAlarmOrList(),
		UveKeys:       MakeUveKeysType(),
		DisplayName:   "",
		UUID:          "",
		ParentType:    "",
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}
