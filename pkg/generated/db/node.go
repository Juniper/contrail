package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertNodeQuery = "insert into `node` (`uuid`,`username`,`type`,`ssh_key`,`private_machine_state`,`private_machine_properties`,`share`,`owner_access`,`owner`,`global_access`,`password`,`parent_uuid`,`parent_type`,`mac_address`,`ipmi_username`,`ipmi_password`,`ipmi_address`,`ip_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`hostname`,`gcp_machine_type`,`gcp_image`,`fq_name`,`display_name`,`aws_instance_type`,`aws_ami`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteNodeQuery = "delete from `node` where uuid = ?"

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
	"ipmi_username",
	"ipmi_password",
	"ipmi_address",
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
	"display_name",
	"aws_instance_type",
	"aws_ami",
	"key_value_pair",
}

// NodeRefFields is db reference fields for Node
var NodeRefFields = map[string][]string{}

// NodeBackRefFields is db back reference fields for Node
var NodeBackRefFields = map[string][]string{}

// NodeParentTypes is possible parents for Node
var NodeParents = []string{}

// CreateNode inserts Node to DB
func (db *DB) createNode(
	ctx context.Context,
	request *models.CreateNodeRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.Node
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNodeQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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
		string(model.GetIpmiUsername()),
		string(model.GetIpmiPassword()),
		string(model.GetIpmiAddress()),
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
		string(model.GetDisplayName()),
		string(model.GetAwsInstanceType()),
		string(model.GetAwsAmi()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "node", model.UUID, model.GetPerms2().GetShare())
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

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["username"]; ok {

		m.Username = schema.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Type = schema.InterfaceToString(value)

	}

	if value, ok := values["ssh_key"]; ok {

		m.SSHKey = schema.InterfaceToString(value)

	}

	if value, ok := values["private_machine_state"]; ok {

		m.PrivateMachineState = schema.InterfaceToString(value)

	}

	if value, ok := values["private_machine_properties"]; ok {

		m.PrivateMachineProperties = schema.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["password"]; ok {

		m.Password = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["mac_address"]; ok {

		m.MacAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["ipmi_username"]; ok {

		m.IpmiUsername = schema.InterfaceToString(value)

	}

	if value, ok := values["ipmi_password"]; ok {

		m.IpmiPassword = schema.InterfaceToString(value)

	}

	if value, ok := values["ipmi_address"]; ok {

		m.IpmiAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["ip_address"]; ok {

		m.IPAddress = schema.InterfaceToString(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["hostname"]; ok {

		m.Hostname = schema.InterfaceToString(value)

	}

	if value, ok := values["gcp_machine_type"]; ok {

		m.GCPMachineType = schema.InterfaceToString(value)

	}

	if value, ok := values["gcp_image"]; ok {

		m.GCPImage = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["aws_instance_type"]; ok {

		m.AwsInstanceType = schema.InterfaceToString(value)

	}

	if value, ok := values["aws_ami"]; ok {

		m.AwsAmi = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListNode lists Node with list spec.
func (db *DB) listNode(ctx context.Context, request *models.ListNodeRequest) (response *models.ListNodeResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "node"
	qb.Fields = NodeFields
	qb.RefFields = NodeRefFields
	qb.BackRefFields = NodeBackRefFields
	result := []*models.Node{}

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
	deleteQuery := deleteNodeQuery
	selectQuery := "select count(uuid) from node where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
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

//CreateNode handle a Create API
func (db *DB) CreateNode(
	ctx context.Context,
	request *models.CreateNodeRequest) (*models.CreateNodeResponse, error) {
	model := request.Node
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
	if err := common.DoInTransaction(
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
		Limit: 1,
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
	if err := common.DoInTransaction(
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
func (db *DB) ListNode(
	ctx context.Context,
	request *models.ListNodeRequest) (response *models.ListNodeResponse, err error) {
	if err := common.DoInTransaction(
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
