package db

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertApplicationPolicySetQuery = "insert into `application_policy_set` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`,`all_applications`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteApplicationPolicySetQuery = "delete from `application_policy_set` where uuid = ?"

// ApplicationPolicySetFields is db columns for ApplicationPolicySet
var ApplicationPolicySetFields = []string{
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
	"display_name",
	"key_value_pair",
	"all_applications",
}

// ApplicationPolicySetRefFields is db reference fields for ApplicationPolicySet
var ApplicationPolicySetRefFields = map[string][]string{

	"global_vrouter_config": {
	// <common.Schema Value>

	},

	"firewall_policy": {
		// <common.Schema Value>
		"sequence",
	},
}

// ApplicationPolicySetBackRefFields is db back reference fields for ApplicationPolicySet
var ApplicationPolicySetBackRefFields = map[string][]string{}

// ApplicationPolicySetParentTypes is possible parents for ApplicationPolicySet
var ApplicationPolicySetParents = []string{

	"policy_management",

	"project",
}

const insertApplicationPolicySetFirewallPolicyQuery = "insert into `ref_application_policy_set_firewall_policy` (`from`, `to` ,`sequence`) values (?, ?,?);"

const insertApplicationPolicySetGlobalVrouterConfigQuery = "insert into `ref_application_policy_set_global_vrouter_config` (`from`, `to` ) values (?, ?);"

// CreateApplicationPolicySet inserts ApplicationPolicySet to DB
func CreateApplicationPolicySet(tx *sql.Tx, model *models.ApplicationPolicySet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertApplicationPolicySetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertApplicationPolicySetQuery,
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
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		bool(model.AllApplications))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtFirewallPolicyRef, err := tx.Prepare(insertApplicationPolicySetFirewallPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing FirewallPolicyRefs create statement failed")
	}
	defer stmtFirewallPolicyRef.Close()
	for _, ref := range model.FirewallPolicyRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeFirewallSequence()
		}

		_, err = stmtFirewallPolicyRef.Exec(model.UUID, ref.UUID, string(ref.Attr.Sequence))
		if err != nil {
			return errors.Wrap(err, "FirewallPolicyRefs create failed")
		}
	}

	stmtGlobalVrouterConfigRef, err := tx.Prepare(insertApplicationPolicySetGlobalVrouterConfigQuery)
	if err != nil {
		return errors.Wrap(err, "preparing GlobalVrouterConfigRefs create statement failed")
	}
	defer stmtGlobalVrouterConfigRef.Close()
	for _, ref := range model.GlobalVrouterConfigRefs {

		_, err = stmtGlobalVrouterConfigRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "GlobalVrouterConfigRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "application_policy_set",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "application_policy_set", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanApplicationPolicySet(values map[string]interface{}) (*models.ApplicationPolicySet, error) {
	m := models.MakeApplicationPolicySet()

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["all_applications"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.AllApplications = castedValue

	}

	if value, ok := values["ref_firewall_policy"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ApplicationPolicySetFirewallPolicyRef{}
			referenceModel.UUID = uuid
			m.FirewallPolicyRefs = append(m.FirewallPolicyRefs, referenceModel)

			attr := models.MakeFirewallSequence()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_global_vrouter_config"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ApplicationPolicySetGlobalVrouterConfigRef{}
			referenceModel.UUID = uuid
			m.GlobalVrouterConfigRefs = append(m.GlobalVrouterConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListApplicationPolicySet lists ApplicationPolicySet with list spec.
func ListApplicationPolicySet(tx *sql.Tx, spec *common.ListSpec) ([]*models.ApplicationPolicySet, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "application_policy_set"
	spec.Fields = ApplicationPolicySetFields
	spec.RefFields = ApplicationPolicySetRefFields
	spec.BackRefFields = ApplicationPolicySetBackRefFields
	result := models.MakeApplicationPolicySetSlice()

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
		m, err := scanApplicationPolicySet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateApplicationPolicySet updates a resource
func UpdateApplicationPolicySet(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateApplicationPolicySetQuery = "update `application_policy_set` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateApplicationPolicySetQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateApplicationPolicySetQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateApplicationPolicySetQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateApplicationPolicySetQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateApplicationPolicySetQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateApplicationPolicySetQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateApplicationPolicySetQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateApplicationPolicySetQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateApplicationPolicySetQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateApplicationPolicySetQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateApplicationPolicySetQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateApplicationPolicySetQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateApplicationPolicySetQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateApplicationPolicySetQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateApplicationPolicySetQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateApplicationPolicySetQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateApplicationPolicySetQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateApplicationPolicySetQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateApplicationPolicySetQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateApplicationPolicySetQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateApplicationPolicySetQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateApplicationPolicySetQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AllApplications", "."); ok {
		updateApplicationPolicySetQuery += "`all_applications` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateApplicationPolicySetQuery += ","
	}

	updateApplicationPolicySetQuery =
		updateApplicationPolicySetQuery[:len(updateApplicationPolicySetQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateApplicationPolicySetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateApplicationPolicySetQuery,
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

// DeleteApplicationPolicySet deletes a resource
func DeleteApplicationPolicySet(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteApplicationPolicySetQuery
	selectQuery := "select count(uuid) from application_policy_set where uuid = ?"
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
