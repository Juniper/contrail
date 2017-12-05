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

const insertDashboardQuery = "insert into `dashboard` (`fq_name`,`container_config`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateDashboardQuery = "update `dashboard` set `fq_name` = ?,`container_config` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteDashboardQuery = "delete from `dashboard` where uuid = ?"
const listDashboardQuery = "select `dashboard`.`fq_name`,`dashboard`.`container_config`,`dashboard`.`user_visible`,`dashboard`.`last_modified`,`dashboard`.`owner`,`dashboard`.`owner_access`,`dashboard`.`other_access`,`dashboard`.`group`,`dashboard`.`group_access`,`dashboard`.`enable`,`dashboard`.`description`,`dashboard`.`created`,`dashboard`.`creator`,`dashboard`.`display_name`,`dashboard`.`key_value_pair`,`dashboard`.`perms2_owner`,`dashboard`.`perms2_owner_access`,`dashboard`.`global_access`,`dashboard`.`share`,`dashboard`.`uuid` from `dashboard`"
const showDashboardQuery = "select `dashboard`.`fq_name`,`dashboard`.`container_config`,`dashboard`.`user_visible`,`dashboard`.`last_modified`,`dashboard`.`owner`,`dashboard`.`owner_access`,`dashboard`.`other_access`,`dashboard`.`group`,`dashboard`.`group_access`,`dashboard`.`enable`,`dashboard`.`description`,`dashboard`.`created`,`dashboard`.`creator`,`dashboard`.`display_name`,`dashboard`.`key_value_pair`,`dashboard`.`perms2_owner`,`dashboard`.`perms2_owner_access`,`dashboard`.`global_access`,`dashboard`.`share`,`dashboard`.`uuid` from `dashboard` where uuid = ?"

func CreateDashboard(tx *sql.Tx, model *models.Dashboard) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDashboardQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.ContainerConfig),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))

	return err
}

func scanDashboard(rows *sql.Rows) (*models.Dashboard, error) {
	m := models.MakeDashboard()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.ContainerConfig,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildDashboardWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["container_config"]; ok {
		results = append(results, "container_config = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListDashboard(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Dashboard, error) {
	result := models.MakeDashboardSlice()
	whereQuery, values := buildDashboardWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listDashboardQuery)
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
		m, _ := scanDashboard(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowDashboard(tx *sql.Tx, uuid string) (*models.Dashboard, error) {
	rows, err := tx.Query(showDashboardQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanDashboard(rows)
	}
	return nil, nil
}

func UpdateDashboard(tx *sql.Tx, uuid string, model *models.Dashboard) error {
	return nil
}

func DeleteDashboard(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteDashboardQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
