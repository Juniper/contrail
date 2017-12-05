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

const insertPolicyManagementQuery = "insert into `policy_management` (`uuid`,`fq_name`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updatePolicyManagementQuery = "update `policy_management` set `uuid` = ?,`fq_name` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deletePolicyManagementQuery = "delete from `policy_management` where uuid = ?"
const listPolicyManagementQuery = "select `policy_management`.`uuid`,`policy_management`.`fq_name`,`policy_management`.`last_modified`,`policy_management`.`owner`,`policy_management`.`owner_access`,`policy_management`.`other_access`,`policy_management`.`group`,`policy_management`.`group_access`,`policy_management`.`enable`,`policy_management`.`description`,`policy_management`.`created`,`policy_management`.`creator`,`policy_management`.`user_visible`,`policy_management`.`display_name`,`policy_management`.`key_value_pair`,`policy_management`.`perms2_owner`,`policy_management`.`perms2_owner_access`,`policy_management`.`global_access`,`policy_management`.`share` from `policy_management`"
const showPolicyManagementQuery = "select `policy_management`.`uuid`,`policy_management`.`fq_name`,`policy_management`.`last_modified`,`policy_management`.`owner`,`policy_management`.`owner_access`,`policy_management`.`other_access`,`policy_management`.`group`,`policy_management`.`group_access`,`policy_management`.`enable`,`policy_management`.`description`,`policy_management`.`created`,`policy_management`.`creator`,`policy_management`.`user_visible`,`policy_management`.`display_name`,`policy_management`.`key_value_pair`,`policy_management`.`perms2_owner`,`policy_management`.`perms2_owner_access`,`policy_management`.`global_access`,`policy_management`.`share` from `policy_management` where uuid = ?"

func CreatePolicyManagement(tx *sql.Tx, model *models.PolicyManagement) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertPolicyManagementQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.UUID),
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
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share))

	return err
}

func scanPolicyManagement(rows *sql.Rows) (*models.PolicyManagement, error) {
	m := models.MakePolicyManagement()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.UUID,
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

func buildPolicyManagementWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	return "where " + strings.Join(results, " and "), values
}

func ListPolicyManagement(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.PolicyManagement, error) {
	result := models.MakePolicyManagementSlice()
	whereQuery, values := buildPolicyManagementWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listPolicyManagementQuery)
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
		m, _ := scanPolicyManagement(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowPolicyManagement(tx *sql.Tx, uuid string) (*models.PolicyManagement, error) {
	rows, err := tx.Query(showPolicyManagementQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanPolicyManagement(rows)
	}
	return nil, nil
}

func UpdatePolicyManagement(tx *sql.Tx, uuid string, model *models.PolicyManagement) error {
	return nil
}

func DeletePolicyManagement(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deletePolicyManagementQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
