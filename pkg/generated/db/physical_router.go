package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
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
func CreatePhysicalRouter(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreatePhysicalRouterRequest) error {
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

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["server_port"]; ok {

		m.TelemetryInfo.ServerPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["server_ip"]; ok {

		m.TelemetryInfo.ServerIP = common.InterfaceToString(value)

	}

	if value, ok := values["resource"]; ok {

		json.Unmarshal(value.([]byte), &m.TelemetryInfo.Resource)

	}

	if value, ok := values["physical_router_vnc_managed"]; ok {

		m.PhysicalRouterVNCManaged = common.InterfaceToBool(value)

	}

	if value, ok := values["physical_router_vendor_name"]; ok {

		m.PhysicalRouterVendorName = common.InterfaceToString(value)

	}

	if value, ok := values["username"]; ok {

		m.PhysicalRouterUserCredentials.Username = common.InterfaceToString(value)

	}

	if value, ok := values["password"]; ok {

		m.PhysicalRouterUserCredentials.Password = common.InterfaceToString(value)

	}

	if value, ok := values["version"]; ok {

		m.PhysicalRouterSNMPCredentials.Version = common.InterfaceToInt64(value)

	}

	if value, ok := values["v3_security_name"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityName = common.InterfaceToString(value)

	}

	if value, ok := values["v3_security_level"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityLevel = common.InterfaceToString(value)

	}

	if value, ok := values["v3_security_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3SecurityEngineID = common.InterfaceToString(value)

	}

	if value, ok := values["v3_privacy_protocol"]; ok {

		m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = common.InterfaceToString(value)

	}

	if value, ok := values["v3_privacy_password"]; ok {

		m.PhysicalRouterSNMPCredentials.V3PrivacyPassword = common.InterfaceToString(value)

	}

	if value, ok := values["v3_engine_time"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineTime = common.InterfaceToInt64(value)

	}

	if value, ok := values["v3_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineID = common.InterfaceToString(value)

	}

	if value, ok := values["v3_engine_boots"]; ok {

		m.PhysicalRouterSNMPCredentials.V3EngineBoots = common.InterfaceToInt64(value)

	}

	if value, ok := values["v3_context_engine_id"]; ok {

		m.PhysicalRouterSNMPCredentials.V3ContextEngineID = common.InterfaceToString(value)

	}

	if value, ok := values["v3_context"]; ok {

		m.PhysicalRouterSNMPCredentials.V3Context = common.InterfaceToString(value)

	}

	if value, ok := values["v3_authentication_protocol"]; ok {

		m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = common.InterfaceToString(value)

	}

	if value, ok := values["v3_authentication_password"]; ok {

		m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = common.InterfaceToString(value)

	}

	if value, ok := values["v2_community"]; ok {

		m.PhysicalRouterSNMPCredentials.V2Community = common.InterfaceToString(value)

	}

	if value, ok := values["timeout"]; ok {

		m.PhysicalRouterSNMPCredentials.Timeout = common.InterfaceToInt64(value)

	}

	if value, ok := values["retries"]; ok {

		m.PhysicalRouterSNMPCredentials.Retries = common.InterfaceToInt64(value)

	}

	if value, ok := values["local_port"]; ok {

		m.PhysicalRouterSNMPCredentials.LocalPort = common.InterfaceToInt64(value)

	}

	if value, ok := values["physical_router_snmp"]; ok {

		m.PhysicalRouterSNMP = common.InterfaceToBool(value)

	}

	if value, ok := values["physical_router_role"]; ok {

		m.PhysicalRouterRole = common.InterfaceToString(value)

	}

	if value, ok := values["physical_router_product_name"]; ok {

		m.PhysicalRouterProductName = common.InterfaceToString(value)

	}

	if value, ok := values["physical_router_management_ip"]; ok {

		m.PhysicalRouterManagementIP = common.InterfaceToString(value)

	}

	if value, ok := values["physical_router_loopback_ip"]; ok {

		m.PhysicalRouterLoopbackIP = common.InterfaceToString(value)

	}

	if value, ok := values["physical_router_lldp"]; ok {

		m.PhysicalRouterLLDP = common.InterfaceToBool(value)

	}

	if value, ok := values["service_port"]; ok {

		json.Unmarshal(value.([]byte), &m.PhysicalRouterJunosServicePorts.ServicePort)

	}

	if value, ok := values["physical_router_image_uri"]; ok {

		m.PhysicalRouterImageURI = common.InterfaceToString(value)

	}

	if value, ok := values["physical_router_dataplane_ip"]; ok {

		m.PhysicalRouterDataplaneIP = common.InterfaceToString(value)

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_bgp_router"]; ok {
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
			referenceModel := &models.PhysicalRouterBGPRouterRef{}
			referenceModel.UUID = uuid
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_router"]; ok {
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
			referenceModel := &models.PhysicalRouterVirtualRouterRef{}
			referenceModel.UUID = uuid
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
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
			referenceModel := &models.PhysicalRouterVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["backref_logical_interface"]; ok {
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
			childModel := models.MakeLogicalInterface()
			m.LogicalInterfaces = append(m.LogicalInterfaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["logical_interface_vlan_tag"]; ok && propertyValue != nil {

				childModel.LogicalInterfaceVlanTag = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["logical_interface_type"]; ok && propertyValue != nil {

				childModel.LogicalInterfaceType = common.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_physical_interface"]; ok {
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
			childModel := models.MakePhysicalInterface()
			m.PhysicalInterfaces = append(m.PhysicalInterfaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

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

			if propertyValue, ok := childResourceMap["ethernet_segment_identifier"]; ok && propertyValue != nil {

				childModel.EthernetSegmentIdentifier = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListPhysicalRouter lists PhysicalRouter with list spec.
func ListPhysicalRouter(ctx context.Context, tx *sql.Tx, request *models.ListPhysicalRouterRequest) (response *models.ListPhysicalRouterResponse, err error) {
	var rows *sql.Rows
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
func UpdatePhysicalRouter(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdatePhysicalRouterRequest,
) error {
	//TODO
	return nil
}

// DeletePhysicalRouter deletes a resource
func DeletePhysicalRouter(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeletePhysicalRouterRequest) error {
	deleteQuery := deletePhysicalRouterQuery
	selectQuery := "select count(uuid) from physical_router where uuid = ?"
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
