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

const insertServiceApplianceQuery = "insert into `service_appliance` (`fq_name`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`key_value_pair`,`username`,`password`,`service_appliance_properties_key_value_pair`,`display_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`service_appliance_ip_address`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceApplianceQuery = "update `service_appliance` set `fq_name` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`key_value_pair` = ?,`username` = ?,`password` = ?,`service_appliance_properties_key_value_pair` = ?,`display_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`service_appliance_ip_address` = ?,`uuid` = ?;"
const deleteServiceApplianceQuery = "delete from `service_appliance` where uuid = ?"
const listServiceApplianceQuery = "select `service_appliance`.`fq_name`,`service_appliance`.`last_modified`,`service_appliance`.`group_access`,`service_appliance`.`owner`,`service_appliance`.`owner_access`,`service_appliance`.`other_access`,`service_appliance`.`group`,`service_appliance`.`enable`,`service_appliance`.`description`,`service_appliance`.`created`,`service_appliance`.`creator`,`service_appliance`.`user_visible`,`service_appliance`.`key_value_pair`,`service_appliance`.`username`,`service_appliance`.`password`,`service_appliance`.`service_appliance_properties_key_value_pair`,`service_appliance`.`display_name`,`service_appliance`.`perms2_owner`,`service_appliance`.`perms2_owner_access`,`service_appliance`.`global_access`,`service_appliance`.`share`,`service_appliance`.`service_appliance_ip_address`,`service_appliance`.`uuid` from `service_appliance`"
const showServiceApplianceQuery = "select `service_appliance`.`fq_name`,`service_appliance`.`last_modified`,`service_appliance`.`group_access`,`service_appliance`.`owner`,`service_appliance`.`owner_access`,`service_appliance`.`other_access`,`service_appliance`.`group`,`service_appliance`.`enable`,`service_appliance`.`description`,`service_appliance`.`created`,`service_appliance`.`creator`,`service_appliance`.`user_visible`,`service_appliance`.`key_value_pair`,`service_appliance`.`username`,`service_appliance`.`password`,`service_appliance`.`service_appliance_properties_key_value_pair`,`service_appliance`.`display_name`,`service_appliance`.`perms2_owner`,`service_appliance`.`perms2_owner_access`,`service_appliance`.`global_access`,`service_appliance`.`share`,`service_appliance`.`service_appliance_ip_address`,`service_appliance`.`uuid` from `service_appliance` where uuid = ?"

const insertServiceAppliancePhysicalInterfaceQuery = "insert into `ref_service_appliance_physical_interface` (`from`, `to` ,`interface_type`) values (?, ?,?);"

func CreateServiceAppliance(tx *sql.Tx, model *models.ServiceAppliance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceApplianceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ServiceApplianceUserCredentials.Username),
		string(model.ServiceApplianceUserCredentials.Password),
		utils.MustJSON(model.ServiceApplianceProperties.KeyValuePair),
		string(model.DisplayName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.ServiceApplianceIPAddress),
		string(model.UUID))

	stmtPhysicalInterfaceRef, err := tx.Prepare(insertServiceAppliancePhysicalInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalInterfaceRef.Close()
	for _, ref := range model.PhysicalInterfaceRefs {
		_, err = stmtPhysicalInterfaceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.InterfaceType))
	}

	return err
}

func scanServiceAppliance(rows *sql.Rows) (*models.ServiceAppliance, error) {
	m := models.MakeServiceAppliance()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonServiceAppliancePropertiesKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&jsonAnnotationsKeyValuePair,
		&m.ServiceApplianceUserCredentials.Username,
		&m.ServiceApplianceUserCredentials.Password,
		&jsonServiceAppliancePropertiesKeyValuePair,
		&m.DisplayName,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.ServiceApplianceIPAddress,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonServiceAppliancePropertiesKeyValuePair), &m.ServiceApplianceProperties.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildServiceApplianceWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["username"]; ok {
		results = append(results, "username = ?")
		values = append(values, value)
	}

	if value, ok := where["password"]; ok {
		results = append(results, "password = ?")
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

	if value, ok := where["service_appliance_ip_address"]; ok {
		results = append(results, "service_appliance_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceAppliance(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceAppliance, error) {
	result := models.MakeServiceApplianceSlice()
	whereQuery, values := buildServiceApplianceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceApplianceQuery)
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
		m, _ := scanServiceAppliance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceAppliance(tx *sql.Tx, uuid string) (*models.ServiceAppliance, error) {
	rows, err := tx.Query(showServiceApplianceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceAppliance(rows)
	}
	return nil, nil
}

func UpdateServiceAppliance(tx *sql.Tx, uuid string, model *models.ServiceAppliance) error {
	return nil
}

func DeleteServiceAppliance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceApplianceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
