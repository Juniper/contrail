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

const insertLogicalInterfaceQuery = "insert into `logical_interface` (`owner`,`owner_access`,`global_access`,`share`,`logical_interface_vlan_tag`,`logical_interface_type`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalInterfaceQuery = "update `logical_interface` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`logical_interface_vlan_tag` = ?,`logical_interface_type` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLogicalInterfaceQuery = "delete from `logical_interface` where uuid = ?"
const listLogicalInterfaceQuery = "select `logical_interface`.`owner`,`logical_interface`.`owner_access`,`logical_interface`.`global_access`,`logical_interface`.`share`,`logical_interface`.`logical_interface_vlan_tag`,`logical_interface`.`logical_interface_type`,`logical_interface`.`uuid`,`logical_interface`.`fq_name`,`logical_interface`.`enable`,`logical_interface`.`description`,`logical_interface`.`created`,`logical_interface`.`creator`,`logical_interface`.`user_visible`,`logical_interface`.`last_modified`,`logical_interface`.`permissions_owner`,`logical_interface`.`permissions_owner_access`,`logical_interface`.`other_access`,`logical_interface`.`group`,`logical_interface`.`group_access`,`logical_interface`.`display_name`,`logical_interface`.`key_value_pair` from `logical_interface`"
const showLogicalInterfaceQuery = "select `logical_interface`.`owner`,`logical_interface`.`owner_access`,`logical_interface`.`global_access`,`logical_interface`.`share`,`logical_interface`.`logical_interface_vlan_tag`,`logical_interface`.`logical_interface_type`,`logical_interface`.`uuid`,`logical_interface`.`fq_name`,`logical_interface`.`enable`,`logical_interface`.`description`,`logical_interface`.`created`,`logical_interface`.`creator`,`logical_interface`.`user_visible`,`logical_interface`.`last_modified`,`logical_interface`.`permissions_owner`,`logical_interface`.`permissions_owner_access`,`logical_interface`.`other_access`,`logical_interface`.`group`,`logical_interface`.`group_access`,`logical_interface`.`display_name`,`logical_interface`.`key_value_pair` from `logical_interface` where uuid = ?"

const insertLogicalInterfaceVirtualMachineInterfaceQuery = "insert into `ref_logical_interface_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

func CreateLogicalInterface(tx *sql.Tx, model *models.LogicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		int(model.LogicalInterfaceVlanTag),
		string(model.LogicalInterfaceType),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLogicalInterfaceVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanLogicalInterface(rows *sql.Rows) (*models.LogicalInterface, error) {
	m := models.MakeLogicalInterface()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.LogicalInterfaceVlanTag,
		&m.LogicalInterfaceType,
		&m.UUID,
		&jsonFQName,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildLogicalInterfaceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["logical_interface_type"]; ok {
		results = append(results, "logical_interface_type = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLogicalInterface(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LogicalInterface, error) {
	result := models.MakeLogicalInterfaceSlice()
	whereQuery, values := buildLogicalInterfaceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLogicalInterfaceQuery)
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
		m, _ := scanLogicalInterface(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLogicalInterface(tx *sql.Tx, uuid string) (*models.LogicalInterface, error) {
	rows, err := tx.Query(showLogicalInterfaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLogicalInterface(rows)
	}
	return nil, nil
}

func UpdateLogicalInterface(tx *sql.Tx, uuid string, model *models.LogicalInterface) error {
	return nil
}

func DeleteLogicalInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLogicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
