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

const insertPhysicalRouterBGPRouterQuery = "insert into `ref_physical_router_bgp_router` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterVirtualRouterQuery = "insert into `ref_physical_router_virtual_router` (`from`, `to` ) values (?, ?);"

const insertPhysicalRouterVirtualNetworkQuery = "insert into `ref_physical_router_virtual_network` (`from`, `to` ) values (?, ?);"

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
func UpdatePhysicalRouter(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	// Prepare statement for updating data
	var updatePhysicalRouterQuery = "update `physical_router` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updatePhysicalRouterQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".TelemetryInfo.ServerPort", "."); ok {
		updatePhysicalRouterQuery += "`server_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".TelemetryInfo.ServerIP", "."); ok {
		updatePhysicalRouterQuery += "`server_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".TelemetryInfo.Resource", "."); ok {
		updatePhysicalRouterQuery += "`resource` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterVNCManaged", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_vnc_managed` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterVendorName", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_vendor_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterUserCredentials.Username", "."); ok {
		updatePhysicalRouterQuery += "`username` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterUserCredentials.Password", "."); ok {
		updatePhysicalRouterQuery += "`password` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.Version", "."); ok {
		updatePhysicalRouterQuery += "`version` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3SecurityName", "."); ok {
		updatePhysicalRouterQuery += "`v3_security_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3SecurityLevel", "."); ok {
		updatePhysicalRouterQuery += "`v3_security_level` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3SecurityEngineID", "."); ok {
		updatePhysicalRouterQuery += "`v3_security_engine_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3PrivacyProtocol", "."); ok {
		updatePhysicalRouterQuery += "`v3_privacy_protocol` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3PrivacyPassword", "."); ok {
		updatePhysicalRouterQuery += "`v3_privacy_password` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3EngineTime", "."); ok {
		updatePhysicalRouterQuery += "`v3_engine_time` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3EngineID", "."); ok {
		updatePhysicalRouterQuery += "`v3_engine_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3EngineBoots", "."); ok {
		updatePhysicalRouterQuery += "`v3_engine_boots` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3ContextEngineID", "."); ok {
		updatePhysicalRouterQuery += "`v3_context_engine_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3Context", "."); ok {
		updatePhysicalRouterQuery += "`v3_context` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3AuthenticationProtocol", "."); ok {
		updatePhysicalRouterQuery += "`v3_authentication_protocol` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V3AuthenticationPassword", "."); ok {
		updatePhysicalRouterQuery += "`v3_authentication_password` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.V2Community", "."); ok {
		updatePhysicalRouterQuery += "`v2_community` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.Timeout", "."); ok {
		updatePhysicalRouterQuery += "`timeout` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.Retries", "."); ok {
		updatePhysicalRouterQuery += "`retries` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMPCredentials.LocalPort", "."); ok {
		updatePhysicalRouterQuery += "`local_port` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterSNMP", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_snmp` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterRole", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_role` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterProductName", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_product_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterManagementIP", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_management_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterLoopbackIP", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_loopback_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterLLDP", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_lldp` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterJunosServicePorts.ServicePort", "."); ok {
		updatePhysicalRouterQuery += "`service_port` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterImageURI", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_image_uri` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PhysicalRouterDataplaneIP", "."); ok {
		updatePhysicalRouterQuery += "`physical_router_dataplane_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updatePhysicalRouterQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updatePhysicalRouterQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updatePhysicalRouterQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updatePhysicalRouterQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updatePhysicalRouterQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updatePhysicalRouterQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updatePhysicalRouterQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updatePhysicalRouterQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updatePhysicalRouterQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updatePhysicalRouterQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updatePhysicalRouterQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updatePhysicalRouterQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updatePhysicalRouterQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updatePhysicalRouterQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updatePhysicalRouterQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updatePhysicalRouterQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updatePhysicalRouterQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updatePhysicalRouterQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updatePhysicalRouterQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updatePhysicalRouterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updatePhysicalRouterQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updatePhysicalRouterQuery += ","
	}

	updatePhysicalRouterQuery =
		updatePhysicalRouterQuery[:len(updatePhysicalRouterQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updatePhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updatePhysicalRouterQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	if value, ok := common.GetValueByPath(model, "VirtualNetworkRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_physical_router_virtual_network` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_physical_router_virtual_network` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref VirtualNetworkRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_physical_router_virtual_network` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing VirtualNetworkRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "VirtualNetworkRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "BGPRouterRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_physical_router_bgp_router` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_physical_router_bgp_router` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref BGPRouterRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_physical_router_bgp_router` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing BGPRouterRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "BGPRouterRefs update failed")
			}
		}
	}

	if value, ok := common.GetValueByPath(model, "VirtualRouterRefs", "."); ok {
		for _, ref := range value.([]interface{}) {
			refQuery := ""
			refValues := make([]interface{}, 0)
			refKeys := make([]string, 0)
			refUUID, ok := common.GetValueByPath(ref.(map[string]interface{}), "UUID", ".")
			if !ok {
				return errors.Wrap(err, "UUID is missing for referred resource. Failed to update Refs")
			}

			refValues = append(refValues, uuid)
			refValues = append(refValues, refUUID)
			operation, ok := common.GetValueByPath(ref.(map[string]interface{}), common.OPERATION, ".")
			switch operation {
			case common.ADD:
				refQuery = "insert into `ref_physical_router_virtual_router` ("
				values := "values("
				for _, value := range refKeys {
					refQuery += "`" + value + "`, "
					values += "?,"
				}
				refQuery += "`from`, `to`) "
				values += "?,?);"
				refQuery += values
			case common.UPDATE:
				refQuery = "update `ref_physical_router_virtual_router` set "
				if len(refKeys) == 0 {
					return errors.Wrap(err, "Failed to update Refs. No Attribute to update for ref VirtualRouterRefs")
				}
				for _, value := range refKeys {
					refQuery += "`" + value + "` = ?,"
				}
				refQuery = refQuery[:len(refQuery)-1] + " where `from` = ? AND `to` = ?;"
			case common.DELETE:
				refQuery = "delete from `ref_physical_router_virtual_router` where `from` = ? AND `to`= ?;"
				refValues = refValues[len(refValues)-2:]
			default:
				return errors.Wrap(err, "Failed to update Refs. Ref operations can be only ADD, UPDATE, DELETE")
			}
			stmt, err := tx.Prepare(refQuery)
			if err != nil {
				return errors.Wrap(err, "preparing VirtualRouterRefs update statement failed")
			}
			_, err = stmt.Exec(refValues...)
			if err != nil {
				return errors.Wrap(err, "VirtualRouterRefs update failed")
			}
		}
	}

	share, ok := common.GetValueByPath(model, ".Perms2.Share", ".")
	if ok {
		err = common.UpdateSharing(tx, "physical_router", string(uuid), share.([]interface{}))
		if err != nil {
			return err
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
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
