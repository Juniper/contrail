package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertDomainQuery = "insert into `domain` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`virtual_network_limit`,`security_group_limit`,`project_limit`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteDomainQuery = "delete from `domain` where uuid = ?"

// DomainFields is db columns for Domain
var DomainFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"virtual_network_limit",
	"security_group_limit",
	"project_limit",
	"display_name",
	"key_value_pair",
}

// DomainRefFields is db reference fields for Domain
var DomainRefFields = map[string][]string{}

// DomainBackRefFields is db back reference fields for Domain
var DomainBackRefFields = map[string][]string{

	"api_access_list": []string{
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"rbac_rule",
		"key_value_pair",
	},

	"namespace": []string{
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"ip_prefix_len",
		"ip_prefix",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
	},

	"project": []string{
		"vxlan_routing",
		"uuid",
		"virtual_router",
		"virtual_network",
		"virtual_machine_interface",
		"virtual_ip",
		"virtual_DNS_record",
		"virtual_DNS",
		"subnet",
		"service_template",
		"service_instance",
		"security_logging_object",
		"security_group_rule",
		"security_group",
		"route_table",
		"network_policy",
		"network_ipam",
		"logical_router",
		"loadbalancer_pool",
		"loadbalancer_member",
		"loadbalancer_healthmonitor",
		"instance_ip",
		"global_vrouter_config",
		"floating_ip_pool",
		"floating_ip",
		"defaults",
		"bgp_router",
		"access_control_list",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
		"alarm_enable",
	},

	"service_template": []string{
		"uuid",
		"vrouter_instance_type",
		"version",
		"service_virtualization_type",
		"service_type",
		"service_scaling",
		"service_mode",
		"ordered_interfaces",
		"interface_type",
		"instance_data",
		"image_name",
		"flavor",
		"availability_zone_enable",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
	},

	"virtual_DNS": []string{
		"reverse_resolution",
		"record_order",
		"next_virtual_DNS",
		"floating_ip_record",
		"external_visible",
		"dynamic_records_from_client",
		"domain_name",
		"default_ttl_seconds",
		"uuid",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"user_visible",
		"permissions_owner_access",
		"permissions_owner",
		"other_access",
		"group_access",
		"group",
		"last_modified",
		"enable",
		"description",
		"creator",
		"created",
		"fq_name",
		"display_name",
		"key_value_pair",
	},
}

// DomainParentTypes is possible parents for Domain
var DomainParents = []string{

	"config_root",
}

// CreateDomain inserts Domain to DB
func (db *DB) createDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.Domain
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertDomainQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		int(model.GetDomainLimits().GetVirtualNetworkLimit()),
		int(model.GetDomainLimits().GetSecurityGroupLimit()),
		int(model.GetDomainLimits().GetProjectLimit()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "domain",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "domain", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanDomain(values map[string]interface{}) (*models.Domain, error) {
	m := models.MakeDomain()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["virtual_network_limit"]; ok {

		m.DomainLimits.VirtualNetworkLimit = schema.InterfaceToInt64(value)

	}

	if value, ok := values["security_group_limit"]; ok {

		m.DomainLimits.SecurityGroupLimit = schema.InterfaceToInt64(value)

	}

	if value, ok := values["project_limit"]; ok {

		m.DomainLimits.ProjectLimit = schema.InterfaceToInt64(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_api_access_list"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeAPIAccessList()
			m.APIAccessLists = append(m.APIAccessLists, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["rbac_rule"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.APIAccessListEntries.RbacRule)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_namespace"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeNamespace()
			m.Namespaces = append(m.Namespaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {

				childModel.NamespaceCidr.IPPrefixLen = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {

				childModel.NamespaceCidr.IPPrefix = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_project"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeProject()
			m.Projects = append(m.Projects, childModel)

			if propertyValue, ok := childResourceMap["vxlan_routing"]; ok && propertyValue != nil {

				childModel.VxlanRouting = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_router"]; ok && propertyValue != nil {

				childModel.Quota.VirtualRouter = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_network"]; ok && propertyValue != nil {

				childModel.Quota.VirtualNetwork = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface"]; ok && propertyValue != nil {

				childModel.Quota.VirtualMachineInterface = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_ip"]; ok && propertyValue != nil {

				childModel.Quota.VirtualIP = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_DNS_record"]; ok && propertyValue != nil {

				childModel.Quota.VirtualDNSRecord = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_DNS"]; ok && propertyValue != nil {

				childModel.Quota.VirtualDNS = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["subnet"]; ok && propertyValue != nil {

				childModel.Quota.Subnet = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_template"]; ok && propertyValue != nil {

				childModel.Quota.ServiceTemplate = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_instance"]; ok && propertyValue != nil {

				childModel.Quota.ServiceInstance = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_logging_object"]; ok && propertyValue != nil {

				childModel.Quota.SecurityLoggingObject = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_group_rule"]; ok && propertyValue != nil {

				childModel.Quota.SecurityGroupRule = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_group"]; ok && propertyValue != nil {

				childModel.Quota.SecurityGroup = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["route_table"]; ok && propertyValue != nil {

				childModel.Quota.RouteTable = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["network_policy"]; ok && propertyValue != nil {

				childModel.Quota.NetworkPolicy = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["network_ipam"]; ok && propertyValue != nil {

				childModel.Quota.NetworkIpam = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["logical_router"]; ok && propertyValue != nil {

				childModel.Quota.LogicalRouter = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_pool"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerPool = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_member"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerMember = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_healthmonitor"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerHealthmonitor = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["instance_ip"]; ok && propertyValue != nil {

				childModel.Quota.InstanceIP = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_vrouter_config"]; ok && propertyValue != nil {

				childModel.Quota.GlobalVrouterConfig = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_pool"]; ok && propertyValue != nil {

				childModel.Quota.FloatingIPPool = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip"]; ok && propertyValue != nil {

				childModel.Quota.FloatingIP = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["defaults"]; ok && propertyValue != nil {

				childModel.Quota.Defaults = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["bgp_router"]; ok && propertyValue != nil {

				childModel.Quota.BGPRouter = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["access_control_list"]; ok && propertyValue != nil {

				childModel.Quota.AccessControlList = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["alarm_enable"]; ok && propertyValue != nil {

				childModel.AlarmEnable = schema.InterfaceToBool(propertyValue)

			}

		}
	}

	if value, ok := values["backref_service_template"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeServiceTemplate()
			m.ServiceTemplates = append(m.ServiceTemplates, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["vrouter_instance_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.VrouterInstanceType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.Version = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_virtualization_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceVirtualizationType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_scaling"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceScaling = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_mode"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceMode = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ordered_interfaces"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.OrderedInterfaces = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["interface_type"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.ServiceTemplateProperties.InterfaceType)

			}

			if propertyValue, ok := childResourceMap["instance_data"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.InstanceData = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["image_name"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ImageName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["flavor"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.Flavor = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["availability_zone_enable"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.AvailabilityZoneEnable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_virtual_DNS"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeVirtualDNS()
			m.VirtualDNSs = append(m.VirtualDNSs, childModel)

			if propertyValue, ok := childResourceMap["reverse_resolution"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.ReverseResolution = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_order"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.RecordOrder = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["next_virtual_DNS"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.NextVirtualDNS = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_record"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.FloatingIPRecord = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["external_visible"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.ExternalVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["dynamic_records_from_client"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DynamicRecordsFromClient = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["domain_name"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DomainName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["default_ttl_seconds"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DefaultTTLSeconds = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListDomain lists Domain with list spec.
func (db *DB) listDomain(ctx context.Context, request *models.ListDomainRequest) (response *models.ListDomainResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "domain"
	qb.Fields = DomainFields
	qb.RefFields = DomainRefFields
	qb.BackRefFields = DomainBackRefFields
	result := []*models.Domain{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
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
		m, err := scanDomain(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListDomainResponse{
		Domains: result,
	}
	return response, nil
}

// UpdateDomain updates a resource
func (db *DB) updateDomain(
	ctx context.Context,
	request *models.UpdateDomainRequest,
) error {
	//TODO
	return nil
}

// DeleteDomain deletes a resource
func (db *DB) deleteDomain(
	ctx context.Context,
	request *models.DeleteDomainRequest) error {
	deleteQuery := deleteDomainQuery
	selectQuery := "select count(uuid) from domain where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateDomain handle a Create API
func (db *DB) CreateDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) (*models.CreateDomainResponse, error) {
	model := request.Domain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createDomain(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateDomainResponse{
		Domain: request.Domain,
	}, nil
}

//UpdateDomain handles a Update request.
func (db *DB) UpdateDomain(
	ctx context.Context,
	request *models.UpdateDomainRequest) (*models.UpdateDomainResponse, error) {
	model := request.Domain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateDomain(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "domain",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateDomainResponse{
		Domain: model,
	}, nil
}

//DeleteDomain delete a resource.
func (db *DB) DeleteDomain(ctx context.Context, request *models.DeleteDomainRequest) (*models.DeleteDomainResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteDomain(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteDomainResponse{
		ID: request.ID,
	}, nil
}

//GetDomain a Get request.
func (db *DB) GetDomain(ctx context.Context, request *models.GetDomainRequest) (response *models.GetDomainResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListDomainRequest{
		Spec: spec,
	}
	var result *models.ListDomainResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listDomain(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Domains) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetDomainResponse{
		Domain: result.Domains[0],
	}
	return response, nil
}

//ListDomain handles a List service Request.
func (db *DB) ListDomain(
	ctx context.Context,
	request *models.ListDomainRequest) (response *models.ListDomainResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listDomain(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
