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

const insertGlobalVrouterConfigQuery = "insert into `global_vrouter_config` (`encapsulation`,`uuid`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`display_name`,`key_value_pair`,`flow_aging_timeout`,`forwarding_mode`,`flow_export_rate`,`vxlan_network_identifier_mode`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`fq_name`,`source_port`,`destination_port`,`destination_ip`,`ip_protocol`,`source_ip`,`hashing_configured`,`linklocal_service_entry`,`enable_security_logging`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateGlobalVrouterConfigQuery = "update `global_vrouter_config` set `encapsulation` = ?,`uuid` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`display_name` = ?,`key_value_pair` = ?,`flow_aging_timeout` = ?,`forwarding_mode` = ?,`flow_export_rate` = ?,`vxlan_network_identifier_mode` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`source_port` = ?,`destination_port` = ?,`destination_ip` = ?,`ip_protocol` = ?,`source_ip` = ?,`hashing_configured` = ?,`linklocal_service_entry` = ?,`enable_security_logging` = ?;"
const deleteGlobalVrouterConfigQuery = "delete from `global_vrouter_config` where uuid = ?"

// GlobalVrouterConfigFields is db columns for GlobalVrouterConfig
var GlobalVrouterConfigFields = []string{
	"encapsulation",
	"uuid",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"owner",
	"display_name",
	"key_value_pair",
	"flow_aging_timeout",
	"forwarding_mode",
	"flow_export_rate",
	"vxlan_network_identifier_mode",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"fq_name",
	"source_port",
	"destination_port",
	"destination_ip",
	"ip_protocol",
	"source_ip",
	"hashing_configured",
	"linklocal_service_entry",
	"enable_security_logging",
}

// GlobalVrouterConfigRefFields is db reference fields for GlobalVrouterConfig
var GlobalVrouterConfigRefFields = map[string][]string{}

// CreateGlobalVrouterConfig inserts GlobalVrouterConfig to DB
func CreateGlobalVrouterConfig(tx *sql.Tx, model *models.GlobalVrouterConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertGlobalVrouterConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertGlobalVrouterConfigQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.EncapsulationPriorities.Encapsulation),
		string(model.UUID),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.FlowAgingTimeoutList.FlowAgingTimeout),
		string(model.ForwardingMode),
		int(model.FlowExportRate),
		string(model.VxlanNetworkIdentifierMode),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.FQName),
		bool(model.EcmpHashingIncludeFields.SourcePort),
		bool(model.EcmpHashingIncludeFields.DestinationPort),
		bool(model.EcmpHashingIncludeFields.DestinationIP),
		bool(model.EcmpHashingIncludeFields.IPProtocol),
		bool(model.EcmpHashingIncludeFields.SourceIP),
		bool(model.EcmpHashingIncludeFields.HashingConfigured),
		utils.MustJSON(model.LinklocalServices.LinklocalServiceEntry),
		bool(model.EnableSecurityLogging))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanGlobalVrouterConfig(values map[string]interface{}) (*models.GlobalVrouterConfig, error) {
	m := models.MakeGlobalVrouterConfig()

	if value, ok := values["encapsulation"]; ok {

		json.Unmarshal(value.([]byte), &m.EncapsulationPriorities.Encapsulation)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

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

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["flow_aging_timeout"]; ok {

		json.Unmarshal(value.([]byte), &m.FlowAgingTimeoutList.FlowAgingTimeout)

	}

	if value, ok := values["forwarding_mode"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ForwardingMode = models.ForwardingModeType(castedValue)

	}

	if value, ok := values["flow_export_rate"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.FlowExportRate = castedValue

	}

	if value, ok := values["vxlan_network_identifier_mode"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VxlanNetworkIdentifierMode = models.VxlanNetworkIdentifierModeType(castedValue)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["source_port"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourcePort = castedValue

	}

	if value, ok := values["destination_port"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationPort = castedValue

	}

	if value, ok := values["destination_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.DestinationIP = castedValue

	}

	if value, ok := values["ip_protocol"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.IPProtocol = castedValue

	}

	if value, ok := values["source_ip"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.SourceIP = castedValue

	}

	if value, ok := values["hashing_configured"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EcmpHashingIncludeFields.HashingConfigured = castedValue

	}

	if value, ok := values["linklocal_service_entry"]; ok {

		json.Unmarshal(value.([]byte), &m.LinklocalServices.LinklocalServiceEntry)

	}

	if value, ok := values["enable_security_logging"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.EnableSecurityLogging = castedValue

	}

	return m, nil
}

// ListGlobalVrouterConfig lists GlobalVrouterConfig with list spec.
func ListGlobalVrouterConfig(tx *sql.Tx, spec *db.ListSpec) ([]*models.GlobalVrouterConfig, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "global_vrouter_config"
	spec.Fields = GlobalVrouterConfigFields
	spec.RefFields = GlobalVrouterConfigRefFields
	result := models.MakeGlobalVrouterConfigSlice()
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
		m, err := scanGlobalVrouterConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowGlobalVrouterConfig shows GlobalVrouterConfig resource
func ShowGlobalVrouterConfig(tx *sql.Tx, uuid string) (*models.GlobalVrouterConfig, error) {
	list, err := ListGlobalVrouterConfig(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateGlobalVrouterConfig updates a resource
func UpdateGlobalVrouterConfig(tx *sql.Tx, uuid string, model *models.GlobalVrouterConfig) error {
	//TODO(nati) support update
	return nil
}

// DeleteGlobalVrouterConfig deletes a resource
func DeleteGlobalVrouterConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteGlobalVrouterConfigQuery)
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
