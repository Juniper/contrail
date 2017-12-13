package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	DefaultCapacityDrives                     string         `json:"default_capacity_drives"`
	IDPerms                                   *IdPermsType   `json:"id_perms"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time"`
	ProvisioningProgress                      int            `json:"provisioning_progress"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members"`
	ParentType                                string         `json:"parent_type"`
	ProvisioningLog                           string         `json:"provisioning_log"`
	DefaultOsdDrives                          string         `json:"default_osd_drives"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end"`
	AdminPassword                             string         `json:"admin_password"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members"`
	ProvisioningState                         string         `json:"provisioning_state"`
	DefaultJournalDrives                      string         `json:"default_journal_drives"`
	PublicIP                                  string         `json:"public_ip"`
	DisplayName                               string         `json:"display_name"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage"`
	FQName                                    []string       `json:"fq_name"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start"`
	ExternalNetCidr                           string         `json:"external_net_cidr"`
	PublicGateway                             string         `json:"public_gateway"`
	Perms2                                    *PermType2     `json:"perms2"`
	ParentUUID                                string         `json:"parent_uuid"`
	ContrailClusterID                         string         `json:"contrail_cluster_id"`
	Annotations                               *KeyValuePairs `json:"annotations"`
	UUID                                      string         `json:"uuid"`
	OpenstackWebui                            string         `json:"openstack_webui"`
}

// String returns json representation of the object
func (model *OpenstackCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster {
	return &OpenstackCluster{
		//TODO(nati): Apply default
		DefaultCapacityDrives:                    "",
		IDPerms:                                  MakeIdPermsType(),
		ProvisioningStartTime:                    "",
		ProvisioningProgress:                     0,
		DefaultStorageAccessBondInterfaceMembers: "",
		ParentType:                                "",
		ProvisioningLog:                           "",
		DefaultOsdDrives:                          "",
		ExternalAllocationPoolEnd:                 "",
		AdminPassword:                             "",
		DefaultPerformanceDrives:                  "",
		DefaultStorageBackendBondInterfaceMembers: "",
		ProvisioningState:                         "",
		DefaultJournalDrives:                      "",
		PublicIP:                                  "",
		DisplayName:                               "",
		ProvisioningProgressStage:                 "",
		FQName: []string{},
		ExternalAllocationPoolStart: "",
		ExternalNetCidr:             "",
		PublicGateway:               "",
		Perms2:                      MakePermType2(),
		ParentUUID:                  "",
		ContrailClusterID:           "",
		Annotations:                 MakeKeyValuePairs(),
		UUID:                        "",
		OpenstackWebui:              "",
	}
}

// InterfaceToOpenstackCluster makes OpenstackCluster from interface
func InterfaceToOpenstackCluster(iData interface{}) *OpenstackCluster {
	data := iData.(map[string]interface{})
	return &OpenstackCluster{
		DefaultCapacityDrives: data["default_capacity_drives"].(string),

		//{"title":"Default Capacity Drives  for Controller Node Role","description":"Drives for capacity oriented applications such as logging for Controller Node Role","default":"sdc,sdd,sde","type":"string","permission":["create","update"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		DefaultStorageAccessBondInterfaceMembers: data["default_storage_access_bond_interface_members"].(string),

		//{"title":"Default Storage Access  Bond Interface Members","description":"Storage Management  Bond Interface Members","default":"ens8f0,ens8f1","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DefaultOsdDrives: data["default_osd_drives"].(string),

		//{"title":"Stoage Drives for Storage Node Role","description":"Drives to use for cloud storage","default":"sdc,sdd,sde","type":"string","permission":["create","update"]}
		ExternalAllocationPoolEnd: data["external_allocation_pool_end"].(string),

		//{"title":"External Allocation pool end","description":"End of the allocation pool range","default":"","type":"string","permission":["create","update"]}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		AdminPassword: data["admin_password"].(string),

		//{"title":"Admin Password","description":"Password for admin openstack account","default":"","type":"string","permission":["create","update"]}
		DefaultPerformanceDrives: data["default_performance_drives"].(string),

		//{"title":"Default Performance Drive  for Controller Node Role","description":"Drives for performance oriented application such as journaling  for Controller Node Role","default":"sdf","type":"string","permission":["create","update"]}
		DefaultStorageBackendBondInterfaceMembers: data["default_storage_backend_bond_interface_members"].(string),

		//{"title":"Default Storage Backend Bond Interface Members","description":"Storage Backend Bond Interface Members","default":"ens9f0,ens9f1","type":"string","permission":["create","update"]}
		DefaultJournalDrives: data["default_journal_drives"].(string),

		//{"title":"Journal Drives  for Storage Node Role","description":"SSD Drives to use for journals","default":"sdf","type":"string","permission":["create","update"]}
		PublicIP: data["public_ip"].(string),

		//{"title":"Public IP","description":"Public Virtual IP (VIP)","default":"","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		ExternalAllocationPoolStart: data["external_allocation_pool_start"].(string),

		//{"title":"External Allocation pool start","description":"Start of the allocation pool range","default":"","type":"string","permission":["create","update"]}
		ExternalNetCidr: data["external_net_cidr"].(string),

		//{"title":"External Network CIDR","description":"Subnet to use for external network","default":"","type":"string","permission":["create","update"]}
		PublicGateway: data["public_gateway"].(string),

		//{"title":"Public Gateway","description":"Gateway for public VIP","default":"","type":"string","permission":["create","update"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ContrailClusterID: data["contrail_cluster_id"].(string),

		//{"title":"Contrail Cluster ID","description":"contrial cluster ID","default":"","type":"string","permission":["create"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		OpenstackWebui: data["openstack_webui"].(string),

		//{"title":"OpenStack WebUI","default":"","type":"string","permission":["create","update"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToOpenstackClusterSlice makes a slice of OpenstackCluster from interface
func InterfaceToOpenstackClusterSlice(data interface{}) []*OpenstackCluster {
	list := data.([]interface{})
	result := MakeOpenstackClusterSlice()
	for _, item := range list {
		result = append(result, InterfaceToOpenstackCluster(item))
	}
	return result
}

// MakeOpenstackClusterSlice() makes a slice of OpenstackCluster
func MakeOpenstackClusterSlice() []*OpenstackCluster {
	return []*OpenstackCluster{}
}
