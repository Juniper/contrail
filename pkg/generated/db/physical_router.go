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

const insertPhysicalRouterQuery = "insert into `physical_router` (`uuid`,`server_port`,`server_ip`,`resource`,`physical_router_vnc_managed`,`physical_router_vendor_name`,`username`,`password`,`version`,`v3_security_name`,`v3_security_level`,`v3_security_engine_id`,`v3_privacy_protocol`,`v3_privacy_password`,`v3_engine_time`,`v3_engine_id`,`v3_engine_boots`,`v3_context_engine_id`,`v3_context`,`v3_authentication_protocol`,`v3_authentication_password`,`v2_community`,`timeout`,`retries`,`local_port`,`physical_router_snmp`,`physical_router_role`,`physical_router_product_name`,`physical_router_management_ip`,`physical_router_loopback_ip`,`physical_router_lldp`,`service_port`,`physical_router_image_uri`,`physical_router_dataplane_ip`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deletePhysicalRouterQuery = "delete from `physical_router` where uuid = ?"

// PhysicalRouterFields is db columns for PhysicalRouter
var PhysicalRouterFields = []string{
	"uuid",
	"server_port",
	"server_ip",
	"resource",
	"physical_router_vnc_managed",
	"physical_router_vendor_name",
	"username",
	"password",
	"version",
	"v3_security_name",
	"v3_security_level",
	"v3_security_engine_id",
	"v3_privacy_protocol",
	"v3_privacy_password",
	"v3_engine_time",
	"v3_engine_id",
	"v3_engine_boots",
	"v3_context_engine_id",
	"v3_context",
	"v3_authentication_protocol",
	"v3_authentication_password",
	"v2_community",
	"timeout",
	"retries",
	"local_port",
	"physical_router_snmp",
	"physical_router_role",
	"physical_router_product_name",
	"physical_router_management_ip",
	"physical_router_loopback_ip",
	"physical_router_lldp",
	"service_port",
	"physical_router_image_uri",
	"physical_router_dataplane_ip",
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

// PhysicalRouterRefFields is db reference fields for PhysicalRouter
var PhysicalRouterRefFields = map[string][]string{

	"virtual_network": []string{
	// <schema.Schema Value>

	},

	"bgp_router": []string{
	// <schema.Schema Value>

	},

	"virtual_router": []string{
	// <schema.Schema Value>

	},
}

// PhysicalRouterBackRefFields is db back reference fields for PhysicalRouter
var PhysicalRouterBackRefFields = map[string][]string{

	"logical_interface": []string{
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"logical_interface_vlan_tag",
		"logical_interface_type",
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
	},

	"physical_interface": []string{
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
		"ethernet_segment_identifier",
		"display_name",
		"key_value_pair",
	},
}

// PhysicalRouterParentTypes is possible parents for PhysicalRouter
var PhysicalRouterParents = []string{

	"global_system_config",

	"location",
}

const insertPhysicalRouterVirtualNetworkQuery = "insert into `ref_physical_router_virtual_network` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterBGPRouterQuery = "insert into `ref_physical_router_bgp_router` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterVirtualRouterQuery = "insert into `ref_physical_router_virtual_router` (`from`, `to` ) values (?, ?);"

// CreatePhysicalRouter inserts PhysicalRouter to DB
func (db *DB) createPhysicalRouter(
	ctx context.Context,
	request *models.CreatePhysicalRouterRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.PhysicalRouter
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertPhysicalRouterQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		int(model.GetTelemetryInfo().GetServerPort()),
		string(model.GetTelemetryInfo().GetServerIP()),
		common.MustJSON(model.GetTelemetryInfo().GetResource()),
		bool(model.GetPhysicalRouterVNCManaged()),
		string(model.GetPhysicalRouterVendorName()),
		string(model.GetPhysicalRouterUserCredentials().GetUsername()),
		string(model.GetPhysicalRouterUserCredentials().GetPassword()),
		int(model.GetPhysicalRouterSNMPCredentials().GetVersion()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3SecurityName()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3SecurityLevel()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3SecurityEngineID()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3PrivacyProtocol()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3PrivacyPassword()),
		int(model.GetPhysicalRouterSNMPCredentials().GetV3EngineTime()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3EngineID()),
		int(model.GetPhysicalRouterSNMPCredentials().GetV3EngineBoots()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3ContextEngineID()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3Context()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3AuthenticationProtocol()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV3AuthenticationPassword()),
		string(model.GetPhysicalRouterSNMPCredentials().GetV2Community()),
		int(model.GetPhysicalRouterSNMPCredentials().GetTimeout()),
		int(model.GetPhysicalRouterSNMPCredentials().GetRetries()),
		int(model.GetPhysicalRouterSNMPCredentials().GetLocalPort()),
		bool(model.GetPhysicalRouterSNMP()),
		string(model.GetPhysicalRouterRole()),
		string(model.GetPhysicalRouterProductName()),
		string(model.GetPhysicalRouterManagementIP()),
		string(model.GetPhysicalRouterLoopbackIP()),
		bool(model.GetPhysicalRouterLLDP()),
		common.MustJSON(model.GetPhysicalRouterJunosServicePorts().GetServicePort()),
		string(model.GetPhysicalRouterImageURI()),
		string(model.GetPhysicalRouterDataplaneIP()),
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

	stmtVirtualRouterRef, err := tx.Prepare(insertPhysicalRouterVirtualRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualRouterRefs create statement failed")
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {

		_, err = stmtVirtualRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertPhysicalRouterVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	stmtBGPRouterRef, err := tx.Prepare(insertPhysicalRouterBGPRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing BGPRouterRefs create statement failed")
	}
	defer stmtBGPRouterRef.Close()
	for _, ref := range model.BGPRouterRefs {

		_, err = stmtBGPRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "BGPRouterRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "physical_router",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "physical_router", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanPhysicalRouter(values map[string]interface{}) (*models.PhysicalRouter, error) {
	m := models.MakePhysicalRouter()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["server_port"]; ok {

		m.TelemetryInfo.ServerPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["server_ip"]; ok {

		m.TelemetryInfo.ServerIP = schema.InterfaceToString(value)

	}

	if value, ok := values["resource"]; ok {

		json.Unmarshal(value.([]byte), &m.TelemetryInfo.Resource)

	}

	if value, ok := values["physical_router_vnc_managed"]; ok {

		m.PhysicalRouterVNCManaged = schema.InterfaceToBool(value)

	}

	if value, ok := values["physical_router_vendor_name"]; ok {

		m.PhysicalRouterVendorName = schema.InterfaceToString(value)

	}

	if value, ok := values["username"]; ok {

		m.PhysicalRouterUserCredentials.Username = schema.InterfaceToString(value)

	}

	if value, ok := values["password"]; ok {

		m.PhysicalRouterUserCredentials.Password = schema.InterfaceToString(value)

	}

	if value, ok := values["version"]; ok {

		m.PhysicalRouterSNMPCredentials.Version = schema.InterfaceToInt64(value)

	}

	if value, ok := values["v3_security_name"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityName = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_security_level"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityLevel = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_security_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityEngineID = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_privacy_protocol"]; ok {

		m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_privacy_password"]; ok {

		m.PhysicalRouterSNMPCredentials.V3PrivacyPassword = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_engine_time"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineTime = schema.InterfaceToInt64(value)

	}

	if value, ok := values["v3_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineID = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_engine_boots"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineBoots = schema.InterfaceToInt64(value)

	}

	if value, ok := values["v3_context_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3ContextEngineID = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_context"]; ok {

		m.PhysicalRouterSNMPCredentials.V3Context = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_authentication_protocol"]; ok {

		m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = schema.InterfaceToString(value)

	}

	if value, ok := values["v3_authentication_password"]; ok {

		m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = schema.InterfaceToString(value)

	}

	if value, ok := values["v2_community"]; ok {

		m.PhysicalRouterSNMPCredentials.V2Community = schema.InterfaceToString(value)

	}

	if value, ok := values["timeout"]; ok {

		m.PhysicalRouterSNMPCredentials.Timeout = schema.InterfaceToInt64(value)

	}

	if value, ok := values["retries"]; ok {

		m.PhysicalRouterSNMPCredentials.Retries = schema.InterfaceToInt64(value)

	}

	if value, ok := values["local_port"]; ok {

		m.PhysicalRouterSNMPCredentials.LocalPort = schema.InterfaceToInt64(value)

	}

	if value, ok := values["physical_router_snmp"]; ok {

		m.PhysicalRouterSNMP = schema.InterfaceToBool(value)

	}

	if value, ok := values["physical_router_role"]; ok {

		m.PhysicalRouterRole = schema.InterfaceToString(value)

	}

	if value, ok := values["physical_router_product_name"]; ok {

		m.PhysicalRouterProductName = schema.InterfaceToString(value)

	}

	if value, ok := values["physical_router_management_ip"]; ok {

		m.PhysicalRouterManagementIP = schema.InterfaceToString(value)

	}

	if value, ok := values["physical_router_loopback_ip"]; ok {

		m.PhysicalRouterLoopbackIP = schema.InterfaceToString(value)

	}

	if value, ok := values["physical_router_lldp"]; ok {

		m.PhysicalRouterLLDP = schema.InterfaceToBool(value)

	}

	if value, ok := values["service_port"]; ok {

		json.Unmarshal(value.([]byte), &m.PhysicalRouterJunosServicePorts.ServicePort)

	}

	if value, ok := values["physical_router_image_uri"]; ok {

		m.PhysicalRouterImageURI = schema.InterfaceToString(value)

	}

	if value, ok := values["physical_router_dataplane_ip"]; ok {

		m.PhysicalRouterDataplaneIP = schema.InterfaceToString(value)

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

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.PhysicalRouterVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.PhysicalRouterBGPRouterRef{}
			referenceModel.UUID = uuid
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_router"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.PhysicalRouterVirtualRouterRef{}
			referenceModel.UUID = uuid
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	if value, ok := values["backref_logical_interface"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeLogicalInterface()
			m.LogicalInterfaces = append(m.LogicalInterfaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["logical_interface_vlan_tag"]; ok && propertyValue != nil {

				childModel.LogicalInterfaceVlanTag = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["logical_interface_type"]; ok && propertyValue != nil {

				childModel.LogicalInterfaceType = schema.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_physical_interface"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakePhysicalInterface()
			m.PhysicalInterfaces = append(m.PhysicalInterfaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["ethernet_segment_identifier"]; ok && propertyValue != nil {

				childModel.EthernetSegmentIdentifier = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListPhysicalRouter lists PhysicalRouter with list spec.
func (db *DB) listPhysicalRouter(ctx context.Context, request *models.ListPhysicalRouterRequest) (response *models.ListPhysicalRouterResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "physical_router"
	qb.Fields = PhysicalRouterFields
	qb.RefFields = PhysicalRouterRefFields
	qb.BackRefFields = PhysicalRouterBackRefFields
	result := []*models.PhysicalRouter{}

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
		m, err := scanPhysicalRouter(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListPhysicalRouterResponse{
		PhysicalRouters: result,
	}
	return response, nil
}

// UpdatePhysicalRouter updates a resource
func (db *DB) updatePhysicalRouter(
	ctx context.Context,
	request *models.UpdatePhysicalRouterRequest,
) error {
	//TODO
	return nil
}

// DeletePhysicalRouter deletes a resource
func (db *DB) deletePhysicalRouter(
	ctx context.Context,
	request *models.DeletePhysicalRouterRequest) error {
	deleteQuery := deletePhysicalRouterQuery
	selectQuery := "select count(uuid) from physical_router where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
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

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreatePhysicalRouter handle a Create API
func (db *DB) CreatePhysicalRouter(
	ctx context.Context,
	request *models.CreatePhysicalRouterRequest) (*models.CreatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createPhysicalRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreatePhysicalRouterResponse{
		PhysicalRouter: request.PhysicalRouter,
	}, nil
}

//UpdatePhysicalRouter handles a Update request.
func (db *DB) UpdatePhysicalRouter(
	ctx context.Context,
	request *models.UpdatePhysicalRouterRequest) (*models.UpdatePhysicalRouterResponse, error) {
	model := request.PhysicalRouter
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updatePhysicalRouter(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "physical_router",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdatePhysicalRouterResponse{
		PhysicalRouter: model,
	}, nil
}

//DeletePhysicalRouter delete a resource.
func (db *DB) DeletePhysicalRouter(ctx context.Context, request *models.DeletePhysicalRouterRequest) (*models.DeletePhysicalRouterResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deletePhysicalRouter(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeletePhysicalRouterResponse{
		ID: request.ID,
	}, nil
}

//GetPhysicalRouter a Get request.
func (db *DB) GetPhysicalRouter(ctx context.Context, request *models.GetPhysicalRouterRequest) (response *models.GetPhysicalRouterResponse, err error) {
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
	listRequest := &models.ListPhysicalRouterRequest{
		Spec: spec,
	}
	var result *models.ListPhysicalRouterResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listPhysicalRouter(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.PhysicalRouters) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetPhysicalRouterResponse{
		PhysicalRouter: result.PhysicalRouters[0],
	}
	return response, nil
}

//ListPhysicalRouter handles a List service Request.
func (db *DB) ListPhysicalRouter(
	ctx context.Context,
	request *models.ListPhysicalRouterRequest) (response *models.ListPhysicalRouterResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listPhysicalRouter(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
