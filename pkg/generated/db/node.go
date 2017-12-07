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

const insertNodeQuery = "insert into `node` (`mac_address`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`private_machine_properties`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`display_name`,`key_value_pair`,`hostname`,`password`,`ssh_key`,`gcp_machine_type`,`private_power_management_password`,`fq_name`,`username`,`aws_ami`,`gcp_image`,`private_machine_state`,`private_power_management_ip`,`private_power_management_username`,`ip_address`,`type`,`aws_instance_type`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNodeQuery = "update `node` set `mac_address` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`private_machine_properties` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?,`hostname` = ?,`password` = ?,`ssh_key` = ?,`gcp_machine_type` = ?,`private_power_management_password` = ?,`fq_name` = ?,`username` = ?,`aws_ami` = ?,`gcp_image` = ?,`private_machine_state` = ?,`private_power_management_ip` = ?,`private_power_management_username` = ?,`ip_address` = ?,`type` = ?,`aws_instance_type` = ?;"
const deleteNodeQuery = "delete from `node` where uuid = ?"

// NodeFields is db columns for Node
var NodeFields = []string{
	"mac_address",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"private_machine_properties",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"uuid",
	"display_name",
	"key_value_pair",
	"hostname",
	"password",
	"ssh_key",
	"gcp_machine_type",
	"private_power_management_password",
	"fq_name",
	"username",
	"aws_ami",
	"gcp_image",
	"private_machine_state",
	"private_power_management_ip",
	"private_power_management_username",
	"ip_address",
	"type",
	"aws_instance_type",
}

// NodeRefFields is db reference fields for Node
var NodeRefFields = map[string][]string{}

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
	_, err = stmt.Exec(string(model.MacAddress),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.PrivateMachineProperties),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Hostname),
		string(model.Password),
		string(model.SSHKey),
		string(model.GCPMachineType),
		string(model.PrivatePowerManagementPassword),
		utils.MustJSON(model.FQName),
		string(model.Username),
		string(model.AwsAmi),
		string(model.GCPImage),
		string(model.PrivateMachineState),
		string(model.PrivatePowerManagementIP),
		string(model.PrivatePowerManagementUsername),
		string(model.IPAddress),
		string(model.Type),
		string(model.AwsInstanceType))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanNode(values map[string]interface{}) (*models.Node, error) {
	m := models.MakeNode()

	if value, ok := values["mac_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.MacAddress = castedValue

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

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["private_machine_properties"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateMachineProperties = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

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

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["hostname"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Hostname = castedValue

	}

	if value, ok := values["password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Password = castedValue

	}

	if value, ok := values["ssh_key"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.SSHKey = castedValue

	}

	if value, ok := values["gcp_machine_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.GCPMachineType = castedValue

	}

	if value, ok := values["private_power_management_password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivatePowerManagementPassword = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["username"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Username = castedValue

	}

	if value, ok := values["aws_ami"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsAmi = castedValue

	}

	if value, ok := values["gcp_image"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.GCPImage = castedValue

	}

	if value, ok := values["private_machine_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateMachineState = castedValue

	}

	if value, ok := values["private_power_management_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivatePowerManagementIP = castedValue

	}

	if value, ok := values["private_power_management_username"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivatePowerManagementUsername = castedValue

	}

	if value, ok := values["ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IPAddress = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Type = castedValue

	}

	if value, ok := values["aws_instance_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsInstanceType = castedValue

	}

	return m, nil
}

// ListNode lists Node with list spec.
func ListNode(tx *sql.Tx, spec *db.ListSpec) ([]*models.Node, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "node"
	spec.Fields = NodeFields
	spec.RefFields = NodeRefFields
	result := models.MakeNodeSlice()
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
		m, err := scanNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowNode shows Node resource
func ShowNode(tx *sql.Tx, uuid string) (*models.Node, error) {
	list, err := ListNode(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateNode updates a resource
func UpdateNode(tx *sql.Tx, uuid string, model *models.Node) error {
	//TODO(nati) support update
	return nil
}

// DeleteNode deletes a resource
func DeleteNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNodeQuery)
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
