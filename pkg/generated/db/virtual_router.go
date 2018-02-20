package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualRouterQuery = "insert into `virtual_router` (`virtual_router_type`,`virtual_router_ip_address`,`virtual_router_dpdk_enabled`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteVirtualRouterQuery = "delete from `virtual_router` where uuid = ?"

// VirtualRouterFields is db columns for VirtualRouter
var VirtualRouterFields = []string{
   "virtual_router_type",
   "virtual_router_ip_address",
   "virtual_router_dpdk_enabled",
   "uuid",
   "share",
   "owner_access",
   "owner",
   "global_access",
   "parent_uuid",
   "parent_type",
   "user_visible",
   "permissions_owner_access",
   "permissions_owner",
   "other_access",
   "group_access",
   "group",
   "last_modified",
   "enable",
   "description",
   "creator",
   "created",
   "fq_name",
   "display_name",
   "key_value_pair",
   
}

// VirtualRouterRefFields is db reference fields for VirtualRouter
var VirtualRouterRefFields = map[string][]string{
   
    "network_ipam": []string{
        // <schema.Schema Value>
        "subnet",
        "allocation_pools",
        
    },
   
    "virtual_machine": []string{
        // <schema.Schema Value>
        
    },
   
}

// VirtualRouterBackRefFields is db back reference fields for VirtualRouter
var VirtualRouterBackRefFields = map[string][]string{
   
   
   "virtual_machine_interface": []string{
        "vrf_assign_rule",
        "vlan_tag_based_bridge_domain",
        "sub_interface_vlan_tag",
        "service_interface_type",
        "local_preference",
        "traffic_direction",
        "udp_port",
        "vtep_dst_mac_address",
        "vtep_dst_ip_address",
        "vni",
        "routing_instance",
        "nic_assisted_mirroring_vlan",
        "nic_assisted_mirroring",
        "nh_mode",
        "juniper_header",
        "encapsulation",
        "analyzer_name",
        "analyzer_mac_address",
        "analyzer_ip_address",
        "mac_address",
        "route",
        "fat_flow_protocol",
        "virtual_machine_interface_disable_policy",
        "dhcp_option",
        "virtual_machine_interface_device_owner",
        "key_value_pair",
        "allowed_address_pair",
        "uuid",
        "port_security_enabled",
        "share",
        "owner_access",
        "owner",
        "global_access",
        "parent_uuid",
        "parent_type",
        "user_visible",
        "permissions_owner_access",
        "permissions_owner",
        "other_access",
        "group_access",
        "group",
        "last_modified",
        "enable",
        "description",
        "creator",
        "created",
        "fq_name",
        "source_port",
        "source_ip",
        "ip_protocol",
        "hashing_configured",
        "destination_port",
        "destination_ip",
        "display_name",
        "annotations_key_value_pair",
        
   },
   
}

// VirtualRouterParentTypes is possible parents for VirtualRouter
var VirtualRouterParents = []string{
   
   "global_system_config",
   
}


const insertVirtualRouterNetworkIpamQuery = "insert into `ref_virtual_router_network_ipam` (`from`, `to` ,`subnet`,`allocation_pools`) values (?, ?,?,?);"

const insertVirtualRouterVirtualMachineQuery = "insert into `ref_virtual_router_virtual_machine` (`from`, `to` ) values (?, ?);"


// CreateVirtualRouter inserts VirtualRouter to DB
func CreateVirtualRouter(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.CreateVirtualRouterRequest) error {
    model := request.VirtualRouter
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualRouterQuery)
	if err != nil {
        return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
    log.WithFields(log.Fields{
        "model": model,
        "query": insertVirtualRouterQuery,
    }).Debug("create query")
    _, err = stmt.ExecContext(ctx, string(model.GetVirtualRouterType()),
    string(model.GetVirtualRouterIPAddress()),
    bool(model.GetVirtualRouterDPDKEnabled()),
    string(model.GetUUID()),
    common.MustJSON(model.GetPerms2().GetShare()),
    int(model.GetPerms2().GetOwnerAccess()),
    string(model.GetPerms2().GetOwner()),
    int(model.GetPerms2().GetGlobalAccess()),
    string(model.GetParentUUID()),
    string(model.GetParentType()),
    bool(model.GetIDPerms().GetUserVisible()),
    int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
    string(model.GetIDPerms().GetPermissions().GetOwner()),
    int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
    int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
    string(model.GetIDPerms().GetPermissions().GetGroup()),
    string(model.GetIDPerms().GetLastModified()),
    bool(model.GetIDPerms().GetEnable()),
    string(model.GetIDPerms().GetDescription()),
    string(model.GetIDPerms().GetCreator()),
    string(model.GetIDPerms().GetCreated()),
    common.MustJSON(model.GetFQName()),
    string(model.GetDisplayName()),
    common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
        return errors.Wrap(err, "create failed")
	}
    
    stmtNetworkIpamRef, err := tx.Prepare(insertVirtualRouterNetworkIpamQuery)
	if err != nil {
        return errors.Wrap(err,"preparing NetworkIpamRefs create statement failed")
	}
    defer stmtNetworkIpamRef.Close()
    for _, ref := range model.NetworkIpamRefs {
       
       if ref.Attr == nil {
           ref.Attr = &models.VirtualRouterNetworkIpamType{}
       }
       
        _, err = stmtNetworkIpamRef.ExecContext(ctx, model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetSubnet()),
    common.MustJSON(ref.Attr.GetAllocationPools()))
	    if err != nil {
            return errors.Wrap(err,"NetworkIpamRefs create failed")
        }
    }
    
    stmtVirtualMachineRef, err := tx.Prepare(insertVirtualRouterVirtualMachineQuery)
	if err != nil {
        return errors.Wrap(err,"preparing VirtualMachineRefs create statement failed")
	}
    defer stmtVirtualMachineRef.Close()
    for _, ref := range model.VirtualMachineRefs {
       
        _, err = stmtVirtualMachineRef.ExecContext(ctx, model.UUID, ref.UUID, )
	    if err != nil {
            return errors.Wrap(err,"VirtualMachineRefs create failed")
        }
    }
    
    metaData := &common.MetaData{
        UUID: model.UUID,
        Type: "virtual_router",
        FQName: model.FQName,
    }
    err = common.CreateMetaData(tx, metaData)
    if err != nil {
        return err
    }
    err = common.CreateSharing(tx, "virtual_router", model.UUID, model.GetPerms2().GetShare())
    if err != nil {
        return err
    }
    log.WithFields(log.Fields{
        "model": model,
    }).Debug("created")
    return nil
}

func scanVirtualRouter(values map[string]interface{} ) (*models.VirtualRouter, error) {
    m := models.MakeVirtualRouter()
    
    if value, ok := values["virtual_router_type"]; ok {
        
            
               m.VirtualRouterType = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["virtual_router_ip_address"]; ok {
        
            
               m.VirtualRouterIPAddress = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["virtual_router_dpdk_enabled"]; ok {
        
            
               m.VirtualRouterDPDKEnabled = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["uuid"]; ok {
        
            
               m.UUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["share"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Perms2.Share)
        
    }
    
    if value, ok := values["owner_access"]; ok {
        
            
               m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["owner"]; ok {
        
            
               m.Perms2.Owner = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["global_access"]; ok {
        
            
               m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["parent_uuid"]; ok {
        
            
               m.ParentUUID = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["parent_type"]; ok {
        
            
               m.ParentType = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["user_visible"]; ok {
        
            
               m.IDPerms.UserVisible = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["permissions_owner_access"]; ok {
        
            
               m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["permissions_owner"]; ok {
        
            
               m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["other_access"]; ok {
        
            
               m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["group_access"]; ok {
        
            
               m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)
            
        
    }
    
    if value, ok := values["group"]; ok {
        
            
               m.IDPerms.Permissions.Group = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["last_modified"]; ok {
        
            
               m.IDPerms.LastModified = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["enable"]; ok {
        
            
               m.IDPerms.Enable = schema.InterfaceToBool(value)
            
        
    }
    
    if value, ok := values["description"]; ok {
        
            
               m.IDPerms.Description = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["creator"]; ok {
        
            
               m.IDPerms.Creator = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["created"]; ok {
        
            
               m.IDPerms.Created = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["fq_name"]; ok {
        
            json.Unmarshal(value.([]byte), &m.FQName)
        
    }
    
    if value, ok := values["display_name"]; ok {
        
            
               m.DisplayName = schema.InterfaceToString(value)
            
        
    }
    
    if value, ok := values["key_value_pair"]; ok {
        
            json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)
        
    }
    
    
    if value, ok := values["ref_network_ipam"]; ok {
        var references []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &references )
        for _, reference := range references {
            referenceMap, ok := reference.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(referenceMap["to"])
            if uuid == "" {
                continue
            }
            referenceModel := &models.VirtualRouterNetworkIpamRef{}
            referenceModel.UUID = uuid
            m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)
            
            attr := models.MakeVirtualRouterNetworkIpamType()
            referenceModel.Attr = attr
            
            
        }
    }
    
    if value, ok := values["ref_virtual_machine"]; ok {
        var references []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &references )
        for _, reference := range references {
            referenceMap, ok := reference.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(referenceMap["to"])
            if uuid == "" {
                continue
            }
            referenceModel := &models.VirtualRouterVirtualMachineRef{}
            referenceModel.UUID = uuid
            m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)
            
        }
    }
    
    
    
    if value, ok := values["backref_virtual_machine_interface"]; ok {
        var childResources []interface{}
        stringValue := schema.InterfaceToString(value)
        json.Unmarshal([]byte("[" + stringValue + "]"), &childResources )
        for _, childResource := range childResources {
            childResourceMap, ok := childResource.(map[string]interface{})
            if !ok {
                continue
            }
            uuid := schema.InterfaceToString(childResourceMap["uuid"])
            if uuid == "" {
                continue
            }
            childModel := models.MakeVirtualMachineInterface()
            m.VirtualMachineInterfaces = append(m.VirtualMachineInterfaces, childModel)

            
                if propertyValue, ok := childResourceMap["vrf_assign_rule"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VRFAssignTable.VRFAssignRule)
                
                }
            
                if propertyValue, ok := childResourceMap["vlan_tag_based_bridge_domain"]; ok && propertyValue != nil {
                
                    
                        childModel.VlanTagBasedBridgeDomain = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["sub_interface_vlan_tag"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["service_interface_type"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.ServiceInterfaceType = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["local_preference"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.LocalPreference = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["traffic_direction"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["udp_port"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vtep_dst_mac_address"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vtep_dst_ip_address"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["vni"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["routing_instance"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nic_assisted_mirroring_vlan"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nic_assisted_mirroring"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["nh_mode"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["juniper_header"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["encapsulation"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_name"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_mac_address"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["analyzer_ip_address"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["mac_address"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceMacAddresses.MacAddress)
                
                }
            
                if propertyValue, ok := childResourceMap["route"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceHostRoutes.Route)
                
                }
            
                if propertyValue, ok := childResourceMap["fat_flow_protocol"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol)
                
                }
            
                if propertyValue, ok := childResourceMap["virtual_machine_interface_disable_policy"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceDisablePolicy = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["dhcp_option"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceDHCPOptionList.DHCPOption)
                
                }
            
                if propertyValue, ok := childResourceMap["virtual_machine_interface_device_owner"]; ok && propertyValue != nil {
                
                    
                        childModel.VirtualMachineInterfaceDeviceOwner = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceBindings.KeyValuePair)
                
                }
            
                if propertyValue, ok := childResourceMap["allowed_address_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)
                
                }
            
                if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {
                
                    
                        childModel.UUID = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["port_security_enabled"]; ok && propertyValue != nil {
                
                    
                        childModel.PortSecurityEnabled = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)
                
                }
            
                if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {
                
                    
                        childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {
                
                    
                        childModel.ParentUUID = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {
                
                    
                        childModel.ParentType = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {
                
                    
                        childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)
                
                }
            
                if propertyValue, ok := childResourceMap["source_port"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.SourcePort = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["source_ip"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.SourceIP = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["ip_protocol"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.IPProtocol = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["hashing_configured"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.HashingConfigured = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["destination_port"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.DestinationPort = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["destination_ip"]; ok && propertyValue != nil {
                
                    
                        childModel.EcmpHashingIncludeFields.DestinationIP = schema.InterfaceToBool(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {
                
                    
                        childModel.DisplayName = schema.InterfaceToString(propertyValue)
                    
                
                }
            
                if propertyValue, ok := childResourceMap["annotations_key_value_pair"]; ok && propertyValue != nil {
                
                    json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)
                
                }
            
        }
    }
    
    return m, nil
}

// ListVirtualRouter lists VirtualRouter with list spec.
func ListVirtualRouter(ctx context.Context, tx *sql.Tx, request *models.ListVirtualRouterRequest) (response *models.ListVirtualRouterResponse, err error) {
    var rows *sql.Rows
    qb := &common.ListQueryBuilder{}
    qb.Auth = common.GetAuthCTX(ctx) 
    spec := request.Spec
    qb.Spec = spec
    qb.Table = "virtual_router"
    qb.Fields = VirtualRouterFields
    qb.RefFields = VirtualRouterRefFields
    qb.BackRefFields = VirtualRouterBackRefFields
    result := []*models.VirtualRouter{}

    if spec.ParentFQName != nil {
       parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
       if err != nil {
           return nil, errors.Wrap(err, "can't find parents")
       }
       spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
    }

    query := qb.BuildQuery()
    columns := qb.Columns
    values := qb.Values
    log.WithFields(log.Fields{
        "listSpec": spec,
        "query": query,
    }).Debug("select query")
    rows, err = tx.QueryContext(ctx, query, values...)
    if err != nil {
        return nil, errors.Wrap(err,"select query failed")
    }
    defer rows.Close()
    if err := rows.Err(); err != nil {
            return nil, errors.Wrap(err, "row error")
    }

    for rows.Next() {
            valuesMap := map[string]interface{}{}
            values := make([]interface{}, len(columns))
            valuesPointers := make([]interface{}, len(columns))
            for _, index := range columns {
                valuesPointers[index] = &values[index]
            }
            if err := rows.Scan(valuesPointers...); err != nil {
                    return nil, errors.Wrap(err, "scan failed")
            }
            for column, index := range columns {
                val := valuesPointers[index].(*interface{})
                valuesMap[column] = *val
            }
            m, err := scanVirtualRouter(valuesMap)
            if err != nil {
                return nil, errors.Wrap(err, "scan row failed")
            }
            result = append(result, m)
    }
    response = &models.ListVirtualRouterResponse{
       VirtualRouters: result,
    }
    return response, nil
}

// UpdateVirtualRouter updates a resource
func UpdateVirtualRouter(
    ctx context.Context, 
    tx *sql.Tx, 
    request *models.UpdateVirtualRouterRequest,
    ) error {
    //TODO
    return nil
}

// DeleteVirtualRouter deletes a resource
func DeleteVirtualRouter(
    ctx context.Context,
    tx *sql.Tx, 
    request *models.DeleteVirtualRouterRequest) error {
    deleteQuery := deleteVirtualRouterQuery
    selectQuery := "select count(uuid) from virtual_router where uuid = ?"
    var err error
    var count int
    uuid := request.ID
    auth := common.GetAuthCTX(ctx)
    if auth.IsAdmin() {
        row := tx.QueryRowContext(ctx, selectQuery, uuid)
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.ExecContext(ctx, deleteQuery, uuid)
    }else{
        deleteQuery += " and owner = ?"
        selectQuery += " and owner = ?"
        row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID() )
        if err != nil {
            return errors.Wrap(err, "not found")
        }
        row.Scan(&count)
        if count == 0 {
           return errors.New("Not found")
        }
        _, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID() )
    }

    if err != nil {
        return errors.Wrap(err, "delete failed")
    }

    err = common.DeleteMetaData(tx, uuid)
    log.WithFields(log.Fields{
        "uuid": uuid,
    }).Debug("deleted")
    return err
}