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

const insertServiceEndpointQuery = "insert into `service_endpoint` (`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceEndpointQuery = "update `service_endpoint` set `fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?;"
const deleteServiceEndpointQuery = "delete from `service_endpoint` where uuid = ?"
const listServiceEndpointQuery = "select `service_endpoint`.`fq_name`,`service_endpoint`.`enable`,`service_endpoint`.`description`,`service_endpoint`.`created`,`service_endpoint`.`creator`,`service_endpoint`.`user_visible`,`service_endpoint`.`last_modified`,`service_endpoint`.`owner`,`service_endpoint`.`owner_access`,`service_endpoint`.`other_access`,`service_endpoint`.`group`,`service_endpoint`.`group_access`,`service_endpoint`.`display_name`,`service_endpoint`.`key_value_pair`,`service_endpoint`.`global_access`,`service_endpoint`.`share`,`service_endpoint`.`perms2_owner`,`service_endpoint`.`perms2_owner_access`,`service_endpoint`.`uuid` from `service_endpoint`"
const showServiceEndpointQuery = "select `service_endpoint`.`fq_name`,`service_endpoint`.`enable`,`service_endpoint`.`description`,`service_endpoint`.`created`,`service_endpoint`.`creator`,`service_endpoint`.`user_visible`,`service_endpoint`.`last_modified`,`service_endpoint`.`owner`,`service_endpoint`.`owner_access`,`service_endpoint`.`other_access`,`service_endpoint`.`group`,`service_endpoint`.`group_access`,`service_endpoint`.`display_name`,`service_endpoint`.`key_value_pair`,`service_endpoint`.`global_access`,`service_endpoint`.`share`,`service_endpoint`.`perms2_owner`,`service_endpoint`.`perms2_owner_access`,`service_endpoint`.`uuid` from `service_endpoint` where uuid = ?"

const insertServiceEndpointServiceConnectionModuleQuery = "insert into `ref_service_endpoint_service_connection_module` (`from`, `to` ) values (?, ?);"

const insertServiceEndpointPhysicalRouterQuery = "insert into `ref_service_endpoint_physical_router` (`from`, `to` ) values (?, ?);"

const insertServiceEndpointServiceObjectQuery = "insert into `ref_service_endpoint_service_object` (`from`, `to` ) values (?, ?);"

func CreateServiceEndpoint(tx *sql.Tx, model *models.ServiceEndpoint) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceEndpointQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID))

	stmtServiceConnectionModuleRef, err := tx.Prepare(insertServiceEndpointServiceConnectionModuleQuery)
	if err != nil {
		return err
	}
	defer stmtServiceConnectionModuleRef.Close()
	for _, ref := range model.ServiceConnectionModuleRefs {
		_, err = stmtServiceConnectionModuleRef.Exec(model.UUID, ref.UUID)
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertServiceEndpointPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {
		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceObjectRef, err := tx.Prepare(insertServiceEndpointServiceObjectQuery)
	if err != nil {
		return err
	}
	defer stmtServiceObjectRef.Close()
	for _, ref := range model.ServiceObjectRefs {
		_, err = stmtServiceObjectRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanServiceEndpoint(rows *sql.Rows) (*models.ServiceEndpoint, error) {
	m := models.MakeServiceEndpoint()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildServiceEndpointWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

func ListServiceEndpoint(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceEndpoint, error) {
	result := models.MakeServiceEndpointSlice()
	whereQuery, values := buildServiceEndpointWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceEndpointQuery)
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
		m, _ := scanServiceEndpoint(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceEndpoint(tx *sql.Tx, uuid string) (*models.ServiceEndpoint, error) {
	rows, err := tx.Query(showServiceEndpointQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceEndpoint(rows)
	}
	return nil, nil
}

func UpdateServiceEndpoint(tx *sql.Tx, uuid string, model *models.ServiceEndpoint) error {
	return nil
}

func DeleteServiceEndpoint(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceEndpointQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
