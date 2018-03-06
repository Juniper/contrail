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

// BaremetalNodeFields is db columns for BaremetalNode
var BaremetalNodeFields = []string{
	"uuid",
	"updated_at",
	"target_provision_state",
	"target_power_state",
	"provision_state",
	"power_state",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"name",
	"maintenance_reason",
	"maintenance",
	"last_error",
	"instance_uuid",
	"vcpus",
	"swap_mb",
	"root_gb",
	"nova_host_id",
	"memory_mb",
	"local_gb",
	"image_source",
	"display_name",
	"capabilities",
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
	"ipmi_username",
	"ipmi_password",
	"ipmi_address",
	"deploy_ramdisk",
	"deploy_kernel",
	"_display_name",
	"created_at",
	"console_enabled",
	"bm_properties_memory_mb",
	"disk_gb",
	"cpu_count",
	"cpu_arch",
	"key_value_pair",
}

// BaremetalNodeRefFields is db reference fields for BaremetalNode
var BaremetalNodeRefFields = map[string][]string{}

// BaremetalNodeBackRefFields is db back reference fields for BaremetalNode
var BaremetalNodeBackRefFields = map[string][]string{}

// BaremetalNodeParentTypes is possible parents for BaremetalNode
var BaremetalNodeParents = []string{}

// CreateBaremetalNode inserts BaremetalNode to DB
// nolint
func (db *DB) createBaremetalNode(
	ctx context.Context,
	request *models.CreateBaremetalNodeRequest) error {
	qb := db.queryBuilders["baremetal_node"]
	tx := GetTransaction(ctx)
	model := request.BaremetalNode
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		string(model.GetUpdatedAt()),
		string(model.GetTargetProvisionState()),
		string(model.GetTargetPowerState()),
		string(model.GetProvisionState()),
		string(model.GetPowerState()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetName()),
		string(model.GetMaintenanceReason()),
		bool(model.GetMaintenance()),
		string(model.GetLastError()),
		string(model.GetInstanceUUID()),
		string(model.GetInstanceInfo().GetVcpus()),
		string(model.GetInstanceInfo().GetSwapMB()),
		string(model.GetInstanceInfo().GetRootGB()),
		string(model.GetInstanceInfo().GetNovaHostID()),
		string(model.GetInstanceInfo().GetMemoryMB()),
		string(model.GetInstanceInfo().GetLocalGB()),
		string(model.GetInstanceInfo().GetImageSource()),
		string(model.GetInstanceInfo().GetDisplayName()),
		string(model.GetInstanceInfo().GetCapabilities()),
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
		string(model.GetDriverInfo().GetIpmiUsername()),
		string(model.GetDriverInfo().GetIpmiPassword()),
		string(model.GetDriverInfo().GetIpmiAddress()),
		string(model.GetDriverInfo().GetDeployRamdisk()),
		string(model.GetDriverInfo().GetDeployKernel()),
		string(model.GetDisplayName()),
		string(model.GetCreatedAt()),
		bool(model.GetConsoleEnabled()),
		int(model.GetBMProperties().GetMemoryMB()),
		int(model.GetBMProperties().GetDiskGB()),
		int(model.GetBMProperties().GetCPUCount()),
		string(model.GetBMProperties().GetCPUArch()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "baremetal_node",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "baremetal_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBaremetalNode(values map[string]interface{}) (*models.BaremetalNode, error) {
	m := models.MakeBaremetalNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["updated_at"]; ok {

		m.UpdatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["target_provision_state"]; ok {

		m.TargetProvisionState = common.InterfaceToString(value)

	}

	if value, ok := values["target_power_state"]; ok {

		m.TargetPowerState = common.InterfaceToString(value)

	}

	if value, ok := values["provision_state"]; ok {

		m.ProvisionState = common.InterfaceToString(value)

	}

	if value, ok := values["power_state"]; ok {

		m.PowerState = common.InterfaceToString(value)

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

	if value, ok := values["name"]; ok {

		m.Name = common.InterfaceToString(value)

	}

	if value, ok := values["maintenance_reason"]; ok {

		m.MaintenanceReason = common.InterfaceToString(value)

	}

	if value, ok := values["maintenance"]; ok {

		m.Maintenance = common.InterfaceToBool(value)

	}

	if value, ok := values["last_error"]; ok {

		m.LastError = common.InterfaceToString(value)

	}

	if value, ok := values["instance_uuid"]; ok {

		m.InstanceUUID = common.InterfaceToString(value)

	}

	if value, ok := values["vcpus"]; ok {

		m.InstanceInfo.Vcpus = common.InterfaceToString(value)

	}

	if value, ok := values["swap_mb"]; ok {

		m.InstanceInfo.SwapMB = common.InterfaceToString(value)

	}

	if value, ok := values["root_gb"]; ok {

		m.InstanceInfo.RootGB = common.InterfaceToString(value)

	}

	if value, ok := values["nova_host_id"]; ok {

		m.InstanceInfo.NovaHostID = common.InterfaceToString(value)

	}

	if value, ok := values["memory_mb"]; ok {

		m.InstanceInfo.MemoryMB = common.InterfaceToString(value)

	}

	if value, ok := values["local_gb"]; ok {

		m.InstanceInfo.LocalGB = common.InterfaceToString(value)

	}

	if value, ok := values["image_source"]; ok {

		m.InstanceInfo.ImageSource = common.InterfaceToString(value)

	}

	if value, ok := values["display_name"]; ok {

		m.InstanceInfo.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["capabilities"]; ok {

		m.InstanceInfo.Capabilities = common.InterfaceToString(value)

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

	if value, ok := values["ipmi_username"]; ok {

		m.DriverInfo.IpmiUsername = common.InterfaceToString(value)

	}

	if value, ok := values["ipmi_password"]; ok {

		m.DriverInfo.IpmiPassword = common.InterfaceToString(value)

	}

	if value, ok := values["ipmi_address"]; ok {

		m.DriverInfo.IpmiAddress = common.InterfaceToString(value)

	}

	if value, ok := values["deploy_ramdisk"]; ok {

		m.DriverInfo.DeployRamdisk = common.InterfaceToString(value)

	}

	if value, ok := values["deploy_kernel"]; ok {

		m.DriverInfo.DeployKernel = common.InterfaceToString(value)

	}

	if value, ok := values["_display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["created_at"]; ok {

		m.CreatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["console_enabled"]; ok {

		m.ConsoleEnabled = common.InterfaceToBool(value)

	}

	if value, ok := values["bm_properties_memory_mb"]; ok {

		m.BMProperties.MemoryMB = common.InterfaceToInt64(value)

	}

	if value, ok := values["disk_gb"]; ok {

		m.BMProperties.DiskGB = common.InterfaceToInt64(value)

	}

	if value, ok := values["cpu_count"]; ok {

		m.BMProperties.CPUCount = common.InterfaceToInt64(value)

	}

	if value, ok := values["cpu_arch"]; ok {

		m.BMProperties.CPUArch = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListBaremetalNode lists BaremetalNode with list spec.
func (db *DB) listBaremetalNode(ctx context.Context, request *models.ListBaremetalNodeRequest) (response *models.ListBaremetalNodeResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["baremetal_node"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.BaremetalNode{}

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
		m, err := scanBaremetalNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBaremetalNodeResponse{
		BaremetalNodes: result,
	}
	return response, nil
}

// UpdateBaremetalNode updates a resource
func (db *DB) updateBaremetalNode(
	ctx context.Context,
	request *models.UpdateBaremetalNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteBaremetalNode deletes a resource
func (db *DB) deleteBaremetalNode(
	ctx context.Context,
	request *models.DeleteBaremetalNodeRequest) error {
	qb := db.queryBuilders["baremetal_node"]

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

//CreateBaremetalNode handle a Create API
// nolint
func (db *DB) CreateBaremetalNode(
	ctx context.Context,
	request *models.CreateBaremetalNodeRequest) (*models.CreateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createBaremetalNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBaremetalNodeResponse{
		BaremetalNode: request.BaremetalNode,
	}, nil
}

//UpdateBaremetalNode handles a Update request.
func (db *DB) UpdateBaremetalNode(
	ctx context.Context,
	request *models.UpdateBaremetalNodeRequest) (*models.UpdateBaremetalNodeResponse, error) {
	model := request.BaremetalNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateBaremetalNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBaremetalNodeResponse{
		BaremetalNode: model,
	}, nil
}

//DeleteBaremetalNode delete a resource.
func (db *DB) DeleteBaremetalNode(ctx context.Context, request *models.DeleteBaremetalNodeRequest) (*models.DeleteBaremetalNodeResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteBaremetalNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBaremetalNodeResponse{
		ID: request.ID,
	}, nil
}

//GetBaremetalNode a Get request.
func (db *DB) GetBaremetalNode(ctx context.Context, request *models.GetBaremetalNodeRequest) (response *models.GetBaremetalNodeResponse, err error) {
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
	listRequest := &models.ListBaremetalNodeRequest{
		Spec: spec,
	}
	var result *models.ListBaremetalNodeResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listBaremetalNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BaremetalNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBaremetalNodeResponse{
		BaremetalNode: result.BaremetalNodes[0],
	}
	return response, nil
}

//ListBaremetalNode handles a List service Request.
// nolint
func (db *DB) ListBaremetalNode(
	ctx context.Context,
	request *models.ListBaremetalNodeRequest) (response *models.ListBaremetalNodeResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listBaremetalNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
