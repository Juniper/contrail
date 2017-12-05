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

const insertServiceConnectionModuleQuery = "insert into `service_connection_module` (`fq_name`,`last_modified`,`owner_access`,`other_access`,`group`,`group_access`,`owner`,`enable`,`description`,`created`,`creator`,`user_visible`,`display_name`,`service_type`,`e2_service`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceConnectionModuleQuery = "update `service_connection_module` set `fq_name` = ?,`last_modified` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`display_name` = ?,`service_type` = ?,`e2_service` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteServiceConnectionModuleQuery = "delete from `service_connection_module` where uuid = ?"
const listServiceConnectionModuleQuery = "select `service_connection_module`.`fq_name`,`service_connection_module`.`last_modified`,`service_connection_module`.`owner_access`,`service_connection_module`.`other_access`,`service_connection_module`.`group`,`service_connection_module`.`group_access`,`service_connection_module`.`owner`,`service_connection_module`.`enable`,`service_connection_module`.`description`,`service_connection_module`.`created`,`service_connection_module`.`creator`,`service_connection_module`.`user_visible`,`service_connection_module`.`display_name`,`service_connection_module`.`service_type`,`service_connection_module`.`e2_service`,`service_connection_module`.`key_value_pair`,`service_connection_module`.`perms2_owner`,`service_connection_module`.`perms2_owner_access`,`service_connection_module`.`global_access`,`service_connection_module`.`share`,`service_connection_module`.`uuid` from `service_connection_module`"
const showServiceConnectionModuleQuery = "select `service_connection_module`.`fq_name`,`service_connection_module`.`last_modified`,`service_connection_module`.`owner_access`,`service_connection_module`.`other_access`,`service_connection_module`.`group`,`service_connection_module`.`group_access`,`service_connection_module`.`owner`,`service_connection_module`.`enable`,`service_connection_module`.`description`,`service_connection_module`.`created`,`service_connection_module`.`creator`,`service_connection_module`.`user_visible`,`service_connection_module`.`display_name`,`service_connection_module`.`service_type`,`service_connection_module`.`e2_service`,`service_connection_module`.`key_value_pair`,`service_connection_module`.`perms2_owner`,`service_connection_module`.`perms2_owner_access`,`service_connection_module`.`global_access`,`service_connection_module`.`share`,`service_connection_module`.`uuid` from `service_connection_module` where uuid = ?"

const insertServiceConnectionModuleServiceObjectQuery = "insert into `ref_service_connection_module_service_object` (`from`, `to` ) values (?, ?);"

func CreateServiceConnectionModule(tx *sql.Tx, model *models.ServiceConnectionModule) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceConnectionModuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.DisplayName),
		string(model.ServiceType),
		string(model.E2Service),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))

	stmtServiceObjectRef, err := tx.Prepare(insertServiceConnectionModuleServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmtServiceObjectRef.Close()
	for _, ref := range model.ServiceObjectRefs {
		_, err = stmtServiceObjectRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanServiceConnectionModule(rows *sql.Rows) (*models.ServiceConnectionModule, error) {
	m := models.MakeServiceConnectionModule()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.DisplayName,
		&m.ServiceType,
		&m.E2Service,
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

func buildServiceConnectionModuleWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["service_type"]; ok {
		results = append(results, "service_type = ?")
		values = append(values, value)
	}

	if value, ok := where["e2_service"]; ok {
		results = append(results, "e2_service = ?")
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

func ListServiceConnectionModule(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceConnectionModule, error) {
	result := models.MakeServiceConnectionModuleSlice()
	whereQuery, values := buildServiceConnectionModuleWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceConnectionModuleQuery)
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
		m, _ := scanServiceConnectionModule(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceConnectionModule(tx *sql.Tx, uuid string) (*models.ServiceConnectionModule, error) {
	rows, err := tx.Query(showServiceConnectionModuleQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceConnectionModule(rows)
	}
	return nil, nil
}

func UpdateServiceConnectionModule(tx *sql.Tx, uuid string, model *models.ServiceConnectionModule) error {
	return nil
}

func DeleteServiceConnectionModule(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceConnectionModuleQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
