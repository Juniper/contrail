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

// NodeFields is db columns for Node
var NodeFields = []string{
	"uuid",
	"username",
	"type",
	"ssh_key",
	"private_machine_state",
	"private_machine_properties",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"password",
	"parent_uuid",
	"parent_type",
	"mac_address",
	"ip_address",
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
	"hostname",
	"gcp_machine_type",
	"gcp_image",
	"fq_name",
	"ipmi_username",
	"ipmi_password",
	"ipmi_address",
	"deploy_ramdisk",
	"deploy_kernel",
	"display_name",
	"configuration_version",
	"memory_mb",
	"disk_gb",
	"cpu_count",
	"cpu_arch",
	"aws_instance_type",
	"aws_ami",
	"key_value_pair",
}

// NodeRefFields is db reference fields for Node
var NodeRefFields = map[string][]string{

	"keypair": []string{
	// <schema.Schema Value>

	},
}

// NodeBackRefFields is db back reference fields for Node
var NodeBackRefFields = map[string][]string{

	"port": []string{
		"uuid",
		"pxe_enabled",
		"share",
		"owner_access",
		"owner",
		"global_access",
		"parent_uuid",
		"parent_type",
		"node_uuid",
		"mac_address",
		"switch_info",
		"switch_id",
		"port_id",
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

// NodeParentTypes is possible parents for Node
var NodeParents = []string{}

// CreateNode inserts Node to DB
// nolint
func (db *DB) createNode(
	ctx context.Context,
	request *models.CreateNodeRequest) error {
	qb := db.queryBuilders["node"]
	tx := GetTransaction(ctx)
	model := request.Node
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		string(model.GetUsername()),
		string(model.GetType()),
		string(model.GetSSHKey()),
		string(model.GetPrivateMachineState()),
		string(model.GetPrivateMachineProperties()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetPassword()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetMacAddress()),
		string(model.GetIPAddress()),
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
		string(model.GetHostname()),
		string(model.GetGCPMachineType()),
		string(model.GetGCPImage()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDriverInfo().GetIpmiUsername()),
		string(model.GetDriverInfo().GetIpmiPassword()),
		string(model.GetDriverInfo().GetIpmiAddress()),
		string(model.GetDriverInfo().GetDeployRamdisk()),
		string(model.GetDriverInfo().GetDeployKernel()),
		string(model.GetDisplayName()),
		int(model.GetConfigurationVersion()),
		int(model.GetBMProperties().GetMemoryMB()),
		int(model.GetBMProperties().GetDiskGB()),
		int(model.GetBMProperties().GetCPUCount()),
		string(model.GetBMProperties().GetCPUArch()),
		string(model.GetAwsInstanceType()),
		string(model.GetAwsAmi()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.KeypairRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("keypair"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "KeypairRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "node",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanNode(values map[string]interface{}) (*models.Node, error) {
	m := models.MakeNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["username"]; ok {

		m.Username = common.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Type = common.InterfaceToString(value)

	}

	if value, ok := values["ssh_key"]; ok {

		m.SSHKey = common.InterfaceToString(value)

	}

	if value, ok := values["private_machine_state"]; ok {

		m.PrivateMachineState = common.InterfaceToString(value)

	}

	if value, ok := values["private_machine_properties"]; ok {

		m.PrivateMachineProperties = common.InterfaceToString(value)

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

	if value, ok := values["password"]; ok {

		m.Password = common.InterfaceToString(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = common.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = common.InterfaceToString(value)

	}

	if value, ok := values["mac_address"]; ok {

		m.MacAddress = common.InterfaceToString(value)

	}

	if value, ok := values["ip_address"]; ok {

		m.IPAddress = common.InterfaceToString(value)

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

	if value, ok := values["hostname"]; ok {

		m.Hostname = common.InterfaceToString(value)

	}

	if value, ok := values["gcp_machine_type"]; ok {

		m.GCPMachineType = common.InterfaceToString(value)

	}

	if value, ok := values["gcp_image"]; ok {

		m.GCPImage = common.InterfaceToString(value)

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

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["memory_mb"]; ok {

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

	if value, ok := values["aws_instance_type"]; ok {

		m.AwsInstanceType = common.InterfaceToString(value)

	}

	if value, ok := values["aws_ami"]; ok {

		m.AwsAmi = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_keypair"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.NodeKeypairRef{}
			referenceModel.UUID = uuid
			m.KeypairRefs = append(m.KeypairRefs, referenceModel)

		}
	}

	if value, ok := values["backref_port"]; ok {
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
			childModel := models.MakePort()
			m.Ports = append(m.Ports, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["pxe_enabled"]; ok && propertyValue != nil {

				childModel.PxeEnabled = common.InterfaceToBool(propertyValue)

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

			if propertyValue, ok := childResourceMap["node_uuid"]; ok && propertyValue != nil {

				childModel.NodeUUID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["mac_address"]; ok && propertyValue != nil {

				childModel.MacAddress = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["switch_info"]; ok && propertyValue != nil {

				childModel.LocalLinkConnection.SwitchInfo = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["switch_id"]; ok && propertyValue != nil {

				childModel.LocalLinkConnection.SwitchID = common.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["port_id"]; ok && propertyValue != nil {

				childModel.LocalLinkConnection.PortID = common.InterfaceToString(propertyValue)

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

// ListNode lists Node with list spec.
func (db *DB) listNode(ctx context.Context, request *models.ListNodeRequest) (response *models.ListNodeResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["node"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.Node{}

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
		m, err := scanNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListNodeResponse{
		Nodes: result,
	}
	return response, nil
}

// UpdateNode updates a resource
func (db *DB) updateNode(
	ctx context.Context,
	request *models.UpdateNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteNode deletes a resource
func (db *DB) deleteNode(
	ctx context.Context,
	request *models.DeleteNodeRequest) error {
	qb := db.queryBuilders["node"]

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

//CreateNode handle a Create API
// nolint
func (db *DB) CreateNode(
	ctx context.Context,
	request *models.CreateNodeRequest) (*models.CreateNodeResponse, error) {
	model := request.Node
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNodeResponse{
		Node: request.Node,
	}, nil
}

//UpdateNode handles a Update request.
func (db *DB) UpdateNode(
	ctx context.Context,
	request *models.UpdateNodeRequest) (*models.UpdateNodeResponse, error) {
	model := request.Node
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNodeResponse{
		Node: model,
	}, nil
}

//DeleteNode delete a resource.
func (db *DB) DeleteNode(ctx context.Context, request *models.DeleteNodeRequest) (*models.DeleteNodeResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNodeResponse{
		ID: request.ID,
	}, nil
}

//GetNode a Get request.
func (db *DB) GetNode(ctx context.Context, request *models.GetNodeRequest) (response *models.GetNodeResponse, err error) {
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
	listRequest := &models.ListNodeRequest{
		Spec: spec,
	}
	var result *models.ListNodeResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Nodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNodeResponse{
		Node: result.Nodes[0],
	}
	return response, nil
}

//ListNode handles a List service Request.
// nolint
func (db *DB) ListNode(
	ctx context.Context,
	request *models.ListNodeRequest) (response *models.ListNodeResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
