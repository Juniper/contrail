package db

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertBridgeDomainQuery = "insert into `bridge_domain` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`mac_move_time_window`,`mac_move_limit_action`,`mac_move_limit`,`mac_limit_action`,`mac_limit`,`mac_learning_enabled`,`mac_aging_time`,`isid`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBridgeDomainQuery = "delete from `bridge_domain` where uuid = ?"

// BridgeDomainFields is db columns for BridgeDomain
var BridgeDomainFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"mac_move_time_window",
	"mac_move_limit_action",
	"mac_move_limit",
	"mac_limit_action",
	"mac_limit",
	"mac_learning_enabled",
	"mac_aging_time",
	"isid",
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

// BridgeDomainRefFields is db reference fields for BridgeDomain
var BridgeDomainRefFields = map[string][]string{}

// BridgeDomainBackRefFields is db back reference fields for BridgeDomain
var BridgeDomainBackRefFields = map[string][]string{}

// BridgeDomainParentTypes is possible parents for BridgeDomain
var BridgeDomainParents = []string{

	"virtual_network",
}

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
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		int(model.MacMoveControl.MacMoveTimeWindow),
		string(model.MacMoveControl.MacMoveLimitAction),
		int(model.MacMoveControl.MacMoveLimit),
		string(model.MacLimitControl.MacLimitAction),
		int(model.MacLimitControl.MacLimit),
		bool(model.MacLearningEnabled),
		int(model.MacAgingTime),
		int(model.Isid),
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "bridge_domain",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "bridge_domain", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBridgeDomain(values map[string]interface{}) (*models.BridgeDomain, error) {
	m := models.MakeBridgeDomain()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

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

	if value, ok := values["mac_move_time_window"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveTimeWindow = models.MACMoveTimeWindow(castedValue)

	}

	if value, ok := values["mac_move_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacMoveControl.MacMoveLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_move_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacMoveControl.MacMoveLimit = castedValue

	}

	if value, ok := values["mac_limit_action"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacLimitControl.MacLimitAction = models.MACLimitExceedActionType(castedValue)

	}

	if value, ok := values["mac_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacLimitControl.MacLimit = castedValue

	}

	if value, ok := values["mac_learning_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.MacLearningEnabled = castedValue

	}

	if value, ok := values["mac_aging_time"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.MacAgingTime = models.MACAgingTime(castedValue)

	}

	if value, ok := values["isid"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Isid = models.IsidType(castedValue)

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
	spec.BackRefFields = BridgeDomainBackRefFields
	result := models.MakeBridgeDomainSlice()

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
		m, err := scanBridgeDomain(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateBridgeDomain updates a resource
func UpdateBridgeDomain(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateBridgeDomainQuery = "update `bridge_domain` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateBridgeDomainQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateBridgeDomainQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateBridgeDomainQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateBridgeDomainQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateBridgeDomainQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateBridgeDomainQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateBridgeDomainQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveTimeWindow", "."); ok {
		updateBridgeDomainQuery += "`mac_move_time_window` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimitAction", "."); ok {
		updateBridgeDomainQuery += "`mac_move_limit_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacMoveControl.MacMoveLimit", "."); ok {
		updateBridgeDomainQuery += "`mac_move_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimitAction", "."); ok {
		updateBridgeDomainQuery += "`mac_limit_action` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLimitControl.MacLimit", "."); ok {
		updateBridgeDomainQuery += "`mac_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacLearningEnabled", "."); ok {
		updateBridgeDomainQuery += "`mac_learning_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacAgingTime", "."); ok {
		updateBridgeDomainQuery += "`mac_aging_time` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Isid", "."); ok {
		updateBridgeDomainQuery += "`isid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateBridgeDomainQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateBridgeDomainQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateBridgeDomainQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateBridgeDomainQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateBridgeDomainQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateBridgeDomainQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateBridgeDomainQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateBridgeDomainQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateBridgeDomainQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateBridgeDomainQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateBridgeDomainQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateBridgeDomainQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateBridgeDomainQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBridgeDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateBridgeDomainQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBridgeDomainQuery += ","
	}

	updateBridgeDomainQuery =
		updateBridgeDomainQuery[:len(updateBridgeDomainQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateBridgeDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateBridgeDomainQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteBridgeDomain deletes a resource
func DeleteBridgeDomain(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteBridgeDomainQuery
	selectQuery := "select count(uuid) from bridge_domain where uuid = ?"
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
