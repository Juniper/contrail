
package models
// GlobalSystemConfig



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propGlobalSystemConfig_plugin_tuning int = iota
    propGlobalSystemConfig_mac_aging_time int = iota
    propGlobalSystemConfig_bgp_always_compare_med int = iota
    propGlobalSystemConfig_graceful_restart_parameters int = iota
    propGlobalSystemConfig_display_name int = iota
    propGlobalSystemConfig_parent_uuid int = iota
    propGlobalSystemConfig_bgpaas_parameters int = iota
    propGlobalSystemConfig_mac_move_control int = iota
    propGlobalSystemConfig_ip_fabric_subnets int = iota
    propGlobalSystemConfig_autonomous_system int = iota
    propGlobalSystemConfig_fq_name int = iota
    propGlobalSystemConfig_alarm_enable int = iota
    propGlobalSystemConfig_ibgp_auto_mesh int = iota
    propGlobalSystemConfig_perms2 int = iota
    propGlobalSystemConfig_uuid int = iota
    propGlobalSystemConfig_config_version int = iota
    propGlobalSystemConfig_mac_limit_control int = iota
    propGlobalSystemConfig_annotations int = iota
    propGlobalSystemConfig_parent_type int = iota
    propGlobalSystemConfig_id_perms int = iota
    propGlobalSystemConfig_user_defined_log_statistics int = iota
)

// GlobalSystemConfig 
type GlobalSystemConfig struct {

    AutonomousSystem AutonomousSystemType `json:"autonomous_system,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    AlarmEnable bool `json:"alarm_enable"`
    MacMoveControl *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
    IPFabricSubnets *SubnetListType `json:"ip_fabric_subnets,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ConfigVersion string `json:"config_version,omitempty"`
    IbgpAutoMesh bool `json:"ibgp_auto_mesh"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    UserDefinedLogStatistics *UserDefinedLogStatList `json:"user_defined_log_statistics,omitempty"`
    MacLimitControl *MACLimitControlType `json:"mac_limit_control,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    BGPAlwaysCompareMed bool `json:"bgp_always_compare_med"`
    GracefulRestartParameters *GracefulRestartParametersType `json:"graceful_restart_parameters,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    BgpaasParameters *BGPaaServiceParametersType `json:"bgpaas_parameters,omitempty"`
    PluginTuning *PluginProperties `json:"plugin_tuning,omitempty"`
    MacAgingTime MACAgingTime `json:"mac_aging_time,omitempty"`

    BGPRouterRefs []*GlobalSystemConfigBGPRouterRef `json:"bgp_router_refs,omitempty"`

    Alarms []*Alarm `json:"alarms,omitempty"`
    AnalyticsNodes []*AnalyticsNode `json:"analytics_nodes,omitempty"`
    APIAccessLists []*APIAccessList `json:"api_access_lists,omitempty"`
    ConfigNodes []*ConfigNode `json:"config_nodes,omitempty"`
    DatabaseNodes []*DatabaseNode `json:"database_nodes,omitempty"`
    GlobalQosConfigs []*GlobalQosConfig `json:"global_qos_configs,omitempty"`
    GlobalVrouterConfigs []*GlobalVrouterConfig `json:"global_vrouter_configs,omitempty"`
    PhysicalRouters []*PhysicalRouter `json:"physical_routers,omitempty"`
    ServiceApplianceSets []*ServiceApplianceSet `json:"service_appliance_sets,omitempty"`
    VirtualRouters []*VirtualRouter `json:"virtual_routers,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// GlobalSystemConfigBGPRouterRef references each other
type GlobalSystemConfigBGPRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *GlobalSystemConfig) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeGlobalSystemConfig makes GlobalSystemConfig
func MakeGlobalSystemConfig() *GlobalSystemConfig{
    return &GlobalSystemConfig{
    //TODO(nati): Apply default
    MacMoveControl: MakeMACMoveLimitControlType(),
        IPFabricSubnets: MakeSubnetListType(),
        AutonomousSystem: MakeAutonomousSystemType(),
        FQName: []string{},
        AlarmEnable: false,
        IbgpAutoMesh: false,
        Perms2: MakePermType2(),
        UUID: "",
        ConfigVersion: "",
        MacLimitControl: MakeMACLimitControlType(),
        Annotations: MakeKeyValuePairs(),
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        UserDefinedLogStatistics: MakeUserDefinedLogStatList(),
        PluginTuning: MakePluginProperties(),
        MacAgingTime: MakeMACAgingTime(),
        BGPAlwaysCompareMed: false,
        GracefulRestartParameters: MakeGracefulRestartParametersType(),
        DisplayName: "",
        ParentUUID: "",
        BgpaasParameters: MakeBGPaaServiceParametersType(),
        
        modified: big.NewInt(0),
    }
}



// MakeGlobalSystemConfigSlice makes a slice of GlobalSystemConfig
func MakeGlobalSystemConfigSlice() []*GlobalSystemConfig {
    return []*GlobalSystemConfig{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *GlobalSystemConfig) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[config_root:0xc42024a780])
    fqn := []string{}
    
    fqn = ConfigRoot{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *GlobalSystemConfig) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-config_root", "_", "-", -1)
}

func (model *GlobalSystemConfig) GetDefaultName() string {
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *GlobalSystemConfig) GetType() string {
    return strings.Replace("global_system_config", "_", "-", -1)
}

func (model *GlobalSystemConfig) GetFQName() []string {
    return model.FQName
}

func (model *GlobalSystemConfig) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *GlobalSystemConfig) GetParentType() string {
    return model.ParentType
}

func (model *GlobalSystemConfig) GetUuid() string {
    return model.UUID
}

func (model *GlobalSystemConfig) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *GlobalSystemConfig) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *GlobalSystemConfig) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *GlobalSystemConfig) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *GlobalSystemConfig) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propGlobalSystemConfig_autonomous_system) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AutonomousSystem); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AutonomousSystem as autonomous_system")
        }
        msg["autonomous_system"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_alarm_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AlarmEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AlarmEnable as alarm_enable")
        }
        msg["alarm_enable"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_mac_move_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveControl as mac_move_control")
        }
        msg["mac_move_control"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_ip_fabric_subnets) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPFabricSubnets); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPFabricSubnets as ip_fabric_subnets")
        }
        msg["ip_fabric_subnets"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_config_version) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConfigVersion); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConfigVersion as config_version")
        }
        msg["config_version"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_ibgp_auto_mesh) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IbgpAutoMesh); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IbgpAutoMesh as ibgp_auto_mesh")
        }
        msg["ibgp_auto_mesh"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_user_defined_log_statistics) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UserDefinedLogStatistics); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UserDefinedLogStatistics as user_defined_log_statistics")
        }
        msg["user_defined_log_statistics"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_mac_limit_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLimitControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLimitControl as mac_limit_control")
        }
        msg["mac_limit_control"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_bgp_always_compare_med) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BGPAlwaysCompareMed); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BGPAlwaysCompareMed as bgp_always_compare_med")
        }
        msg["bgp_always_compare_med"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_graceful_restart_parameters) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GracefulRestartParameters); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GracefulRestartParameters as graceful_restart_parameters")
        }
        msg["graceful_restart_parameters"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_bgpaas_parameters) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasParameters); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasParameters as bgpaas_parameters")
        }
        msg["bgpaas_parameters"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_plugin_tuning) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PluginTuning); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PluginTuning as plugin_tuning")
        }
        msg["plugin_tuning"] = &val
    }
    
    if model.modified.Bit(propGlobalSystemConfig_mac_aging_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAgingTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAgingTime as mac_aging_time")
        }
        msg["mac_aging_time"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *GlobalSystemConfig) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *GlobalSystemConfig) UpdateReferences() error {
    return nil
}


