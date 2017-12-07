package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertBridgeDomainQuery = "insert into `bridge_domain` (`isid`,`mac_learning_enabled`,`mac_move_time_window`,`mac_move_limit`,`mac_move_limit_action`,`owner`,`owner_access`,`global_access`,`share`,`mac_aging_time`,`mac_limit`,`mac_limit_action`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBridgeDomainQuery = "update `bridge_domain` set `isid` = ?,`mac_learning_enabled` = ?,`mac_move_time_window` = ?,`mac_move_limit` = ?,`mac_move_limit_action` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`mac_aging_time` = ?,`mac_limit` = ?,`mac_limit_action` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteBridgeDomainQuery = "delete from `bridge_domain` where uuid = ?"

// BridgeDomainFields is db columns for BridgeDomain
var BridgeDomainFields = []string{
	"isid",
	"mac_learning_enabled",
	"mac_move_time_window",
	"mac_move_limit",
	"mac_move_limit_action",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"mac_aging_time",
	"mac_limit",
	"mac_limit_action",
	"uuid",
	"fq_name",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"display_name",
	"key_value_pair",
}

// BridgeDomainRefFields is db reference fields for BridgeDomain
var BridgeDomainRefFields = map[string][]string{}

// CreateBridgeDomain inserts BridgeDomain to DB
func CreateBridgeDomain(tx *sql.Tx, model *models.BridgeDomain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBridgeDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBridgeDomainQuery,
	}).Debug("create query")
	_, err = stmt.Exec(int(model.Isid),
		bool(model.MacLearningEnabled),
		int(model.MacMoveControl.MacMoveTimeWindow),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacMoveControl.MacMoveLimitAction),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		int(model.MacAgingTime),
		int(model.MacLimitControl.MacLimit),
		string(model.MacLimitControl.MacLimitAction),
		string(model.UUID),
		common.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanBridgeDomain(values map[string]interface{}) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

	if value, ok := values["isid"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Isid = models.IsidType(castedValue)

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MacLearningEnabled = castedValue

	}

	if value, ok := values["mac_move_time_window"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

	}

	if value, ok := values["mac_move_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveLimit = castedValue

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

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

	if value, ok := values["mac_aging_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacAgingTime = models.MACAgingTime(castedValue)

	}

	if value, ok := values["mac_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacLimitControl.MacLimit = castedValue

	}

	if value, ok := values["mac_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

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

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

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

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListBridgeDomain lists BridgeDomain with list spec.
func ListBridgeDomain(tx *sql.Tx, spec *common.ListSpec) ([]*models.BridgeDomain, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "bridge_domain"
	spec.Fields = BridgeDomainFields
	spec.RefFields = BridgeDomainRefFields
	result := models.MakeBridgeDomainSlice()
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
		m, err := scanBridgeDomain(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowBridgeDomain shows BridgeDomain resource
func ShowBridgeDomain(tx *sql.Tx, uuid string) (*models.BridgeDomain, error) {
	list, err := ListBridgeDomain(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateBridgeDomain updates a resource
func UpdateBridgeDomain(tx *sql.Tx, uuid string, model *models.BridgeDomain) error {
	//TODO(nati) support update
	return nil
}

// DeleteBridgeDomain deletes a resource
func DeleteBridgeDomain(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBridgeDomainQuery)
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
