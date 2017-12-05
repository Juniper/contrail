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

const insertAliasIPPoolQuery = "insert into `alias_ip_pool` (`uuid`,`fq_name`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAliasIPPoolQuery = "update `alias_ip_pool` set `uuid` = ?,`fq_name` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteAliasIPPoolQuery = "delete from `alias_ip_pool` where uuid = ?"
const listAliasIPPoolQuery = "select `alias_ip_pool`.`uuid`,`alias_ip_pool`.`fq_name`,`alias_ip_pool`.`last_modified`,`alias_ip_pool`.`group`,`alias_ip_pool`.`group_access`,`alias_ip_pool`.`owner`,`alias_ip_pool`.`owner_access`,`alias_ip_pool`.`other_access`,`alias_ip_pool`.`enable`,`alias_ip_pool`.`description`,`alias_ip_pool`.`created`,`alias_ip_pool`.`creator`,`alias_ip_pool`.`user_visible`,`alias_ip_pool`.`display_name`,`alias_ip_pool`.`key_value_pair`,`alias_ip_pool`.`perms2_owner`,`alias_ip_pool`.`perms2_owner_access`,`alias_ip_pool`.`global_access`,`alias_ip_pool`.`share` from `alias_ip_pool`"
const showAliasIPPoolQuery = "select `alias_ip_pool`.`uuid`,`alias_ip_pool`.`fq_name`,`alias_ip_pool`.`last_modified`,`alias_ip_pool`.`group`,`alias_ip_pool`.`group_access`,`alias_ip_pool`.`owner`,`alias_ip_pool`.`owner_access`,`alias_ip_pool`.`other_access`,`alias_ip_pool`.`enable`,`alias_ip_pool`.`description`,`alias_ip_pool`.`created`,`alias_ip_pool`.`creator`,`alias_ip_pool`.`user_visible`,`alias_ip_pool`.`display_name`,`alias_ip_pool`.`key_value_pair`,`alias_ip_pool`.`perms2_owner`,`alias_ip_pool`.`perms2_owner_access`,`alias_ip_pool`.`global_access`,`alias_ip_pool`.`share` from `alias_ip_pool` where uuid = ?"

func CreateAliasIPPool(tx *sql.Tx, model *models.AliasIPPool) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAliasIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))

	return err
}

func scanAliasIPPool(rows *sql.Rows) (*models.AliasIPPool, error) {
	m := models.MakeAliasIPPool()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildAliasIPPoolWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListAliasIPPool(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.AliasIPPool, error) {
	result := models.MakeAliasIPPoolSlice()
	whereQuery, values := buildAliasIPPoolWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listAliasIPPoolQuery)
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
		m, _ := scanAliasIPPool(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowAliasIPPool(tx *sql.Tx, uuid string) (*models.AliasIPPool, error) {
	rows, err := tx.Query(showAliasIPPoolQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanAliasIPPool(rows)
	}
	return nil, nil
}

func UpdateAliasIPPool(tx *sql.Tx, uuid string, model *models.AliasIPPool) error {
	return nil
}

func DeleteAliasIPPool(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAliasIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
