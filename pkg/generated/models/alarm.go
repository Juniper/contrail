package models
// Alarm



import "encoding/json"

// Alarm 
//proteus:generate
type Alarm struct {

    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    AlarmRules *AlarmOrList `json:"alarm_rules,omitempty"`
    UveKeys *UveKeysType `json:"uve_keys,omitempty"`
    AlarmSeverity AlarmSeverity `json:"alarm_severity,omitempty"`


}



// String returns json representation of the object
func (model *Alarm) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

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
        AlarmSeverity: MakeAlarmSeverity(),
        
    }
}



// MakeAlarmSlice() makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
    return []*Alarm{}
}
