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

const insertServiceInstanceQuery = "insert into `service_instance` (`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`display_name`,`key_value_pair`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`service_instance_bindings`,`auto_policy`,`management_virtual_network`,`virtual_router_id`,`right_ip_address`,`availability_zone`,`left_virtual_network`,`auto_scale`,`max_instances`,`left_ip_address`,`ha_mode`,`interface_list`,`right_virtual_network`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`service_instance_bindings` = ?,`auto_policy` = ?,`management_virtual_network` = ?,`virtual_router_id` = ?,`right_ip_address` = ?,`availability_zone` = ?,`left_virtual_network` = ?,`auto_scale` = ?,`max_instances` = ?,`left_ip_address` = ?,`ha_mode` = ?,`interface_list` = ?,`right_virtual_network` = ?,`uuid` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"
const listServiceInstanceQuery = "select `service_instance`.`fq_name`,`service_instance`.`enable`,`service_instance`.`description`,`service_instance`.`created`,`service_instance`.`creator`,`service_instance`.`user_visible`,`service_instance`.`last_modified`,`service_instance`.`group`,`service_instance`.`group_access`,`service_instance`.`owner`,`service_instance`.`owner_access`,`service_instance`.`other_access`,`service_instance`.`display_name`,`service_instance`.`key_value_pair`,`service_instance`.`perms2_owner_access`,`service_instance`.`global_access`,`service_instance`.`share`,`service_instance`.`perms2_owner`,`service_instance`.`service_instance_bindings`,`service_instance`.`auto_policy`,`service_instance`.`management_virtual_network`,`service_instance`.`virtual_router_id`,`service_instance`.`right_ip_address`,`service_instance`.`availability_zone`,`service_instance`.`left_virtual_network`,`service_instance`.`auto_scale`,`service_instance`.`max_instances`,`service_instance`.`left_ip_address`,`service_instance`.`ha_mode`,`service_instance`.`interface_list`,`service_instance`.`right_virtual_network`,`service_instance`.`uuid` from `service_instance`"
const showServiceInstanceQuery = "select `service_instance`.`fq_name`,`service_instance`.`enable`,`service_instance`.`description`,`service_instance`.`created`,`service_instance`.`creator`,`service_instance`.`user_visible`,`service_instance`.`last_modified`,`service_instance`.`group`,`service_instance`.`group_access`,`service_instance`.`owner`,`service_instance`.`owner_access`,`service_instance`.`other_access`,`service_instance`.`display_name`,`service_instance`.`key_value_pair`,`service_instance`.`perms2_owner_access`,`service_instance`.`global_access`,`service_instance`.`share`,`service_instance`.`perms2_owner`,`service_instance`.`service_instance_bindings`,`service_instance`.`auto_policy`,`service_instance`.`management_virtual_network`,`service_instance`.`virtual_router_id`,`service_instance`.`right_ip_address`,`service_instance`.`availability_zone`,`service_instance`.`left_virtual_network`,`service_instance`.`auto_scale`,`service_instance`.`max_instances`,`service_instance`.`left_ip_address`,`service_instance`.`ha_mode`,`service_instance`.`interface_list`,`service_instance`.`right_virtual_network`,`service_instance`.`uuid` from `service_instance` where uuid = ?"

const insertServiceInstanceServiceTemplateQuery = "insert into `ref_service_instance_service_template` (`from`, `to` ) values (?, ?);"

const insertServiceInstanceInstanceIPQuery = "insert into `ref_service_instance_instance_ip` (`from`, `to` ) values (?, ?);"

func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
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
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		utils.MustJSON(model.ServiceInstanceBindings),
		bool(model.ServiceInstanceProperties.AutoPolicy),
		string(model.ServiceInstanceProperties.ManagementVirtualNetwork),
		string(model.ServiceInstanceProperties.VirtualRouterID),
		string(model.ServiceInstanceProperties.RightIPAddress),
		string(model.ServiceInstanceProperties.AvailabilityZone),
		string(model.ServiceInstanceProperties.LeftVirtualNetwork),
		bool(model.ServiceInstanceProperties.ScaleOut.AutoScale),
		int(model.ServiceInstanceProperties.ScaleOut.MaxInstances),
		string(model.ServiceInstanceProperties.LeftIPAddress),
		string(model.ServiceInstanceProperties.HaMode),
		utils.MustJSON(model.ServiceInstanceProperties.InterfaceList),
		string(model.ServiceInstanceProperties.RightVirtualNetwork),
		string(model.UUID))

	stmtServiceTemplateRef, err := tx.Prepare(insertServiceInstanceServiceTemplateQuery)
	if err != nil {
		return err
	}
	defer stmtServiceTemplateRef.Close()
	for _, ref := range model.ServiceTemplateRefs {
		_, err = stmtServiceTemplateRef.Exec(model.UUID, ref.UUID)
	}

	stmtInstanceIPRef, err := tx.Prepare(insertServiceInstanceInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmtInstanceIPRef.Close()
	for _, ref := range model.InstanceIPRefs {
		_, err = stmtInstanceIPRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanServiceInstance(rows *sql.Rows) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonServiceInstanceBindings string

	var jsonServiceInstancePropertiesInterfaceList string

	if err := rows.Scan(&jsonFQName,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&jsonServiceInstanceBindings,
		&m.ServiceInstanceProperties.AutoPolicy,
		&m.ServiceInstanceProperties.ManagementVirtualNetwork,
		&m.ServiceInstanceProperties.VirtualRouterID,
		&m.ServiceInstanceProperties.RightIPAddress,
		&m.ServiceInstanceProperties.AvailabilityZone,
		&m.ServiceInstanceProperties.LeftVirtualNetwork,
		&m.ServiceInstanceProperties.ScaleOut.AutoScale,
		&m.ServiceInstanceProperties.ScaleOut.MaxInstances,
		&m.ServiceInstanceProperties.LeftIPAddress,
		&m.ServiceInstanceProperties.HaMode,
		&jsonServiceInstancePropertiesInterfaceList,
		&m.ServiceInstanceProperties.RightVirtualNetwork,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonServiceInstanceBindings), &m.ServiceInstanceBindings)

	json.Unmarshal([]byte(jsonServiceInstancePropertiesInterfaceList), &m.ServiceInstanceProperties.InterfaceList)

	return m, nil
}

func buildServiceInstanceWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["management_virtual_network"]; ok {
		results = append(results, "management_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_router_id"]; ok {
		results = append(results, "virtual_router_id = ?")
		values = append(values, value)
	}

	if value, ok := where["right_ip_address"]; ok {
		results = append(results, "right_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["availability_zone"]; ok {
		results = append(results, "availability_zone = ?")
		values = append(values, value)
	}

	if value, ok := where["left_virtual_network"]; ok {
		results = append(results, "left_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["left_ip_address"]; ok {
		results = append(results, "left_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["ha_mode"]; ok {
		results = append(results, "ha_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["right_virtual_network"]; ok {
		results = append(results, "right_virtual_network = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListServiceInstance(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ServiceInstance, error) {
	result := models.MakeServiceInstanceSlice()
	whereQuery, values := buildServiceInstanceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listServiceInstanceQuery)
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
		m, _ := scanServiceInstance(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowServiceInstance(tx *sql.Tx, uuid string) (*models.ServiceInstance, error) {
	rows, err := tx.Query(showServiceInstanceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanServiceInstance(rows)
	}
	return nil, nil
}

func UpdateServiceInstance(tx *sql.Tx, uuid string, model *models.ServiceInstance) error {
	return nil
}

func DeleteServiceInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
