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

const insertServiceGroupQuery = "insert into `service_group` (`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`service_group_firewall_service_list`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceGroupQuery = "update `service_group` set `uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`service_group_firewall_service_list` = ?;"
const deleteServiceGroupQuery = "delete from `service_group` where uuid = ?"
const listServiceGroupQuery = "select `service_group`.`uuid`,`service_group`.`fq_name`,`service_group`.`created`,`service_group`.`creator`,`service_group`.`user_visible`,`service_group`.`last_modified`,`service_group`.`group_access`,`service_group`.`owner`,`service_group`.`owner_access`,`service_group`.`other_access`,`service_group`.`group`,`service_group`.`enable`,`service_group`.`description`,`service_group`.`display_name`,`service_group`.`key_value_pair`,`service_group`.`global_access`,`service_group`.`share`,`service_group`.`perms2_owner`,`service_group`.`perms2_owner_access`,`service_group`.`service_group_firewall_service_list` from `service_group`"
const showServiceGroupQuery = "select `service_group`.`uuid`,`service_group`.`fq_name`,`service_group`.`created`,`service_group`.`creator`,`service_group`.`user_visible`,`service_group`.`last_modified`,`service_group`.`group_access`,`service_group`.`owner`,`service_group`.`owner_access`,`service_group`.`other_access`,`service_group`.`group`,`service_group`.`enable`,`service_group`.`description`,`service_group`.`display_name`,`service_group`.`key_value_pair`,`service_group`.`global_access`,`service_group`.`share`,`service_group`.`perms2_owner`,`service_group`.`perms2_owner_access`,`service_group`.`service_group_firewall_service_list` from `service_group` where uuid = ?"

func CreateServiceGroup(tx *sql.Tx, model *models.ServiceGroup) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.ServiceGroupFirewallServiceList))

	return err
}

func scanServiceGroup(rows *sql.Rows) (*models.ServiceGroup, error) {
	m := models.MakeServiceGroup()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonServiceGroupFirewallServiceList string

	if err := rows.Scan(&m.UUID,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonServiceGroupFirewallServiceList); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonServiceGroupFirewallServiceList), &m.ServiceGroupFirewallServiceList)

	return m, nil
}

func buildServiceGroupWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

func ListServiceGroup(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceGroup, error) {
	result := models.MakeServiceGroupSlice()
	whereQuery, values := buildServiceGroupWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceGroupQuery)
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
		m, _ := scanServiceGroup(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceGroup(tx *sql.Tx, uuid string) (*models.ServiceGroup, error) {
	rows, err := tx.Query(showServiceGroupQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceGroup(rows)
	}
	return nil, nil
}

func UpdateServiceGroup(tx *sql.Tx, uuid string, model *models.ServiceGroup) error {
	return nil
}

func DeleteServiceGroup(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceGroupQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
