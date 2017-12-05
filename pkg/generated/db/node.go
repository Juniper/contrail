package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertNodeQuery = "insert into `node` (`type`,`aws_instance_type`,`private_power_management_password`,`private_power_management_username`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`ip_address`,`password`,`username`,`gcp_machine_type`,`private_power_management_ip`,`fq_name`,`mac_address`,`ssh_key`,`aws_ami`,`gcp_image`,`private_machine_state`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`hostname`,`private_machine_properties`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNodeQuery = "update `node` set `type` = ?,`aws_instance_type` = ?,`private_power_management_password` = ?,`private_power_management_username` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`ip_address` = ?,`password` = ?,`username` = ?,`gcp_machine_type` = ?,`private_power_management_ip` = ?,`fq_name` = ?,`mac_address` = ?,`ssh_key` = ?,`aws_ami` = ?,`gcp_image` = ?,`private_machine_state` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`hostname` = ?,`private_machine_properties` = ?;"
const deleteNodeQuery = "delete from `node` where uuid = ?"
const listNodeQuery = "select `node`.`type`,`node`.`aws_instance_type`,`node`.`private_power_management_password`,`node`.`private_power_management_username`,`node`.`enable`,`node`.`description`,`node`.`created`,`node`.`creator`,`node`.`user_visible`,`node`.`last_modified`,`node`.`owner`,`node`.`owner_access`,`node`.`other_access`,`node`.`group`,`node`.`group_access`,`node`.`ip_address`,`node`.`password`,`node`.`username`,`node`.`gcp_machine_type`,`node`.`private_power_management_ip`,`node`.`fq_name`,`node`.`mac_address`,`node`.`ssh_key`,`node`.`aws_ami`,`node`.`gcp_image`,`node`.`private_machine_state`,`node`.`display_name`,`node`.`key_value_pair`,`node`.`perms2_owner_access`,`node`.`global_access`,`node`.`share`,`node`.`perms2_owner`,`node`.`uuid`,`node`.`hostname`,`node`.`private_machine_properties` from `node`"
const showNodeQuery = "select `node`.`type`,`node`.`aws_instance_type`,`node`.`private_power_management_password`,`node`.`private_power_management_username`,`node`.`enable`,`node`.`description`,`node`.`created`,`node`.`creator`,`node`.`user_visible`,`node`.`last_modified`,`node`.`owner`,`node`.`owner_access`,`node`.`other_access`,`node`.`group`,`node`.`group_access`,`node`.`ip_address`,`node`.`password`,`node`.`username`,`node`.`gcp_machine_type`,`node`.`private_power_management_ip`,`node`.`fq_name`,`node`.`mac_address`,`node`.`ssh_key`,`node`.`aws_ami`,`node`.`gcp_image`,`node`.`private_machine_state`,`node`.`display_name`,`node`.`key_value_pair`,`node`.`perms2_owner_access`,`node`.`global_access`,`node`.`share`,`node`.`perms2_owner`,`node`.`uuid`,`node`.`hostname`,`node`.`private_machine_properties` from `node` where uuid = ?"

func CreateNode(tx *sql.Tx, model *models.Node) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Type),
		string(model.AwsInstanceType),
		string(model.PrivatePowerManagementPassword),
		string(model.PrivatePowerManagementUsername),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IPAddress),
		string(model.Password),
		string(model.Username),
		string(model.GCPMachineType),
		string(model.PrivatePowerManagementIP),
		utils.MustJSON(model.FQName),
		string(model.MacAddress),
		string(model.SSHKey),
		string(model.AwsAmi),
		string(model.GCPImage),
		string(model.PrivateMachineState),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		string(model.Hostname),
		string(model.PrivateMachineProperties))

	return err
}

func scanNode(rows *sql.Rows) (*models.Node, error) {
	m := models.MakeNode()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.Type,
		&m.AwsInstanceType,
		&m.PrivatePowerManagementPassword,
		&m.PrivatePowerManagementUsername,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IPAddress,
		&m.Password,
		&m.Username,
		&m.GCPMachineType,
		&m.PrivatePowerManagementIP,
		&jsonFQName,
		&m.MacAddress,
		&m.SSHKey,
		&m.AwsAmi,
		&m.GCPImage,
		&m.PrivateMachineState,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&m.Hostname,
		&m.PrivateMachineProperties); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildNodeWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["type"]; ok {
		results = append(results, "type = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_instance_type"]; ok {
		results = append(results, "aws_instance_type = ?")
		values = append(values, value)
	}

	if value, ok := where["private_power_management_password"]; ok {
		results = append(results, "private_power_management_password = ?")
		values = append(values, value)
	}

	if value, ok := where["private_power_management_username"]; ok {
		results = append(results, "private_power_management_username = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_address"]; ok {
		results = append(results, "ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["password"]; ok {
		results = append(results, "password = ?")
		values = append(values, value)
	}

	if value, ok := where["username"]; ok {
		results = append(results, "username = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_machine_type"]; ok {
		results = append(results, "gcp_machine_type = ?")
		values = append(values, value)
	}

	if value, ok := where["private_power_management_ip"]; ok {
		results = append(results, "private_power_management_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["mac_address"]; ok {
		results = append(results, "mac_address = ?")
		values = append(values, value)
	}

	if value, ok := where["ssh_key"]; ok {
		results = append(results, "ssh_key = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_ami"]; ok {
		results = append(results, "aws_ami = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_image"]; ok {
		results = append(results, "gcp_image = ?")
		values = append(values, value)
	}

	if value, ok := where["private_machine_state"]; ok {
		results = append(results, "private_machine_state = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["hostname"]; ok {
		results = append(results, "hostname = ?")
		values = append(values, value)
	}

	if value, ok := where["private_machine_properties"]; ok {
		results = append(results, "private_machine_properties = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListNode(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Node, error) {
	result := models.MakeNodeSlice()
	whereQuery, values := buildNodeWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNodeQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNode(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNode(tx *sql.Tx, uuid string) (*models.Node, error) {
	rows, err := tx.Query(showNodeQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNode(rows)
	}
	return nil, nil
}

func UpdateNode(tx *sql.Tx, uuid string, model *models.Node) error {
	return nil
}

func DeleteNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNodeQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
