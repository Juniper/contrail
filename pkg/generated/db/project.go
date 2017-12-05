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

const insertProjectQuery = "insert into `project` (`vxlan_routing`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`display_name`,`alarm_enable`,`network_ipam`,`bgp_router`,`virtual_ip`,`virtual_router`,`access_control_list`,`loadbalancer_healthmonitor`,`route_table`,`subnet`,`virtual_DNS_record`,`service_instance`,`floating_ip_pool`,`logical_router`,`instance_ip`,`global_vrouter_config`,`security_group_rule`,`loadbalancer_member`,`virtual_network`,`virtual_DNS`,`security_group`,`floating_ip`,`virtual_machine_interface`,`network_policy`,`security_logging_object`,`defaults`,`loadbalancer_pool`,`service_template`,`fq_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProjectQuery = "update `project` set `vxlan_routing` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`display_name` = ?,`alarm_enable` = ?,`network_ipam` = ?,`bgp_router` = ?,`virtual_ip` = ?,`virtual_router` = ?,`access_control_list` = ?,`loadbalancer_healthmonitor` = ?,`route_table` = ?,`subnet` = ?,`virtual_DNS_record` = ?,`service_instance` = ?,`floating_ip_pool` = ?,`logical_router` = ?,`instance_ip` = ?,`global_vrouter_config` = ?,`security_group_rule` = ?,`loadbalancer_member` = ?,`virtual_network` = ?,`virtual_DNS` = ?,`security_group` = ?,`floating_ip` = ?,`virtual_machine_interface` = ?,`network_policy` = ?,`security_logging_object` = ?,`defaults` = ?,`loadbalancer_pool` = ?,`service_template` = ?,`fq_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?;"
const deleteProjectQuery = "delete from `project` where uuid = ?"
const listProjectQuery = "select `project`.`vxlan_routing`,`project`.`user_visible`,`project`.`last_modified`,`project`.`group_access`,`project`.`owner`,`project`.`owner_access`,`project`.`other_access`,`project`.`group`,`project`.`enable`,`project`.`description`,`project`.`created`,`project`.`creator`,`project`.`display_name`,`project`.`alarm_enable`,`project`.`network_ipam`,`project`.`bgp_router`,`project`.`virtual_ip`,`project`.`virtual_router`,`project`.`access_control_list`,`project`.`loadbalancer_healthmonitor`,`project`.`route_table`,`project`.`subnet`,`project`.`virtual_DNS_record`,`project`.`service_instance`,`project`.`floating_ip_pool`,`project`.`logical_router`,`project`.`instance_ip`,`project`.`global_vrouter_config`,`project`.`security_group_rule`,`project`.`loadbalancer_member`,`project`.`virtual_network`,`project`.`virtual_DNS`,`project`.`security_group`,`project`.`floating_ip`,`project`.`virtual_machine_interface`,`project`.`network_policy`,`project`.`security_logging_object`,`project`.`defaults`,`project`.`loadbalancer_pool`,`project`.`service_template`,`project`.`fq_name`,`project`.`key_value_pair`,`project`.`perms2_owner`,`project`.`perms2_owner_access`,`project`.`global_access`,`project`.`share`,`project`.`uuid` from `project`"
const showProjectQuery = "select `project`.`vxlan_routing`,`project`.`user_visible`,`project`.`last_modified`,`project`.`group_access`,`project`.`owner`,`project`.`owner_access`,`project`.`other_access`,`project`.`group`,`project`.`enable`,`project`.`description`,`project`.`created`,`project`.`creator`,`project`.`display_name`,`project`.`alarm_enable`,`project`.`network_ipam`,`project`.`bgp_router`,`project`.`virtual_ip`,`project`.`virtual_router`,`project`.`access_control_list`,`project`.`loadbalancer_healthmonitor`,`project`.`route_table`,`project`.`subnet`,`project`.`virtual_DNS_record`,`project`.`service_instance`,`project`.`floating_ip_pool`,`project`.`logical_router`,`project`.`instance_ip`,`project`.`global_vrouter_config`,`project`.`security_group_rule`,`project`.`loadbalancer_member`,`project`.`virtual_network`,`project`.`virtual_DNS`,`project`.`security_group`,`project`.`floating_ip`,`project`.`virtual_machine_interface`,`project`.`network_policy`,`project`.`security_logging_object`,`project`.`defaults`,`project`.`loadbalancer_pool`,`project`.`service_template`,`project`.`fq_name`,`project`.`key_value_pair`,`project`.`perms2_owner`,`project`.`perms2_owner_access`,`project`.`global_access`,`project`.`share`,`project`.`uuid` from `project` where uuid = ?"

const insertProjectAliasIPPoolQuery = "insert into `ref_project_alias_ip_pool` (`from`, `to` ) values (?, ?);"

const insertProjectNamespaceQuery = "insert into `ref_project_namespace` (`from`, `to` ,`ip_prefix`,`ip_prefix_len`) values (?, ?,?,?);"

const insertProjectApplicationPolicySetQuery = "insert into `ref_project_application_policy_set` (`from`, `to` ) values (?, ?);"

const insertProjectFloatingIPPoolQuery = "insert into `ref_project_floating_ip_pool` (`from`, `to` ) values (?, ?);"

func CreateProject(tx *sql.Tx, model *models.Project) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(bool(model.VxlanRouting),
		bool(model.IDPerms.UserVisible),
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
		string(model.DisplayName),
		bool(model.AlarmEnable),
		int(model.Quota.NetworkIpam),
		int(model.Quota.BGPRouter),
		int(model.Quota.VirtualIP),
		int(model.Quota.VirtualRouter),
		int(model.Quota.AccessControlList),
		int(model.Quota.LoadbalancerHealthmonitor),
		int(model.Quota.RouteTable),
		int(model.Quota.Subnet),
		int(model.Quota.VirtualDNSRecord),
		int(model.Quota.ServiceInstance),
		int(model.Quota.FloatingIPPool),
		int(model.Quota.LogicalRouter),
		int(model.Quota.InstanceIP),
		int(model.Quota.GlobalVrouterConfig),
		int(model.Quota.SecurityGroupRule),
		int(model.Quota.LoadbalancerMember),
		int(model.Quota.VirtualNetwork),
		int(model.Quota.VirtualDNS),
		int(model.Quota.SecurityGroup),
		int(model.Quota.FloatingIP),
		int(model.Quota.VirtualMachineInterface),
		int(model.Quota.NetworkPolicy),
		int(model.Quota.SecurityLoggingObject),
		int(model.Quota.Defaults),
		int(model.Quota.LoadbalancerPool),
		int(model.Quota.ServiceTemplate),
		utils.MustJSON(model.FQName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.UUID))

	stmtAliasIPPoolRef, err := tx.Prepare(insertProjectAliasIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmtAliasIPPoolRef.Close()
	for _, ref := range model.AliasIPPoolRefs {
		_, err = stmtAliasIPPoolRef.Exec(model.UUID, ref.UUID)
	}

	stmtNamespaceRef, err := tx.Prepare(insertProjectNamespaceQuery)
	if err != nil {
		return err
	}
	defer stmtNamespaceRef.Close()
	for _, ref := range model.NamespaceRefs {
		_, err = stmtNamespaceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.IPPrefix),
			int(ref.Attr.IPPrefixLen))
	}

	stmtApplicationPolicySetRef, err := tx.Prepare(insertProjectApplicationPolicySetQuery)
	if err != nil {
		return err
	}
	defer stmtApplicationPolicySetRef.Close()
	for _, ref := range model.ApplicationPolicySetRefs {
		_, err = stmtApplicationPolicySetRef.Exec(model.UUID, ref.UUID)
	}

	stmtFloatingIPPoolRef, err := tx.Prepare(insertProjectFloatingIPPoolQuery)
	if err != nil {
		return err
	}
	defer stmtFloatingIPPoolRef.Close()
	for _, ref := range model.FloatingIPPoolRefs {
		_, err = stmtFloatingIPPoolRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanProject(rows *sql.Rows) (*models.Project, error) {
	m := models.MakeProject()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&m.VxlanRouting,
		&m.IDPerms.UserVisible,
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
		&m.DisplayName,
		&m.AlarmEnable,
		&m.Quota.NetworkIpam,
		&m.Quota.BGPRouter,
		&m.Quota.VirtualIP,
		&m.Quota.VirtualRouter,
		&m.Quota.AccessControlList,
		&m.Quota.LoadbalancerHealthmonitor,
		&m.Quota.RouteTable,
		&m.Quota.Subnet,
		&m.Quota.VirtualDNSRecord,
		&m.Quota.ServiceInstance,
		&m.Quota.FloatingIPPool,
		&m.Quota.LogicalRouter,
		&m.Quota.InstanceIP,
		&m.Quota.GlobalVrouterConfig,
		&m.Quota.SecurityGroupRule,
		&m.Quota.LoadbalancerMember,
		&m.Quota.VirtualNetwork,
		&m.Quota.VirtualDNS,
		&m.Quota.SecurityGroup,
		&m.Quota.FloatingIP,
		&m.Quota.VirtualMachineInterface,
		&m.Quota.NetworkPolicy,
		&m.Quota.SecurityLoggingObject,
		&m.Quota.Defaults,
		&m.Quota.LoadbalancerPool,
		&m.Quota.ServiceTemplate,
		&jsonFQName,
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

func buildProjectWhereQuery(where map[string]interface{}) (string, []interface{}) {
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

	return "where " + strings.Join(results, " and "), values
}

func ListProject(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Project, error) {
	result := models.MakeProjectSlice()
	whereQuery, values := buildProjectWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listProjectQuery)
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
		m, _ := scanProject(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowProject(tx *sql.Tx, uuid string) (*models.Project, error) {
	rows, err := tx.Query(showProjectQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanProject(rows)
	}
	return nil, nil
}

func UpdateProject(tx *sql.Tx, uuid string, model *models.Project) error {
	return nil
}

func DeleteProject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteProjectQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
