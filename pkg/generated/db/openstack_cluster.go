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

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`contrail_cluster_id`,`external_allocation_pool_end`,`external_net_cidr`,`provisioning_state`,`provisioning_progress`,`provisioning_progress_stage`,`default_performance_drives`,`external_allocation_pool_start`,`owner`,`owner_access`,`global_access`,`share`,`fq_name`,`display_name`,`admin_password`,`default_storage_backend_bond_interface_members`,`key_value_pair`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`enable`,`description`,`provisioning_log`,`public_gateway`,`public_ip`,`uuid`,`default_capacity_drives`,`default_journal_drives`,`default_osd_drives`,`default_storage_access_bond_interface_members`,`openstack_webui`,`provisioning_start_time`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `contrail_cluster_id` = ?,`external_allocation_pool_end` = ?,`external_net_cidr` = ?,`provisioning_state` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`default_performance_drives` = ?,`external_allocation_pool_start` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`fq_name` = ?,`display_name` = ?,`admin_password` = ?,`default_storage_backend_bond_interface_members` = ?,`key_value_pair` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`provisioning_log` = ?,`public_gateway` = ?,`public_ip` = ?,`uuid` = ?,`default_capacity_drives` = ?,`default_journal_drives` = ?,`default_osd_drives` = ?,`default_storage_access_bond_interface_members` = ?,`openstack_webui` = ?,`provisioning_start_time` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"
const listOpenstackClusterQuery = "select `openstack_cluster`.`contrail_cluster_id`,`openstack_cluster`.`external_allocation_pool_end`,`openstack_cluster`.`external_net_cidr`,`openstack_cluster`.`provisioning_state`,`openstack_cluster`.`provisioning_progress`,`openstack_cluster`.`provisioning_progress_stage`,`openstack_cluster`.`default_performance_drives`,`openstack_cluster`.`external_allocation_pool_start`,`openstack_cluster`.`owner`,`openstack_cluster`.`owner_access`,`openstack_cluster`.`global_access`,`openstack_cluster`.`share`,`openstack_cluster`.`fq_name`,`openstack_cluster`.`display_name`,`openstack_cluster`.`admin_password`,`openstack_cluster`.`default_storage_backend_bond_interface_members`,`openstack_cluster`.`key_value_pair`,`openstack_cluster`.`created`,`openstack_cluster`.`creator`,`openstack_cluster`.`user_visible`,`openstack_cluster`.`last_modified`,`openstack_cluster`.`group`,`openstack_cluster`.`group_access`,`openstack_cluster`.`permissions_owner`,`openstack_cluster`.`permissions_owner_access`,`openstack_cluster`.`other_access`,`openstack_cluster`.`enable`,`openstack_cluster`.`description`,`openstack_cluster`.`provisioning_log`,`openstack_cluster`.`public_gateway`,`openstack_cluster`.`public_ip`,`openstack_cluster`.`uuid`,`openstack_cluster`.`default_capacity_drives`,`openstack_cluster`.`default_journal_drives`,`openstack_cluster`.`default_osd_drives`,`openstack_cluster`.`default_storage_access_bond_interface_members`,`openstack_cluster`.`openstack_webui`,`openstack_cluster`.`provisioning_start_time` from `openstack_cluster`"
const showOpenstackClusterQuery = "select `openstack_cluster`.`contrail_cluster_id`,`openstack_cluster`.`external_allocation_pool_end`,`openstack_cluster`.`external_net_cidr`,`openstack_cluster`.`provisioning_state`,`openstack_cluster`.`provisioning_progress`,`openstack_cluster`.`provisioning_progress_stage`,`openstack_cluster`.`default_performance_drives`,`openstack_cluster`.`external_allocation_pool_start`,`openstack_cluster`.`owner`,`openstack_cluster`.`owner_access`,`openstack_cluster`.`global_access`,`openstack_cluster`.`share`,`openstack_cluster`.`fq_name`,`openstack_cluster`.`display_name`,`openstack_cluster`.`admin_password`,`openstack_cluster`.`default_storage_backend_bond_interface_members`,`openstack_cluster`.`key_value_pair`,`openstack_cluster`.`created`,`openstack_cluster`.`creator`,`openstack_cluster`.`user_visible`,`openstack_cluster`.`last_modified`,`openstack_cluster`.`group`,`openstack_cluster`.`group_access`,`openstack_cluster`.`permissions_owner`,`openstack_cluster`.`permissions_owner_access`,`openstack_cluster`.`other_access`,`openstack_cluster`.`enable`,`openstack_cluster`.`description`,`openstack_cluster`.`provisioning_log`,`openstack_cluster`.`public_gateway`,`openstack_cluster`.`public_ip`,`openstack_cluster`.`uuid`,`openstack_cluster`.`default_capacity_drives`,`openstack_cluster`.`default_journal_drives`,`openstack_cluster`.`default_osd_drives`,`openstack_cluster`.`default_storage_access_bond_interface_members`,`openstack_cluster`.`openstack_webui`,`openstack_cluster`.`provisioning_start_time` from `openstack_cluster` where uuid = ?"

func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.ContrailClusterID),
		string(model.ExternalAllocationPoolEnd),
		string(model.ExternalNetCidr),
		string(model.ProvisioningState),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.DefaultPerformanceDrives),
		string(model.ExternalAllocationPoolStart),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.FQName),
		string(model.DisplayName),
		string(model.AdminPassword),
		string(model.DefaultStorageBackendBondInterfaceMembers),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.ProvisioningLog),
		string(model.PublicGateway),
		string(model.PublicIP),
		string(model.UUID),
		string(model.DefaultCapacityDrives),
		string(model.DefaultJournalDrives),
		string(model.DefaultOsdDrives),
		string(model.DefaultStorageAccessBondInterfaceMembers),
		string(model.OpenstackWebui),
		string(model.ProvisioningStartTime))

	return err
}

func scanOpenstackCluster(rows *sql.Rows) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	var jsonPerms2Share string

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.ContrailClusterID,
		&m.ExternalAllocationPoolEnd,
		&m.ExternalNetCidr,
		&m.ProvisioningState,
		&m.ProvisioningProgress,
		&m.ProvisioningProgressStage,
		&m.DefaultPerformanceDrives,
		&m.ExternalAllocationPoolStart,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonFQName,
		&m.DisplayName,
		&m.AdminPassword,
		&m.DefaultStorageBackendBondInterfaceMembers,
		&jsonAnnotationsKeyValuePair,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.ProvisioningLog,
		&m.PublicGateway,
		&m.PublicIP,
		&m.UUID,
		&m.DefaultCapacityDrives,
		&m.DefaultJournalDrives,
		&m.DefaultOsdDrives,
		&m.DefaultStorageAccessBondInterfaceMembers,
		&m.OpenstackWebui,
		&m.ProvisioningStartTime); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildOpenstackClusterWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["contrail_cluster_id"]; ok {
		results = append(results, "contrail_cluster_id = ?")
		values = append(values, value)
	}

	if value, ok := where["external_allocation_pool_end"]; ok {
		results = append(results, "external_allocation_pool_end = ?")
		values = append(values, value)
	}

	if value, ok := where["external_net_cidr"]; ok {
		results = append(results, "external_net_cidr = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_state"]; ok {
		results = append(results, "provisioning_state = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_progress_stage"]; ok {
		results = append(results, "provisioning_progress_stage = ?")
		values = append(values, value)
	}

	if value, ok := where["default_performance_drives"]; ok {
		results = append(results, "default_performance_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["external_allocation_pool_start"]; ok {
		results = append(results, "external_allocation_pool_start = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["admin_password"]; ok {
		results = append(results, "admin_password = ?")
		values = append(values, value)
	}

	if value, ok := where["default_storage_backend_bond_interface_members"]; ok {
		results = append(results, "default_storage_backend_bond_interface_members = ?")
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

	if value, ok := where["provisioning_log"]; ok {
		results = append(results, "provisioning_log = ?")
		values = append(values, value)
	}

	if value, ok := where["public_gateway"]; ok {
		results = append(results, "public_gateway = ?")
		values = append(values, value)
	}

	if value, ok := where["public_ip"]; ok {
		results = append(results, "public_ip = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["default_capacity_drives"]; ok {
		results = append(results, "default_capacity_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_journal_drives"]; ok {
		results = append(results, "default_journal_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_osd_drives"]; ok {
		results = append(results, "default_osd_drives = ?")
		values = append(values, value)
	}

	if value, ok := where["default_storage_access_bond_interface_members"]; ok {
		results = append(results, "default_storage_access_bond_interface_members = ?")
		values = append(values, value)
	}

	if value, ok := where["openstack_webui"]; ok {
		results = append(results, "openstack_webui = ?")
		values = append(values, value)
	}

	if value, ok := where["provisioning_start_time"]; ok {
		results = append(results, "provisioning_start_time = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListOpenstackCluster(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.OpenstackCluster, error) {
	result := models.MakeOpenstackClusterSlice()
	whereQuery, values := buildOpenstackClusterWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listOpenstackClusterQuery)
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
		m, _ := scanOpenstackCluster(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowOpenstackCluster(tx *sql.Tx, uuid string) (*models.OpenstackCluster, error) {
	rows, err := tx.Query(showOpenstackClusterQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanOpenstackCluster(rows)
	}
	return nil, nil
}

func UpdateOpenstackCluster(tx *sql.Tx, uuid string, model *models.OpenstackCluster) error {
	return nil
}

func DeleteOpenstackCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackClusterQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
