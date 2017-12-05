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

const insertSubnetQuery = "insert into `subnet` (`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`display_name`,`ip_prefix`,`ip_prefix_len`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateSubnetQuery = "update `subnet` set `key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`display_name` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?;"
const deleteSubnetQuery = "delete from `subnet` where uuid = ?"
const listSubnetQuery = "select `subnet`.`key_value_pair`,`subnet`.`owner_access`,`subnet`.`global_access`,`subnet`.`share`,`subnet`.`owner`,`subnet`.`uuid`,`subnet`.`fq_name`,`subnet`.`enable`,`subnet`.`description`,`subnet`.`created`,`subnet`.`creator`,`subnet`.`user_visible`,`subnet`.`last_modified`,`subnet`.`permissions_owner_access`,`subnet`.`other_access`,`subnet`.`group`,`subnet`.`group_access`,`subnet`.`permissions_owner`,`subnet`.`display_name`,`subnet`.`ip_prefix`,`subnet`.`ip_prefix_len` from `subnet`"
const showSubnetQuery = "select `subnet`.`key_value_pair`,`subnet`.`owner_access`,`subnet`.`global_access`,`subnet`.`share`,`subnet`.`owner`,`subnet`.`uuid`,`subnet`.`fq_name`,`subnet`.`enable`,`subnet`.`description`,`subnet`.`created`,`subnet`.`creator`,`subnet`.`user_visible`,`subnet`.`last_modified`,`subnet`.`permissions_owner_access`,`subnet`.`other_access`,`subnet`.`group`,`subnet`.`group_access`,`subnet`.`permissions_owner`,`subnet`.`display_name`,`subnet`.`ip_prefix`,`subnet`.`ip_prefix_len` from `subnet` where uuid = ?"

const insertSubnetVirtualMachineInterfaceQuery = "insert into `ref_subnet_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

func CreateSubnet(tx *sql.Tx, model *models.Subnet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertSubnetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		string(model.DisplayName),
		string(model.SubnetIPPrefix.IPPrefix),
		int(model.SubnetIPPrefix.IPPrefixLen))

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertSubnetVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanSubnet(rows *sql.Rows) (*models.Subnet, error) {
	m := models.MakeSubnet()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.DisplayName,
		&m.SubnetIPPrefix.IPPrefix,
		&m.SubnetIPPrefix.IPPrefixLen); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildSubnetWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListSubnet(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Subnet, error) {
	result := models.MakeSubnetSlice()
	whereQuery, values := buildSubnetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listSubnetQuery)
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
		m, _ := scanSubnet(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowSubnet(tx *sql.Tx, uuid string) (*models.Subnet, error) {
	rows, err := tx.Query(showSubnetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanSubnet(rows)
	}
	return nil, nil
}

func UpdateSubnet(tx *sql.Tx, uuid string, model *models.Subnet) error {
	return nil
}

func DeleteSubnet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteSubnetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
