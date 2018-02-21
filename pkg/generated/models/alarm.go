
package models
// Alarm



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAlarm_parent_type int = iota
    propAlarm_id_perms int = iota
    propAlarm_perms2 int = iota
    propAlarm_alarm_rules int = iota
    propAlarm_alarm_severity int = iota
    propAlarm_fq_name int = iota
    propAlarm_display_name int = iota
    propAlarm_annotations int = iota
    propAlarm_uuid int = iota
    propAlarm_parent_uuid int = iota
    propAlarm_uve_keys int = iota
)

// Alarm 
type Alarm struct {

    AlarmRules *AlarmOrList `json:"alarm_rules,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    UveKeys *UveKeysType `json:"uve_keys,omitempty"`
    AlarmSeverity AlarmSeverity `json:"alarm_severity,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
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
    AlarmRules: MakeAlarmOrList(),
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        UveKeys: MakeUveKeysType(),
        AlarmSeverity: MakeAlarmSeverity(),
        FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAlarmSlice makes a slice of Alarm
func MakeAlarmSlice() []*Alarm {
    return []*Alarm{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Alarm) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc420182a00 project:0xc420182aa0])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *Alarm) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Alarm) GetDefaultName() string {
    return strings.Replace("default-alarm", "_", "-", -1)
}

func (model *Alarm) GetType() string {
    return strings.Replace("alarm", "_", "-", -1)
}

func (model *Alarm) GetFQName() []string {
    return model.FQName
}

func (model *Alarm) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Alarm) GetParentType() string {
    return model.ParentType
}

func (model *Alarm) GetUuid() string {
    return model.UUID
}

func (model *Alarm) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Alarm) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Alarm) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Alarm) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Alarm) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAlarm_alarm_rules) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AlarmRules); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AlarmRules as alarm_rules")
        }
        msg["alarm_rules"] = &val
    }
    
    if model.modified.Bit(propAlarm_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propAlarm_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAlarm_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAlarm_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAlarm_uve_keys) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UveKeys); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UveKeys as uve_keys")
        }
        msg["uve_keys"] = &val
    }
    
    if model.modified.Bit(propAlarm_alarm_severity) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AlarmSeverity); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AlarmSeverity as alarm_severity")
        }
        msg["alarm_severity"] = &val
    }
    
    if model.modified.Bit(propAlarm_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAlarm_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propAlarm_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAlarm_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Alarm) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Alarm) UpdateReferences() error {
    return nil
}


