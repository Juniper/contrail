package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertSecurityLoggingObjectQuery = "insert into `security_logging_object` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`rule`,`security_logging_object_rate`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSecurityLoggingObjectQuery = "update `security_logging_object` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`rule` = ?,`security_logging_object_rate` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteSecurityLoggingObjectQuery = "delete from `security_logging_object` where uuid = ?"

// SecurityLoggingObjectFields is db columns for SecurityLoggingObject
var SecurityLoggingObjectFields = []string{
	"key_value_pair",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"uuid",
	"fq_name",
	"rule",
	"security_logging_object_rate",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"enable",
	"description",
	"display_name",
}

// SecurityLoggingObjectRefFields is db reference fields for SecurityLoggingObject
var SecurityLoggingObjectRefFields = map[string][]string{

	"security_group": {
		// <common.Schema Value>
		"rule",
	},

	"network_policy": {
	// <common.Schema Value>

	},
}

const insertSecurityLoggingObjectSecurityGroupQuery = "insert into `ref_security_logging_object_security_group` (`from`, `to` ,`rule`) values (?, ?,?);"

const insertSecurityLoggingObjectNetworkPolicyQuery = "insert into `ref_security_logging_object_network_policy` (`from`, `to` ) values (?, ?);"

// CreateSecurityLoggingObject inserts SecurityLoggingObject to DB
func CreateSecurityLoggingObject(tx *sql.Tx, model *models.SecurityLoggingObject) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSecurityLoggingObjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertSecurityLoggingObjectQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.UUID),
		common.MustJSON(model.FQName),
		common.MustJSON(model.SecurityLoggingObjectRules.Rule),
		int(model.SecurityLoggingObjectRate),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtSecurityGroupRef, err := tx.Prepare(insertSecurityLoggingObjectSecurityGroupQuery)
	if err != nil {
		return errors.Wrap(err, "preparing SecurityGroupRefs create statement failed")
	}
	defer stmtSecurityGroupRef.Close()
	for _, ref := range model.SecurityGroupRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeSecurityLoggingObjectRuleListType()
		}

		_, err = stmtSecurityGroupRef.Exec(model.UUID, ref.UUID, common.MustJSON(ref.Attr.Rule))
		if err != nil {
			return errors.Wrap(err, "SecurityGroupRefs create failed")
		}
	}

	stmtNetworkPolicyRef, err := tx.Prepare(insertSecurityLoggingObjectNetworkPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkPolicyRefs create statement failed")
	}
	defer stmtNetworkPolicyRef.Close()
	for _, ref := range model.NetworkPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeSecurityLoggingObjectRuleListType()
		}

		_, err = stmtNetworkPolicyRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NetworkPolicyRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanSecurityLoggingObject(values map[string]interface{}) (*models.SecurityLoggingObject, error) {
	m := models.MakeSecurityLoggingObject()

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["rule"]; ok {

		json.Unmarshal(value.([]byte), &m.SecurityLoggingObjectRules.Rule)

	}

	if value, ok := values["security_logging_object_rate"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.SecurityLoggingObjectRate = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["ref_security_group"]; ok {
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
			referenceModel := &models.SecurityLoggingObjectSecurityGroupRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.SecurityGroupRefs = append(m.SecurityGroupRefs, referenceModel)

			attr := models.MakeSecurityLoggingObjectRuleListType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_network_policy"]; ok {
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
			referenceModel := &models.SecurityLoggingObjectNetworkPolicyRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.NetworkPolicyRefs = append(m.NetworkPolicyRefs, referenceModel)

			attr := models.MakeSecurityLoggingObjectRuleListType()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListSecurityLoggingObject lists SecurityLoggingObject with list spec.
func ListSecurityLoggingObject(tx *sql.Tx, spec *common.ListSpec) ([]*models.SecurityLoggingObject, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "security_logging_object"
	spec.Fields = SecurityLoggingObjectFields
	spec.RefFields = SecurityLoggingObjectRefFields
	result := models.MakeSecurityLoggingObjectSlice()
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
		m, err := scanSecurityLoggingObject(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowSecurityLoggingObject shows SecurityLoggingObject resource
func ShowSecurityLoggingObject(tx *sql.Tx, uuid string) (*models.SecurityLoggingObject, error) {
	list, err := ListSecurityLoggingObject(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateSecurityLoggingObject updates a resource
func UpdateSecurityLoggingObject(tx *sql.Tx, uuid string, model *models.SecurityLoggingObject) error {
	//TODO(nati) support update
	return nil
}

// DeleteSecurityLoggingObject deletes a resource
func DeleteSecurityLoggingObject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSecurityLoggingObjectQuery)
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
