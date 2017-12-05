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

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`bgpaas_session_attributes`,`bgpaas_suppress_route_advertisement`,`bgpaas_ip_address`,`display_name`,`global_access`,`share`,`owner`,`owner_access`,`key_value_pair`,`uuid`,`bgpaas_shared`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`autonomous_system`,`fq_name`,`creator`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`enable`,`description`,`created`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateBGPAsAServiceQuery = "update `bgp_as_a_service` set `bgpaas_session_attributes` = ?,`bgpaas_suppress_route_advertisement` = ?,`bgpaas_ip_address` = ?,`display_name` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`key_value_pair` = ?,`uuid` = ?,`bgpaas_shared` = ?,`bgpaas_ipv4_mapped_ipv6_nexthop` = ?,`autonomous_system` = ?,`fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?;"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"
const listBGPAsAServiceQuery = "select `bgp_as_a_service`.`bgpaas_session_attributes`,`bgp_as_a_service`.`bgpaas_suppress_route_advertisement`,`bgp_as_a_service`.`bgpaas_ip_address`,`bgp_as_a_service`.`display_name`,`bgp_as_a_service`.`global_access`,`bgp_as_a_service`.`share`,`bgp_as_a_service`.`owner`,`bgp_as_a_service`.`owner_access`,`bgp_as_a_service`.`key_value_pair`,`bgp_as_a_service`.`uuid`,`bgp_as_a_service`.`bgpaas_shared`,`bgp_as_a_service`.`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgp_as_a_service`.`autonomous_system`,`bgp_as_a_service`.`fq_name`,`bgp_as_a_service`.`creator`,`bgp_as_a_service`.`user_visible`,`bgp_as_a_service`.`last_modified`,`bgp_as_a_service`.`other_access`,`bgp_as_a_service`.`group`,`bgp_as_a_service`.`group_access`,`bgp_as_a_service`.`permissions_owner`,`bgp_as_a_service`.`permissions_owner_access`,`bgp_as_a_service`.`enable`,`bgp_as_a_service`.`description`,`bgp_as_a_service`.`created` from `bgp_as_a_service`"
const showBGPAsAServiceQuery = "select `bgp_as_a_service`.`bgpaas_session_attributes`,`bgp_as_a_service`.`bgpaas_suppress_route_advertisement`,`bgp_as_a_service`.`bgpaas_ip_address`,`bgp_as_a_service`.`display_name`,`bgp_as_a_service`.`global_access`,`bgp_as_a_service`.`share`,`bgp_as_a_service`.`owner`,`bgp_as_a_service`.`owner_access`,`bgp_as_a_service`.`key_value_pair`,`bgp_as_a_service`.`uuid`,`bgp_as_a_service`.`bgpaas_shared`,`bgp_as_a_service`.`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgp_as_a_service`.`autonomous_system`,`bgp_as_a_service`.`fq_name`,`bgp_as_a_service`.`creator`,`bgp_as_a_service`.`user_visible`,`bgp_as_a_service`.`last_modified`,`bgp_as_a_service`.`other_access`,`bgp_as_a_service`.`group`,`bgp_as_a_service`.`group_access`,`bgp_as_a_service`.`permissions_owner`,`bgp_as_a_service`.`permissions_owner_access`,`bgp_as_a_service`.`enable`,`bgp_as_a_service`.`description`,`bgp_as_a_service`.`created` from `bgp_as_a_service` where uuid = ?"

const insertBGPAsAServiceVirtualMachineInterfaceQuery = "insert into `ref_bgp_as_a_service_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertBGPAsAServiceServiceHealthCheckQuery = "insert into `ref_bgp_as_a_service_service_health_check` (`from`, `to` ) values (?, ?);"

func CreateBGPAsAService(tx *sql.Tx, model *models.BGPAsAService) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.BgpaasSessionAttributes),
		bool(model.BgpaasSuppressRouteAdvertisement),
		string(model.BgpaasIPAddress),
		string(model.DisplayName),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		bool(model.BgpaasShared),
		bool(model.BgpaasIpv4MappedIpv6Nexthop),
		int(model.AutonomousSystem),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created))

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertBGPAsAServiceVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertBGPAsAServiceServiceHealthCheckQuery)
	if err != nil {
		return err
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {
		_, err = stmtServiceHealthCheckRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanBGPAsAService(rows *sql.Rows) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

	var jsonPerms2Share string

	var jsonAnnotationsKeyValuePair string

	var jsonFQName string

	if err := rows.Scan(&m.BgpaasSessionAttributes,
		&m.BgpaasSuppressRouteAdvertisement,
		&m.BgpaasIPAddress,
		&m.DisplayName,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&jsonAnnotationsKeyValuePair,
		&m.UUID,
		&m.BgpaasShared,
		&m.BgpaasIpv4MappedIpv6Nexthop,
		&m.AutonomousSystem,
		&jsonFQName,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildBGPAsAServiceWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["bgpaas_session_attributes"]; ok {
		results = append(results, "bgpaas_session_attributes = ?")
		values = append(values, value)
	}

	if value, ok := where["bgpaas_ip_address"]; ok {
		results = append(results, "bgpaas_ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
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

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListBGPAsAService(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.BGPAsAService, error) {
	result := models.MakeBGPAsAServiceSlice()
	whereQuery, values := buildBGPAsAServiceWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listBGPAsAServiceQuery)
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
		m, _ := scanBGPAsAService(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowBGPAsAService(tx *sql.Tx, uuid string) (*models.BGPAsAService, error) {
	rows, err := tx.Query(showBGPAsAServiceQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanBGPAsAService(rows)
	}
	return nil, nil
}

func UpdateBGPAsAService(tx *sql.Tx, uuid string, model *models.BGPAsAService) error {
	return nil
}

func DeleteBGPAsAService(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteBGPAsAServiceQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
