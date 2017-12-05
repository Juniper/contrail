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

const insertBGPRouterQuery = "insert into `bgp_router` (`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPRouterQuery = "update `bgp_router` set `last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteBGPRouterQuery = "delete from `bgp_router` where uuid = ?"
const listBGPRouterQuery = "select `bgp_router`.`last_modified`,`bgp_router`.`owner`,`bgp_router`.`owner_access`,`bgp_router`.`other_access`,`bgp_router`.`group`,`bgp_router`.`group_access`,`bgp_router`.`enable`,`bgp_router`.`description`,`bgp_router`.`created`,`bgp_router`.`creator`,`bgp_router`.`user_visible`,`bgp_router`.`display_name`,`bgp_router`.`key_value_pair`,`bgp_router`.`perms2_owner_access`,`bgp_router`.`global_access`,`bgp_router`.`share`,`bgp_router`.`perms2_owner`,`bgp_router`.`uuid`,`bgp_router`.`fq_name` from `bgp_router`"
const showBGPRouterQuery = "select `bgp_router`.`last_modified`,`bgp_router`.`owner`,`bgp_router`.`owner_access`,`bgp_router`.`other_access`,`bgp_router`.`group`,`bgp_router`.`group_access`,`bgp_router`.`enable`,`bgp_router`.`description`,`bgp_router`.`created`,`bgp_router`.`creator`,`bgp_router`.`user_visible`,`bgp_router`.`display_name`,`bgp_router`.`key_value_pair`,`bgp_router`.`perms2_owner_access`,`bgp_router`.`global_access`,`bgp_router`.`share`,`bgp_router`.`perms2_owner`,`bgp_router`.`uuid`,`bgp_router`.`fq_name` from `bgp_router` where uuid = ?"

func CreateBGPRouter(tx *sql.Tx, model *models.BGPRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.IDPerms.LastModified),
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
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName))

	return err
}

func scanBGPRouter(rows *sql.Rows) (*models.BGPRouter, error) {
	m := models.MakeBGPRouter()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.IDPerms.LastModified,
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
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildBGPRouterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

func ListBGPRouter(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BGPRouter, error) {
	result := models.MakeBGPRouterSlice()
	whereQuery, values := buildBGPRouterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBGPRouterQuery)
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
		m, _ := scanBGPRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPRouter(tx *sql.Tx, uuid string) (*models.BGPRouter, error) {
	rows, err := tx.Query(showBGPRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPRouter(rows)
	}
	return nil, nil
}

func UpdateBGPRouter(tx *sql.Tx, uuid string, model *models.BGPRouter) error {
	return nil
}

func DeleteBGPRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
