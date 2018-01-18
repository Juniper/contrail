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

// ForwardingClassParentTypes is possible parents for ForwardingClass
var ForwardingClassParents = []string{

	"global_qos_config",
}

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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "forwarding_class",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "forwarding_class", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
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
	spec.Fields = ForwardingClassFields
	spec.RefFields = ForwardingClassRefFields
	spec.BackRefFields = ForwardingClassBackRefFields
	result := models.MakeForwardingClassSlice()

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
		m, err := scanForwardingClass(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateForwardingClass updates a resource
func UpdateForwardingClass(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateForwardingClassQuery = "update `forwarding_class` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateForwardingClassQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateForwardingClassQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateForwardingClassQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateForwardingClassQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateForwardingClassQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateForwardingClassQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateForwardingClassQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateForwardingClassQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateForwardingClassQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateForwardingClassQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateForwardingClassQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateForwardingClassQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateForwardingClassQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateForwardingClassQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateForwardingClassQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateForwardingClassQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateForwardingClassQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateForwardingClassQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateForwardingClassQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ForwardingClassVlanPriority", "."); ok {
		updateForwardingClassQuery += "`forwarding_class_vlan_priority` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ForwardingClassMPLSExp", "."); ok {
		updateForwardingClassQuery += "`forwarding_class_mpls_exp` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ForwardingClassID", "."); ok {
		updateForwardingClassQuery += "`forwarding_class_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ForwardingClassDSCP", "."); ok {
		updateForwardingClassQuery += "`forwarding_class_dscp` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateForwardingClassQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateForwardingClassQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateForwardingClassQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateForwardingClassQuery += ","
	}

	updateForwardingClassQuery =
		updateForwardingClassQuery[:len(updateForwardingClassQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateForwardingClassQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateForwardingClassQuery,
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

// DeleteForwardingClass deletes a resource
func DeleteForwardingClass(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteForwardingClassQuery
	selectQuery := "select count(uuid) from forwarding_class where uuid = ?"
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
