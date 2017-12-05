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

const insertE2ServiceProviderQuery = "insert into `e2_service_provider` (`e2_service_provider_promiscuous`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateE2ServiceProviderQuery = "update `e2_service_provider` set `e2_service_provider_promiscuous` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?;"
const deleteE2ServiceProviderQuery = "delete from `e2_service_provider` where uuid = ?"
const listE2ServiceProviderQuery = "select `e2_service_provider`.`e2_service_provider_promiscuous`,`e2_service_provider`.`uuid`,`e2_service_provider`.`fq_name`,`e2_service_provider`.`enable`,`e2_service_provider`.`description`,`e2_service_provider`.`created`,`e2_service_provider`.`creator`,`e2_service_provider`.`user_visible`,`e2_service_provider`.`last_modified`,`e2_service_provider`.`other_access`,`e2_service_provider`.`group`,`e2_service_provider`.`group_access`,`e2_service_provider`.`owner`,`e2_service_provider`.`owner_access`,`e2_service_provider`.`display_name`,`e2_service_provider`.`key_value_pair`,`e2_service_provider`.`perms2_owner_access`,`e2_service_provider`.`global_access`,`e2_service_provider`.`share`,`e2_service_provider`.`perms2_owner` from `e2_service_provider`"
const showE2ServiceProviderQuery = "select `e2_service_provider`.`e2_service_provider_promiscuous`,`e2_service_provider`.`uuid`,`e2_service_provider`.`fq_name`,`e2_service_provider`.`enable`,`e2_service_provider`.`description`,`e2_service_provider`.`created`,`e2_service_provider`.`creator`,`e2_service_provider`.`user_visible`,`e2_service_provider`.`last_modified`,`e2_service_provider`.`other_access`,`e2_service_provider`.`group`,`e2_service_provider`.`group_access`,`e2_service_provider`.`owner`,`e2_service_provider`.`owner_access`,`e2_service_provider`.`display_name`,`e2_service_provider`.`key_value_pair`,`e2_service_provider`.`perms2_owner_access`,`e2_service_provider`.`global_access`,`e2_service_provider`.`share`,`e2_service_provider`.`perms2_owner` from `e2_service_provider` where uuid = ?"

const insertE2ServiceProviderPeeringPolicyQuery = "insert into `ref_e2_service_provider_peering_policy` (`from`, `to` ) values (?, ?);"

const insertE2ServiceProviderPhysicalRouterQuery = "insert into `ref_e2_service_provider_physical_router` (`from`, `to` ) values (?, ?);"

func CreateE2ServiceProvider(tx *sql.Tx, model *models.E2ServiceProvider) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertE2ServiceProviderQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.E2ServiceProviderPromiscuous),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner))

	stmtPhysicalRouterRef, err := tx.Prepare(insertE2ServiceProviderPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {
		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtPeeringPolicyRef, err := tx.Prepare(insertE2ServiceProviderPeeringPolicyQuery)
	if err != nil {
		return err
	}
	defer stmtPeeringPolicyRef.Close()
	for _, ref := range model.PeeringPolicyRefs {
		_, err = stmtPeeringPolicyRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanE2ServiceProvider(rows *sql.Rows) (*models.E2ServiceProvider, error) {
	m := models.MakeE2ServiceProvider()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.E2ServiceProviderPromiscuous,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildE2ServiceProviderWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	return "where " + strings.Join(results, " and "), values
}

func ListE2ServiceProvider(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.E2ServiceProvider, error) {
	result := models.MakeE2ServiceProviderSlice()
	whereQuery, values := buildE2ServiceProviderWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listE2ServiceProviderQuery)
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
		m, _ := scanE2ServiceProvider(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowE2ServiceProvider(tx *sql.Tx, uuid string) (*models.E2ServiceProvider, error) {
	rows, err := tx.Query(showE2ServiceProviderQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanE2ServiceProvider(rows)
	}
	return nil, nil
}

func UpdateE2ServiceProvider(tx *sql.Tx, uuid string, model *models.E2ServiceProvider) error {
	return nil
}

func DeleteE2ServiceProvider(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteE2ServiceProviderQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
