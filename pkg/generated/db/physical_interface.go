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

const insertPhysicalInterfaceQuery = "insert into `physical_interface` (`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`enable`,`description`,`created`,`display_name`,`ethernet_segment_identifier`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePhysicalInterfaceQuery = "update `physical_interface` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`ethernet_segment_identifier` = ?,`key_value_pair` = ?;"
const deletePhysicalInterfaceQuery = "delete from `physical_interface` where uuid = ?"
const listPhysicalInterfaceQuery = "select `physical_interface`.`owner`,`physical_interface`.`owner_access`,`physical_interface`.`global_access`,`physical_interface`.`share`,`physical_interface`.`uuid`,`physical_interface`.`fq_name`,`physical_interface`.`creator`,`physical_interface`.`user_visible`,`physical_interface`.`last_modified`,`physical_interface`.`permissions_owner_access`,`physical_interface`.`other_access`,`physical_interface`.`group`,`physical_interface`.`group_access`,`physical_interface`.`permissions_owner`,`physical_interface`.`enable`,`physical_interface`.`description`,`physical_interface`.`created`,`physical_interface`.`display_name`,`physical_interface`.`ethernet_segment_identifier`,`physical_interface`.`key_value_pair` from `physical_interface`"
const showPhysicalInterfaceQuery = "select `physical_interface`.`owner`,`physical_interface`.`owner_access`,`physical_interface`.`global_access`,`physical_interface`.`share`,`physical_interface`.`uuid`,`physical_interface`.`fq_name`,`physical_interface`.`creator`,`physical_interface`.`user_visible`,`physical_interface`.`last_modified`,`physical_interface`.`permissions_owner_access`,`physical_interface`.`other_access`,`physical_interface`.`group`,`physical_interface`.`group_access`,`physical_interface`.`permissions_owner`,`physical_interface`.`enable`,`physical_interface`.`description`,`physical_interface`.`created`,`physical_interface`.`display_name`,`physical_interface`.`ethernet_segment_identifier`,`physical_interface`.`key_value_pair` from `physical_interface` where uuid = ?"

const insertPhysicalInterfacePhysicalInterfaceQuery = "insert into `ref_physical_interface_physical_interface` (`from`, `to` ) values (?, ?);"

func CreatePhysicalInterface(tx *sql.Tx, model *models.PhysicalInterface) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPhysicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		string(model.EthernetSegmentIdentifier),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertPhysicalInterfacePhysicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {
		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanPhysicalInterface(rows *sql.Rows) (*models.PhysicalInterface, error) {
	m := models.MakePhysicalInterface()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.DisplayName,
		&m.EthernetSegmentIdentifier,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildPhysicalInterfaceWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["ethernet_segment_identifier"]; ok {
		results = append(results, "ethernet_segment_identifier = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListPhysicalInterface(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PhysicalInterface, error) {
	result := models.MakePhysicalInterfaceSlice()
	whereQuery, values := buildPhysicalInterfaceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPhysicalInterfaceQuery)
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
		m, _ := scanPhysicalInterface(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPhysicalInterface(tx *sql.Tx, uuid string) (*models.PhysicalInterface, error) {
	rows, err := tx.Query(showPhysicalInterfaceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPhysicalInterface(rows)
	}
	return nil, nil
}

func UpdatePhysicalInterface(tx *sql.Tx, uuid string, model *models.PhysicalInterface) error {
	return nil
}

func DeletePhysicalInterface(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePhysicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
