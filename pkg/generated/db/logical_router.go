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

const insertLogicalRouterQuery = "insert into `logical_router` (`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`display_name`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`uuid`,`fq_name`,`vxlan_network_identifier`,`route_target`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLogicalRouterQuery = "update `logical_router` set `user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`uuid` = ?,`fq_name` = ?,`vxlan_network_identifier` = ?,`route_target` = ?;"
const deleteLogicalRouterQuery = "delete from `logical_router` where uuid = ?"
const listLogicalRouterQuery = "select `logical_router`.`user_visible`,`logical_router`.`last_modified`,`logical_router`.`owner`,`logical_router`.`owner_access`,`logical_router`.`other_access`,`logical_router`.`group`,`logical_router`.`group_access`,`logical_router`.`enable`,`logical_router`.`description`,`logical_router`.`created`,`logical_router`.`creator`,`logical_router`.`display_name`,`logical_router`.`key_value_pair`,`logical_router`.`global_access`,`logical_router`.`share`,`logical_router`.`perms2_owner`,`logical_router`.`perms2_owner_access`,`logical_router`.`uuid`,`logical_router`.`fq_name`,`logical_router`.`vxlan_network_identifier`,`logical_router`.`route_target` from `logical_router`"
const showLogicalRouterQuery = "select `logical_router`.`user_visible`,`logical_router`.`last_modified`,`logical_router`.`owner`,`logical_router`.`owner_access`,`logical_router`.`other_access`,`logical_router`.`group`,`logical_router`.`group_access`,`logical_router`.`enable`,`logical_router`.`description`,`logical_router`.`created`,`logical_router`.`creator`,`logical_router`.`display_name`,`logical_router`.`key_value_pair`,`logical_router`.`global_access`,`logical_router`.`share`,`logical_router`.`perms2_owner`,`logical_router`.`perms2_owner_access`,`logical_router`.`uuid`,`logical_router`.`fq_name`,`logical_router`.`vxlan_network_identifier`,`logical_router`.`route_target` from `logical_router` where uuid = ?"

const insertLogicalRouterPhysicalRouterQuery = "insert into `ref_logical_router_physical_router` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterBGPVPNQuery = "insert into `ref_logical_router_bgpvpn` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterRouteTargetQuery = "insert into `ref_logical_router_route_target` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterVirtualMachineInterfaceQuery = "insert into `ref_logical_router_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterServiceInstanceQuery = "insert into `ref_logical_router_service_instance` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterRouteTableQuery = "insert into `ref_logical_router_route_table` (`from`, `to` ) values (?, ?);"

const insertLogicalRouterVirtualNetworkQuery = "insert into `ref_logical_router_virtual_network` (`from`, `to` ) values (?, ?);"

func CreateLogicalRouter(tx *sql.Tx, model *models.LogicalRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLogicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.VxlanNetworkIdentifier),
		utils.MustJSON(model.ConfiguredRouteTargetList.RouteTarget))

	stmtRouteTargetRef, err := tx.Prepare(insertLogicalRouterRouteTargetQuery)
	if err != nil {
		return err
	}
	defer stmtRouteTargetRef.Close()
	for _, ref := range model.RouteTargetRefs {
		_, err = stmtRouteTargetRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertLogicalRouterVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtServiceInstanceRef, err := tx.Prepare(insertLogicalRouterServiceInstanceQuery)
	if err != nil {
		return err
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {
		_, err = stmtServiceInstanceRef.Exec(model.UUID, ref.UUID)
	}

	stmtRouteTableRef, err := tx.Prepare(insertLogicalRouterRouteTableQuery)
	if err != nil {
		return err
	}
	defer stmtRouteTableRef.Close()
	for _, ref := range model.RouteTableRefs {
		_, err = stmtRouteTableRef.Exec(model.UUID, ref.UUID)
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertLogicalRouterVirtualNetworkQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {
		_, err = stmtVirtualNetworkRef.Exec(model.UUID, ref.UUID)
	}

	stmtPhysicalRouterRef, err := tx.Prepare(insertLogicalRouterPhysicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {
		_, err = stmtPhysicalRouterRef.Exec(model.UUID, ref.UUID)
	}

	stmtBGPVPNRef, err := tx.Prepare(insertLogicalRouterBGPVPNQuery)
	if err != nil {
		return err
	}
	defer stmtBGPVPNRef.Close()
	for _, ref := range model.BGPVPNRefs {
		_, err = stmtBGPVPNRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanLogicalRouter(rows *sql.Rows) (*models.LogicalRouter, error) {
	m := models.MakeLogicalRouter()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	var jsonConfiguredRouteTargetListRouteTarget string

	if err := rows.Scan(&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.UUID,
		&jsonFQName,
		&m.VxlanNetworkIdentifier,
		&jsonConfiguredRouteTargetListRouteTarget); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonConfiguredRouteTargetListRouteTarget), &m.ConfiguredRouteTargetList.RouteTarget)

	return m, nil
}

func buildLogicalRouterWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	if value, ok := where["vxlan_network_identifier"]; ok {
		results = append(results, "vxlan_network_identifier = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLogicalRouter(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.LogicalRouter, error) {
	result := models.MakeLogicalRouterSlice()
	whereQuery, values := buildLogicalRouterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLogicalRouterQuery)
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
		m, _ := scanLogicalRouter(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLogicalRouter(tx *sql.Tx, uuid string) (*models.LogicalRouter, error) {
	rows, err := tx.Query(showLogicalRouterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLogicalRouter(rows)
	}
	return nil, nil
}

func UpdateLogicalRouter(tx *sql.Tx, uuid string, model *models.LogicalRouter) error {
	return nil
}

func DeleteLogicalRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLogicalRouterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
