package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertBaremetalPortQuery = "insert into `baremetal_port` (`uuid`,`switch_info`,`switch_id`,`pxe_enabled`,`port_id`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`node`,`mac_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBaremetalPortQuery = "delete from `baremetal_port` where uuid = ?"

// BaremetalPortFields is db columns for BaremetalPort
var BaremetalPortFields = []string{
	"uuid",
	"switch_info",
	"switch_id",
	"pxe_enabled",
	"port_id",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"node",
	"mac_address",
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

// BaremetalPortRefFields is db reference fields for BaremetalPort
var BaremetalPortRefFields = map[string][]string{}

// BaremetalPortBackRefFields is db back reference fields for BaremetalPort
var BaremetalPortBackRefFields = map[string][]string{}

// BaremetalPortParentTypes is possible parents for BaremetalPort
var BaremetalPortParents = []string{}

// CreateBaremetalPort inserts BaremetalPort to DB
func CreateBaremetalPort(tx *sql.Tx, model *models.BaremetalPort) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBaremetalPortQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBaremetalPortQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		string(model.SwitchInfo),
		string(model.SwitchID),
		bool(model.PxeEnabled),
		string(model.PortID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.Node),
		string(model.MacAddress),
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
		Type:   "baremetal_port",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "baremetal_port", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBaremetalPort(values map[string]interface{}) (*models.BaremetalPort, error) {
	m := models.MakeBaremetalPort()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["switch_info"]; ok {

		castedValue := common.InterfaceToString(value)

		m.SwitchInfo = castedValue

	}

	if value, ok := values["switch_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.SwitchID = castedValue

	}

	if value, ok := values["pxe_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.PxeEnabled = castedValue

	}

	if value, ok := values["port_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PortID = castedValue

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

	if value, ok := values["node"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Node = castedValue

	}

	if value, ok := values["mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacAddress = castedValue

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

// ListBaremetalPort lists BaremetalPort with list spec.
func ListBaremetalPort(tx *sql.Tx, spec *common.ListSpec) ([]*models.BaremetalPort, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "baremetal_port"
	spec.Fields = BaremetalPortFields
	spec.RefFields = BaremetalPortRefFields
	spec.BackRefFields = BaremetalPortBackRefFields
	result := models.MakeBaremetalPortSlice()

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
		m, err := scanBaremetalPort(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateBaremetalPort updates a resource
func UpdateBaremetalPort(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateBaremetalPortQuery = "update `baremetal_port` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateBaremetalPortQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SwitchInfo", "."); ok {
		updateBaremetalPortQuery += "`switch_info` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SwitchID", "."); ok {
		updateBaremetalPortQuery += "`switch_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PxeEnabled", "."); ok {
		updateBaremetalPortQuery += "`pxe_enabled` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PortID", "."); ok {
		updateBaremetalPortQuery += "`port_id` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateBaremetalPortQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateBaremetalPortQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateBaremetalPortQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateBaremetalPortQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateBaremetalPortQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateBaremetalPortQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Node", "."); ok {
		updateBaremetalPortQuery += "`node` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacAddress", "."); ok {
		updateBaremetalPortQuery += "`mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateBaremetalPortQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateBaremetalPortQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateBaremetalPortQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateBaremetalPortQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateBaremetalPortQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateBaremetalPortQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateBaremetalPortQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateBaremetalPortQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateBaremetalPortQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateBaremetalPortQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateBaremetalPortQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateBaremetalPortQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateBaremetalPortQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateBaremetalPortQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateBaremetalPortQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateBaremetalPortQuery += ","
	}

	updateBaremetalPortQuery =
		updateBaremetalPortQuery[:len(updateBaremetalPortQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateBaremetalPortQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateBaremetalPortQuery,
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

// DeleteBaremetalPort deletes a resource
func DeleteBaremetalPort(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteBaremetalPortQuery
	selectQuery := "select count(uuid) from baremetal_port where uuid = ?"
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
