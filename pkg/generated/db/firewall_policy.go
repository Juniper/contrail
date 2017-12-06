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

const insertFirewallPolicyQuery = "insert into `firewall_policy` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFirewallPolicyQuery = "update `firewall_policy` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`display_name` = ?;"
const deleteFirewallPolicyQuery = "delete from `firewall_policy` where uuid = ?"

// FirewallPolicyFields is db columns for FirewallPolicy
var FirewallPolicyFields = []string{
	"key_value_pair",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"uuid",
	"fq_name",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"display_name",
}

// FirewallPolicyRefFields is db reference fields for FirewallPolicy
var FirewallPolicyRefFields = map[string][]string{

	"firewall_rule": {
	// <utils.Schema Value>

	},

	"security_logging_object": {
	// <utils.Schema Value>

	},
}

const insertFirewallPolicyFirewallRuleQuery = "insert into `ref_firewall_policy_firewall_rule` (`from`, `to` ) values (?, ?);"

const insertFirewallPolicySecurityLoggingObjectQuery = "insert into `ref_firewall_policy_security_logging_object` (`from`, `to` ) values (?, ?);"

// CreateFirewallPolicy inserts FirewallPolicy to DB
func CreateFirewallPolicy(tx *sql.Tx, model *models.FirewallPolicy) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFirewallPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertFirewallPolicyQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.DisplayName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtFirewallRuleRef, err := tx.Prepare(insertFirewallPolicyFirewallRuleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing FirewallRuleRefs create statement failed")
	}
	defer stmtFirewallRuleRef.Close()
	for _, ref := range model.FirewallRuleRefs {
		_, err = stmtFirewallRuleRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "FirewallRuleRefs create failed")
		}
	}

	stmtSecurityLoggingObjectRef, err := tx.Prepare(insertFirewallPolicySecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityLoggingObjectRefs create statement failed")
	}
	defer stmtSecurityLoggingObjectRef.Close()
	for _, ref := range model.SecurityLoggingObjectRefs {
		_, err = stmtSecurityLoggingObjectRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "SecurityLoggingObjectRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanFirewallPolicy(values map[string]interface{}) (*models.FirewallPolicy, error) {
	m := models.MakeFirewallPolicy()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["ref_firewall_rule"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.FirewallPolicyFirewallRuleRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.FirewallRuleRefs = append(m.FirewallRuleRefs, referenceModel)

			attr := models.MakeFirewallSequence()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_security_logging_object"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.FirewallPolicySecurityLoggingObjectRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.SecurityLoggingObjectRefs = append(m.SecurityLoggingObjectRefs, referenceModel)

		}
	}

	return m, nil
}

// ListFirewallPolicy lists FirewallPolicy with list spec.
func ListFirewallPolicy(tx *sql.Tx, spec *db.ListSpec) ([]*models.FirewallPolicy, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "firewall_policy"
	spec.Fields = FirewallPolicyFields
	spec.RefFields = FirewallPolicyRefFields
	result := models.MakeFirewallPolicySlice()
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
		m, err := scanFirewallPolicy(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowFirewallPolicy shows FirewallPolicy resource
func ShowFirewallPolicy(tx *sql.Tx, uuid string) (*models.FirewallPolicy, error) {
	list, err := ListFirewallPolicy(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateFirewallPolicy updates a resource
func UpdateFirewallPolicy(tx *sql.Tx, uuid string, model *models.FirewallPolicy) error {
	//TODO(nati) support update
	return nil
}

// DeleteFirewallPolicy deletes a resource
func DeleteFirewallPolicy(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFirewallPolicyQuery)
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
