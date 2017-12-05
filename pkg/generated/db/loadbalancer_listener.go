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

const insertLoadbalancerListenerQuery = "insert into `loadbalancer_listener` (`key_value_pair`,`default_tls_container`,`protocol`,`connection_limit`,`admin_state`,`sni_containers`,`protocol_port`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerListenerQuery = "update `loadbalancer_listener` set `key_value_pair` = ?,`default_tls_container` = ?,`protocol` = ?,`connection_limit` = ?,`admin_state` = ?,`sni_containers` = ?,`protocol_port` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?;"
const deleteLoadbalancerListenerQuery = "delete from `loadbalancer_listener` where uuid = ?"
const listLoadbalancerListenerQuery = "select `loadbalancer_listener`.`key_value_pair`,`loadbalancer_listener`.`default_tls_container`,`loadbalancer_listener`.`protocol`,`loadbalancer_listener`.`connection_limit`,`loadbalancer_listener`.`admin_state`,`loadbalancer_listener`.`sni_containers`,`loadbalancer_listener`.`protocol_port`,`loadbalancer_listener`.`owner_access`,`loadbalancer_listener`.`global_access`,`loadbalancer_listener`.`share`,`loadbalancer_listener`.`owner`,`loadbalancer_listener`.`uuid`,`loadbalancer_listener`.`fq_name`,`loadbalancer_listener`.`created`,`loadbalancer_listener`.`creator`,`loadbalancer_listener`.`user_visible`,`loadbalancer_listener`.`last_modified`,`loadbalancer_listener`.`group_access`,`loadbalancer_listener`.`permissions_owner`,`loadbalancer_listener`.`permissions_owner_access`,`loadbalancer_listener`.`other_access`,`loadbalancer_listener`.`group`,`loadbalancer_listener`.`enable`,`loadbalancer_listener`.`description`,`loadbalancer_listener`.`display_name` from `loadbalancer_listener`"
const showLoadbalancerListenerQuery = "select `loadbalancer_listener`.`key_value_pair`,`loadbalancer_listener`.`default_tls_container`,`loadbalancer_listener`.`protocol`,`loadbalancer_listener`.`connection_limit`,`loadbalancer_listener`.`admin_state`,`loadbalancer_listener`.`sni_containers`,`loadbalancer_listener`.`protocol_port`,`loadbalancer_listener`.`owner_access`,`loadbalancer_listener`.`global_access`,`loadbalancer_listener`.`share`,`loadbalancer_listener`.`owner`,`loadbalancer_listener`.`uuid`,`loadbalancer_listener`.`fq_name`,`loadbalancer_listener`.`created`,`loadbalancer_listener`.`creator`,`loadbalancer_listener`.`user_visible`,`loadbalancer_listener`.`last_modified`,`loadbalancer_listener`.`group_access`,`loadbalancer_listener`.`permissions_owner`,`loadbalancer_listener`.`permissions_owner_access`,`loadbalancer_listener`.`other_access`,`loadbalancer_listener`.`group`,`loadbalancer_listener`.`enable`,`loadbalancer_listener`.`description`,`loadbalancer_listener`.`display_name` from `loadbalancer_listener` where uuid = ?"

const insertLoadbalancerListenerLoadbalancerQuery = "insert into `ref_loadbalancer_listener_loadbalancer` (`from`, `to` ) values (?, ?);"

func CreateLoadbalancerListener(tx *sql.Tx, model *models.LoadbalancerListener) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerListenerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.LoadbalancerListenerProperties.DefaultTLSContainer),
		string(model.LoadbalancerListenerProperties.Protocol),
		int(model.LoadbalancerListenerProperties.ConnectionLimit),
		bool(model.LoadbalancerListenerProperties.AdminState),
		utils.MustJSON(model.LoadbalancerListenerProperties.SniContainers),
		int(model.LoadbalancerListenerProperties.ProtocolPort),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.DisplayName))

	stmtLoadbalancerRef, err := tx.Prepare(insertLoadbalancerListenerLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmtLoadbalancerRef.Close()
	for _, ref := range model.LoadbalancerRefs {
		_, err = stmtLoadbalancerRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanLoadbalancerListener(rows *sql.Rows) (*models.LoadbalancerListener, error) {
	m := models.MakeLoadbalancerListener()

	var jsonAnnotationsKeyValuePair string

	var jsonLoadbalancerListenerPropertiesSniContainers string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonAnnotationsKeyValuePair,
		&m.LoadbalancerListenerProperties.DefaultTLSContainer,
		&m.LoadbalancerListenerProperties.Protocol,
		&m.LoadbalancerListenerProperties.ConnectionLimit,
		&m.LoadbalancerListenerProperties.AdminState,
		&jsonLoadbalancerListenerPropertiesSniContainers,
		&m.LoadbalancerListenerProperties.ProtocolPort,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonLoadbalancerListenerPropertiesSniContainers), &m.LoadbalancerListenerProperties.SniContainers)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildLoadbalancerListenerWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["default_tls_container"]; ok {
		results = append(results, "default_tls_container = ?")
		values = append(values, value)
	}

	if value, ok := where["protocol"]; ok {
		results = append(results, "protocol = ?")
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

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancerListener(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerListener, error) {
	result := models.MakeLoadbalancerListenerSlice()
	whereQuery, values := buildLoadbalancerListenerWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerListenerQuery)
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
		m, _ := scanLoadbalancerListener(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerListener(tx *sql.Tx, uuid string) (*models.LoadbalancerListener, error) {
	rows, err := tx.Query(showLoadbalancerListenerQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerListener(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerListener(tx *sql.Tx, uuid string, model *models.LoadbalancerListener) error {
	return nil
}

func DeleteLoadbalancerListener(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerListenerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
