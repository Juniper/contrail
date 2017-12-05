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

const insertInstanceIPQuery = "insert into `instance_ip` (`ip_prefix`,`ip_prefix_len`,`instance_ip_address`,`fq_name`,`owner_access`,`global_access`,`share`,`owner`,`instance_ip_mode`,`service_instance_ip`,`uuid`,`service_health_check_ip`,`subnet_uuid`,`instance_ip_local_ip`,`instance_ip_secondary`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`instance_ip_family`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateInstanceIPQuery = "update `instance_ip` set `ip_prefix` = ?,`ip_prefix_len` = ?,`instance_ip_address` = ?,`fq_name` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`instance_ip_mode` = ?,`service_instance_ip` = ?,`uuid` = ?,`service_health_check_ip` = ?,`subnet_uuid` = ?,`instance_ip_local_ip` = ?,`instance_ip_secondary` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`instance_ip_family` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"
const listInstanceIPQuery = "select `instance_ip`.`ip_prefix`,`instance_ip`.`ip_prefix_len`,`instance_ip`.`instance_ip_address`,`instance_ip`.`fq_name`,`instance_ip`.`owner_access`,`instance_ip`.`global_access`,`instance_ip`.`share`,`instance_ip`.`owner`,`instance_ip`.`instance_ip_mode`,`instance_ip`.`service_instance_ip`,`instance_ip`.`uuid`,`instance_ip`.`service_health_check_ip`,`instance_ip`.`subnet_uuid`,`instance_ip`.`instance_ip_local_ip`,`instance_ip`.`instance_ip_secondary`,`instance_ip`.`last_modified`,`instance_ip`.`group_access`,`instance_ip`.`permissions_owner`,`instance_ip`.`permissions_owner_access`,`instance_ip`.`other_access`,`instance_ip`.`group`,`instance_ip`.`enable`,`instance_ip`.`description`,`instance_ip`.`created`,`instance_ip`.`creator`,`instance_ip`.`user_visible`,`instance_ip`.`instance_ip_family`,`instance_ip`.`display_name`,`instance_ip`.`key_value_pair` from `instance_ip`"
const showInstanceIPQuery = "select `instance_ip`.`ip_prefix`,`instance_ip`.`ip_prefix_len`,`instance_ip`.`instance_ip_address`,`instance_ip`.`fq_name`,`instance_ip`.`owner_access`,`instance_ip`.`global_access`,`instance_ip`.`share`,`instance_ip`.`owner`,`instance_ip`.`instance_ip_mode`,`instance_ip`.`service_instance_ip`,`instance_ip`.`uuid`,`instance_ip`.`service_health_check_ip`,`instance_ip`.`subnet_uuid`,`instance_ip`.`instance_ip_local_ip`,`instance_ip`.`instance_ip_secondary`,`instance_ip`.`last_modified`,`instance_ip`.`group_access`,`instance_ip`.`permissions_owner`,`instance_ip`.`permissions_owner_access`,`instance_ip`.`other_access`,`instance_ip`.`group`,`instance_ip`.`enable`,`instance_ip`.`description`,`instance_ip`.`created`,`instance_ip`.`creator`,`instance_ip`.`user_visible`,`instance_ip`.`instance_ip_family`,`instance_ip`.`display_name`,`instance_ip`.`key_value_pair` from `instance_ip` where uuid = ?"

const insertInstanceIPPhysicalRouterQuery = "insert into `ref_instance_ip_physical_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualRouterQuery = "insert into `ref_instance_ip_virtual_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPNetworkIpamQuery = "insert into `ref_instance_ip_network_ipam` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualNetworkQuery = "insert into `ref_instance_ip_virtual_network` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualMachineInterfaceQuery = "insert into `ref_instance_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

func CreateInstanceIP(tx *sql.Tx, model *models.InstanceIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.SecondaryIPTrackingIP.IPPrefix),
		int(model.SecondaryIPTrackingIP.IPPrefixLen),
		string(model.InstanceIPAddress),
		utils.MustJSON(model.FQName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.InstanceIPMode),
		bool(model.ServiceInstanceIP),
		string(model.UUID),
		bool(model.ServiceHealthCheckIP),
		string(model.SubnetUUID),
		bool(model.InstanceIPLocalIP),
		bool(model.InstanceIPSecondary),
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
		string(model.InstanceIPFamily),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtNetworkIpamRef, err := tx.Prepare(insertInstanceIPNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {
		_, err = stmtNetworkIpamRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertInstanceIPVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertInstanceIPVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertInstanceIPPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {
		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualRouterRef, err := tx.Prepare(insertInstanceIPVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {
		_, err = stmtVirtualRouterRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanInstanceIP(rows *sql.Rows) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	var jsonFQName string

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.SecondaryIPTrackingIP.IPPrefix,
		&m.SecondaryIPTrackingIP.IPPrefixLen,
		&m.InstanceIPAddress,
		&jsonFQName,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.InstanceIPMode,
		&m.ServiceInstanceIP,
		&m.UUID,
		&m.ServiceHealthCheckIP,
		&m.SubnetUUID,
		&m.InstanceIPLocalIP,
		&m.InstanceIPSecondary,
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
		&m.InstanceIPFamily,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildInstanceIPWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_ip_address"]; ok {
		results = append(results, "instance_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_ip_mode"]; ok {
		results = append(results, "instance_ip_mode = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["subnet_uuid"]; ok {
		results = append(results, "subnet_uuid = ?")
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

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["instance_ip_family"]; ok {
		results = append(results, "instance_ip_family = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListInstanceIP(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.InstanceIP, error) {
	result := models.MakeInstanceIPSlice()
	whereQuery, values := buildInstanceIPWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listInstanceIPQuery)
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
		m, _ := scanInstanceIP(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowInstanceIP(tx *sql.Tx, uuid string) (*models.InstanceIP, error) {
	rows, err := tx.Query(showInstanceIPQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanInstanceIP(rows)
	}
	return nil, nil
}

func UpdateInstanceIP(tx *sql.Tx, uuid string, model *models.InstanceIP) error {
	return nil
}

func DeleteInstanceIP(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteInstanceIPQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
