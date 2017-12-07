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

const insertApplicationPolicySetQuery = "insert into `application_policy_set` (`uuid`,`fq_name`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`all_applications`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateApplicationPolicySetQuery = "update `application_policy_set` set `uuid` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`all_applications` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteApplicationPolicySetQuery = "delete from `application_policy_set` where uuid = ?"

// ApplicationPolicySetFields is db columns for ApplicationPolicySet
var ApplicationPolicySetFields = []string{
	"uuid",
	"fq_name",
	"other_access",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"display_name",
	"all_applications",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
}

// ApplicationPolicySetRefFields is db reference fields for ApplicationPolicySet
var ApplicationPolicySetRefFields = map[string][]string{

	"global_vrouter_config": {
	// <utils.Schema Value>

	},

	"firewall_policy": {
		// <utils.Schema Value>
		"sequence",
	},
}

const insertApplicationPolicySetGlobalVrouterConfigQuery = "insert into `ref_application_policy_set_global_vrouter_config` (`from`, `to` ) values (?, ?);"

const insertApplicationPolicySetFirewallPolicyQuery = "insert into `ref_application_policy_set_firewall_policy` (`from`, `to` ,`sequence`) values (?, ?,?);"

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
		utils.MustJSON(model.FQName),
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
		string(model.DisplayName),
		bool(model.AllApplications),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtFirewallPolicyRef, err := tx.Prepare(insertApplicationPolicySetFirewallPolicyQuery)
	if err != nil {
		return errors.Wrap(err, "preparing FirewallPolicyRefs create statement failed")
	}
	defer stmtFirewallPolicyRef.Close()
	for _, ref := range model.FirewallPolicyRefs {
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

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanApplicationPolicySet(values map[string]interface{}) (*models.ApplicationPolicySet, error) {
	m := models.MakeApplicationPolicySet()

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

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

	if value, ok := values["all_applications"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.AllApplications = castedValue

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

	if value, ok := values["ref_firewall_policy"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ApplicationPolicySetFirewallPolicyRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.FirewallPolicyRefs = append(m.FirewallPolicyRefs, referenceModel)

			attr := models.MakeFirewallSequence()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_global_vrouter_config"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ApplicationPolicySetGlobalVrouterConfigRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.GlobalVrouterConfigRefs = append(m.GlobalVrouterConfigRefs, referenceModel)

		}
	}

	return m, nil
}

// ListApplicationPolicySet lists ApplicationPolicySet with list spec.
func ListApplicationPolicySet(tx *sql.Tx, spec *db.ListSpec) ([]*models.ApplicationPolicySet, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "application_policy_set"
	spec.Fields = ApplicationPolicySetFields
	spec.RefFields = ApplicationPolicySetRefFields
	result := models.MakeApplicationPolicySetSlice()
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
		m, err := scanApplicationPolicySet(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowApplicationPolicySet shows ApplicationPolicySet resource
func ShowApplicationPolicySet(tx *sql.Tx, uuid string) (*models.ApplicationPolicySet, error) {
	list, err := ListApplicationPolicySet(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateApplicationPolicySet updates a resource
func UpdateApplicationPolicySet(tx *sql.Tx, uuid string, model *models.ApplicationPolicySet) error {
	//TODO(nati) support update
	return nil
}

// DeleteApplicationPolicySet deletes a resource
func DeleteApplicationPolicySet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteApplicationPolicySetQuery)
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
