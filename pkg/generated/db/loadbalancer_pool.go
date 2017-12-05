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

const insertLoadbalancerPoolQuery = "insert into `loadbalancer_pool` (`status`,`protocol`,`subnet_id`,`session_persistence`,`admin_state`,`persistence_cookie_name`,`status_description`,`loadbalancer_method`,`key_value_pair`,`loadbalancer_pool_provider`,`owner`,`owner_access`,`global_access`,`share`,`uuid`,`fq_name`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`annotations_key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerPoolQuery = "update `loadbalancer_pool` set `status` = ?,`protocol` = ?,`subnet_id` = ?,`session_persistence` = ?,`admin_state` = ?,`persistence_cookie_name` = ?,`status_description` = ?,`loadbalancer_method` = ?,`key_value_pair` = ?,`loadbalancer_pool_provider` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`annotations_key_value_pair` = ?;"
const deleteLoadbalancerPoolQuery = "delete from `loadbalancer_pool` where uuid = ?"
const listLoadbalancerPoolQuery = "select `loadbalancer_pool`.`status`,`loadbalancer_pool`.`protocol`,`loadbalancer_pool`.`subnet_id`,`loadbalancer_pool`.`session_persistence`,`loadbalancer_pool`.`admin_state`,`loadbalancer_pool`.`persistence_cookie_name`,`loadbalancer_pool`.`status_description`,`loadbalancer_pool`.`loadbalancer_method`,`loadbalancer_pool`.`key_value_pair`,`loadbalancer_pool`.`loadbalancer_pool_provider`,`loadbalancer_pool`.`owner`,`loadbalancer_pool`.`owner_access`,`loadbalancer_pool`.`global_access`,`loadbalancer_pool`.`share`,`loadbalancer_pool`.`uuid`,`loadbalancer_pool`.`fq_name`,`loadbalancer_pool`.`last_modified`,`loadbalancer_pool`.`permissions_owner`,`loadbalancer_pool`.`permissions_owner_access`,`loadbalancer_pool`.`other_access`,`loadbalancer_pool`.`group`,`loadbalancer_pool`.`group_access`,`loadbalancer_pool`.`enable`,`loadbalancer_pool`.`description`,`loadbalancer_pool`.`created`,`loadbalancer_pool`.`creator`,`loadbalancer_pool`.`user_visible`,`loadbalancer_pool`.`display_name`,`loadbalancer_pool`.`annotations_key_value_pair` from `loadbalancer_pool`"
const showLoadbalancerPoolQuery = "select `loadbalancer_pool`.`status`,`loadbalancer_pool`.`protocol`,`loadbalancer_pool`.`subnet_id`,`loadbalancer_pool`.`session_persistence`,`loadbalancer_pool`.`admin_state`,`loadbalancer_pool`.`persistence_cookie_name`,`loadbalancer_pool`.`status_description`,`loadbalancer_pool`.`loadbalancer_method`,`loadbalancer_pool`.`key_value_pair`,`loadbalancer_pool`.`loadbalancer_pool_provider`,`loadbalancer_pool`.`owner`,`loadbalancer_pool`.`owner_access`,`loadbalancer_pool`.`global_access`,`loadbalancer_pool`.`share`,`loadbalancer_pool`.`uuid`,`loadbalancer_pool`.`fq_name`,`loadbalancer_pool`.`last_modified`,`loadbalancer_pool`.`permissions_owner`,`loadbalancer_pool`.`permissions_owner_access`,`loadbalancer_pool`.`other_access`,`loadbalancer_pool`.`group`,`loadbalancer_pool`.`group_access`,`loadbalancer_pool`.`enable`,`loadbalancer_pool`.`description`,`loadbalancer_pool`.`created`,`loadbalancer_pool`.`creator`,`loadbalancer_pool`.`user_visible`,`loadbalancer_pool`.`display_name`,`loadbalancer_pool`.`annotations_key_value_pair` from `loadbalancer_pool` where uuid = ?"

const insertLoadbalancerPoolServiceApplianceSetQuery = "insert into `ref_loadbalancer_pool_service_appliance_set` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_pool_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerListenerQuery = "insert into `ref_loadbalancer_pool_loadbalancer_listener` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolServiceInstanceQuery = "insert into `ref_loadbalancer_pool_service_instance` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerPoolLoadbalancerHealthmonitorQuery = "insert into `ref_loadbalancer_pool_loadbalancer_healthmonitor` (`from`, `to` ) values (?, ?);"

func CreateLoadbalancerPool(tx *sql.Tx, model *models.LoadbalancerPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.LoadbalancerPoolProperties.Status),
		string(model.LoadbalancerPoolProperties.Protocol),
		string(model.LoadbalancerPoolProperties.SubnetID),
		string(model.LoadbalancerPoolProperties.SessionPersistence),
		bool(model.LoadbalancerPoolProperties.AdminState),
		string(model.LoadbalancerPoolProperties.PersistenceCookieName),
		string(model.LoadbalancerPoolProperties.StatusDescription),
		string(model.LoadbalancerPoolProperties.LoadbalancerMethod),
		utils.MustJSON(model.LoadbalancerPoolCustomAttributes.KeyValuePair),
		string(model.LoadbalancerPoolProvider),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtLoadbalancerHealthmonitorRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerHealthmonitorQuery)
	if err != nil {
		return err
	}
	defer stmtLoadbalancerHealthmonitorRef.Close()
	for _, ref := range model.LoadbalancerHealthmonitorRefs {
		_, err = stmtLoadbalancerHealthmonitorRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerPoolServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {
		_, err = stmtServiceApplianceSetRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerPoolVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtLoadbalancerListenerRef, err := tx.Prepare(insertLoadbalancerPoolLoadbalancerListenerQuery)
	if err != nil {
		return err
	}
	defer stmtLoadbalancerListenerRef.Close()
	for _, ref := range model.LoadbalancerListenerRefs {
		_, err = stmtLoadbalancerListenerRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerPoolServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanLoadbalancerPool(rows *sql.Rows) (*models.LoadbalancerPool, error) {
	m := models.MakeLoadbalancerPool()

	var jsonLoadbalancerPoolCustomAttributesKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.LoadbalancerPoolProperties.Status,
		&m.LoadbalancerPoolProperties.Protocol,
		&m.LoadbalancerPoolProperties.SubnetID,
		&m.LoadbalancerPoolProperties.SessionPersistence,
		&m.LoadbalancerPoolProperties.AdminState,
		&m.LoadbalancerPoolProperties.PersistenceCookieName,
		&m.LoadbalancerPoolProperties.StatusDescription,
		&m.LoadbalancerPoolProperties.LoadbalancerMethod,
		&jsonLoadbalancerPoolCustomAttributesKeyValuePair,
		&m.LoadbalancerPoolProvider,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonLoadbalancerPoolCustomAttributesKeyValuePair), &m.LoadbalancerPoolCustomAttributes.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildLoadbalancerPoolWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
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

	if value, ok := where["session_persistence"]; ok {
		results = append(results, "session_persistence = ?")
		values = append(values, value)
	}

	if value, ok := where["persistence_cookie_name"]; ok {
		results = append(results, "persistence_cookie_name = ?")
		values = append(values, value)
	}

	if value, ok := where["status_description"]; ok {
		results = append(results, "status_description = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_method"]; ok {
		results = append(results, "loadbalancer_method = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_pool_provider"]; ok {
		results = append(results, "loadbalancer_pool_provider = ?")
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

func ListLoadbalancerPool(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerPool, error) {
	result := models.MakeLoadbalancerPoolSlice()
	whereQuery, values := buildLoadbalancerPoolWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerPoolQuery)
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
		m, _ := scanLoadbalancerPool(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerPool(tx *sql.Tx, uuid string) (*models.LoadbalancerPool, error) {
	rows, err := tx.Query(showLoadbalancerPoolQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerPool(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerPool(tx *sql.Tx, uuid string, model *models.LoadbalancerPool) error {
	return nil
}

func DeleteLoadbalancerPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
