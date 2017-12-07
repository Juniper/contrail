package models

// OpenstackCluster

import "encoding/json"

// OpenstackCluster
type OpenstackCluster struct {
	AdminPassword                             string         `json:"admin_password"`
	DefaultJournalDrives                      string         `json:"default_journal_drives"`
	DefaultOsdDrives                          string         `json:"default_osd_drives"`
	Annotations                               *KeyValuePairs `json:"annotations"`
	Perms2                                    *PermType2     `json:"perms2"`
	ProvisioningLog                           string         `json:"provisioning_log"`
	DefaultPerformanceDrives                  string         `json:"default_performance_drives"`
	ExternalAllocationPoolEnd                 string         `json:"external_allocation_pool_end"`
	ExternalNetCidr                           string         `json:"external_net_cidr"`
	PublicIP                                  string         `json:"public_ip"`
	IDPerms                                   *IdPermsType   `json:"id_perms"`
	ProvisioningProgress                      int            `json:"provisioning_progress"`
	ProvisioningState                         string         `json:"provisioning_state"`
	ProvisioningStartTime                     string         `json:"provisioning_start_time"`
	DefaultCapacityDrives                     string         `json:"default_capacity_drives"`
	DefaultStorageAccessBondInterfaceMembers  string         `json:"default_storage_access_bond_interface_members"`
	OpenstackWebui                            string         `json:"openstack_webui"`
	PublicGateway                             string         `json:"public_gateway"`
	DisplayName                               string         `json:"display_name"`
	ProvisioningProgressStage                 string         `json:"provisioning_progress_stage"`
	ContrailClusterID                         string         `json:"contrail_cluster_id"`
	DefaultStorageBackendBondInterfaceMembers string         `json:"default_storage_backend_bond_interface_members"`
	ExternalAllocationPoolStart               string         `json:"external_allocation_pool_start"`
	FQName                                    []string       `json:"fq_name"`
	UUID                                      string         `json:"uuid"`
}

// OpenstackCluster parents relation object

// String returns json representation of the object
func (model *OpenstackCluster) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackCluster makes OpenstackCluster
func MakeOpenstackCluster() *OpenstackCluster {
	return &OpenstackCluster{
		//TODO(nati): Apply default
		ExternalAllocationPoolStart: "",
		FQName: []string{},
		UUID:   "",
		ProvisioningProgressStage:                 "",
		ContrailClusterID:                         "",
		DefaultStorageBackendBondInterfaceMembers: "",
		DefaultOsdDrives:                          "",
		AdminPassword:                             "",
		DefaultJournalDrives:                      "",
		ExternalNetCidr:                           "",
		PublicIP:                                  "",
		IDPerms:                                   MakeIdPermsType(),
		Annotations:                               MakeKeyValuePairs(),
		Perms2:                                    MakePermType2(),
		ProvisioningLog:                           "",
		DefaultPerformanceDrives:                  "",
		ExternalAllocationPoolEnd:                 "",
		ProvisioningProgress:                      0,
		ProvisioningState:                         "",
		OpenstackWebui:                            "",
		PublicGateway:                             "",
		DisplayName:                               "",
		ProvisioningStartTime:                     "",
		DefaultCapacityDrives:                     "",
		DefaultStorageAccessBondInterfaceMembers:  "",
	}
}

// InterfaceToOpenstackCluster makes OpenstackCluster from interface
func InterfaceToOpenstackCluster(iData interface{}) *OpenstackCluster {
	data := iData.(map[string]interface{})
	return &OpenstackCluster{
		ProvisioningLog: data["provisioning_log"].(string),

		//{"Title":"Provisioning Log","Description":"","SQL":"text","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_log","Item":null,"GoName":"ProvisioningLog","GoType":"string","GoPremitive":true}
		DefaultPerformanceDrives: data["default_performance_drives"].(string),

		//{"Title":"Default Performance Drive  for Controller Node Role","Description":"Drives for performance oriented application such as journaling  for Controller Node Role","SQL":"varchar(255)","Default":"sdf","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_performance_drives","Item":null,"GoName":"DefaultPerformanceDrives","GoType":"string","GoPremitive":true}
		ExternalAllocationPoolEnd: data["external_allocation_pool_end"].(string),

		//{"Title":"External Allocation pool end","Description":"End of the allocation pool range","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_allocation_pool_end","Item":null,"GoName":"ExternalAllocationPoolEnd","GoType":"string","GoPremitive":true}
		ExternalNetCidr: data["external_net_cidr"].(string),

		//{"Title":"External Network CIDR","Description":"Subnet to use for external network","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_net_cidr","Item":null,"GoName":"ExternalNetCidr","GoType":"string","GoPremitive":true}
		PublicIP: data["public_ip"].(string),

		//{"Title":"Public IP","Description":"Public Virtual IP (VIP)","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"public_ip","Item":null,"GoName":"PublicIP","GoType":"string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"Title":"Provisioning Progress","Description":"","SQL":"int","Default":0,"Operation":"","Presence":"","Type":"integer","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress","Item":null,"GoName":"ProvisioningProgress","GoType":"int","GoPremitive":true}
		ProvisioningState: data["provisioning_state"].(string),

		//{"Title":"Provisioning Status","Description":"","SQL":"varchar(255)","Default":"CREATED","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_state","Item":null,"GoName":"ProvisioningState","GoType":"string","GoPremitive":true}
		DefaultCapacityDrives: data["default_capacity_drives"].(string),

		//{"Title":"Default Capacity Drives  for Controller Node Role","Description":"Drives for capacity oriented applications such as logging for Controller Node Role","SQL":"varchar(255)","Default":"sdc,sdd,sde","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_capacity_drives","Item":null,"GoName":"DefaultCapacityDrives","GoType":"string","GoPremitive":true}
		DefaultStorageAccessBondInterfaceMembers: data["default_storage_access_bond_interface_members"].(string),

		//{"Title":"Default Storage Access  Bond Interface Members","Description":"Storage Management  Bond Interface Members","SQL":"varchar(255)","Default":"ens8f0,ens8f1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_storage_access_bond_interface_members","Item":null,"GoName":"DefaultStorageAccessBondInterfaceMembers","GoType":"string","GoPremitive":true}
		OpenstackWebui: data["openstack_webui"].(string),

		//{"Title":"OpenStack WebUI","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"openstack_webui","Item":null,"GoName":"OpenstackWebui","GoType":"string","GoPremitive":true}
		PublicGateway: data["public_gateway"].(string),

		//{"Title":"Public Gateway","Description":"Gateway for public VIP","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"public_gateway","Item":null,"GoName":"PublicGateway","GoType":"string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"Title":"Time provisioning started","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_start_time","Item":null,"GoName":"ProvisioningStartTime","GoType":"string","GoPremitive":true}
		ContrailClusterID: data["contrail_cluster_id"].(string),

		//{"Title":"Contrail Cluster ID","Description":"contrial cluster ID","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"contrail_cluster_id","Item":null,"GoName":"ContrailClusterID","GoType":"string","GoPremitive":true}
		DefaultStorageBackendBondInterfaceMembers: data["default_storage_backend_bond_interface_members"].(string),

		//{"Title":"Default Storage Backend Bond Interface Members","Description":"Storage Backend Bond Interface Members","SQL":"varchar(255)","Default":"ens9f0,ens9f1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_storage_backend_bond_interface_members","Item":null,"GoName":"DefaultStorageBackendBondInterfaceMembers","GoType":"string","GoPremitive":true}
		ExternalAllocationPoolStart: data["external_allocation_pool_start"].(string),

		//{"Title":"External Allocation pool start","Description":"Start of the allocation pool range","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"external_allocation_pool_start","Item":null,"GoName":"ExternalAllocationPoolStart","GoType":"string","GoPremitive":true}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"Title":"Provisioning Progress Stage","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress_stage","Item":null,"GoName":"ProvisioningProgressStage","GoType":"string","GoPremitive":true}
		AdminPassword: data["admin_password"].(string),

		//{"Title":"Admin Password","Description":"Password for admin openstack account","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"admin_password","Item":null,"GoName":"AdminPassword","GoType":"string","GoPremitive":true}
		DefaultJournalDrives: data["default_journal_drives"].(string),

		//{"Title":"Journal Drives  for Storage Node Role","Description":"SSD Drives to use for journals","SQL":"varchar(255)","Default":"sdf","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_journal_drives","Item":null,"GoName":"DefaultJournalDrives","GoType":"string","GoPremitive":true}
		DefaultOsdDrives: data["default_osd_drives"].(string),

		//{"Title":"Stoage Drives for Storage Node Role","Description":"Drives to use for cloud storage","SQL":"varchar(255)","Default":"sdc,sdd,sde","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_osd_drives","Item":null,"GoName":"DefaultOsdDrives","GoType":"string","GoPremitive":true}

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
