package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertOpenstackClusterQuery = "insert into `openstack_cluster` (`uuid`,`public_ip`,`public_gateway`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`openstack_webui`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`external_net_cidr`,`external_allocation_pool_start`,`external_allocation_pool_end`,`display_name`,`default_storage_backend_bond_interface_members`,`default_storage_access_bond_interface_members`,`default_performance_drives`,`default_osd_drives`,`default_journal_drives`,`default_capacity_drives`,`contrail_cluster_id`,`key_value_pair`,`admin_password`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteOpenstackClusterQuery = "delete from `openstack_cluster` where uuid = ?"

// OpenstackClusterFields is db columns for OpenstackCluster
var OpenstackClusterFields = []string{
	"uuid",
	"public_ip",
	"public_gateway",
	"provisioning_state",
	"provisioning_start_time",
	"provisioning_progress_stage",
	"provisioning_progress",
	"provisioning_log",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"openstack_webui",
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
	"external_net_cidr",
	"external_allocation_pool_start",
	"external_allocation_pool_end",
	"display_name",
	"default_storage_backend_bond_interface_members",
	"default_storage_access_bond_interface_members",
	"default_performance_drives",
	"default_osd_drives",
	"default_journal_drives",
	"default_capacity_drives",
	"contrail_cluster_id",
	"key_value_pair",
	"admin_password",
}

// OpenstackClusterRefFields is db reference fields for OpenstackCluster
var OpenstackClusterRefFields = map[string][]string{}

// OpenstackClusterBackRefFields is db back reference fields for OpenstackCluster
var OpenstackClusterBackRefFields = map[string][]string{}

// OpenstackClusterParentTypes is possible parents for OpenstackCluster
var OpenstackClusterParents = []string{}

// CreateOpenstackCluster inserts OpenstackCluster to DB
func CreateOpenstackCluster(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateOpenstackClusterRequest) error {
	model := request.OpenstackCluster
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
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetPublicIP()),
		string(model.GetPublicGateway()),
		string(model.GetProvisioningState()),
		string(model.GetProvisioningStartTime()),
		string(model.GetProvisioningProgressStage()),
		int(model.GetProvisioningProgress()),
		string(model.GetProvisioningLog()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetOpenstackWebui()),
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
		string(model.GetExternalNetCidr()),
		string(model.GetExternalAllocationPoolStart()),
		string(model.GetExternalAllocationPoolEnd()),
		string(model.GetDisplayName()),
		string(model.GetDefaultStorageBackendBondInterfaceMembers()),
		string(model.GetDefaultStorageAccessBondInterfaceMembers()),
		string(model.GetDefaultPerformanceDrives()),
		string(model.GetDefaultOsdDrives()),
		string(model.GetDefaultJournalDrives()),
		string(model.GetDefaultCapacityDrives()),
		string(model.GetContrailClusterID()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()),
		string(model.GetAdminPassword()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "openstack_cluster",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "openstack_cluster", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanOpenstackCluster(values map[string]interface{}) (*models.OpenstackCluster, error) {
	m := models.MakeOpenstackCluster()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["public_ip"]; ok {

		m.PublicIP = common.InterfaceToString(value)

	}

	if value, ok := values["public_gateway"]; ok {

		m.PublicGateway = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_state"]; ok {

		m.ProvisioningState = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		m.ProvisioningStartTime = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		m.ProvisioningProgressStage = common.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress"]; ok {

		m.ProvisioningProgress = common.InterfaceToInt64(value)

	}

	if value, ok := values["provisioning_log"]; ok {

		m.ProvisioningLog = common.InterfaceToString(value)

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

	if value, ok := values["openstack_webui"]; ok {

		m.OpenstackWebui = common.InterfaceToString(value)

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

	if value, ok := values["external_net_cidr"]; ok {

		m.ExternalNetCidr = common.InterfaceToString(value)

	}

	if value, ok := values["external_allocation_pool_start"]; ok {

		m.ExternalAllocationPoolStart = common.InterfaceToString(value)

	}

	if value, ok := values["external_allocation_pool_end"]; ok {

		m.ExternalAllocationPoolEnd = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["default_storage_backend_bond_interface_members"]; ok {

		m.DefaultStorageBackendBondInterfaceMembers = common.InterfaceToString(value)

	}

	if value, ok := values["default_storage_access_bond_interface_members"]; ok {

		m.DefaultStorageAccessBondInterfaceMembers = common.InterfaceToString(value)

	}

	if value, ok := values["default_performance_drives"]; ok {

		m.DefaultPerformanceDrives = common.InterfaceToString(value)

	}

	if value, ok := values["default_osd_drives"]; ok {

		m.DefaultOsdDrives = common.InterfaceToString(value)

	}

	if value, ok := values["default_journal_drives"]; ok {

		m.DefaultJournalDrives = common.InterfaceToString(value)

	}

	if value, ok := values["default_capacity_drives"]; ok {

		m.DefaultCapacityDrives = common.InterfaceToString(value)

	}

	if value, ok := values["contrail_cluster_id"]; ok {

		m.ContrailClusterID = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["admin_password"]; ok {

		m.AdminPassword = common.InterfaceToString(value)

	}

	return m, nil
}

// ListOpenstackCluster lists OpenstackCluster with list spec.
func ListOpenstackCluster(ctx context.Context, tx *sql.Tx, request *models.ListOpenstackClusterRequest) (response *models.ListOpenstackClusterResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "openstack_cluster"
	qb.Fields = OpenstackClusterFields
	qb.RefFields = OpenstackClusterRefFields
	qb.BackRefFields = OpenstackClusterBackRefFields
	result := []*models.OpenstackCluster{}

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
		m, err := scanOpenstackCluster(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListOpenstackClusterResponse{
		OpenstackClusters: result,
	}
	return response, nil
}

// UpdateOpenstackCluster updates a resource
func UpdateOpenstackCluster(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateOpenstackClusterRequest,
) error {
	//TODO
	return nil
}

// DeleteOpenstackCluster deletes a resource
func DeleteOpenstackCluster(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteOpenstackClusterRequest) error {
	deleteQuery := deleteOpenstackClusterQuery
	selectQuery := "select count(uuid) from openstack_cluster where uuid = ?"
	var err error
	var count int
	uuid := request.ID
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
