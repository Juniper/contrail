package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertPhysicalRouterQuery = "insert into `physical_router` (`uuid`,`server_port`,`server_ip`,`resource`,`physical_router_vnc_managed`,`physical_router_vendor_name`,`username`,`password`,`version`,`v3_security_name`,`v3_security_level`,`v3_security_engine_id`,`v3_privacy_protocol`,`v3_privacy_password`,`v3_engine_time`,`v3_engine_id`,`v3_engine_boots`,`v3_context_engine_id`,`v3_context`,`v3_authentication_protocol`,`v3_authentication_password`,`v2_community`,`timeout`,`retries`,`local_port`,`physical_router_snmp`,`physical_router_role`,`physical_router_product_name`,`physical_router_management_ip`,`physical_router_loopback_ip`,`physical_router_lldp`,`service_port`,`physical_router_image_uri`,`physical_router_dataplane_ip`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `uuid` = ?,`server_port` = ?,`server_ip` = ?,`resource` = ?,`physical_router_vnc_managed` = ?,`physical_router_vendor_name` = ?,`username` = ?,`password` = ?,`version` = ?,`v3_security_name` = ?,`v3_security_level` = ?,`v3_security_engine_id` = ?,`v3_privacy_protocol` = ?,`v3_privacy_password` = ?,`v3_engine_time` = ?,`v3_engine_id` = ?,`v3_engine_boots` = ?,`v3_context_engine_id` = ?,`v3_context` = ?,`v3_authentication_protocol` = ?,`v3_authentication_password` = ?,`v2_community` = ?,`timeout` = ?,`retries` = ?,`local_port` = ?,`physical_router_snmp` = ?,`physical_router_role` = ?,`physical_router_product_name` = ?,`physical_router_management_ip` = ?,`physical_router_loopback_ip` = ?,`physical_router_lldp` = ?,`service_port` = ?,`physical_router_image_uri` = ?,`physical_router_dataplane_ip` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
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

	"virtual_network": {
	// <common.Schema Value>

	},

	"bgp_router": {
	// <common.Schema Value>

	},

	"virtual_router": {
	// <common.Schema Value>

	},
}

// PhysicalRouterBackRefFields is db back reference fields for PhysicalRouter
var PhysicalRouterBackRefFields = map[string][]string{

	"logical_interface": {
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

	"physical_interface": {
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
func CreatePhysicalRouter(tx *sql.Tx, model *models.PhysicalRouter) error {
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
	_, err = stmt.Exec(string(model.UUID),
		int(model.TelemetryInfo.ServerPort),
		string(model.TelemetryInfo.ServerIP),
		common.MustJSON(model.TelemetryInfo.Resource),
		bool(model.PhysicalRouterVNCManaged),
		string(model.PhysicalRouterVendorName),
		string(model.PhysicalRouterUserCredentials.Username),
		string(model.PhysicalRouterUserCredentials.Password),
		int(model.PhysicalRouterSNMPCredentials.Version),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityName),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityLevel),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityEngineID),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyPassword),
		int(model.PhysicalRouterSNMPCredentials.V3EngineTime),
		string(model.PhysicalRouterSNMPCredentials.V3EngineID),
		int(model.PhysicalRouterSNMPCredentials.V3EngineBoots),
		string(model.PhysicalRouterSNMPCredentials.V3ContextEngineID),
		string(model.PhysicalRouterSNMPCredentials.V3Context),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword),
		string(model.PhysicalRouterSNMPCredentials.V2Community),
		int(model.PhysicalRouterSNMPCredentials.Timeout),
		int(model.PhysicalRouterSNMPCredentials.Retries),
		int(model.PhysicalRouterSNMPCredentials.LocalPort),
		bool(model.PhysicalRouterSNMP),
		string(model.PhysicalRouterRole),
		string(model.PhysicalRouterProductName),
		string(model.PhysicalRouterManagementIP),
		string(model.PhysicalRouterLoopbackIP),
		bool(model.PhysicalRouterLLDP),
		common.MustJSON(model.PhysicalRouterJunosServicePorts.ServicePort),
		string(model.PhysicalRouterImageURI),
		string(model.PhysicalRouterDataplaneIP),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertPhysicalRouterVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
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

		_, err = stmtBGPRouterRef.Exec(model.UUID, ref.UUID)
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

		_, err = stmtVirtualRouterRef.Exec(model.UUID, ref.UUID)
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
	err = common.CreateSharing(tx, "physical_router", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["server_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.TelemetryInfo.ServerPort = castedValue

	}

	if value, ok := values["server_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.TelemetryInfo.ServerIP = castedValue

	}

	if value, ok := values["resource"]; ok {

		json.Unmarshal(value.([]byte), &m.TelemetryInfo.Resource)

	}

	if value, ok := values["physical_router_vnc_managed"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PhysicalRouterVNCManaged = castedValue

	}

	if value, ok := values["physical_router_vendor_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterVendorName = castedValue

	}

	if value, ok := values["username"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterUserCredentials.Username = castedValue

	}

	if value, ok := values["password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterUserCredentials.Password = castedValue

	}

	if value, ok := values["version"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Version = castedValue

	}

	if value, ok := values["v3_security_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityName = castedValue

	}

	if value, ok := values["v3_security_level"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityLevel = castedValue

	}

	if value, ok := values["v3_security_engine_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityEngineID = castedValue

	}

	if value, ok := values["v3_privacy_protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = castedValue

	}

	if value, ok := values["v3_privacy_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3PrivacyPassword = castedValue

	}

	if value, ok := values["v3_engine_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.V3EngineTime = castedValue

	}

	if value, ok := values["v3_engine_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3EngineID = castedValue

	}

	if value, ok := values["v3_engine_boots"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.V3EngineBoots = castedValue

	}

	if value, ok := values["v3_context_engine_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3ContextEngineID = castedValue

	}

	if value, ok := values["v3_context"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3Context = castedValue

	}

	if value, ok := values["v3_authentication_protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = castedValue

	}

	if value, ok := values["v3_authentication_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = castedValue

	}

	if value, ok := values["v2_community"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V2Community = castedValue

	}

	if value, ok := values["timeout"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Timeout = castedValue

	}

	if value, ok := values["retries"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Retries = castedValue

	}

	if value, ok := values["local_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.LocalPort = castedValue

	}

	if value, ok := values["physical_router_snmp"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PhysicalRouterSNMP = castedValue

	}

	if value, ok := values["physical_router_role"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterRole = models.PhysicalRouterRole(castedValue)

	}

	if value, ok := values["physical_router_product_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterProductName = castedValue

	}

	if value, ok := values["physical_router_management_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterManagementIP = castedValue

	}

	if value, ok := values["physical_router_loopback_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterLoopbackIP = castedValue

	}

	if value, ok := values["physical_router_lldp"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PhysicalRouterLLDP = castedValue

	}

	if value, ok := values["service_port"]; ok {

		json.Unmarshal(value.([]byte), &m.PhysicalRouterJunosServicePorts.ServicePort)

	}

	if value, ok := values["physical_router_image_uri"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterImageURI = castedValue

	}

	if value, ok := values["physical_router_dataplane_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PhysicalRouterDataplaneIP = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["logical_interface_vlan_tag"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.LogicalInterfaceVlanTag = castedValue

			}

			if propertyValue, ok := childResourceMap["logical_interface_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.LogicalInterfaceType = models.LogicalInterfaceType(castedValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["ethernet_segment_identifier"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.EthernetSegmentIdentifier = castedValue

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListPhysicalRouter lists PhysicalRouter with list spec.
func ListPhysicalRouter(tx *sql.Tx, spec *common.ListSpec) ([]*models.PhysicalRouter, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "physical_router"
	spec.Fields = PhysicalRouterFields
	spec.RefFields = PhysicalRouterRefFields
	spec.BackRefFields = PhysicalRouterBackRefFields
	result := models.MakePhysicalRouterSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
	return result, nil
}

// UpdatePhysicalRouter updates a resource
func UpdatePhysicalRouter(tx *sql.Tx, uuid string, model *models.PhysicalRouter) error {
	//TODO(nati) support update
	return nil
}

// DeletePhysicalRouter deletes a resource
func DeletePhysicalRouter(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deletePhysicalRouterQuery
	selectQuery := "select count(uuid) from physical_router where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
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
