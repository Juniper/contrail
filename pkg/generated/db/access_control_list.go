package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertAccessControlListQuery = "insert into `access_control_list` (`key_value_pair`,`access_control_list_hash`,`dynamic`,`acl_rule`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAccessControlListQuery = "update `access_control_list` set `key_value_pair` = ?,`access_control_list_hash` = ?,`dynamic` = ?,`acl_rule` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?;"
const deleteAccessControlListQuery = "delete from `access_control_list` where uuid = ?"

// AccessControlListFields is db columns for AccessControlList
var AccessControlListFields = []string{
	"key_value_pair",
	"access_control_list_hash",
	"dynamic",
	"acl_rule",
	"owner_access",
	"global_access",
	"share",
	"owner",
	"uuid",
	"fq_name",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"display_name",
}

// AccessControlListRefFields is db reference fields for AccessControlList
var AccessControlListRefFields = map[string][]string{}

// CreateAccessControlList inserts AccessControlList to DB
func CreateAccessControlList(tx *sql.Tx, model *models.AccessControlList) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAccessControlListQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertAccessControlListQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.Annotations.KeyValuePair),
		common.MustJSON(model.AccessControlListHash),
		bool(model.AccessControlListEntries.Dynamic),
		common.MustJSON(model.AccessControlListEntries.ACLRule),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		common.MustJSON(model.FQName),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanAccessControlList(values map[string]interface{}) (*models.AccessControlList, error) {
	m := models.MakeAccessControlList()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["access_control_list_hash"]; ok {

		json.Unmarshal(value.([]byte), &m.AccessControlListHash)

	}

	if value, ok := values["dynamic"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.AccessControlListEntries.Dynamic = castedValue

	}

	if value, ok := values["acl_rule"]; ok {

		json.Unmarshal(value.([]byte), &m.AccessControlListEntries.ACLRule)

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

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	return m, nil
}

// ListAccessControlList lists AccessControlList with list spec.
func ListAccessControlList(tx *sql.Tx, spec *common.ListSpec) ([]*models.AccessControlList, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "access_control_list"
	spec.Fields = AccessControlListFields
	spec.RefFields = AccessControlListRefFields
	result := models.MakeAccessControlListSlice()
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
		m, err := scanAccessControlList(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowAccessControlList shows AccessControlList resource
func ShowAccessControlList(tx *sql.Tx, uuid string) (*models.AccessControlList, error) {
	list, err := ListAccessControlList(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateAccessControlList updates a resource
func UpdateAccessControlList(tx *sql.Tx, uuid string, model *models.AccessControlList) error {
	//TODO(nati) support update
	return nil
}

// DeleteAccessControlList deletes a resource
func DeleteAccessControlList(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAccessControlListQuery)
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
