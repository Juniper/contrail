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

const insertLoadbalancerQuery = "insert into `loadbalancer` (`global_access`,`share`,`owner`,`owner_access`,`vip_subnet_id`,`operating_status`,`status`,`provisioning_status`,`admin_state`,`vip_address`,`loadbalancer_provider`,`uuid`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLoadbalancerQuery = "update `loadbalancer` set `global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`vip_subnet_id` = ?,`operating_status` = ?,`status` = ?,`provisioning_status` = ?,`admin_state` = ?,`vip_address` = ?,`loadbalancer_provider` = ?,`uuid` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteLoadbalancerQuery = "delete from `loadbalancer` where uuid = ?"
const listLoadbalancerQuery = "select `loadbalancer`.`global_access`,`loadbalancer`.`share`,`loadbalancer`.`owner`,`loadbalancer`.`owner_access`,`loadbalancer`.`vip_subnet_id`,`loadbalancer`.`operating_status`,`loadbalancer`.`status`,`loadbalancer`.`provisioning_status`,`loadbalancer`.`admin_state`,`loadbalancer`.`vip_address`,`loadbalancer`.`loadbalancer_provider`,`loadbalancer`.`uuid`,`loadbalancer`.`fq_name`,`loadbalancer`.`created`,`loadbalancer`.`creator`,`loadbalancer`.`user_visible`,`loadbalancer`.`last_modified`,`loadbalancer`.`group_access`,`loadbalancer`.`permissions_owner`,`loadbalancer`.`permissions_owner_access`,`loadbalancer`.`other_access`,`loadbalancer`.`group`,`loadbalancer`.`enable`,`loadbalancer`.`description`,`loadbalancer`.`display_name`,`loadbalancer`.`key_value_pair` from `loadbalancer`"
const showLoadbalancerQuery = "select `loadbalancer`.`global_access`,`loadbalancer`.`share`,`loadbalancer`.`owner`,`loadbalancer`.`owner_access`,`loadbalancer`.`vip_subnet_id`,`loadbalancer`.`operating_status`,`loadbalancer`.`status`,`loadbalancer`.`provisioning_status`,`loadbalancer`.`admin_state`,`loadbalancer`.`vip_address`,`loadbalancer`.`loadbalancer_provider`,`loadbalancer`.`uuid`,`loadbalancer`.`fq_name`,`loadbalancer`.`created`,`loadbalancer`.`creator`,`loadbalancer`.`user_visible`,`loadbalancer`.`last_modified`,`loadbalancer`.`group_access`,`loadbalancer`.`permissions_owner`,`loadbalancer`.`permissions_owner_access`,`loadbalancer`.`other_access`,`loadbalancer`.`group`,`loadbalancer`.`enable`,`loadbalancer`.`description`,`loadbalancer`.`display_name`,`loadbalancer`.`key_value_pair` from `loadbalancer` where uuid = ?"

const insertLoadbalancerVirtualMachineInterfaceQuery = "insert into `ref_loadbalancer_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerServiceInstanceQuery = "insert into `ref_loadbalancer_service_instance` (`from`, `to` ) values (?, ?);"

const insertLoadbalancerServiceApplianceSetQuery = "insert into `ref_loadbalancer_service_appliance_set` (`from`, `to` ) values (?, ?);"

func CreateLoadbalancer(tx *sql.Tx, model *models.Loadbalancer) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.LoadbalancerProperties.VipSubnetID),
		string(model.LoadbalancerProperties.OperatingStatus),
		string(model.LoadbalancerProperties.Status),
		string(model.LoadbalancerProperties.ProvisioningStatus),
		bool(model.LoadbalancerProperties.AdminState),
		string(model.LoadbalancerProperties.VipAddress),
		string(model.LoadbalancerProvider),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtServiceInstanceRef, err := tx.Prepare(insertLoadbalancerServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceApplianceSetRef, err := tx.Prepare(insertLoadbalancerServiceApplianceSetQuery)
	if err != nil {
		return err
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {
		_, err = stmtServiceApplianceSetRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLoadbalancerVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanLoadbalancer(rows *sql.Rows) (*models.Loadbalancer, error) {
	m := models.MakeLoadbalancer()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.LoadbalancerProperties.VipSubnetID,
		&m.LoadbalancerProperties.OperatingStatus,
		&m.LoadbalancerProperties.Status,
		&m.LoadbalancerProperties.ProvisioningStatus,
		&m.LoadbalancerProperties.AdminState,
		&m.LoadbalancerProperties.VipAddress,
		&m.LoadbalancerProvider,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildLoadbalancerWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["vip_subnet_id"]; ok {
		results = append(results, "vip_subnet_id = ?")
		values = append(values, value)
	}

	if value, ok := where["operating_status"]; ok {
		results = append(results, "operating_status = ?")
		values = append(values, value)
	}

	if value, ok := where["status"]; ok {
		results = append(results, "status = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_status"]; ok {
		results = append(results, "provisioning_status = ?")
		values = append(values, value)
	}

	if value, ok := where["vip_address"]; ok {
		results = append(results, "vip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["loadbalancer_provider"]; ok {
		results = append(results, "loadbalancer_provider = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLoadbalancer(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Loadbalancer, error) {
	result := models.MakeLoadbalancerSlice()
	whereQuery, values := buildLoadbalancerWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLoadbalancerQuery)
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
		m, _ := scanLoadbalancer(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLoadbalancer(tx *sql.Tx, uuid string) (*models.Loadbalancer, error) {
	rows, err := tx.Query(showLoadbalancerQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLoadbalancer(rows)
	}
	return nil, nil
}

func UpdateLoadbalancer(tx *sql.Tx, uuid string, model *models.Loadbalancer) error {
	return nil
}

func DeleteLoadbalancer(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLoadbalancerQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
