package models


// MakeAlarm makes Alarm
func MakeAlarm() *Alarm{
    return &Alarm{
    //TODO(nati): Apply default
    UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        AlarmRules: MakeAlarmOrList(),
        UveKeys: MakeUveKeysType(),
        AlarmSeverity: 0,
        
    }
}

// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
    return []*Alarm{}
}


