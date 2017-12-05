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

const insertFloatingIPQuery = "insert into `floating_ip` (`floating_ip_port_mappings`,`floating_ip_is_virtual_ip`,`floating_ip_fixed_ip_address`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`uuid`,`floating_ip_address_family`,`floating_ip_address`,`floating_ip_port_mappings_enable`,`floating_ip_traffic_direction`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateFloatingIPQuery = "update `floating_ip` set `floating_ip_port_mappings` = ?,`floating_ip_is_virtual_ip` = ?,`floating_ip_fixed_ip_address` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`uuid` = ?,`floating_ip_address_family` = ?,`floating_ip_address` = ?,`floating_ip_port_mappings_enable` = ?,`floating_ip_traffic_direction` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`fq_name` = ?;"
const deleteFloatingIPQuery = "delete from `floating_ip` where uuid = ?"
const listFloatingIPQuery = "select `floating_ip`.`floating_ip_port_mappings`,`floating_ip`.`floating_ip_is_virtual_ip`,`floating_ip`.`floating_ip_fixed_ip_address`,`floating_ip`.`enable`,`floating_ip`.`description`,`floating_ip`.`created`,`floating_ip`.`creator`,`floating_ip`.`user_visible`,`floating_ip`.`last_modified`,`floating_ip`.`owner_access`,`floating_ip`.`other_access`,`floating_ip`.`group`,`floating_ip`.`group_access`,`floating_ip`.`owner`,`floating_ip`.`uuid`,`floating_ip`.`floating_ip_address_family`,`floating_ip`.`floating_ip_address`,`floating_ip`.`floating_ip_port_mappings_enable`,`floating_ip`.`floating_ip_traffic_direction`,`floating_ip`.`display_name`,`floating_ip`.`key_value_pair`,`floating_ip`.`perms2_owner_access`,`floating_ip`.`global_access`,`floating_ip`.`share`,`floating_ip`.`perms2_owner`,`floating_ip`.`fq_name` from `floating_ip`"
const showFloatingIPQuery = "select `floating_ip`.`floating_ip_port_mappings`,`floating_ip`.`floating_ip_is_virtual_ip`,`floating_ip`.`floating_ip_fixed_ip_address`,`floating_ip`.`enable`,`floating_ip`.`description`,`floating_ip`.`created`,`floating_ip`.`creator`,`floating_ip`.`user_visible`,`floating_ip`.`last_modified`,`floating_ip`.`owner_access`,`floating_ip`.`other_access`,`floating_ip`.`group`,`floating_ip`.`group_access`,`floating_ip`.`owner`,`floating_ip`.`uuid`,`floating_ip`.`floating_ip_address_family`,`floating_ip`.`floating_ip_address`,`floating_ip`.`floating_ip_port_mappings_enable`,`floating_ip`.`floating_ip_traffic_direction`,`floating_ip`.`display_name`,`floating_ip`.`key_value_pair`,`floating_ip`.`perms2_owner_access`,`floating_ip`.`global_access`,`floating_ip`.`share`,`floating_ip`.`perms2_owner`,`floating_ip`.`fq_name` from `floating_ip` where uuid = ?"

const insertFloatingIPProjectQuery = "insert into `ref_floating_ip_project` (`from`, `to` ) values (?, ?);"

const insertFloatingIPVirtualMachineInterfaceQuery = "insert into `ref_floating_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

func CreateFloatingIP(tx *sql.Tx, model *models.FloatingIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FloatingIPPortMappings),
		bool(model.FloatingIPIsVirtualIP),
		string(model.FloatingIPFixedIPAddress),
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
		string(model.UUID),
		string(model.FloatingIPAddressFamily),
		string(model.FloatingIPAddress),
		bool(model.FloatingIPPortMappingsEnable),
		string(model.FloatingIPTrafficDirection),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		utils.MustJSON(model.FQName))

	stmtProjectRef, err := tx.Prepare(insertFloatingIPProjectQuery)
	if err != nil {
		return err
	}
	defer stmtProjectRef.Close()
	for _, ref := range model.ProjectRefs {
		_, err = stmtProjectRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertFloatingIPVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanFloatingIP(rows *sql.Rows) (*models.FloatingIP, error) {
	m := models.MakeFloatingIP()

	var jsonFloatingIPPortMappings string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonFloatingIPPortMappings,
		&m.FloatingIPIsVirtualIP,
		&m.FloatingIPFixedIPAddress,
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
		&m.UUID,
		&m.FloatingIPAddressFamily,
		&m.FloatingIPAddress,
		&m.FloatingIPPortMappingsEnable,
		&m.FloatingIPTrafficDirection,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFloatingIPPortMappings), &m.FloatingIPPortMappings)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildFloatingIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["floating_ip_fixed_ip_address"]; ok {
		results = append(results, "floating_ip_fixed_ip_address = ?")
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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_address_family"]; ok {
		results = append(results, "floating_ip_address_family = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_address"]; ok {
		results = append(results, "floating_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["floating_ip_traffic_direction"]; ok {
		results = append(results, "floating_ip_traffic_direction = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListFloatingIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.FloatingIP, error) {
	result := models.MakeFloatingIPSlice()
	whereQuery, values := buildFloatingIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listFloatingIPQuery)
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
		m, _ := scanFloatingIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowFloatingIP(tx *sql.Tx, uuid string) (*models.FloatingIP, error) {
	rows, err := tx.Query(showFloatingIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanFloatingIP(rows)
	}
	return nil, nil
}

func UpdateFloatingIP(tx *sql.Tx, uuid string, model *models.FloatingIP) error {
	return nil
}

func DeleteFloatingIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
