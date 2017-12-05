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

const insertVirtualMachineQuery = "insert into `virtual_machine` (`key_value_pair`,`global_access`,`share`,`owner`,`owner_access`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualMachineQuery = "update `virtual_machine` set `key_value_pair` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?;"
const deleteVirtualMachineQuery = "delete from `virtual_machine` where uuid = ?"
const listVirtualMachineQuery = "select `virtual_machine`.`key_value_pair`,`virtual_machine`.`global_access`,`virtual_machine`.`share`,`virtual_machine`.`owner`,`virtual_machine`.`owner_access`,`virtual_machine`.`uuid`,`virtual_machine`.`fq_name`,`virtual_machine`.`creator`,`virtual_machine`.`user_visible`,`virtual_machine`.`last_modified`,`virtual_machine`.`other_access`,`virtual_machine`.`group`,`virtual_machine`.`group_access`,`virtual_machine`.`permissions_owner`,`virtual_machine`.`permissions_owner_access`,`virtual_machine`.`enable`,`virtual_machine`.`description`,`virtual_machine`.`created`,`virtual_machine`.`display_name` from `virtual_machine`"
const showVirtualMachineQuery = "select `virtual_machine`.`key_value_pair`,`virtual_machine`.`global_access`,`virtual_machine`.`share`,`virtual_machine`.`owner`,`virtual_machine`.`owner_access`,`virtual_machine`.`uuid`,`virtual_machine`.`fq_name`,`virtual_machine`.`creator`,`virtual_machine`.`user_visible`,`virtual_machine`.`last_modified`,`virtual_machine`.`other_access`,`virtual_machine`.`group`,`virtual_machine`.`group_access`,`virtual_machine`.`permissions_owner`,`virtual_machine`.`permissions_owner_access`,`virtual_machine`.`enable`,`virtual_machine`.`description`,`virtual_machine`.`created`,`virtual_machine`.`display_name` from `virtual_machine` where uuid = ?"

const insertVirtualMachineServiceInstanceQuery = "insert into `ref_virtual_machine_service_instance` (`from`, `to` ) values (?, ?);"

func CreateVirtualMachine(tx *sql.Tx, model *models.VirtualMachine) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualMachineQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName))

	stmtServiceInstanceRef, err := tx.Prepare(insertVirtualMachineServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanVirtualMachine(rows *sql.Rows) (*models.VirtualMachine, error) {
	m := models.MakeVirtualMachine()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildVirtualMachineWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualMachine(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualMachine, error) {
	result := models.MakeVirtualMachineSlice()
	whereQuery, values := buildVirtualMachineWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualMachineQuery)
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
		m, _ := scanVirtualMachine(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualMachine(tx *sql.Tx, uuid string) (*models.VirtualMachine, error) {
	rows, err := tx.Query(showVirtualMachineQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualMachine(rows)
	}
	return nil, nil
}

func UpdateVirtualMachine(tx *sql.Tx, uuid string, model *models.VirtualMachine) error {
	return nil
}

func DeleteVirtualMachine(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualMachineQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
