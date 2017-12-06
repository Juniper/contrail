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

const insertDsaRuleQuery = "insert into `dsa_rule` (`fq_name`,`subscriber`,`ip_prefix_len`,`ip_prefix`,`ep_version`,`ep_id`,`ep_type`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDsaRuleQuery = "update `dsa_rule` set `fq_name` = ?,`subscriber` = ?,`ip_prefix_len` = ?,`ip_prefix` = ?,`ep_version` = ?,`ep_id` = ?,`ep_type` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?;"
const deleteDsaRuleQuery = "delete from `dsa_rule` where uuid = ?"

// DsaRuleFields is db columns for DsaRule
var DsaRuleFields = []string{
	"fq_name",
	"subscriber",
	"ip_prefix_len",
	"ip_prefix",
	"ep_version",
	"ep_id",
	"ep_type",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"display_name",
	"key_value_pair",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"uuid",
}

// DsaRuleRefFields is db reference fields for DsaRule
var DsaRuleRefFields = map[string][]string{}

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
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		utils.MustJSON(model.DsaRuleEntry.Subscriber),
		int(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen),
		string(model.DsaRuleEntry.Publisher.EpPrefix.IPPrefix),
		string(model.DsaRuleEntry.Publisher.EpVersion),
		string(model.DsaRuleEntry.Publisher.EpID),
		string(model.DsaRuleEntry.Publisher.EpType),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanDsaRule(values map[string]interface{}) (*models.DsaRule, error) {
	m := models.MakeDsaRule()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["subscriber"]; ok {

		json.Unmarshal(value.([]byte), &m.DsaRuleEntry.Subscriber)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefixLen = castedValue

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpPrefix.IPPrefix = castedValue

	}

	if value, ok := values["ep_version"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpVersion = castedValue

	}

	if value, ok := values["ep_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpID = castedValue

	}

	if value, ok := values["ep_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DsaRuleEntry.Publisher.EpType = castedValue

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

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	return m, nil
}

// ListDsaRule lists DsaRule with list spec.
func ListDsaRule(tx *sql.Tx, spec *db.ListSpec) ([]*models.DsaRule, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "dsa_rule"
	spec.Fields = DsaRuleFields
	spec.RefFields = DsaRuleRefFields
	result := models.MakeDsaRuleSlice()
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
		m, err := scanDsaRule(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowDsaRule shows DsaRule resource
func ShowDsaRule(tx *sql.Tx, uuid string) (*models.DsaRule, error) {
	list, err := ListDsaRule(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateDsaRule updates a resource
func UpdateDsaRule(tx *sql.Tx, uuid string, model *models.DsaRule) error {
	//TODO(nati) support update
	return nil
}

// DeleteDsaRule deletes a resource
func DeleteDsaRule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDsaRuleQuery)
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
