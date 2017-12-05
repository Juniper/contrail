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

const insertDomainQuery = "insert into `domain` (`owner`,`owner_access`,`global_access`,`share`,`project_limit`,`virtual_network_limit`,`security_group_limit`,`uuid`,`fq_name`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDomainQuery = "update `domain` set `owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`project_limit` = ?,`virtual_network_limit` = ?,`security_group_limit` = ?,`uuid` = ?,`fq_name` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteDomainQuery = "delete from `domain` where uuid = ?"
const listDomainQuery = "select `domain`.`owner`,`domain`.`owner_access`,`domain`.`global_access`,`domain`.`share`,`domain`.`project_limit`,`domain`.`virtual_network_limit`,`domain`.`security_group_limit`,`domain`.`uuid`,`domain`.`fq_name`,`domain`.`other_access`,`domain`.`group`,`domain`.`group_access`,`domain`.`permissions_owner`,`domain`.`permissions_owner_access`,`domain`.`enable`,`domain`.`description`,`domain`.`created`,`domain`.`creator`,`domain`.`user_visible`,`domain`.`last_modified`,`domain`.`display_name`,`domain`.`key_value_pair` from `domain`"
const showDomainQuery = "select `domain`.`owner`,`domain`.`owner_access`,`domain`.`global_access`,`domain`.`share`,`domain`.`project_limit`,`domain`.`virtual_network_limit`,`domain`.`security_group_limit`,`domain`.`uuid`,`domain`.`fq_name`,`domain`.`other_access`,`domain`.`group`,`domain`.`group_access`,`domain`.`permissions_owner`,`domain`.`permissions_owner_access`,`domain`.`enable`,`domain`.`description`,`domain`.`created`,`domain`.`creator`,`domain`.`user_visible`,`domain`.`last_modified`,`domain`.`display_name`,`domain`.`key_value_pair` from `domain` where uuid = ?"

func CreateDomain(tx *sql.Tx, model *models.Domain) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		int(model.DomainLimits.ProjectLimit),
		int(model.DomainLimits.VirtualNetworkLimit),
		int(model.DomainLimits.SecurityGroupLimit),
		string(model.UUID),
		utils.MustJSON(model.FQName),
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
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	return err
}

func scanDomain(rows *sql.Rows) (*models.Domain, error) {
	m := models.MakeDomain()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.DomainLimits.ProjectLimit,
		&m.DomainLimits.VirtualNetworkLimit,
		&m.DomainLimits.SecurityGroupLimit,
		&m.UUID,
		&jsonFQName,
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
		&m.IDPerms.LastModified,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildDomainWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListDomain(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Domain, error) {
	result := models.MakeDomainSlice()
	whereQuery, values := buildDomainWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDomainQuery)
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
		m, _ := scanDomain(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDomain(tx *sql.Tx, uuid string) (*models.Domain, error) {
	rows, err := tx.Query(showDomainQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDomain(rows)
	}
	return nil, nil
}

func UpdateDomain(tx *sql.Tx, uuid string, model *models.Domain) error {
	return nil
}

func DeleteDomain(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDomainQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
