package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLocationQuery = "insert into `location` (`uuid`,`type`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`private_redhat_subscription_user`,`private_redhat_subscription_pasword`,`private_redhat_subscription_key`,`private_redhat_pool_id`,`private_ospd_vm_vcpus`,`private_ospd_vm_ram_mb`,`private_ospd_vm_name`,`private_ospd_vm_disk_gb`,`private_ospd_user_password`,`private_ospd_user_name`,`private_ospd_package_url`,`private_ntp_hosts`,`private_dns_servers`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`gcp_subnet`,`gcp_region`,`gcp_asn`,`gcp_account_info`,`fq_name`,`display_name`,`aws_subnet`,`aws_secret_key`,`aws_region`,`aws_access_key`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `uuid` = ?,`type` = ?,`provisioning_state` = ?,`provisioning_start_time` = ?,`provisioning_progress_stage` = ?,`provisioning_progress` = ?,`provisioning_log` = ?,`private_redhat_subscription_user` = ?,`private_redhat_subscription_pasword` = ?,`private_redhat_subscription_key` = ?,`private_redhat_pool_id` = ?,`private_ospd_vm_vcpus` = ?,`private_ospd_vm_ram_mb` = ?,`private_ospd_vm_name` = ?,`private_ospd_vm_disk_gb` = ?,`private_ospd_user_password` = ?,`private_ospd_user_name` = ?,`private_ospd_package_url` = ?,`private_ntp_hosts` = ?,`private_dns_servers` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`gcp_subnet` = ?,`gcp_region` = ?,`gcp_asn` = ?,`gcp_account_info` = ?,`fq_name` = ?,`display_name` = ?,`aws_subnet` = ?,`aws_secret_key` = ?,`aws_region` = ?,`aws_access_key` = ?,`key_value_pair` = ?;"
const deleteLocationQuery = "delete from `location` where uuid = ?"

// LocationFields is db columns for Location
var LocationFields = []string{
	"uuid",
	"type",
	"provisioning_state",
	"provisioning_start_time",
	"provisioning_progress_stage",
	"provisioning_progress",
	"provisioning_log",
	"private_redhat_subscription_user",
	"private_redhat_subscription_pasword",
	"private_redhat_subscription_key",
	"private_redhat_pool_id",
	"private_ospd_vm_vcpus",
	"private_ospd_vm_ram_mb",
	"private_ospd_vm_name",
	"private_ospd_vm_disk_gb",
	"private_ospd_user_password",
	"private_ospd_user_name",
	"private_ospd_package_url",
	"private_ntp_hosts",
	"private_dns_servers",
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
	"gcp_subnet",
	"gcp_region",
	"gcp_asn",
	"gcp_account_info",
	"fq_name",
	"display_name",
	"aws_subnet",
	"aws_secret_key",
	"aws_region",
	"aws_access_key",
	"key_value_pair",
}

// LocationRefFields is db reference fields for Location
var LocationRefFields = map[string][]string{}

// LocationBackRefFields is db back reference fields for Location
var LocationBackRefFields = map[string][]string{

	"physical_router": {
		"uuid",
		"server_port",
		"server_ip",
		"resource",
		"physical_router_vnc_managed",
		"physical_router_vendor_name",
		"username",
		"password",
		"version",
		"v3_security_name",
		"v3_security_level",
		"v3_security_engine_id",
		"v3_privacy_protocol",
		"v3_privacy_password",
		"v3_engine_time",
		"v3_engine_id",
		"v3_engine_boots",
		"v3_context_engine_id",
		"v3_context",
		"v3_authentication_protocol",
		"v3_authentication_password",
		"v2_community",
		"timeout",
		"retries",
		"local_port",
		"physical_router_snmp",
		"physical_router_role",
		"physical_router_product_name",
		"physical_router_management_ip",
		"physical_router_loopback_ip",
		"physical_router_lldp",
		"service_port",
		"physical_router_image_uri",
		"physical_router_dataplane_ip",
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

// CreateLocation inserts Location to DB
func CreateLocation(tx *sql.Tx, model *models.Location) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLocationQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		string(model.Type),
		string(model.ProvisioningState),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningProgressStage),
		int(model.ProvisioningProgress),
		string(model.ProvisioningLog),
		string(model.PrivateRedhatSubscriptionUser),
		string(model.PrivateRedhatSubscriptionPasword),
		string(model.PrivateRedhatSubscriptionKey),
		string(model.PrivateRedhatPoolID),
		string(model.PrivateOspdVMVcpus),
		string(model.PrivateOspdVMRAMMB),
		string(model.PrivateOspdVMName),
		string(model.PrivateOspdVMDiskGB),
		string(model.PrivateOspdUserPassword),
		string(model.PrivateOspdUserName),
		string(model.PrivateOspdPackageURL),
		string(model.PrivateNTPHosts),
		string(model.PrivateDNSServers),
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
		string(model.GCPSubnet),
		string(model.GCPRegion),
		int(model.GCPAsn),
		string(model.GCPAccountInfo),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.AwsSubnet),
		string(model.AwsSecretKey),
		string(model.AwsRegion),
		string(model.AwsAccessKey),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanLocation(values map[string]interface{}) (*models.Location, error) {
	m := models.MakeLocation()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Type = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["private_redhat_subscription_user"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateRedhatSubscriptionUser = castedValue

	}

	if value, ok := values["private_redhat_subscription_pasword"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateRedhatSubscriptionPasword = castedValue

	}

	if value, ok := values["private_redhat_subscription_key"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateRedhatSubscriptionKey = castedValue

	}

	if value, ok := values["private_redhat_pool_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateRedhatPoolID = castedValue

	}

	if value, ok := values["private_ospd_vm_vcpus"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdVMVcpus = castedValue

	}

	if value, ok := values["private_ospd_vm_ram_mb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdVMRAMMB = castedValue

	}

	if value, ok := values["private_ospd_vm_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdVMName = castedValue

	}

	if value, ok := values["private_ospd_vm_disk_gb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdVMDiskGB = castedValue

	}

	if value, ok := values["private_ospd_user_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdUserPassword = castedValue

	}

	if value, ok := values["private_ospd_user_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdUserName = castedValue

	}

	if value, ok := values["private_ospd_package_url"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateOspdPackageURL = castedValue

	}

	if value, ok := values["private_ntp_hosts"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateNTPHosts = castedValue

	}

	if value, ok := values["private_dns_servers"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PrivateDNSServers = castedValue

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

	if value, ok := values["gcp_subnet"]; ok {

		castedValue := common.InterfaceToString(value)

		m.GCPSubnet = castedValue

	}

	if value, ok := values["gcp_region"]; ok {

		castedValue := common.InterfaceToString(value)

		m.GCPRegion = castedValue

	}

	if value, ok := values["gcp_asn"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.GCPAsn = castedValue

	}

	if value, ok := values["gcp_account_info"]; ok {

		castedValue := common.InterfaceToString(value)

		m.GCPAccountInfo = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["aws_subnet"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsSubnet = castedValue

	}

	if value, ok := values["aws_secret_key"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsSecretKey = castedValue

	}

	if value, ok := values["aws_region"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsRegion = castedValue

	}

	if value, ok := values["aws_access_key"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AwsAccessKey = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["backref_physical_router"]; ok {
		var childResources []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			if childResourceMap["uuid"] == "" {
				continue
			}
			childModel := models.MakePhysicalRouter()
			m.PhysicalRouters = append(m.PhysicalRouters, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.UUID = castedValue

			}

			if propertyValue, ok := childResourceMap["server_port"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.TelemetryInfo.ServerPort = castedValue

			}

			if propertyValue, ok := childResourceMap["server_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.TelemetryInfo.ServerIP = castedValue

			}

			if propertyValue, ok := childResourceMap["resource"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.TelemetryInfo.Resource)

			}

			if propertyValue, ok := childResourceMap["physical_router_vnc_managed"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.PhysicalRouterVNCManaged = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_vendor_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterVendorName = castedValue

			}

			if propertyValue, ok := childResourceMap["username"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterUserCredentials.Username = castedValue

			}

			if propertyValue, ok := childResourceMap["password"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterUserCredentials.Password = castedValue

			}

			if propertyValue, ok := childResourceMap["version"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.Version = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_security_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3SecurityName = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_security_level"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3SecurityLevel = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_security_engine_id"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3SecurityEngineID = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_privacy_protocol"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3PrivacyProtocol = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_privacy_password"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3PrivacyPassword = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_engine_time"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3EngineTime = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_engine_id"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3EngineID = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_engine_boots"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3EngineBoots = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_context_engine_id"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3ContextEngineID = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_context"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3Context = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_authentication_protocol"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3AuthenticationProtocol = castedValue

			}

			if propertyValue, ok := childResourceMap["v3_authentication_password"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V3AuthenticationPassword = castedValue

			}

			if propertyValue, ok := childResourceMap["v2_community"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.V2Community = castedValue

			}

			if propertyValue, ok := childResourceMap["timeout"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.Timeout = castedValue

			}

			if propertyValue, ok := childResourceMap["retries"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.Retries = castedValue

			}

			if propertyValue, ok := childResourceMap["local_port"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToInt(propertyValue)

				childModel.PhysicalRouterSNMPCredentials.LocalPort = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_snmp"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.PhysicalRouterSNMP = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_role"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterRole = models.PhysicalRouterRole(castedValue)

			}

			if propertyValue, ok := childResourceMap["physical_router_product_name"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterProductName = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_management_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterManagementIP = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_loopback_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterLoopbackIP = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_lldp"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToBool(propertyValue)

				childModel.PhysicalRouterLLDP = castedValue

			}

			if propertyValue, ok := childResourceMap["service_port"]; ok && propertyValue != nil {

				json.Unmarshal(common.InterfaceToBytes(propertyValue), &childModel.PhysicalRouterJunosServicePorts.ServicePort)

			}

			if propertyValue, ok := childResourceMap["physical_router_image_uri"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterImageURI = castedValue

			}

			if propertyValue, ok := childResourceMap["physical_router_dataplane_ip"]; ok && propertyValue != nil {

				castedValue := common.InterfaceToString(propertyValue)

				childModel.PhysicalRouterDataplaneIP = castedValue

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

// ListLocation lists Location with list spec.
func ListLocation(tx *sql.Tx, spec *common.ListSpec) ([]*models.Location, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "location"
	spec.Fields = LocationFields
	spec.RefFields = LocationRefFields
	spec.BackRefFields = LocationBackRefFields
	result := models.MakeLocationSlice()
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
		m, err := scanLocation(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateLocation updates a resource
func UpdateLocation(tx *sql.Tx, uuid string, model *models.Location) error {
	//TODO(nati) support update
	return nil
}

// DeleteLocation deletes a resource
func DeleteLocation(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteLocationQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
