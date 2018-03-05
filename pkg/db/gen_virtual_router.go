// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

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
	"configuration_version",
	"key_value_pair",
}

// VirtualRouterRefFields is db reference fields for VirtualRouter
var VirtualRouterRefFields = map[string][]string{

	"virtual_machine": []string{
	// <schema.Schema Value>

	},

	"network_ipam": []string{
		// <schema.Schema Value>
		"subnet",
		"allocation_pools",
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
		"configuration_version",
		"annotations_key_value_pair",
	},
}

// VirtualRouterParentTypes is possible parents for VirtualRouter
var VirtualRouterParents = []string{

	"global_system_config",
}

// CreateVirtualRouter inserts VirtualRouter to DB
// nolint
func (db *DB) createVirtualRouter(
	ctx context.Context,
	request *models.CreateVirtualRouterRequest) error {
	qb := db.queryBuilders["virtual_router"]
	tx := GetTransaction(ctx)
	model := request.VirtualRouter
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetVirtualRouterType()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.NetworkIpamRefs {

		if ref.Attr == nil {
			ref.Attr = &models.VirtualRouterNetworkIpamType{}
		}

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("network_ipam"), model.UUID, ref.UUID, common.MustJSON(ref.Attr.GetSubnet()),
			common.MustJSON(ref.Attr.GetAllocationPools()))
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	for _, ref := range model.VirtualMachineRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_machine"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "virtual_router",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "virtual_router", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanVirtualRouter(values map[string]interface{}) (*models.VirtualRouter, error) {
	m := models.MakeVirtualRouter()

	if value, ok := values["virtual_router_type"]; ok {

		m.VirtualRouterType = common.InterfaceToString(value)

	}

	if value, ok := values["virtual_router_ip_address"]; ok {

		m.VirtualRouterIPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["virtual_router_dpdk_enabled"]; ok {

		m.VirtualRouterDPDKEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_network_ipam"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
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
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeVirtualMachineInterface()
			m.VirtualMachineInterfaces = append(m.VirtualMachineInterfaces, childModel)

			if propertyValue, ok := childResourceMap["vrf_assign_rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VRFAssignTable.VRFAssignRule)

			}

			if propertyValue, ok := childResourceMap["vlan_tag_based_bridge_domain"]; ok && propertyValue != nil {

				childModel.VlanTagBasedBridgeDomain = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["sub_interface_vlan_tag"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.SubInterfaceVlanTag = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_interface_type"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.ServiceInterfaceType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["local_preference"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.LocalPreference = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["traffic_direction"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.TrafficDirection = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["udp_port"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.UDPPort = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["vtep_dst_mac_address"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTMacAddress = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["vtep_dst_ip_address"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.VtepDSTIPAddress = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["vni"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.StaticNHHeader.Vni = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["routing_instance"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.RoutingInstance = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["nic_assisted_mirroring_vlan"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroringVlan = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["nic_assisted_mirroring"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NicAssistedMirroring = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["nh_mode"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.NHMode = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["juniper_header"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.JuniperHeader = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["encapsulation"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.Encapsulation = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["analyzer_name"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["analyzer_mac_address"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerMacAddress = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["analyzer_ip_address"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceProperties.InterfaceMirror.MirrorTo.AnalyzerIPAddress = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_address"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceMacAddresses.MacAddress)

			}

			if propertyValue, ok := childResourceMap["route"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceHostRoutes.Route)

			}

			if propertyValue, ok := childResourceMap["fat_flow_protocol"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceFatFlowProtocols.FatFlowProtocol)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface_disable_policy"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceDisablePolicy = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["dhcp_option"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceDHCPOptionList.DHCPOption)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface_device_owner"]; ok && propertyValue != nil {

				childModel.VirtualMachineInterfaceDeviceOwner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceBindings.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["allowed_address_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.VirtualMachineInterfaceAllowedAddressPairs.AllowedAddressPair)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["port_security_enabled"]; ok && propertyValue != nil {

				childModel.PortSecurityEnabled = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["source_port"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.SourcePort = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["source_ip"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.SourceIP = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_protocol"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.IPProtocol = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["hashing_configured"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.HashingConfigured = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["destination_port"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.DestinationPort = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["destination_ip"]; ok && propertyValue != nil {

				childModel.EcmpHashingIncludeFields.DestinationIP = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["annotations_key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListVirtualRouter lists VirtualRouter with list spec.
func (db *DB) listVirtualRouter(ctx context.Context, request *models.ListVirtualRouterRequest) (response *models.ListVirtualRouterResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["virtual_router"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.VirtualRouter{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
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
func (db *DB) updateVirtualRouter(
	ctx context.Context,
	request *models.UpdateVirtualRouterRequest,
) error {
	//TODO
	return nil
}

// DeleteVirtualRouter deletes a resource
func (db *DB) deleteVirtualRouter(
	ctx context.Context,
	request *models.DeleteVirtualRouterRequest) error {
	qb := db.queryBuilders["virtual_router"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
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
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateVirtualRouter handle a Create API
// nolint
func (db *DB) CreateVirtualRouter(
	ctx context.Context,
	request *models.CreateVirtualRouterRequest) (*models.CreateVirtualRouterResponse, error) {
	model := request.VirtualRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createVirtualRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateVirtualRouterResponse{
		VirtualRouter: request.VirtualRouter,
	}, nil
}

//UpdateVirtualRouter handles a Update request.
func (db *DB) UpdateVirtualRouter(
	ctx context.Context,
	request *models.UpdateVirtualRouterRequest) (*models.UpdateVirtualRouterResponse, error) {
	model := request.VirtualRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateVirtualRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "virtual_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateVirtualRouterResponse{
		VirtualRouter: model,
	}, nil
}

//DeleteVirtualRouter delete a resource.
func (db *DB) DeleteVirtualRouter(ctx context.Context, request *models.DeleteVirtualRouterRequest) (*models.DeleteVirtualRouterResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteVirtualRouter(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteVirtualRouterResponse{
		ID: request.ID,
	}, nil
}

//GetVirtualRouter a Get request.
func (db *DB) GetVirtualRouter(ctx context.Context, request *models.GetVirtualRouterRequest) (response *models.GetVirtualRouterResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListVirtualRouterRequest{
		Spec: spec,
	}
	var result *models.ListVirtualRouterResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listVirtualRouter(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.VirtualRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetVirtualRouterResponse{
		VirtualRouter: result.VirtualRouters[0],
	}
	return response, nil
}

//ListVirtualRouter handles a List service Request.
// nolint
func (db *DB) ListVirtualRouter(
	ctx context.Context,
	request *models.ListVirtualRouterRequest) (response *models.ListVirtualRouterResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listVirtualRouter(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
