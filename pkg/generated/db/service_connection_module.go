package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceConnectionModuleQuery = "insert into `service_connection_module` (`display_name`,`key_value_pair`,`service_type`,`e2_service`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceConnectionModuleQuery = "update `service_connection_module` set `display_name` = ?,`key_value_pair` = ?,`service_type` = ?,`e2_service` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?;"
const deleteServiceConnectionModuleQuery = "delete from `service_connection_module` where uuid = ?"

// ServiceConnectionModuleFields is db columns for ServiceConnectionModule
var ServiceConnectionModuleFields = []string{
	"display_name",
	"key_value_pair",
	"service_type",
	"e2_service",
	"owner_access",
	"global_access",
	"share",
	"owner",
	"uuid",
	"fq_name",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"enable",
	"description",
}

// ServiceConnectionModuleRefFields is db reference fields for ServiceConnectionModule
var ServiceConnectionModuleRefFields = map[string][]string{

	"service_object": {
	// <common.Schema Value>

	},
}

const insertServiceConnectionModuleServiceObjectQuery = "insert into `ref_service_connection_module_service_object` (`from`, `to` ) values (?, ?);"

// CreateServiceConnectionModule inserts ServiceConnectionModule to DB
func CreateServiceConnectionModule(tx *sql.Tx, model *models.ServiceConnectionModule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceConnectionModuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceConnectionModuleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.ServiceType),
		string(model.E2Service),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		common.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceObjectRef, err := tx.Prepare(insertServiceConnectionModuleServiceObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceObjectRefs create statement failed")
	}
	defer stmtServiceObjectRef.Close()
	for _, ref := range model.ServiceObjectRefs {

		_, err = stmtServiceObjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceObjectRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceConnectionModule(values map[string]interface{}) (*models.ServiceConnectionModule, error) {
	m := models.MakeServiceConnectionModule()

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["service_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceType = models.ServiceConnectionType(castedValue)

	}

	if value, ok := values["e2_service"]; ok {

		castedValue := common.InterfaceToString(value)

		m.E2Service = models.E2servicetype(castedValue)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["ref_service_object"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ServiceConnectionModuleServiceObjectRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceObjectRefs = append(m.ServiceObjectRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceConnectionModule lists ServiceConnectionModule with list spec.
func ListServiceConnectionModule(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceConnectionModule, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_connection_module"
	spec.Fields = ServiceConnectionModuleFields
	spec.RefFields = ServiceConnectionModuleRefFields
	result := models.MakeServiceConnectionModuleSlice()
	query, columns, values := common.BuildListQuery(spec)
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
		m, err := scanServiceConnectionModule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceConnectionModule shows ServiceConnectionModule resource
func ShowServiceConnectionModule(tx *sql.Tx, uuid string) (*models.ServiceConnectionModule, error) {
	list, err := ListServiceConnectionModule(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceConnectionModule updates a resource
func UpdateServiceConnectionModule(tx *sql.Tx, uuid string, model *models.ServiceConnectionModule) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceConnectionModule deletes a resource
func DeleteServiceConnectionModule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceConnectionModuleQuery)
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
