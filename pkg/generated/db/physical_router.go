package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertPhysicalRouterQuery = "insert into `physical_router` (`physical_router_role`,`username`,`password`,`global_access`,`share`,`owner`,`owner_access`,`display_name`,`fq_name`,`key_value_pair`,`physical_router_management_ip`,`physical_router_product_name`,`physical_router_lldp`,`physical_router_snmp`,`service_port`,`physical_router_loopback_ip`,`physical_router_dataplane_ip`,`uuid`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`v3_security_level`,`v3_engine_id`,`v3_security_name`,`v3_context`,`v3_privacy_password`,`v3_context_engine_id`,`v3_engine_boots`,`timeout`,`v3_engine_time`,`v2_community`,`version`,`v3_privacy_protocol`,`v3_authentication_protocol`,`v3_authentication_password`,`retries`,`v3_security_engine_id`,`local_port`,`physical_router_vendor_name`,`physical_router_vnc_managed`,`physical_router_image_uri`,`resource`,`server_port`,`server_ip`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalRouterQuery = "update `physical_router` set `physical_router_role` = ?,`username` = ?,`password` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`display_name` = ?,`fq_name` = ?,`key_value_pair` = ?,`physical_router_management_ip` = ?,`physical_router_product_name` = ?,`physical_router_lldp` = ?,`physical_router_snmp` = ?,`service_port` = ?,`physical_router_loopback_ip` = ?,`physical_router_dataplane_ip` = ?,`uuid` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`v3_security_level` = ?,`v3_engine_id` = ?,`v3_security_name` = ?,`v3_context` = ?,`v3_privacy_password` = ?,`v3_context_engine_id` = ?,`v3_engine_boots` = ?,`timeout` = ?,`v3_engine_time` = ?,`v2_community` = ?,`version` = ?,`v3_privacy_protocol` = ?,`v3_authentication_protocol` = ?,`v3_authentication_password` = ?,`retries` = ?,`v3_security_engine_id` = ?,`local_port` = ?,`physical_router_vendor_name` = ?,`physical_router_vnc_managed` = ?,`physical_router_image_uri` = ?,`resource` = ?,`server_port` = ?,`server_ip` = ?;"
const deletePhysicalRouterQuery = "delete from `physical_router` where uuid = ?"

// PhysicalRouterFields is db columns for PhysicalRouter
var PhysicalRouterFields = []string{
	"physical_router_role",
	"username",
	"password",
	"global_access",
	"share",
	"owner",
	"owner_access",
	"display_name",
	"fq_name",
	"key_value_pair",
	"physical_router_management_ip",
	"physical_router_product_name",
	"physical_router_lldp",
	"physical_router_snmp",
	"service_port",
	"physical_router_loopback_ip",
	"physical_router_dataplane_ip",
	"uuid",
	"last_modified",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"v3_security_level",
	"v3_engine_id",
	"v3_security_name",
	"v3_context",
	"v3_privacy_password",
	"v3_context_engine_id",
	"v3_engine_boots",
	"timeout",
	"v3_engine_time",
	"v2_community",
	"version",
	"v3_privacy_protocol",
	"v3_authentication_protocol",
	"v3_authentication_password",
	"retries",
	"v3_security_engine_id",
	"local_port",
	"physical_router_vendor_name",
	"physical_router_vnc_managed",
	"physical_router_image_uri",
	"resource",
	"server_port",
	"server_ip",
}

// PhysicalRouterRefFields is db reference fields for PhysicalRouter
var PhysicalRouterRefFields = map[string][]string{

	"virtual_network": {
	// <utils.Schema Value>

	},

	"bgp_router": {
	// <utils.Schema Value>

	},

	"virtual_router": {
	// <utils.Schema Value>

	},
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
	_, err = stmt.Exec(string(model.PhysicalRouterRole),
		string(model.PhysicalRouterUserCredentials.Username),
		string(model.PhysicalRouterUserCredentials.Password),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.DisplayName),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.PhysicalRouterManagementIP),
		string(model.PhysicalRouterProductName),
		bool(model.PhysicalRouterLLDP),
		bool(model.PhysicalRouterSNMP),
		utils.MustJSON(model.PhysicalRouterJunosServicePorts.ServicePort),
		string(model.PhysicalRouterLoopbackIP),
		string(model.PhysicalRouterDataplaneIP),
		string(model.UUID),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityLevel),
		string(model.PhysicalRouterSNMPCredentials.V3EngineID),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityName),
		string(model.PhysicalRouterSNMPCredentials.V3Context),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyPassword),
		string(model.PhysicalRouterSNMPCredentials.V3ContextEngineID),
		int(model.PhysicalRouterSNMPCredentials.V3EngineBoots),
		int(model.PhysicalRouterSNMPCredentials.Timeout),
		int(model.PhysicalRouterSNMPCredentials.V3EngineTime),
		string(model.PhysicalRouterSNMPCredentials.V2Community),
		int(model.PhysicalRouterSNMPCredentials.Version),
		string(model.PhysicalRouterSNMPCredentials.V3PrivacyProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol),
		string(model.PhysicalRouterSNMPCredentials.V3AuthenticationPassword),
		int(model.PhysicalRouterSNMPCredentials.Retries),
		string(model.PhysicalRouterSNMPCredentials.V3SecurityEngineID),
		int(model.PhysicalRouterSNMPCredentials.LocalPort),
		string(model.PhysicalRouterVendorName),
		bool(model.PhysicalRouterVNCManaged),
		string(model.PhysicalRouterImageURI),
		utils.MustJSON(model.TelemetryInfo.Resource),
		int(model.TelemetryInfo.ServerPort),
		string(model.TelemetryInfo.ServerIP))
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanPhysicalRouter(values map[string]interface{}) (*models.PhysicalRouter, error) {
	m := models.MakePhysicalRouter()

	if value, ok := values["physical_router_role"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterRole = models.PhysicalRouterRole(castedValue)

	}

	if value, ok := values["username"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterUserCredentials.Username = castedValue

	}

	if value, ok := values["password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterUserCredentials.Password = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["physical_router_management_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterManagementIP = castedValue

	}

	if value, ok := values["physical_router_product_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterProductName = castedValue

	}

	if value, ok := values["physical_router_lldp"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.PhysicalRouterLLDP = castedValue

	}

	if value, ok := values["physical_router_snmp"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.PhysicalRouterSNMP = castedValue

	}

	if value, ok := values["service_port"]; ok {

		json.Unmarshal(value.([]byte), &m.PhysicalRouterJunosServicePorts.ServicePort)

	}

	if value, ok := values["physical_router_loopback_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterLoopbackIP = castedValue

	}

	if value, ok := values["physical_router_dataplane_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterDataplaneIP = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["v3_security_level"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityLevel = castedValue

	}

	if value, ok := values["v3_engine_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3EngineID = castedValue

	}

	if value, ok := values["v3_security_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityName = castedValue

	}

	if value, ok := values["v3_context"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3Context = castedValue

	}

	if value, ok := values["v3_privacy_password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3PrivacyPassword = castedValue

	}

	if value, ok := values["v3_context_engine_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3ContextEngineID = castedValue

	}

	if value, ok := values["v3_engine_boots"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.V3EngineBoots = castedValue

	}

	if value, ok := values["timeout"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Timeout = castedValue

	}

	if value, ok := values["v3_engine_time"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.V3EngineTime = castedValue

	}

	if value, ok := values["v2_community"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V2Community = castedValue

	}

	if value, ok := values["version"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Version = castedValue

	}

	if value, ok := values["v3_privacy_protocol"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = castedValue

	}

	if value, ok := values["v3_authentication_protocol"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = castedValue

	}

	if value, ok := values["v3_authentication_password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = castedValue

	}

	if value, ok := values["retries"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.Retries = castedValue

	}

	if value, ok := values["v3_security_engine_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterSNMPCredentials.V3SecurityEngineID = castedValue

	}

	if value, ok := values["local_port"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.PhysicalRouterSNMPCredentials.LocalPort = castedValue

	}

	if value, ok := values["physical_router_vendor_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterVendorName = castedValue

	}

	if value, ok := values["physical_router_vnc_managed"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.PhysicalRouterVNCManaged = castedValue

	}

	if value, ok := values["physical_router_image_uri"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PhysicalRouterImageURI = castedValue

	}

	if value, ok := values["resource"]; ok {

		json.Unmarshal(value.([]byte), &m.TelemetryInfo.Resource)

	}

	if value, ok := values["server_port"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.TelemetryInfo.ServerPort = castedValue

	}

	if value, ok := values["server_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.TelemetryInfo.ServerIP = castedValue

	}

	if value, ok := values["ref_virtual_network"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.PhysicalRouterVirtualNetworkRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_bgp_router"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.PhysicalRouterBGPRouterRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.BGPRouterRefs = append(m.BGPRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_router"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.PhysicalRouterVirtualRouterRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	return m, nil
}

// ListPhysicalRouter lists PhysicalRouter with list spec.
func ListPhysicalRouter(tx *sql.Tx, spec *db.ListSpec) ([]*models.PhysicalRouter, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "physical_router"
	spec.Fields = PhysicalRouterFields
	spec.RefFields = PhysicalRouterRefFields
	result := models.MakePhysicalRouterSlice()
	query, columns, values := db.BuildListQuery(spec)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanPhysicalRouter(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowPhysicalRouter shows PhysicalRouter resource
func ShowPhysicalRouter(tx *sql.Tx, uuid string) (*models.PhysicalRouter, error) {
	list, err := ListPhysicalRouter(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdatePhysicalRouter updates a resource
func UpdatePhysicalRouter(tx *sql.Tx, uuid string, model *models.PhysicalRouter) error {
	//TODO(nati) support update
	return nil
}

// DeletePhysicalRouter deletes a resource
func DeletePhysicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
