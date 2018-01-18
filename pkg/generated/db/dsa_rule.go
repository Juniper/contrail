package db

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertDsaRuleQuery = "insert into `dsa_rule` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`subscriber`,`ep_version`,`ep_type`,`ip_prefix_len`,`ip_prefix`,`ep_id`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteDsaRuleQuery = "delete from `dsa_rule` where uuid = ?"

// DsaRuleFields is db columns for DsaRule
var DsaRuleFields = []string{
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
	"subscriber",
	"ep_version",
	"ep_type",
	"ip_prefix_len",
	"ip_prefix",
	"ep_id",
	"display_name",
	"key_value_pair",
}

// DsaRuleRefFields is db reference fields for DsaRule
var DsaRuleRefFields = map[string][]string{}

// DsaRuleBackRefFields is db back reference fields for DsaRule
var DsaRuleBackRefFields = map[string][]string{}

// DsaRuleParentTypes is possible parents for DsaRule
var DsaRuleParents = []string{

	"discovery_service_assignment",
}

// CreateDsaRule inserts DsaRule to DB
func CreateDsaRule(tx *sql.Tx, model *models.DsaRule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDsaRuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertDsaRuleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
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
		common.MustJSON(model.DsaRuleEntry.Subscriber),
		string(model.DsaRuleEntry.Publisher.EpVersion),
		string(model.DsaRuleEntry.Publisher.EpType),
		int(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen),
		string(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix),
		string(model.DsaRuleEntry.Publisher.EpID),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "dsa_rule",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "dsa_rule", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanDsaRule(values map[string]interface{}) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

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

	if value, ok := values["subscriber"]; ok {

		json.Unmarshal(value.([]byte), &m.DsaRuleEntry.Subscriber)

	}

	if value, ok := values["ep_version"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpVersion = castedValue

	}

	if value, ok := values["ep_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpType = castedValue

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix = castedValue

	}

	if value, ok := values["ep_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpID = castedValue

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

// ListDsaRule lists DsaRule with list spec.
func ListDsaRule(tx *sql.Tx, spec *common.ListSpec) ([]*models.DsaRule, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "dsa_rule"
	spec.Fields = DsaRuleFields
	spec.RefFields = DsaRuleRefFields
	spec.BackRefFields = DsaRuleBackRefFields
	result := models.MakeDsaRuleSlice()

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
		m, err := scanDsaRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateDsaRule updates a resource
func UpdateDsaRule(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateDsaRuleQuery = "update `dsa_rule` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateDsaRuleQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateDsaRuleQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateDsaRuleQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateDsaRuleQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateDsaRuleQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateDsaRuleQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateDsaRuleQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateDsaRuleQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateDsaRuleQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateDsaRuleQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateDsaRuleQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateDsaRuleQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateDsaRuleQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateDsaRuleQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateDsaRuleQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateDsaRuleQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateDsaRuleQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateDsaRuleQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateDsaRuleQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Subscriber", "."); ok {
		updateDsaRuleQuery += "`subscriber` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Publisher.EpVersion", "."); ok {
		updateDsaRuleQuery += "`ep_version` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Publisher.EpType", "."); ok {
		updateDsaRuleQuery += "`ep_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen", "."); ok {
		updateDsaRuleQuery += "`ip_prefix_len` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Publisher.EpPrefix.IPPrefix", "."); ok {
		updateDsaRuleQuery += "`ip_prefix` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DsaRuleEntry.Publisher.EpID", "."); ok {
		updateDsaRuleQuery += "`ep_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateDsaRuleQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDsaRuleQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateDsaRuleQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDsaRuleQuery += ","
	}

	updateDsaRuleQuery =
		updateDsaRuleQuery[:len(updateDsaRuleQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateDsaRuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateDsaRuleQuery,
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

// DeleteDsaRule deletes a resource
func DeleteDsaRule(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteDsaRuleQuery
	selectQuery := "select count(uuid) from dsa_rule where uuid = ?"
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
