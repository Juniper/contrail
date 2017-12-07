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

const insertLoadbalancerMemberQuery = "insert into `loadbalancer_member` (`status`,`status_description`,`weight`,`admin_state`,`address`,`protocol_port`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerMemberQuery = "update `loadbalancer_member` set `status` = ?,`status_description` = ?,`weight` = ?,`admin_state` = ?,`address` = ?,`protocol_port` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteLoadbalancerMemberQuery = "delete from `loadbalancer_member` where uuid = ?"

// LoadbalancerMemberFields is db columns for LoadbalancerMember
var LoadbalancerMemberFields = []string{
	"status",
	"status_description",
	"weight",
	"admin_state",
	"address",
	"protocol_port",
	"user_visible",
	"last_modified",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"enable",
	"description",
	"created",
	"creator",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"uuid",
	"fq_name",
}

// LoadbalancerMemberRefFields is db reference fields for LoadbalancerMember
var LoadbalancerMemberRefFields = map[string][]string{}

// CreateLoadbalancerMember inserts LoadbalancerMember to DB
func CreateLoadbalancerMember(tx *sql.Tx, model *models.LoadbalancerMember) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerMemberQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerMemberQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.LoadbalancerMemberProperties.Status),
		string(model.LoadbalancerMemberProperties.StatusDescription),
		int(model.LoadbalancerMemberProperties.Weight),
		bool(model.LoadbalancerMemberProperties.AdminState),
		string(model.LoadbalancerMemberProperties.Address),
		int(model.LoadbalancerMemberProperties.ProtocolPort),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLoadbalancerMember(values map[string]interface{}) (*models.LoadbalancerMember, error) {
	m := models.MakeLoadbalancerMember()

	if value, ok := values["status"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerMemberProperties.Status = castedValue

	}

	if value, ok := values["status_description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerMemberProperties.StatusDescription = castedValue

	}

	if value, ok := values["weight"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.LoadbalancerMemberProperties.Weight = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.LoadbalancerMemberProperties.AdminState = castedValue

	}

	if value, ok := values["address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.LoadbalancerMemberProperties.Address = models.IpAddressType(castedValue)

	}

	if value, ok := values["protocol_port"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.LoadbalancerMemberProperties.ProtocolPort = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

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

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	return m, nil
}

// ListLoadbalancerMember lists LoadbalancerMember with list spec.
func ListLoadbalancerMember(tx *sql.Tx, spec *db.ListSpec) ([]*models.LoadbalancerMember, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "loadbalancer_member"
	spec.Fields = LoadbalancerMemberFields
	spec.RefFields = LoadbalancerMemberRefFields
	result := models.MakeLoadbalancerMemberSlice()
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
		m, err := scanLoadbalancerMember(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLoadbalancerMember shows LoadbalancerMember resource
func ShowLoadbalancerMember(tx *sql.Tx, uuid string) (*models.LoadbalancerMember, error) {
	list, err := ListLoadbalancerMember(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLoadbalancerMember updates a resource
func UpdateLoadbalancerMember(tx *sql.Tx, uuid string, model *models.LoadbalancerMember) error {
	//TODO(nati) support update
	return nil
}

// DeleteLoadbalancerMember deletes a resource
func DeleteLoadbalancerMember(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerMemberQuery)
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
