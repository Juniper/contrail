package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertProjectQuery = "insert into `project` (`vxlan_routing`,`owner_access`,`global_access`,`share`,`owner`,`fq_name`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`alarm_enable`,`loadbalancer_pool`,`service_template`,`network_policy`,`network_ipam`,`bgp_router`,`instance_ip`,`service_instance`,`subnet`,`global_vrouter_config`,`security_group`,`loadbalancer_member`,`floating_ip`,`loadbalancer_healthmonitor`,`logical_router`,`virtual_network`,`virtual_DNS_record`,`security_group_rule`,`virtual_machine_interface`,`route_table`,`virtual_router`,`virtual_DNS`,`virtual_ip`,`access_control_list`,`security_logging_object`,`defaults`,`floating_ip_pool`,`display_name`,`key_value_pair`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProjectQuery = "update `project` set `vxlan_routing` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`fq_name` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`alarm_enable` = ?,`loadbalancer_pool` = ?,`service_template` = ?,`network_policy` = ?,`network_ipam` = ?,`bgp_router` = ?,`instance_ip` = ?,`service_instance` = ?,`subnet` = ?,`global_vrouter_config` = ?,`security_group` = ?,`loadbalancer_member` = ?,`floating_ip` = ?,`loadbalancer_healthmonitor` = ?,`logical_router` = ?,`virtual_network` = ?,`virtual_DNS_record` = ?,`security_group_rule` = ?,`virtual_machine_interface` = ?,`route_table` = ?,`virtual_router` = ?,`virtual_DNS` = ?,`virtual_ip` = ?,`access_control_list` = ?,`security_logging_object` = ?,`defaults` = ?,`floating_ip_pool` = ?,`display_name` = ?,`key_value_pair` = ?,`uuid` = ?;"
const deleteProjectQuery = "delete from `project` where uuid = ?"

// ProjectFields is db columns for Project
var ProjectFields = []string{
	"vxlan_routing",
	"owner_access",
	"global_access",
	"share",
	"owner",
	"fq_name",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"alarm_enable",
	"loadbalancer_pool",
	"service_template",
	"network_policy",
	"network_ipam",
	"bgp_router",
	"instance_ip",
	"service_instance",
	"subnet",
	"global_vrouter_config",
	"security_group",
	"loadbalancer_member",
	"floating_ip",
	"loadbalancer_healthmonitor",
	"logical_router",
	"virtual_network",
	"virtual_DNS_record",
	"security_group_rule",
	"virtual_machine_interface",
	"route_table",
	"virtual_router",
	"virtual_DNS",
	"virtual_ip",
	"access_control_list",
	"security_logging_object",
	"defaults",
	"floating_ip_pool",
	"display_name",
	"key_value_pair",
	"uuid",
}

// ProjectRefFields is db reference fields for Project
var ProjectRefFields = map[string][]string{

	"application_policy_set": {
	// <common.Schema Value>

	},

	"floating_ip_pool": {
	// <common.Schema Value>

	},

	"alias_ip_pool": {
	// <common.Schema Value>

	},

	"namespace": {
		// <common.Schema Value>
		"ip_prefix",
		"ip_prefix_len",
	},
}

const insertProjectAliasIPPoolQuery = "insert into `ref_project_alias_ip_pool` (`from`, `to` ) values (?, ?);"

const insertProjectNamespaceQuery = "insert into `ref_project_namespace` (`from`, `to` ,`ip_prefix`,`ip_prefix_len`) values (?, ?,?,?);"

const insertProjectApplicationPolicySetQuery = "insert into `ref_project_application_policy_set` (`from`, `to` ) values (?, ?);"

const insertProjectFloatingIPPoolQuery = "insert into `ref_project_floating_ip_pool` (`from`, `to` ) values (?, ?);"

// CreateProject inserts Project to DB
func CreateProject(tx *sql.Tx, model *models.Project) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertProjectQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.VxlanRouting),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		common.MustJSON(model.FQName),
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
		string(model.IDPerms.LastModified),
		bool(model.AlarmEnable),
		int(model.Quota.LoadbalancerPool),
		int(model.Quota.ServiceTemplate),
		int(model.Quota.NetworkPolicy),
		int(model.Quota.NetworkIpam),
		int(model.Quota.BGPRouter),
		int(model.Quota.InstanceIP),
		int(model.Quota.ServiceInstance),
		int(model.Quota.Subnet),
		int(model.Quota.GlobalVrouterConfig),
		int(model.Quota.SecurityGroup),
		int(model.Quota.LoadbalancerMember),
		int(model.Quota.FloatingIP),
		int(model.Quota.LoadbalancerHealthmonitor),
		int(model.Quota.LogicalRouter),
		int(model.Quota.VirtualNetwork),
		int(model.Quota.VirtualDNSRecord),
		int(model.Quota.SecurityGroupRule),
		int(model.Quota.VirtualMachineInterface),
		int(model.Quota.RouteTable),
		int(model.Quota.VirtualRouter),
		int(model.Quota.VirtualDNS),
		int(model.Quota.VirtualIP),
		int(model.Quota.AccessControlList),
		int(model.Quota.SecurityLoggingObject),
		int(model.Quota.Defaults),
		int(model.Quota.FloatingIPPool),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtAliasIPPoolRef, err := tx.Prepare(insertProjectAliasIPPoolQuery)
	if err != nil {
		return errors.Wrap(err, "preparing AliasIPPoolRefs create statement failed")
	}
	defer stmtAliasIPPoolRef.Close()
	for _, ref := range model.AliasIPPoolRefs {

		_, err = stmtAliasIPPoolRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "AliasIPPoolRefs create failed")
		}
	}

	stmtNamespaceRef, err := tx.Prepare(insertProjectNamespaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NamespaceRefs create statement failed")
	}
	defer stmtNamespaceRef.Close()
	for _, ref := range model.NamespaceRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeSubnetType()
		}

		_, err = stmtNamespaceRef.Exec(model.UUID, ref.UUID, string(ref.Attr.IPPrefix),
			int(ref.Attr.IPPrefixLen))
		if err != nil {
			return errors.Wrap(err, "NamespaceRefs create failed")
		}
	}

	stmtApplicationPolicySetRef, err := tx.Prepare(insertProjectApplicationPolicySetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ApplicationPolicySetRefs create statement failed")
	}
	defer stmtApplicationPolicySetRef.Close()
	for _, ref := range model.ApplicationPolicySetRefs {

		_, err = stmtApplicationPolicySetRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ApplicationPolicySetRefs create failed")
		}
	}

	stmtFloatingIPPoolRef, err := tx.Prepare(insertProjectFloatingIPPoolQuery)
	if err != nil {
		return errors.Wrap(err, "preparing FloatingIPPoolRefs create statement failed")
	}
	defer stmtFloatingIPPoolRef.Close()
	for _, ref := range model.FloatingIPPoolRefs {

		_, err = stmtFloatingIPPoolRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "FloatingIPPoolRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanProject(values map[string]interface{}) (*models.Project, error) {
	m := models.MakeProject()

	if value, ok := values["vxlan_routing"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VxlanRouting = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["alarm_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.AlarmEnable = castedValue

	}

	if value, ok := values["loadbalancer_pool"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.LoadbalancerPool = castedValue

	}

	if value, ok := values["service_template"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.ServiceTemplate = castedValue

	}

	if value, ok := values["network_policy"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.NetworkPolicy = castedValue

	}

	if value, ok := values["network_ipam"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.NetworkIpam = castedValue

	}

	if value, ok := values["bgp_router"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.BGPRouter = castedValue

	}

	if value, ok := values["instance_ip"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.InstanceIP = castedValue

	}

	if value, ok := values["service_instance"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.ServiceInstance = castedValue

	}

	if value, ok := values["subnet"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.Subnet = castedValue

	}

	if value, ok := values["global_vrouter_config"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.GlobalVrouterConfig = castedValue

	}

	if value, ok := values["security_group"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.SecurityGroup = castedValue

	}

	if value, ok := values["loadbalancer_member"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.LoadbalancerMember = castedValue

	}

	if value, ok := values["floating_ip"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.FloatingIP = castedValue

	}

	if value, ok := values["loadbalancer_healthmonitor"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.LoadbalancerHealthmonitor = castedValue

	}

	if value, ok := values["logical_router"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.LogicalRouter = castedValue

	}

	if value, ok := values["virtual_network"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualNetwork = castedValue

	}

	if value, ok := values["virtual_DNS_record"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualDNSRecord = castedValue

	}

	if value, ok := values["security_group_rule"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.SecurityGroupRule = castedValue

	}

	if value, ok := values["virtual_machine_interface"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualMachineInterface = castedValue

	}

	if value, ok := values["route_table"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.RouteTable = castedValue

	}

	if value, ok := values["virtual_router"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualRouter = castedValue

	}

	if value, ok := values["virtual_DNS"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualDNS = castedValue

	}

	if value, ok := values["virtual_ip"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.VirtualIP = castedValue

	}

	if value, ok := values["access_control_list"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.AccessControlList = castedValue

	}

	if value, ok := values["security_logging_object"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.SecurityLoggingObject = castedValue

	}

	if value, ok := values["defaults"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.Defaults = castedValue

	}

	if value, ok := values["floating_ip_pool"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Quota.FloatingIPPool = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["ref_namespace"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ProjectNamespaceRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.NamespaceRefs = append(m.NamespaceRefs, referenceModel)

			attr := models.MakeSubnetType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_application_policy_set"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ProjectApplicationPolicySetRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ApplicationPolicySetRefs = append(m.ApplicationPolicySetRefs, referenceModel)

		}
	}

	if value, ok := values["ref_floating_ip_pool"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ProjectFloatingIPPoolRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.FloatingIPPoolRefs = append(m.FloatingIPPoolRefs, referenceModel)

		}
	}

	if value, ok := values["ref_alias_ip_pool"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.ProjectAliasIPPoolRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.AliasIPPoolRefs = append(m.AliasIPPoolRefs, referenceModel)

		}
	}

	return m, nil
}

// ListProject lists Project with list spec.
func ListProject(tx *sql.Tx, spec *common.ListSpec) ([]*models.Project, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "project"
	spec.Fields = ProjectFields
	spec.RefFields = ProjectRefFields
	result := models.MakeProjectSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanProject(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowProject shows Project resource
func ShowProject(tx *sql.Tx, uuid string) (*models.Project, error) {
	list, err := ListProject(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateProject updates a resource
func UpdateProject(tx *sql.Tx, uuid string, model *models.Project) error {
	//TODO(nati) support update
	return nil
}

// DeleteProject deletes a resource
func DeleteProject(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteProjectQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
