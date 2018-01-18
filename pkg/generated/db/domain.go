package db

import (
	"database/sql"
	"encoding/json"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
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

	"api_access_list": {
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

	"namespace": {
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

	"project": {
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

	"service_template": {
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

	"virtual_DNS": {
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
func CreateDomain(tx *sql.Tx, model *models.Domain) error {
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
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		int(model.DomainLimits.VirtualNetworkLimit),
		int(model.DomainLimits.SecurityGroupLimit),
		int(model.DomainLimits.ProjectLimit),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
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
	err = common.CreateSharing(tx, "domain", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["virtual_network_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.DomainLimits.VirtualNetworkLimit = castedValue

	}

	if value, ok := values["security_group_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.DomainLimits.SecurityGroupLimit = castedValue

	}

	if value, ok := values["project_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.DomainLimits.ProjectLimit = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["ip_prefix_len"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.NamespaceCidr.IPPrefixLen = castedValue

			}

			if propertyValue, ok := childResourceMap["ip_prefix"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.NamespaceCidr.IPPrefix = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VxlanRouting = castedValue

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_router"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualRouter = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_network"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualNetwork = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_machine_interface"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualMachineInterface = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualIP = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_DNS_record"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualDNSRecord = castedValue

			}

			if propertyValue, ok := childResourceMap["virtual_DNS"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.VirtualDNS = castedValue

			}

			if propertyValue, ok := childResourceMap["subnet"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.Subnet = castedValue

			}

			if propertyValue, ok := childResourceMap["service_template"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.ServiceTemplate = castedValue

			}

			if propertyValue, ok := childResourceMap["service_instance"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.ServiceInstance = castedValue

			}

			if propertyValue, ok := childResourceMap["security_logging_object"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.SecurityLoggingObject = castedValue

			}

			if propertyValue, ok := childResourceMap["security_group_rule"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.SecurityGroupRule = castedValue

			}

			if propertyValue, ok := childResourceMap["security_group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.SecurityGroup = castedValue

			}

			if propertyValue, ok := childResourceMap["route_table"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.RouteTable = castedValue

			}

			if propertyValue, ok := childResourceMap["network_policy"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.NetworkPolicy = castedValue

			}

			if propertyValue, ok := childResourceMap["network_ipam"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.NetworkIpam = castedValue

			}

			if propertyValue, ok := childResourceMap["logical_router"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.LogicalRouter = castedValue

			}

			if propertyValue, ok := childResourceMap["loadbalancer_pool"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.LoadbalancerPool = castedValue

			}

			if propertyValue, ok := childResourceMap["loadbalancer_member"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.LoadbalancerMember = castedValue

			}

			if propertyValue, ok := childResourceMap["loadbalancer_healthmonitor"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.LoadbalancerHealthmonitor = castedValue

			}

			if propertyValue, ok := childResourceMap["instance_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.InstanceIP = castedValue

			}

			if propertyValue, ok := childResourceMap["global_vrouter_config"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.GlobalVrouterConfig = castedValue

			}

			if propertyValue, ok := childResourceMap["floating_ip_pool"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.FloatingIPPool = castedValue

			}

			if propertyValue, ok := childResourceMap["floating_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.FloatingIP = castedValue

			}

			if propertyValue, ok := childResourceMap["defaults"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.Defaults = castedValue

			}

			if propertyValue, ok := childResourceMap["bgp_router"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.BGPRouter = castedValue

			}

			if propertyValue, ok := childResourceMap["access_control_list"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Quota.AccessControlList = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

			if propertyValue, ok := childResourceMap["alarm_enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.AlarmEnable = castedValue

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

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["vrouter_instance_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.VrouterInstanceType = models.VRouterInstanceType(castedValue)

			}

			if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.ServiceTemplateProperties.Version = castedValue

			}

			if propertyValue, ok := childResourceMap["service_virtualization_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.ServiceVirtualizationType = models.ServiceVirtualizationType(castedValue)

			}

			if propertyValue, ok := childResourceMap["service_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.ServiceType = models.ServiceType(castedValue)

			}

			if propertyValue, ok := childResourceMap["service_scaling"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.ServiceTemplateProperties.ServiceScaling = castedValue

			}

			if propertyValue, ok := childResourceMap["service_mode"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.ServiceMode = models.ServiceModeType(castedValue)

			}

			if propertyValue, ok := childResourceMap["ordered_interfaces"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.ServiceTemplateProperties.OrderedInterfaces = castedValue

			}

			if propertyValue, ok := childResourceMap["interface_type"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.ServiceTemplateProperties.InterfaceType)

			}

			if propertyValue, ok := childResourceMap["instance_data"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.InstanceData = castedValue

			}

			if propertyValue, ok := childResourceMap["image_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.ImageName = castedValue

			}

			if propertyValue, ok := childResourceMap["flavor"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ServiceTemplateProperties.Flavor = castedValue

			}

			if propertyValue, ok := childResourceMap["availability_zone_enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.ServiceTemplateProperties.AvailabilityZoneEnable = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

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

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualDNSData.ReverseResolution = castedValue

			}

			if propertyValue, ok := childResourceMap["record_order"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualDNSData.RecordOrder = models.DnsRecordOrderType(castedValue)

			}

			if propertyValue, ok := childResourceMap["next_virtual_DNS"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualDNSData.NextVirtualDNS = castedValue

			}

			if propertyValue, ok := childResourceMap["floating_ip_record"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualDNSData.FloatingIPRecord = models.FloatingIpDnsNotation(castedValue)

			}

			if propertyValue, ok := childResourceMap["external_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualDNSData.ExternalVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["dynamic_records_from_client"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.VirtualDNSData.DynamicRecordsFromClient = castedValue

			}

			if propertyValue, ok := childResourceMap["domain_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.VirtualDNSData.DomainName = castedValue

			}

			if propertyValue, ok := childResourceMap["default_ttl_seconds"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.VirtualDNSData.DefaultTTLSeconds = castedValue

			}

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.Perms2.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.Perms2.GlobalAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentUUID = castedValue

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.ParentType = castedValue

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.UserVisible = castedValue

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Owner = castedValue

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Permissions.Group = castedValue

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.LastModified = castedValue

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.IDPerms.Enable = castedValue

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Description = castedValue

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Creator = castedValue

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.IDPerms.Created = castedValue

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.DisplayName = castedValue

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListDomain lists Domain with list spec.
func ListDomain(tx *sql.Tx, spec *common.ListSpec) ([]*models.Domain, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "domain"
	spec.Fields = DomainFields
	spec.RefFields = DomainRefFields
	spec.BackRefFields = DomainBackRefFields
	result := models.MakeDomainSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanDomain(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateDomain updates a resource
func UpdateDomain(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateDomainQuery = "update `domain` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateDomainQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateDomainQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateDomainQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateDomainQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateDomainQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateDomainQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateDomainQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateDomainQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateDomainQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateDomainQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateDomainQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateDomainQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateDomainQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateDomainQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateDomainQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateDomainQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateDomainQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateDomainQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateDomainQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DomainLimits.VirtualNetworkLimit", "."); ok {
		updateDomainQuery += "`virtual_network_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DomainLimits.SecurityGroupLimit", "."); ok {
		updateDomainQuery += "`security_group_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DomainLimits.ProjectLimit", "."); ok {
		updateDomainQuery += "`project_limit` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateDomainQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateDomainQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateDomainQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateDomainQuery += ","
	}

	updateDomainQuery =
		updateDomainQuery[:len(updateDomainQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateDomainQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateDomainQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteDomain deletes a resource
func DeleteDomain(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteDomainQuery
	selectQuery := "select count(uuid) from domain where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
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
