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

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`contrail_cluster_id`,`default_capacity_drives`,`default_storage_backend_bond_interface_members`,`external_allocation_pool_start`,`public_gateway`,`public_ip`,`uuid`,`display_name`,`admin_password`,`default_osd_drives`,`default_performance_drives`,`provisioning_log`,`provisioning_progress`,`provisioning_progress_stage`,`key_value_pair`,`global_access`,`share`,`perms2_owner`,`perms2_owner_access`,`provisioning_start_time`,`provisioning_state`,`default_journal_drives`,`default_storage_access_bond_interface_members`,`openstack_webui`,`external_allocation_pool_end`,`external_net_cidr`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackClusterQuery = "update `openstack_cluster` set `fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`contrail_cluster_id` = ?,`default_capacity_drives` = ?,`default_storage_backend_bond_interface_members` = ?,`external_allocation_pool_start` = ?,`public_gateway` = ?,`public_ip` = ?,`uuid` = ?,`display_name` = ?,`admin_password` = ?,`default_osd_drives` = ?,`default_performance_drives` = ?,`provisioning_log` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`key_value_pair` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`default_journal_drives` = ?,`default_storage_access_bond_interface_members` = ?,`openstack_webui` = ?,`external_allocation_pool_end` = ?,`external_net_cidr` = ?;"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"

// OpenstackClusterFields is db columns for OpenstackCluster
var OpenstackClusterFields = []string{
	"fq_name",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"contrail_cluster_id",
	"default_capacity_drives",
	"default_storage_backend_bond_interface_members",
	"external_allocation_pool_start",
	"public_gateway",
	"public_ip",
	"uuid",
	"display_name",
	"admin_password",
	"default_osd_drives",
	"default_performance_drives",
	"provisioning_log",
	"provisioning_progress",
	"provisioning_progress_stage",
	"key_value_pair",
	"global_access",
	"share",
	"perms2_owner",
	"perms2_owner_access",
	"provisioning_start_time",
	"provisioning_state",
	"default_journal_drives",
	"default_storage_access_bond_interface_members",
	"openstack_webui",
	"external_allocation_pool_end",
	"external_net_cidr",
}

// OpenstackClusterRefFields is db reference fields for OpenstackCluster
var OpenstackClusterRefFields = map[string][]string{}

// CreateOpenstackCluster inserts OpenstackCluster to DB
func CreateOpenstackCluster(tx *sql.Tx, model *models.OpenstackCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackClusterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertOpenstackClusterQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.ContrailClusterID),
		string(model.DefaultCapacityDrives),
		string(model.DefaultStorageBackendBondInterfaceMembers),
		string(model.ExternalAllocationPoolStart),
		string(model.PublicGateway),
		string(model.PublicIP),
		string(model.UUID),
		string(model.DisplayName),
		string(model.AdminPassword),
		string(model.DefaultOsdDrives),
		string(model.DefaultPerformanceDrives),
		string(model.ProvisioningLog),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.DefaultJournalDrives),
		string(model.DefaultStorageAccessBondInterfaceMembers),
		string(model.OpenstackWebui),
		string(model.ExternalAllocationPoolEnd),
		string(model.ExternalNetCidr))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanOpenstackCluster(values map[string]interface{}) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

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

	if value, ok := values["contrail_cluster_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ContrailClusterID = castedValue

	}

	if value, ok := values["default_capacity_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultCapacityDrives = castedValue

	}

	if value, ok := values["default_storage_backend_bond_interface_members"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultStorageBackendBondInterfaceMembers = castedValue

	}

	if value, ok := values["external_allocation_pool_start"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ExternalAllocationPoolStart = castedValue

	}

	if value, ok := values["public_gateway"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PublicGateway = castedValue

	}

	if value, ok := values["public_ip"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.PublicIP = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["admin_password"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.AdminPassword = castedValue

	}

	if value, ok := values["default_osd_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultOsdDrives = castedValue

	}

	if value, ok := values["default_performance_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultPerformanceDrives = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["default_journal_drives"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultJournalDrives = castedValue

	}

	if value, ok := values["default_storage_access_bond_interface_members"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultStorageAccessBondInterfaceMembers = castedValue

	}

	if value, ok := values["openstack_webui"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.OpenstackWebui = castedValue

	}

	if value, ok := values["external_allocation_pool_end"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ExternalAllocationPoolEnd = castedValue

	}

	if value, ok := values["external_net_cidr"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ExternalNetCidr = castedValue

	}

	return m, nil
}

// ListOpenstackCluster lists OpenstackCluster with list spec.
func ListOpenstackCluster(tx *sql.Tx, spec *db.ListSpec) ([]*models.OpenstackCluster, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "openstack_cluster"
	spec.Fields = OpenstackClusterFields
	spec.RefFields = OpenstackClusterRefFields
	result := models.MakeOpenstackClusterSlice()
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
		m, err := scanOpenstackCluster(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowOpenstackCluster shows OpenstackCluster resource
func ShowOpenstackCluster(tx *sql.Tx, uuid string) (*models.OpenstackCluster, error) {
	list, err := ListOpenstackCluster(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateOpenstackCluster updates a resource
func UpdateOpenstackCluster(tx *sql.Tx, uuid string, model *models.OpenstackCluster) error {
	//TODO(nati) support update
	return nil
}

// DeleteOpenstackCluster deletes a resource
func DeleteOpenstackCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackClusterQuery)
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
