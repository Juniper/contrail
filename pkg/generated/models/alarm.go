package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeAlarm makes Alarm
func MakeAlarm() *Alarm {
	return &Alarm{
		//TODO(nati): Apply default
		UUID:          "",
		ParentUUID:    "",
		ParentType:    "",
		FQName:        []string{},
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Perms2:        MakePermType2(),
		AlarmRules:    MakeAlarmOrList(),
		UveKeys:       MakeUveKeysType(),
		AlarmSeverity: 0,
	}
}

// MakeAlarm makes Alarm
func InterfaceToAlarm(i interface{}) *Alarm {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Alarm{
		//TODO(nati): Apply default
		UUID:          schema.InterfaceToString(m["uuid"]),
		ParentUUID:    schema.InterfaceToString(m["parent_uuid"]),
		ParentType:    schema.InterfaceToString(m["parent_type"]),
		FQName:        schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:       InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:   schema.InterfaceToString(m["display_name"]),
		Annotations:   InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:        InterfaceToPermType2(m["perms2"]),
		AlarmRules:    InterfaceToAlarmOrList(m["alarm_rules"]),
		UveKeys:       InterfaceToUveKeysType(m["uve_keys"]),
		AlarmSeverity: schema.InterfaceToInt64(m["alarm_severity"]),
	}
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
	return []*Alarm{}
}

// InterfaceToAlarmSlice() makes a slice of Alarm
func InterfaceToAlarmSlice(i interface{}) []*Alarm {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Alarm{}
	for _, item := range list {
		result = append(result, InterfaceToAlarm(item))
	}
	return result
}
