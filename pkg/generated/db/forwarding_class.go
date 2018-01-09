package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertForwardingClassQuery = "insert into `forwarding_class` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`forwarding_class_vlan_priority`,`forwarding_class_mpls_exp`,`forwarding_class_id`,`forwarding_class_dscp`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateForwardingClassQuery = "update `forwarding_class` set `uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`forwarding_class_vlan_priority` = ?,`forwarding_class_mpls_exp` = ?,`forwarding_class_id` = ?,`forwarding_class_dscp` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteForwardingClassQuery = "delete from `forwarding_class` where uuid = ?"

// ForwardingClassFields is db columns for ForwardingClass
var ForwardingClassFields = []string{
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
	"forwarding_class_vlan_priority",
	"forwarding_class_mpls_exp",
	"forwarding_class_id",
	"forwarding_class_dscp",
	"display_name",
	"key_value_pair",
}

// ForwardingClassRefFields is db reference fields for ForwardingClass
var ForwardingClassRefFields = map[string][]string{

	"qos_queue": {
	// <common.Schema Value>

	},
}

// ForwardingClassBackRefFields is db back reference fields for ForwardingClass
var ForwardingClassBackRefFields = map[string][]string{}

const insertForwardingClassQosQueueQuery = "insert into `ref_forwarding_class_qos_queue` (`from`, `to` ) values (?, ?);"

// CreateForwardingClass inserts ForwardingClass to DB
func CreateForwardingClass(tx *sql.Tx, model *models.ForwardingClass) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertForwardingClassQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertForwardingClassQuery,
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
		int(model.ForwardingClassVlanPriority),
		int(model.ForwardingClassMPLSExp),
		int(model.ForwardingClassID),
		int(model.ForwardingClassDSCP),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtQosQueueRef, err := tx.Prepare(insertForwardingClassQosQueueQuery)
	if err != nil {
		return errors.Wrap(err, "preparing QosQueueRefs create statement failed")
	}
	defer stmtQosQueueRef.Close()
	for _, ref := range model.QosQueueRefs {

		_, err = stmtQosQueueRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "QosQueueRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanForwardingClass(values map[string]interface{}) (*models.ForwardingClass, error) {
	m := models.MakeForwardingClass()

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

	if value, ok := values["forwarding_class_vlan_priority"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ForwardingClassVlanPriority = models.VlanPriorityType(castedValue)

	}

	if value, ok := values["forwarding_class_mpls_exp"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ForwardingClassMPLSExp = models.MplsExpType(castedValue)

	}

	if value, ok := values["forwarding_class_id"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ForwardingClassID = models.ForwardingClassId(castedValue)

	}

	if value, ok := values["forwarding_class_dscp"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ForwardingClassDSCP = models.DscpValueType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_qos_queue"]; ok {
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
			referenceModel := &models.ForwardingClassQosQueueRef{}
			referenceModel.UUID = uuid
			m.QosQueueRefs = append(m.QosQueueRefs, referenceModel)

		}
	}

	return m, nil
}

// ListForwardingClass lists ForwardingClass with list spec.
func ListForwardingClass(tx *sql.Tx, spec *common.ListSpec) ([]*models.ForwardingClass, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "forwarding_class"
	if spec.Fields == nil {
		spec.Fields = ForwardingClassFields
	}
	spec.RefFields = ForwardingClassRefFields
	spec.BackRefFields = ForwardingClassBackRefFields
	result := models.MakeForwardingClassSlice()
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
		m, err := scanForwardingClass(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateForwardingClass updates a resource
func UpdateForwardingClass(tx *sql.Tx, uuid string, model *models.ForwardingClass) error {
	//TODO(nati) support update
	return nil
}

// DeleteForwardingClass deletes a resource
func DeleteForwardingClass(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteForwardingClassQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
