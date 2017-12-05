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

const insertLocationQuery = "insert into `location` (`private_ospd_vm_name`,`display_name`,`provisioning_state`,`private_redhat_subscription_key`,`provisioning_start_time`,`private_ospd_vm_ram_mb`,`private_redhat_subscription_user`,`gcp_account_info`,`uuid`,`provisioning_progress`,`private_dns_servers`,`private_ntp_hosts`,`private_ospd_package_url`,`provisioning_log`,`private_ospd_vm_vcpus`,`aws_region`,`provisioning_progress_stage`,`type`,`aws_access_key`,`aws_secret_key`,`aws_subnet`,`share`,`owner`,`owner_access`,`global_access`,`private_ospd_vm_disk_gb`,`private_redhat_pool_id`,`gcp_subnet`,`gcp_asn`,`gcp_region`,`fq_name`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`key_value_pair`,`private_ospd_user_name`,`private_ospd_user_password`,`private_redhat_subscription_pasword`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateLocationQuery = "update `location` set `private_ospd_vm_name` = ?,`display_name` = ?,`provisioning_state` = ?,`private_redhat_subscription_key` = ?,`provisioning_start_time` = ?,`private_ospd_vm_ram_mb` = ?,`private_redhat_subscription_user` = ?,`gcp_account_info` = ?,`uuid` = ?,`provisioning_progress` = ?,`private_dns_servers` = ?,`private_ntp_hosts` = ?,`private_ospd_package_url` = ?,`provisioning_log` = ?,`private_ospd_vm_vcpus` = ?,`aws_region` = ?,`provisioning_progress_stage` = ?,`type` = ?,`aws_access_key` = ?,`aws_secret_key` = ?,`aws_subnet` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`private_ospd_vm_disk_gb` = ?,`private_redhat_pool_id` = ?,`gcp_subnet` = ?,`gcp_asn` = ?,`gcp_region` = ?,`fq_name` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`key_value_pair` = ?,`private_ospd_user_name` = ?,`private_ospd_user_password` = ?,`private_redhat_subscription_pasword` = ?;"
const deleteLocationQuery = "delete from `location` where uuid = ?"
const listLocationQuery = "select `location`.`private_ospd_vm_name`,`location`.`display_name`,`location`.`provisioning_state`,`location`.`private_redhat_subscription_key`,`location`.`provisioning_start_time`,`location`.`private_ospd_vm_ram_mb`,`location`.`private_redhat_subscription_user`,`location`.`gcp_account_info`,`location`.`uuid`,`location`.`provisioning_progress`,`location`.`private_dns_servers`,`location`.`private_ntp_hosts`,`location`.`private_ospd_package_url`,`location`.`provisioning_log`,`location`.`private_ospd_vm_vcpus`,`location`.`aws_region`,`location`.`provisioning_progress_stage`,`location`.`type`,`location`.`aws_access_key`,`location`.`aws_secret_key`,`location`.`aws_subnet`,`location`.`share`,`location`.`owner`,`location`.`owner_access`,`location`.`global_access`,`location`.`private_ospd_vm_disk_gb`,`location`.`private_redhat_pool_id`,`location`.`gcp_subnet`,`location`.`gcp_asn`,`location`.`gcp_region`,`location`.`fq_name`,`location`.`created`,`location`.`creator`,`location`.`user_visible`,`location`.`last_modified`,`location`.`permissions_owner`,`location`.`permissions_owner_access`,`location`.`other_access`,`location`.`group`,`location`.`group_access`,`location`.`enable`,`location`.`description`,`location`.`key_value_pair`,`location`.`private_ospd_user_name`,`location`.`private_ospd_user_password`,`location`.`private_redhat_subscription_pasword` from `location`"
const showLocationQuery = "select `location`.`private_ospd_vm_name`,`location`.`display_name`,`location`.`provisioning_state`,`location`.`private_redhat_subscription_key`,`location`.`provisioning_start_time`,`location`.`private_ospd_vm_ram_mb`,`location`.`private_redhat_subscription_user`,`location`.`gcp_account_info`,`location`.`uuid`,`location`.`provisioning_progress`,`location`.`private_dns_servers`,`location`.`private_ntp_hosts`,`location`.`private_ospd_package_url`,`location`.`provisioning_log`,`location`.`private_ospd_vm_vcpus`,`location`.`aws_region`,`location`.`provisioning_progress_stage`,`location`.`type`,`location`.`aws_access_key`,`location`.`aws_secret_key`,`location`.`aws_subnet`,`location`.`share`,`location`.`owner`,`location`.`owner_access`,`location`.`global_access`,`location`.`private_ospd_vm_disk_gb`,`location`.`private_redhat_pool_id`,`location`.`gcp_subnet`,`location`.`gcp_asn`,`location`.`gcp_region`,`location`.`fq_name`,`location`.`created`,`location`.`creator`,`location`.`user_visible`,`location`.`last_modified`,`location`.`permissions_owner`,`location`.`permissions_owner_access`,`location`.`other_access`,`location`.`group`,`location`.`group_access`,`location`.`enable`,`location`.`description`,`location`.`key_value_pair`,`location`.`private_ospd_user_name`,`location`.`private_ospd_user_password`,`location`.`private_redhat_subscription_pasword` from `location` where uuid = ?"

func CreateLocation(tx *sql.Tx, model *models.Location) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLocationQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.PrivateOspdVMName),
		string(model.DisplayName),
		string(model.ProvisioningState),
		string(model.PrivateRedhatSubscriptionKey),
		string(model.ProvisioningStartTime),
		string(model.PrivateOspdVMRAMMB),
		string(model.PrivateRedhatSubscriptionUser),
		string(model.GCPAccountInfo),
		string(model.UUID),
		int(model.ProvisioningProgress),
		string(model.PrivateDNSServers),
		string(model.PrivateNTPHosts),
		string(model.PrivateOspdPackageURL),
		string(model.ProvisioningLog),
		string(model.PrivateOspdVMVcpus),
		string(model.AwsRegion),
		string(model.ProvisioningProgressStage),
		string(model.Type),
		string(model.AwsAccessKey),
		string(model.AwsSecretKey),
		string(model.AwsSubnet),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.PrivateOspdVMDiskGB),
		string(model.PrivateRedhatPoolID),
		string(model.GCPSubnet),
		int(model.GCPAsn),
		string(model.GCPRegion),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.PrivateOspdUserName),
		string(model.PrivateOspdUserPassword),
		string(model.PrivateRedhatSubscriptionPasword))

	return err
}

func scanLocation(rows *sql.Rows) (*models.Location, error) {
	m := models.MakeLocation()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.PrivateOspdVMName,
		&m.DisplayName,
		&m.ProvisioningState,
		&m.PrivateRedhatSubscriptionKey,
		&m.ProvisioningStartTime,
		&m.PrivateOspdVMRAMMB,
		&m.PrivateRedhatSubscriptionUser,
		&m.GCPAccountInfo,
		&m.UUID,
		&m.ProvisioningProgress,
		&m.PrivateDNSServers,
		&m.PrivateNTPHosts,
		&m.PrivateOspdPackageURL,
		&m.ProvisioningLog,
		&m.PrivateOspdVMVcpus,
		&m.AwsRegion,
		&m.ProvisioningProgressStage,
		&m.Type,
		&m.AwsAccessKey,
		&m.AwsSecretKey,
		&m.AwsSubnet,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.PrivateOspdVMDiskGB,
		&m.PrivateRedhatPoolID,
		&m.GCPSubnet,
		&m.GCPAsn,
		&m.GCPRegion,
		&jsonFQName,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&jsonAnnotationsKeyValuePair,
		&m.PrivateOspdUserName,
		&m.PrivateOspdUserPassword,
		&m.PrivateRedhatSubscriptionPasword); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildLocationWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["private_ospd_vm_name"]; ok {
		results = append(results, "private_ospd_vm_name = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_key"]; ok {
		results = append(results, "private_redhat_subscription_key = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_ram_mb"]; ok {
		results = append(results, "private_ospd_vm_ram_mb = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_user"]; ok {
		results = append(results, "private_redhat_subscription_user = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_account_info"]; ok {
		results = append(results, "gcp_account_info = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["private_dns_servers"]; ok {
		results = append(results, "private_dns_servers = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ntp_hosts"]; ok {
		results = append(results, "private_ntp_hosts = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_package_url"]; ok {
		results = append(results, "private_ospd_package_url = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_vcpus"]; ok {
		results = append(results, "private_ospd_vm_vcpus = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_region"]; ok {
		results = append(results, "aws_region = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["type"]; ok {
		results = append(results, "type = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_access_key"]; ok {
		results = append(results, "aws_access_key = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_secret_key"]; ok {
		results = append(results, "aws_secret_key = ?")
		values = append(values, value)
	}

	if value, ok := where["aws_subnet"]; ok {
		results = append(results, "aws_subnet = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_vm_disk_gb"]; ok {
		results = append(results, "private_ospd_vm_disk_gb = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_pool_id"]; ok {
		results = append(results, "private_redhat_pool_id = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_subnet"]; ok {
		results = append(results, "gcp_subnet = ?")
		values = append(values, value)
	}

	if value, ok := where["gcp_region"]; ok {
		results = append(results, "gcp_region = ?")
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

	if value, ok := where["private_ospd_user_name"]; ok {
		results = append(results, "private_ospd_user_name = ?")
		values = append(values, value)
	}

	if value, ok := where["private_ospd_user_password"]; ok {
		results = append(results, "private_ospd_user_password = ?")
		values = append(values, value)
	}

	if value, ok := where["private_redhat_subscription_pasword"]; ok {
		results = append(results, "private_redhat_subscription_pasword = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListLocation(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.Location, error) {
	result := models.MakeLocationSlice()
	whereQuery, values := buildLocationWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listLocationQuery)
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
		m, _ := scanLocation(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowLocation(tx *sql.Tx, uuid string) (*models.Location, error) {
	rows, err := tx.Query(showLocationQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanLocation(rows)
	}
	return nil, nil
}

func UpdateLocation(tx *sql.Tx, uuid string, model *models.Location) error {
	return nil
}

func DeleteLocation(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteLocationQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
