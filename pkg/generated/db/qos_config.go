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

const insertQosConfigQuery = "insert into `qos_config` (`fq_name`,`display_name`,`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`mpls_exp_entries`,`vlan_priority_entries`,`default_forwarding_class_id`,`dscp_entries`,`uuid`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`qos_config_type`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateQosConfigQuery = "update `qos_config` set `fq_name` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`mpls_exp_entries` = ?,`vlan_priority_entries` = ?,`default_forwarding_class_id` = ?,`dscp_entries` = ?,`uuid` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`qos_config_type` = ?;"
const deleteQosConfigQuery = "delete from `qos_config` where uuid = ?"

// QosConfigFields is db columns for QosConfig
var QosConfigFields = []string{
	"fq_name",
	"display_name",
	"key_value_pair",
	"global_access",
	"share",
	"owner",
	"owner_access",
	"mpls_exp_entries",
	"vlan_priority_entries",
	"default_forwarding_class_id",
	"dscp_entries",
	"uuid",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"qos_config_type",
}

// QosConfigRefFields is db reference fields for QosConfig
var QosConfigRefFields = map[string][]string{

	"global_system_config": {
	// <utils.Schema Value>

	},
}

const insertQosConfigGlobalSystemConfigQuery = "insert into `ref_qos_config_global_system_config` (`from`, `to` ) values (?, ?);"

// CreateQosConfig inserts QosConfig to DB
func CreateQosConfig(tx *sql.Tx, model *models.QosConfig) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertQosConfigQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.MPLSExpEntries),
		utils.MustJSON(model.VlanPriorityEntries),
		int(model.DefaultForwardingClassID),
		utils.MustJSON(model.DSCPEntries),
		string(model.UUID),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.QosConfigType))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtGlobalSystemConfigRef, err := tx.Prepare(insertQosConfigGlobalSystemConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing GlobalSystemConfigRefs create statement failed")
	}
	defer stmtGlobalSystemConfigRef.Close()
	for _, ref := range model.GlobalSystemConfigRefs {
		_, err = stmtGlobalSystemConfigRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "GlobalSystemConfigRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanQosConfig(values map[string]interface{}) (*models.QosConfig, error) {
	m := models.MakeQosConfig()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["mpls_exp_entries"]; ok {

		json.Unmarshal(value.([]byte), &m.MPLSExpEntries)

	}

	if value, ok := values["vlan_priority_entries"]; ok {

		json.Unmarshal(value.([]byte), &m.VlanPriorityEntries)

	}

	if value, ok := values["default_forwarding_class_id"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.DefaultForwardingClassID = models.ForwardingClassId(castedValue)

	}

	if value, ok := values["dscp_entries"]; ok {

		json.Unmarshal(value.([]byte), &m.DSCPEntries)

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

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["qos_config_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.QosConfigType = models.QosConfigType(castedValue)

	}

	if value, ok := values["ref_global_system_config"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.QosConfigGlobalSystemConfigRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.GlobalSystemConfigRefs = append(m.GlobalSystemConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListQosConfig lists QosConfig with list spec.
func ListQosConfig(tx *sql.Tx, spec *db.ListSpec) ([]*models.QosConfig, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "qos_config"
	spec.Fields = QosConfigFields
	spec.RefFields = QosConfigRefFields
	result := models.MakeQosConfigSlice()
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
		m, err := scanQosConfig(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowQosConfig shows QosConfig resource
func ShowQosConfig(tx *sql.Tx, uuid string) (*models.QosConfig, error) {
	list, err := ListQosConfig(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateQosConfig updates a resource
func UpdateQosConfig(tx *sql.Tx, uuid string, model *models.QosConfig) error {
	//TODO(nati) support update
	return nil
}

// DeleteQosConfig deletes a resource
func DeleteQosConfig(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteQosConfigQuery)
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
