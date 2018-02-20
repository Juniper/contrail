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

const insertBaremetalNodeQuery = "insert into `baremetal_node` (`uuid`,`updated_at`,`target_provision_state`,`target_power_state`,`provision_state`,`power_state`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`maintenance_reason`,`maintenance`,`last_error`,`instance_uuid`,`vcpus`,`swap_mb`,`root_gb`,`nova_host_id`,`memory_mb`,`local_gb`,`image_source`,`display_name`,`capabilities`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`ipmi_username`,`ipmi_password`,`ipmi_address`,`deploy_ramdisk`,`deploy_kernel`,`_display_name`,`created_at`,`console_enabled`,`bm_properties_memory_mb`,`disk_gb`,`cpu_count`,`cpu_arch`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBaremetalNodeQuery = "delete from `baremetal_node` where uuid = ?"

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
func CreateBaremetalNode(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateBaremetalNodeRequest) error {
	model := request.BaremetalNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBaremetalNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBaremetalNodeQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.UUID),
		string(model.UpdatedAt),
		string(model.TargetProvisionState),
		string(model.TargetPowerState),
		string(model.ProvisionState),
		string(model.PowerState),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		string(model.Name),
		string(model.MaintenanceReason),
		bool(model.Maintenance),
		string(model.LastError),
		string(model.InstanceUUID),
		string(model.InstanceInfo.Vcpus),
		string(model.InstanceInfo.SwapMB),
		string(model.InstanceInfo.RootGB),
		string(model.InstanceInfo.NovaHostID),
		string(model.InstanceInfo.MemoryMB),
		string(model.InstanceInfo.LocalGB),
		string(model.InstanceInfo.ImageSource),
		string(model.InstanceInfo.DisplayName),
		string(model.InstanceInfo.Capabilities),
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
		string(model.DriverInfo.IpmiUsername),
		string(model.DriverInfo.IpmiPassword),
		string(model.DriverInfo.IpmiAddress),
		string(model.DriverInfo.DeployRamdisk),
		string(model.DriverInfo.DeployKernel),
		string(model.DisplayName),
		string(model.CreatedAt),
		bool(model.ConsoleEnabled),
		int(model.BMProperties.MemoryMB),
		int(model.BMProperties.DiskGB),
		int(model.BMProperties.CPUCount),
		string(model.BMProperties.CPUArch),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "baremetal_node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "baremetal_node", model.UUID, model.Perms2.Share)
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

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["updated_at"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UpdatedAt = castedValue

	}

	if value, ok := values["target_provision_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.TargetProvisionState = castedValue

	}

	if value, ok := values["target_power_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.TargetPowerState = castedValue

	}

	if value, ok := values["provision_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ProvisionState = castedValue

	}

	if value, ok := values["power_state"]; ok {

		castedValue := common.InterfaceToString(value)

		m.PowerState = castedValue

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

	if value, ok := values["name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Name = castedValue

	}

	if value, ok := values["maintenance_reason"]; ok {

		castedValue := common.InterfaceToString(value)

		m.MaintenanceReason = castedValue

	}

	if value, ok := values["maintenance"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.Maintenance = castedValue

	}

	if value, ok := values["last_error"]; ok {

		castedValue := common.InterfaceToString(value)

		m.LastError = castedValue

	}

	if value, ok := values["instance_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceUUID = castedValue

	}

	if value, ok := values["vcpus"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.Vcpus = castedValue

	}

	if value, ok := values["swap_mb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.SwapMB = castedValue

	}

	if value, ok := values["root_gb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.RootGB = castedValue

	}

	if value, ok := values["nova_host_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.NovaHostID = castedValue

	}

	if value, ok := values["memory_mb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.MemoryMB = castedValue

	}

	if value, ok := values["local_gb"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.LocalGB = castedValue

	}

	if value, ok := values["image_source"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.ImageSource = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.DisplayName = castedValue

	}

	if value, ok := values["capabilities"]; ok {

		castedValue := common.InterfaceToString(value)

		m.InstanceInfo.Capabilities = castedValue

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

	if value, ok := values["ipmi_username"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DriverInfo.IpmiUsername = castedValue

	}

	if value, ok := values["ipmi_password"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DriverInfo.IpmiPassword = castedValue

	}

	if value, ok := values["ipmi_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DriverInfo.IpmiAddress = castedValue

	}

	if value, ok := values["deploy_ramdisk"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DriverInfo.DeployRamdisk = castedValue

	}

	if value, ok := values["deploy_kernel"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DriverInfo.DeployKernel = castedValue

	}

	if value, ok := values["_display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["created_at"]; ok {

		castedValue := common.InterfaceToString(value)

		m.CreatedAt = castedValue

	}

	if value, ok := values["console_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ConsoleEnabled = castedValue

	}

	if value, ok := values["bm_properties_memory_mb"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.BMProperties.MemoryMB = castedValue

	}

	if value, ok := values["disk_gb"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.BMProperties.DiskGB = castedValue

	}

	if value, ok := values["cpu_count"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.BMProperties.CPUCount = castedValue

	}

	if value, ok := values["cpu_arch"]; ok {

		castedValue := common.InterfaceToString(value)

		m.BMProperties.CPUArch = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListBaremetalNode lists BaremetalNode with list spec.
func ListBaremetalNode(ctx context.Context, tx *sql.Tx, request *models.ListBaremetalNodeRequest) (response *models.ListBaremetalNodeResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "baremetal_node"
	qb.Fields = BaremetalNodeFields
	qb.RefFields = BaremetalNodeRefFields
	qb.BackRefFields = BaremetalNodeBackRefFields
	result := models.MakeBaremetalNodeSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
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
func UpdateBaremetalNode(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateBaremetalNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteBaremetalNode deletes a resource
func DeleteBaremetalNode(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteBaremetalNodeRequest) error {
	deleteQuery := deleteBaremetalNodeQuery
	selectQuery := "select count(uuid) from baremetal_node where uuid = ?"
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
