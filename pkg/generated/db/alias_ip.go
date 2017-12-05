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

const insertAliasIPQuery = "insert into `alias_ip` (`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`alias_ip_address`,`alias_ip_address_family`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPQuery = "update `alias_ip` set `owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`alias_ip_address` = ?,`alias_ip_address_family` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteAliasIPQuery = "delete from `alias_ip` where uuid = ?"
const listAliasIPQuery = "select `alias_ip`.`owner`,`alias_ip`.`owner_access`,`alias_ip`.`other_access`,`alias_ip`.`group`,`alias_ip`.`group_access`,`alias_ip`.`enable`,`alias_ip`.`description`,`alias_ip`.`created`,`alias_ip`.`creator`,`alias_ip`.`user_visible`,`alias_ip`.`last_modified`,`alias_ip`.`display_name`,`alias_ip`.`key_value_pair`,`alias_ip`.`perms2_owner_access`,`alias_ip`.`global_access`,`alias_ip`.`share`,`alias_ip`.`perms2_owner`,`alias_ip`.`alias_ip_address`,`alias_ip`.`alias_ip_address_family`,`alias_ip`.`uuid`,`alias_ip`.`fq_name` from `alias_ip`"
const showAliasIPQuery = "select `alias_ip`.`owner`,`alias_ip`.`owner_access`,`alias_ip`.`other_access`,`alias_ip`.`group`,`alias_ip`.`group_access`,`alias_ip`.`enable`,`alias_ip`.`description`,`alias_ip`.`created`,`alias_ip`.`creator`,`alias_ip`.`user_visible`,`alias_ip`.`last_modified`,`alias_ip`.`display_name`,`alias_ip`.`key_value_pair`,`alias_ip`.`perms2_owner_access`,`alias_ip`.`global_access`,`alias_ip`.`share`,`alias_ip`.`perms2_owner`,`alias_ip`.`alias_ip_address`,`alias_ip`.`alias_ip_address_family`,`alias_ip`.`uuid`,`alias_ip`.`fq_name` from `alias_ip` where uuid = ?"

const insertAliasIPProjectQuery = "insert into `ref_alias_ip_project` (`from`, `to` ) values (?, ?);"

const insertAliasIPVirtualMachineInterfaceQuery = "insert into `ref_alias_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

func CreateAliasIP(tx *sql.Tx, model *models.AliasIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.AliasIPAddress),
		string(model.AliasIPAddressFamily),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	stmtProjectRef, err := tx.Prepare(insertAliasIPProjectQuery)
	if err != nil {
		return err
	}
	defer stmtProjectRef.Close()
	for _, ref := range model.ProjectRefs {
		_, err = stmtProjectRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertAliasIPVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanAliasIP(rows *sql.Rows) (*models.AliasIP, error) {
	m := models.MakeAliasIP()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.AliasIPAddress,
		&m.AliasIPAddressFamily,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildAliasIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["alias_ip_address"]; ok {
		results = append(results, "alias_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["alias_ip_address_family"]; ok {
		results = append(results, "alias_ip_address_family = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListAliasIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AliasIP, error) {
	result := models.MakeAliasIPSlice()
	whereQuery, values := buildAliasIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAliasIPQuery)
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
		m, _ := scanAliasIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAliasIP(tx *sql.Tx, uuid string) (*models.AliasIP, error) {
	rows, err := tx.Query(showAliasIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAliasIP(rows)
	}
	return nil, nil
}

func UpdateAliasIP(tx *sql.Tx, uuid string, model *models.AliasIP) error {
	return nil
}

func DeleteAliasIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAliasIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
