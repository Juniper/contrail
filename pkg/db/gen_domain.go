// nolint
package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

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
	"configuration_version",
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
		"configuration_version",
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
		"configuration_version",
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
		"configuration_version",
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
		"configuration_version",
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
		"configuration_version",
		"key_value_pair",
	},
}

// DomainParentTypes is possible parents for Domain
var DomainParents = []string{

	"config_root",
}

// CreateDomain inserts Domain to DB
// nolint
func (db *DB) createDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) error {
	qb := db.queryBuilders["domain"]
	tx := GetTransaction(ctx)
	model := request.Domain
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "domain",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "domain", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = common.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = common.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = common.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = common.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = common.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = common.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = common.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = common.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["virtual_network_limit"]; ok {

		m.DomainLimits.VirtualNetworkLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["security_group_limit"]; ok {

		m.DomainLimits.SecurityGroupLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["project_limit"]; ok {

		m.DomainLimits.ProjectLimit = common.InterfaceToInt64(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_api_access_list"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeAPIAccessList()
			m.APIAccessLists = append(m.APIAccessLists, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["rbac_rule"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.APIAccessListEntries.RbacRule)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_namespace"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeNamespace()
			m.Namespaces = append(m.Namespaces, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {

				childModel.NamespaceCidr.IPPrefixLen = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {

				childModel.NamespaceCidr.IPPrefix = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_project"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeProject()
			m.Projects = append(m.Projects, childModel)

			if propertyValue, ok := childResourceMap["vxlan_routing"]; ok && propertyValue != nil {

				childModel.VxlanRouting = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_router"]; ok && propertyValue != nil {

				childModel.Quota.VirtualRouter = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_network"]; ok && propertyValue != nil {

				childModel.Quota.VirtualNetwork = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface"]; ok && propertyValue != nil {

				childModel.Quota.VirtualMachineInterface = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_ip"]; ok && propertyValue != nil {

				childModel.Quota.VirtualIP = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_DNS_record"]; ok && propertyValue != nil {

				childModel.Quota.VirtualDNSRecord = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["virtual_DNS"]; ok && propertyValue != nil {

				childModel.Quota.VirtualDNS = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["subnet"]; ok && propertyValue != nil {

				childModel.Quota.Subnet = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_template"]; ok && propertyValue != nil {

				childModel.Quota.ServiceTemplate = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_instance"]; ok && propertyValue != nil {

				childModel.Quota.ServiceInstance = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_logging_object"]; ok && propertyValue != nil {

				childModel.Quota.SecurityLoggingObject = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_group_rule"]; ok && propertyValue != nil {

				childModel.Quota.SecurityGroupRule = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["security_group"]; ok && propertyValue != nil {

				childModel.Quota.SecurityGroup = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["route_table"]; ok && propertyValue != nil {

				childModel.Quota.RouteTable = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["network_policy"]; ok && propertyValue != nil {

				childModel.Quota.NetworkPolicy = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["network_ipam"]; ok && propertyValue != nil {

				childModel.Quota.NetworkIpam = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["logical_router"]; ok && propertyValue != nil {

				childModel.Quota.LogicalRouter = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_pool"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerPool = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_member"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerMember = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["loadbalancer_healthmonitor"]; ok && propertyValue != nil {

				childModel.Quota.LoadbalancerHealthmonitor = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["instance_ip"]; ok && propertyValue != nil {

				childModel.Quota.InstanceIP = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_vrouter_config"]; ok && propertyValue != nil {

				childModel.Quota.GlobalVrouterConfig = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_pool"]; ok && propertyValue != nil {

				childModel.Quota.FloatingIPPool = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip"]; ok && propertyValue != nil {

				childModel.Quota.FloatingIP = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["defaults"]; ok && propertyValue != nil {

				childModel.Quota.Defaults = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["bgp_router"]; ok && propertyValue != nil {

				childModel.Quota.BGPRouter = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["access_control_list"]; ok && propertyValue != nil {

				childModel.Quota.AccessControlList = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["alarm_enable"]; ok && propertyValue != nil {

				childModel.AlarmEnable = common.InterfaceToBool(propertyValue)

			}

		}
	}

	if value, ok := values["backref_service_template"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeServiceTemplate()
			m.ServiceTemplates = append(m.ServiceTemplates, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["vrouter_instance_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.VrouterInstanceType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.Version = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_virtualization_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceVirtualizationType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_type"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_scaling"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceScaling = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["service_mode"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ServiceMode = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["ordered_interfaces"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.OrderedInterfaces = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["interface_type"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ServiceTemplateProperties.InterfaceType)

			}

			if propertyValue, ok := childResourceMap["instance_data"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.InstanceData = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["image_name"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.ImageName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["flavor"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.Flavor = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["availability_zone_enable"]; ok && propertyValue != nil {

				childModel.ServiceTemplateProperties.AvailabilityZoneEnable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	if value, ok := values["backref_virtual_DNS"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeVirtualDNS()
			m.VirtualDNSs = append(m.VirtualDNSs, childModel)

			if propertyValue, ok := childResourceMap["reverse_resolution"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.ReverseResolution = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["record_order"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.RecordOrder = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["next_virtual_DNS"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.NextVirtualDNS = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_record"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.FloatingIPRecord = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["external_visible"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.ExternalVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["dynamic_records_from_client"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DynamicRecordsFromClient = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["domain_name"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DomainName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["default_ttl_seconds"]; ok && propertyValue != nil {

				childModel.VirtualDNSData.DefaultTTLSeconds = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = common.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["configuration_version"]; ok && propertyValue != nil {

				childModel.ConfigurationVersion = common.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListDomain lists Domain with list spec.
func (db *DB) listDomain(ctx context.Context, request *models.ListDomainRequest) (response *models.ListDomainResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["domain"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.Domain{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
	qb := db.queryBuilders["domain"]

	selectQuery := qb.SelectForDeleteQuery()
	deleteQuery := qb.DeleteQuery()

	var err error
	var count int
	uuid := request.ID
	tx := GetTransaction(ctx)
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

	err = db.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateDomain handle a Create API
// nolint
func (db *DB) CreateDomain(
	ctx context.Context,
	request *models.CreateDomainRequest) (*models.CreateDomainResponse, error) {
	model := request.Domain
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
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
	if err := DoInTransaction(
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
	if err := DoInTransaction(
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
	if err := DoInTransaction(
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
// nolint
func (db *DB) ListDomain(
	ctx context.Context,
	request *models.ListDomainRequest) (response *models.ListDomainResponse, err error) {
	if err := DoInTransaction(
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
