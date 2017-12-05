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

const insertVirtualIPQuery = "insert into `virtual_ip` (`key_value_pair`,`owner`,`owner_access`,`global_access`,`share`,`persistence_cookie_name`,`admin_state`,`status`,`connection_limit`,`persistence_type`,`address`,`protocol_port`,`protocol`,`subnet_id`,`status_description`,`uuid`,`fq_name`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `key_value_pair` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`persistence_cookie_name` = ?,`admin_state` = ?,`status` = ?,`connection_limit` = ?,`persistence_type` = ?,`address` = ?,`protocol_port` = ?,`protocol` = ?,`subnet_id` = ?,`status_description` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip` where uuid = ?"
const listVirtualIPQuery = "select `virtual_ip`.`key_value_pair`,`virtual_ip`.`owner`,`virtual_ip`.`owner_access`,`virtual_ip`.`global_access`,`virtual_ip`.`share`,`virtual_ip`.`persistence_cookie_name`,`virtual_ip`.`admin_state`,`virtual_ip`.`status`,`virtual_ip`.`connection_limit`,`virtual_ip`.`persistence_type`,`virtual_ip`.`address`,`virtual_ip`.`protocol_port`,`virtual_ip`.`protocol`,`virtual_ip`.`subnet_id`,`virtual_ip`.`status_description`,`virtual_ip`.`uuid`,`virtual_ip`.`fq_name`,`virtual_ip`.`last_modified`,`virtual_ip`.`other_access`,`virtual_ip`.`group`,`virtual_ip`.`group_access`,`virtual_ip`.`permissions_owner`,`virtual_ip`.`permissions_owner_access`,`virtual_ip`.`enable`,`virtual_ip`.`description`,`virtual_ip`.`created`,`virtual_ip`.`creator`,`virtual_ip`.`user_visible`,`virtual_ip`.`display_name` from `virtual_ip`"
const showVirtualIPQuery = "select `virtual_ip`.`key_value_pair`,`virtual_ip`.`owner`,`virtual_ip`.`owner_access`,`virtual_ip`.`global_access`,`virtual_ip`.`share`,`virtual_ip`.`persistence_cookie_name`,`virtual_ip`.`admin_state`,`virtual_ip`.`status`,`virtual_ip`.`connection_limit`,`virtual_ip`.`persistence_type`,`virtual_ip`.`address`,`virtual_ip`.`protocol_port`,`virtual_ip`.`protocol`,`virtual_ip`.`subnet_id`,`virtual_ip`.`status_description`,`virtual_ip`.`uuid`,`virtual_ip`.`fq_name`,`virtual_ip`.`last_modified`,`virtual_ip`.`other_access`,`virtual_ip`.`group`,`virtual_ip`.`group_access`,`virtual_ip`.`permissions_owner`,`virtual_ip`.`permissions_owner_access`,`virtual_ip`.`enable`,`virtual_ip`.`description`,`virtual_ip`.`created`,`virtual_ip`.`creator`,`virtual_ip`.`user_visible`,`virtual_ip`.`display_name` from `virtual_ip` where uuid = ?"

const insertVirtualIPVirtualMachineInterfaceQuery = "insert into `ref_virtual_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertVirtualIPLoadbalancerPoolQuery = "insert into `ref_virtual_ip_loadbalancer_pool` (`from`, `to` ) values (?, ?);"

func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.VirtualIPProperties.PersistenceCookieName),
		bool(model.VirtualIPProperties.AdminState),
		string(model.VirtualIPProperties.Status),
		int(model.VirtualIPProperties.ConnectionLimit),
		string(model.VirtualIPProperties.PersistenceType),
		string(model.VirtualIPProperties.Address),
		int(model.VirtualIPProperties.ProtocolPort),
		string(model.VirtualIPProperties.Protocol),
		string(model.VirtualIPProperties.SubnetID),
		string(model.VirtualIPProperties.StatusDescription),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName))

	stmtLoadbalancerPoolRef, err := tx.Prepare(insertVirtualIPLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmtLoadbalancerPoolRef.Close()
	for _, ref := range model.LoadbalancerPoolRefs {
		_, err = stmtLoadbalancerPoolRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualIPVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanVirtualIP(rows *sql.Rows) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.VirtualIPProperties.PersistenceCookieName,
		&m.VirtualIPProperties.AdminState,
		&m.VirtualIPProperties.Status,
		&m.VirtualIPProperties.ConnectionLimit,
		&m.VirtualIPProperties.PersistenceType,
		&m.VirtualIPProperties.Address,
		&m.VirtualIPProperties.ProtocolPort,
		&m.VirtualIPProperties.Protocol,
		&m.VirtualIPProperties.SubnetID,
		&m.VirtualIPProperties.StatusDescription,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildVirtualIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["persistence_cookie_name"]; ok {
		results = append(results, "persistence_cookie_name = ?")
		values = append(values, value)
	}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["persistence_type"]; ok {
		results = append(results, "persistence_type = ?")
		values = append(values, value)
	}

	if value, ok := where["address"]; ok {
		results = append(results, "address = ?")
		values = append(values, value)
	}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
		values = append(values, value)
	}

	if value, ok := where["subnet_id"]; ok {
		results = append(results, "subnet_id = ?")
		values = append(values, value)
	}

	if value, ok := where["status_description"]; ok {
		results = append(results, "status_description = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListVirtualIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.VirtualIP, error) {
	result := models.MakeVirtualIPSlice()
	whereQuery, values := buildVirtualIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listVirtualIPQuery)
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
		m, _ := scanVirtualIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowVirtualIP(tx *sql.Tx, uuid string) (*models.VirtualIP, error) {
	rows, err := tx.Query(showVirtualIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanVirtualIP(rows)
	}
	return nil, nil
}

func UpdateVirtualIP(tx *sql.Tx, uuid string, model *models.VirtualIP) error {
	return nil
}

func DeleteVirtualIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
