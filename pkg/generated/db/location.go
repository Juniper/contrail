package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLocationQuery = "insert into `location` (`private_ospd_user_name`,`gcp_account_info`,`gcp_subnet`,`uuid`,`private_ntp_hosts`,`aws_subnet`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`display_name`,`provisioning_start_time`,`private_dns_servers`,`private_ospd_vm_ram_mb`,`private_redhat_subscription_user`,`aws_secret_key`,`private_ospd_package_url`,`private_ospd_user_password`,`private_redhat_subscription_key`,`gcp_region`,`type`,`key_value_pair`,`provisioning_state`,`provisioning_progress`,`private_redhat_subscription_pasword`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`provisioning_progress_stage`,`provisioning_log`,`private_ospd_vm_disk_gb`,`private_ospd_vm_name`,`private_ospd_vm_vcpus`,`private_redhat_pool_id`,`fq_name`,`gcp_asn`,`aws_access_key`,`aws_region`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `private_ospd_user_name` = ?,`gcp_account_info` = ?,`gcp_subnet` = ?,`uuid` = ?,`private_ntp_hosts` = ?,`aws_subnet` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`provisioning_start_time` = ?,`private_dns_servers` = ?,`private_ospd_vm_ram_mb` = ?,`private_redhat_subscription_user` = ?,`aws_secret_key` = ?,`private_ospd_package_url` = ?,`private_ospd_user_password` = ?,`private_redhat_subscription_key` = ?,`gcp_region` = ?,`type` = ?,`key_value_pair` = ?,`provisioning_state` = ?,`provisioning_progress` = ?,`private_redhat_subscription_pasword` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_progress_stage` = ?,`provisioning_log` = ?,`private_ospd_vm_disk_gb` = ?,`private_ospd_vm_name` = ?,`private_ospd_vm_vcpus` = ?,`private_redhat_pool_id` = ?,`fq_name` = ?,`gcp_asn` = ?,`aws_access_key` = ?,`aws_region` = ?;"
const deleteLocationQuery = "delete from `location` where uuid = ?"

// LocationFields is db columns for Location
var LocationFields = []string{
	"private_ospd_user_name",
	"gcp_account_info",
	"gcp_subnet",
	"uuid",
	"private_ntp_hosts",
	"aws_subnet",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"display_name",
	"provisioning_start_time",
	"private_dns_servers",
	"private_ospd_vm_ram_mb",
	"private_redhat_subscription_user",
	"aws_secret_key",
	"private_ospd_package_url",
	"private_ospd_user_password",
	"private_redhat_subscription_key",
	"gcp_region",
	"type",
	"key_value_pair",
	"provisioning_state",
	"provisioning_progress",
	"private_redhat_subscription_pasword",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"provisioning_progress_stage",
	"provisioning_log",
	"private_ospd_vm_disk_gb",
	"private_ospd_vm_name",
	"private_ospd_vm_vcpus",
	"private_redhat_pool_id",
	"fq_name",
	"gcp_asn",
	"aws_access_key",
	"aws_region",
}

// LocationRefFields is db reference fields for Location
var LocationRefFields = map[string][]string{}

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
	_, err = stmt.Exec(string(model.PrivateOspdUserName),
		string(model.GCPAccountInfo),
		string(model.GCPSubnet),
		string(model.UUID),
		string(model.PrivateNTPHosts),
		string(model.AwsSubnet),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.DisplayName),
		string(model.ProvisioningStartTime),
		string(model.PrivateDNSServers),
		string(model.PrivateOspdVMRAMMB),
		string(model.PrivateRedhatSubscriptionUser),
		string(model.AwsSecretKey),
		string(model.PrivateOspdPackageURL),
		string(model.PrivateOspdUserPassword),
		string(model.PrivateRedhatSubscriptionKey),
		string(model.GCPRegion),
		string(model.Type),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningState),
		int(model.ProvisioningProgress),
		string(model.PrivateRedhatSubscriptionPasword),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.ProvisioningProgressStage),
		string(model.ProvisioningLog),
		string(model.PrivateOspdVMDiskGB),
		string(model.PrivateOspdVMName),
		string(model.PrivateOspdVMVcpus),
		string(model.PrivateRedhatPoolID),
		utils.MustJSON(model.FQName),
		int(model.GCPAsn),
		string(model.AwsAccessKey),
		string(model.AwsRegion))
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

	if value, ok := values["private_ospd_user_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdUserName = castedValue

	}

	if value, ok := values["gcp_account_info"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.GCPAccountInfo = castedValue

	}

	if value, ok := values["gcp_subnet"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.GCPSubnet = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["private_ntp_hosts"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateNTPHosts = castedValue

	}

	if value, ok := values["aws_subnet"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsSubnet = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["private_dns_servers"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateDNSServers = castedValue

	}

	if value, ok := values["private_ospd_vm_ram_mb"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdVMRAMMB = castedValue

	}

	if value, ok := values["private_redhat_subscription_user"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateRedhatSubscriptionUser = castedValue

	}

	if value, ok := values["aws_secret_key"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsSecretKey = castedValue

	}

	if value, ok := values["private_ospd_package_url"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdPackageURL = castedValue

	}

	if value, ok := values["private_ospd_user_password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdUserPassword = castedValue

	}

	if value, ok := values["private_redhat_subscription_key"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateRedhatSubscriptionKey = castedValue

	}

	if value, ok := values["gcp_region"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.GCPRegion = castedValue

	}

	if value, ok := values["type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Type = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["private_redhat_subscription_pasword"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateRedhatSubscriptionPasword = castedValue

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["private_ospd_vm_disk_gb"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdVMDiskGB = castedValue

	}

	if value, ok := values["private_ospd_vm_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdVMName = castedValue

	}

	if value, ok := values["private_ospd_vm_vcpus"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateOspdVMVcpus = castedValue

	}

	if value, ok := values["private_redhat_pool_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PrivateRedhatPoolID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["gcp_asn"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.GCPAsn = castedValue

	}

	if value, ok := values["aws_access_key"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsAccessKey = castedValue

	}

	if value, ok := values["aws_region"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AwsRegion = castedValue

	}

	return m, nil
}

// ListLocation lists Location with list spec.
func ListLocation(tx *sql.Tx, spec *db.ListSpec) ([]*models.Location, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "location"
	spec.Fields = LocationFields
	spec.RefFields = LocationRefFields
	result := models.MakeLocationSlice()
	query, columns, values := db.BuildListQuery(spec)
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
		m, err := scanLocation(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowLocation shows Location resource
func ShowLocation(tx *sql.Tx, uuid string) (*models.Location, error) {
	list, err := ListLocation(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateLocation updates a resource
func UpdateLocation(tx *sql.Tx, uuid string, model *models.Location) error {
	//TODO(nati) support update
	return nil
}

// DeleteLocation deletes a resource
func DeleteLocation(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLocationQuery)
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
