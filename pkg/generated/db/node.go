package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertNodeQuery = "insert into `node` (`uuid`,`username`,`type`,`ssh_key`,`private_power_management_username`,`private_power_management_password`,`private_power_management_ip`,`private_machine_state`,`private_machine_properties`,`share`,`owner_access`,`owner`,`global_access`,`password`,`parent_uuid`,`parent_type`,`mac_address`,`ip_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`hostname`,`gcp_machine_type`,`gcp_image`,`fq_name`,`display_name`,`aws_instance_type`,`aws_ami`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteNodeQuery = "delete from `node` where uuid = ?"

// NodeFields is db columns for Node
var NodeFields = []string{
	"uuid",
	"username",
	"type",
	"ssh_key",
	"private_power_management_username",
	"private_power_management_password",
	"private_power_management_ip",
	"private_machine_state",
	"private_machine_properties",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"password",
	"parent_uuid",
	"parent_type",
	"mac_address",
	"ip_address",
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
	"hostname",
	"gcp_machine_type",
	"gcp_image",
	"fq_name",
	"display_name",
	"aws_instance_type",
	"aws_ami",
	"key_value_pair",
}

// NodeRefFields is db reference fields for Node
var NodeRefFields = map[string][]string{}

// NodeBackRefFields is db back reference fields for Node
var NodeBackRefFields = map[string][]string{}

// NodeParentTypes is possible parents for Node
var NodeParents = []string{}

// CreateNode inserts Node to DB
func CreateNode(tx *sql.Tx, model *models.Node) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNodeQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		string(model.Username),
		string(model.Type),
		string(model.SSHKey),
		string(model.PrivatePowerManagementUsername),
		string(model.PrivatePowerManagementPassword),
		string(model.PrivatePowerManagementIP),
		string(model.PrivateMachineState),
		string(model.PrivateMachineProperties),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.Password),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.MacAddress),
		string(model.IPAddress),
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
		string(model.Hostname),
		string(model.GCPMachineType),
		string(model.GCPImage),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.AwsInstanceType),
		string(model.AwsAmi),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "node", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanNode(values map[string]interface{}) (*models.Node, error) {
	m := models.MakeNode()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["username"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Username = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Type = castedValue

	}

	if value, ok := values["ssh_key"]; ok {

		castedValue := common.InterfaceToString(value)

		m.SSHKey = castedValue

	}

	if value, ok := values["private_power_management_username"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivatePowerManagementUsername = castedValue

	}

	if value, ok := values["private_power_management_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivatePowerManagementPassword = castedValue

	}

	if value, ok := values["private_power_management_ip"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivatePowerManagementIP = castedValue

	}

	if value, ok := values["private_machine_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateMachineState = castedValue

	}

	if value, ok := values["private_machine_properties"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateMachineProperties = castedValue

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

	if value, ok := values["password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Password = castedValue

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["mac_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MacAddress = castedValue

	}

	if value, ok := values["ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IPAddress = castedValue

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

	if value, ok := values["hostname"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Hostname = castedValue

	}

	if value, ok := values["gcp_machine_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.GCPMachineType = castedValue

	}

	if value, ok := values["gcp_image"]; ok {

		castedValue := common.InterfaceToString(value)

		m.GCPImage = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["aws_instance_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsInstanceType = castedValue

	}

	if value, ok := values["aws_ami"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsAmi = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListNode lists Node with list spec.
func ListNode(tx *sql.Tx, spec *common.ListSpec) ([]*models.Node, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "node"
	spec.Fields = NodeFields
	spec.RefFields = NodeRefFields
	spec.BackRefFields = NodeBackRefFields
	result := models.MakeNodeSlice()

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
		m, err := scanNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateNode updates a resource
func UpdateNode(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateNodeQuery = "update `node` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateNodeQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Username", "."); ok {
		updateNodeQuery += "`username` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Type", "."); ok {
		updateNodeQuery += "`type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".SSHKey", "."); ok {
		updateNodeQuery += "`ssh_key` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PrivatePowerManagementUsername", "."); ok {
		updateNodeQuery += "`private_power_management_username` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PrivatePowerManagementPassword", "."); ok {
		updateNodeQuery += "`private_power_management_password` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PrivatePowerManagementIP", "."); ok {
		updateNodeQuery += "`private_power_management_ip` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PrivateMachineState", "."); ok {
		updateNodeQuery += "`private_machine_state` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".PrivateMachineProperties", "."); ok {
		updateNodeQuery += "`private_machine_properties` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateNodeQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateNodeQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateNodeQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateNodeQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Password", "."); ok {
		updateNodeQuery += "`password` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateNodeQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateNodeQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".MacAddress", "."); ok {
		updateNodeQuery += "`mac_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IPAddress", "."); ok {
		updateNodeQuery += "`ip_address` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateNodeQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateNodeQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateNodeQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateNodeQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateNodeQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateNodeQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateNodeQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateNodeQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateNodeQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateNodeQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateNodeQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Hostname", "."); ok {
		updateNodeQuery += "`hostname` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".GCPMachineType", "."); ok {
		updateNodeQuery += "`gcp_machine_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".GCPImage", "."); ok {
		updateNodeQuery += "`gcp_image` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateNodeQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateNodeQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AwsInstanceType", "."); ok {
		updateNodeQuery += "`aws_instance_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".AwsAmi", "."); ok {
		updateNodeQuery += "`aws_ami` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateNodeQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateNodeQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateNodeQuery += ","
	}

	updateNodeQuery =
		updateNodeQuery[:len(updateNodeQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateNodeQuery,
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

// DeleteNode deletes a resource
func DeleteNode(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteNodeQuery
	selectQuery := "select count(uuid) from node where uuid = ?"
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
