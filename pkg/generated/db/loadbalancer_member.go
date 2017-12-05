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

const insertLoadbalancerMemberQuery = "insert into `loadbalancer_member` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`protocol_port`,`status`,`status_description`,`weight`,`admin_state`,`address`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerMemberQuery = "update `loadbalancer_member` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`protocol_port` = ?,`status` = ?,`status_description` = ?,`weight` = ?,`admin_state` = ?,`address` = ?;"
const deleteLoadbalancerMemberQuery = "delete from `loadbalancer_member` where uuid = ?"
const listLoadbalancerMemberQuery = "select `loadbalancer_member`.`uuid`,`loadbalancer_member`.`fq_name`,`loadbalancer_member`.`description`,`loadbalancer_member`.`created`,`loadbalancer_member`.`creator`,`loadbalancer_member`.`user_visible`,`loadbalancer_member`.`last_modified`,`loadbalancer_member`.`owner_access`,`loadbalancer_member`.`other_access`,`loadbalancer_member`.`group`,`loadbalancer_member`.`group_access`,`loadbalancer_member`.`owner`,`loadbalancer_member`.`enable`,`loadbalancer_member`.`display_name`,`loadbalancer_member`.`key_value_pair`,`loadbalancer_member`.`global_access`,`loadbalancer_member`.`share`,`loadbalancer_member`.`perms2_owner`,`loadbalancer_member`.`perms2_owner_access`,`loadbalancer_member`.`protocol_port`,`loadbalancer_member`.`status`,`loadbalancer_member`.`status_description`,`loadbalancer_member`.`weight`,`loadbalancer_member`.`admin_state`,`loadbalancer_member`.`address` from `loadbalancer_member`"
const showLoadbalancerMemberQuery = "select `loadbalancer_member`.`uuid`,`loadbalancer_member`.`fq_name`,`loadbalancer_member`.`description`,`loadbalancer_member`.`created`,`loadbalancer_member`.`creator`,`loadbalancer_member`.`user_visible`,`loadbalancer_member`.`last_modified`,`loadbalancer_member`.`owner_access`,`loadbalancer_member`.`other_access`,`loadbalancer_member`.`group`,`loadbalancer_member`.`group_access`,`loadbalancer_member`.`owner`,`loadbalancer_member`.`enable`,`loadbalancer_member`.`display_name`,`loadbalancer_member`.`key_value_pair`,`loadbalancer_member`.`global_access`,`loadbalancer_member`.`share`,`loadbalancer_member`.`perms2_owner`,`loadbalancer_member`.`perms2_owner_access`,`loadbalancer_member`.`protocol_port`,`loadbalancer_member`.`status`,`loadbalancer_member`.`status_description`,`loadbalancer_member`.`weight`,`loadbalancer_member`.`admin_state`,`loadbalancer_member`.`address` from `loadbalancer_member` where uuid = ?"

func CreateLoadbalancerMember(tx *sql.Tx, model *models.LoadbalancerMember) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
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
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.LoadbalancerMemberProperties.ProtocolPort),
		string(model.LoadbalancerMemberProperties.Status),
		string(model.LoadbalancerMemberProperties.StatusDescription),
		int(model.LoadbalancerMemberProperties.Weight),
		bool(model.LoadbalancerMemberProperties.AdminState),
		string(model.LoadbalancerMemberProperties.Address))

	return err
}

func scanLoadbalancerMember(rows *sql.Rows) (*models.LoadbalancerMember, error) {
	m := models.MakeLoadbalancerMember()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.LoadbalancerMemberProperties.ProtocolPort,
		&m.LoadbalancerMemberProperties.Status,
		&m.LoadbalancerMemberProperties.StatusDescription,
		&m.LoadbalancerMemberProperties.Weight,
		&m.LoadbalancerMemberProperties.AdminState,
		&m.LoadbalancerMemberProperties.Address); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildLoadbalancerMemberWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["status_description"]; ok {
		results = append(results, "status_description = ?")
		values = append(values, value)
	}

	if value, ok := where["address"]; ok {
		results = append(results, "address = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancerMember(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LoadbalancerMember, error) {
	result := models.MakeLoadbalancerMemberSlice()
	whereQuery, values := buildLoadbalancerMemberWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerMemberQuery)
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
		m, _ := scanLoadbalancerMember(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancerMember(tx *sql.Tx, uuid string) (*models.LoadbalancerMember, error) {
	rows, err := tx.Query(showLoadbalancerMemberQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancerMember(rows)
	}
	return nil, nil
}

func UpdateLoadbalancerMember(tx *sql.Tx, uuid string, model *models.LoadbalancerMember) error {
	return nil
}

func DeleteLoadbalancerMember(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
