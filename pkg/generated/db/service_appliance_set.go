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

const insertServiceApplianceSetQuery = "insert into `service_appliance_set` (`key_value_pair`,`service_appliance_driver`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`annotations_key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`service_appliance_ha_mode`,`uuid`,`fq_name`,`display_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceSetQuery = "update `service_appliance_set` set `key_value_pair` = ?,`service_appliance_driver` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`annotations_key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`service_appliance_ha_mode` = ?,`uuid` = ?,`fq_name` = ?,`display_name` = ?;"
const deleteServiceApplianceSetQuery = "delete from `service_appliance_set` where uuid = ?"
const listServiceApplianceSetQuery = "select `service_appliance_set`.`key_value_pair`,`service_appliance_set`.`service_appliance_driver`,`service_appliance_set`.`enable`,`service_appliance_set`.`description`,`service_appliance_set`.`created`,`service_appliance_set`.`creator`,`service_appliance_set`.`user_visible`,`service_appliance_set`.`last_modified`,`service_appliance_set`.`group`,`service_appliance_set`.`group_access`,`service_appliance_set`.`owner`,`service_appliance_set`.`owner_access`,`service_appliance_set`.`other_access`,`service_appliance_set`.`annotations_key_value_pair`,`service_appliance_set`.`share`,`service_appliance_set`.`perms2_owner`,`service_appliance_set`.`perms2_owner_access`,`service_appliance_set`.`global_access`,`service_appliance_set`.`service_appliance_ha_mode`,`service_appliance_set`.`uuid`,`service_appliance_set`.`fq_name`,`service_appliance_set`.`display_name` from `service_appliance_set`"
const showServiceApplianceSetQuery = "select `service_appliance_set`.`key_value_pair`,`service_appliance_set`.`service_appliance_driver`,`service_appliance_set`.`enable`,`service_appliance_set`.`description`,`service_appliance_set`.`created`,`service_appliance_set`.`creator`,`service_appliance_set`.`user_visible`,`service_appliance_set`.`last_modified`,`service_appliance_set`.`group`,`service_appliance_set`.`group_access`,`service_appliance_set`.`owner`,`service_appliance_set`.`owner_access`,`service_appliance_set`.`other_access`,`service_appliance_set`.`annotations_key_value_pair`,`service_appliance_set`.`share`,`service_appliance_set`.`perms2_owner`,`service_appliance_set`.`perms2_owner_access`,`service_appliance_set`.`global_access`,`service_appliance_set`.`service_appliance_ha_mode`,`service_appliance_set`.`uuid`,`service_appliance_set`.`fq_name`,`service_appliance_set`.`display_name` from `service_appliance_set` where uuid = ?"

func CreateServiceApplianceSet(tx *sql.Tx, model *models.ServiceApplianceSet) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.ServiceApplianceSetProperties.KeyValuePair),
		string(model.ServiceApplianceDriver),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.ServiceApplianceHaMode),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.DisplayName))

	return err
}

func scanServiceApplianceSet(rows *sql.Rows) (*models.ServiceApplianceSet, error) {
	m := models.MakeServiceApplianceSet()

	var jsonServiceApplianceSetPropertiesKeyValuePair string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&jsonServiceApplianceSetPropertiesKeyValuePair,
		&m.ServiceApplianceDriver,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.ServiceApplianceHaMode,
		&m.UUID,
		&jsonFQName,
		&m.DisplayName); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonServiceApplianceSetPropertiesKeyValuePair), &m.ServiceApplianceSetProperties.KeyValuePair)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildServiceApplianceSetWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["service_appliance_driver"]; ok {
		results = append(results, "service_appliance_driver = ?")
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

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["service_appliance_ha_mode"]; ok {
		results = append(results, "service_appliance_ha_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceApplianceSet(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceApplianceSet, error) {
	result := models.MakeServiceApplianceSetSlice()
	whereQuery, values := buildServiceApplianceSetWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceApplianceSetQuery)
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
		m, _ := scanServiceApplianceSet(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceApplianceSet(tx *sql.Tx, uuid string) (*models.ServiceApplianceSet, error) {
	rows, err := tx.Query(showServiceApplianceSetQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceApplianceSet(rows)
	}
	return nil, nil
}

func UpdateServiceApplianceSet(tx *sql.Tx, uuid string, model *models.ServiceApplianceSet) error {
	return nil
}

func DeleteServiceApplianceSet(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
