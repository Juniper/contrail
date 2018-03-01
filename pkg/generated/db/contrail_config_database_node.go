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

const insertContrailConfigDatabaseNodeQuery = "insert into `contrail_config_database_node` (`uuid`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailConfigDatabaseNodeQuery = "delete from `contrail_config_database_node` where uuid = ?"

// ContrailConfigDatabaseNodeFields is db columns for ContrailConfigDatabaseNode
var ContrailConfigDatabaseNodeFields = []string{
	"uuid",
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
}

// ContrailConfigDatabaseNodeRefFields is db reference fields for ContrailConfigDatabaseNode
var ContrailConfigDatabaseNodeRefFields = map[string][]string{

	"node": []string{
	// <schema.Schema Value>

	},
}

// ContrailConfigDatabaseNodeBackRefFields is db back reference fields for ContrailConfigDatabaseNode
var ContrailConfigDatabaseNodeBackRefFields = map[string][]string{}

// ContrailConfigDatabaseNodeParentTypes is possible parents for ContrailConfigDatabaseNode
var ContrailConfigDatabaseNodeParents = []string{

	"contrail_cluster",
}

const insertContrailConfigDatabaseNodeNodeQuery = "insert into `ref_contrail_config_database_node_node` (`from`, `to` ) values (?, ?);"

// CreateContrailConfigDatabaseNode inserts ContrailConfigDatabaseNode to DB
func (db *DB) createContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.CreateContrailConfigDatabaseNodeRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ContrailConfigDatabaseNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailConfigDatabaseNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailConfigDatabaseNodeQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
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
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtNodeRef, err := tx.Prepare(insertContrailConfigDatabaseNodeNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NodeRefs create statement failed")
	}
	defer stmtNodeRef.Close()
	for _, ref := range model.NodeRefs {

		_, err = stmtNodeRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NodeRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "contrail_config_database_node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "contrail_config_database_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanContrailConfigDatabaseNode(values map[string]interface{}) (*models.ContrailConfigDatabaseNode, error) {
	m := models.MakeContrailConfigDatabaseNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_state"]; ok {

		m.ProvisioningState = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		m.ProvisioningStartTime = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		m.ProvisioningProgressStage = schema.InterfaceToString(value)

	}

	if value, ok := values["provisioning_progress"]; ok {

		m.ProvisioningProgress = schema.InterfaceToInt64(value)

	}

	if value, ok := values["provisioning_log"]; ok {

		m.ProvisioningLog = schema.InterfaceToString(value)

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

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

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

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_node"]; ok {
		var references []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ContrailConfigDatabaseNodeNodeRef{}
			referenceModel.UUID = uuid
			m.NodeRefs = append(m.NodeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListContrailConfigDatabaseNode lists ContrailConfigDatabaseNode with list spec.
func (db *DB) listContrailConfigDatabaseNode(ctx context.Context, request *models.ListContrailConfigDatabaseNodeRequest) (response *models.ListContrailConfigDatabaseNodeResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "contrail_config_database_node"
	qb.Fields = ContrailConfigDatabaseNodeFields
	qb.RefFields = ContrailConfigDatabaseNodeRefFields
	qb.BackRefFields = ContrailConfigDatabaseNodeBackRefFields
	result := []*models.ContrailConfigDatabaseNode{}

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
		m, err := scanContrailConfigDatabaseNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListContrailConfigDatabaseNodeResponse{
		ContrailConfigDatabaseNodes: result,
	}
	return response, nil
}

// UpdateContrailConfigDatabaseNode updates a resource
func (db *DB) updateContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.UpdateContrailConfigDatabaseNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteContrailConfigDatabaseNode deletes a resource
func (db *DB) deleteContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.DeleteContrailConfigDatabaseNodeRequest) error {
	deleteQuery := deleteContrailConfigDatabaseNodeQuery
	selectQuery := "select count(uuid) from contrail_config_database_node where uuid = ?"
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

//CreateContrailConfigDatabaseNode handle a Create API
func (db *DB) CreateContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.CreateContrailConfigDatabaseNodeRequest) (*models.CreateContrailConfigDatabaseNodeResponse, error) {
	model := request.ContrailConfigDatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createContrailConfigDatabaseNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_database_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailConfigDatabaseNodeResponse{
		ContrailConfigDatabaseNode: request.ContrailConfigDatabaseNode,
	}, nil
}

//UpdateContrailConfigDatabaseNode handles a Update request.
func (db *DB) UpdateContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.UpdateContrailConfigDatabaseNodeRequest) (*models.UpdateContrailConfigDatabaseNodeResponse, error) {
	model := request.ContrailConfigDatabaseNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateContrailConfigDatabaseNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_config_database_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailConfigDatabaseNodeResponse{
		ContrailConfigDatabaseNode: model,
	}, nil
}

//DeleteContrailConfigDatabaseNode delete a resource.
func (db *DB) DeleteContrailConfigDatabaseNode(ctx context.Context, request *models.DeleteContrailConfigDatabaseNodeRequest) (*models.DeleteContrailConfigDatabaseNodeResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteContrailConfigDatabaseNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailConfigDatabaseNodeResponse{
		ID: request.ID,
	}, nil
}

//GetContrailConfigDatabaseNode a Get request.
func (db *DB) GetContrailConfigDatabaseNode(ctx context.Context, request *models.GetContrailConfigDatabaseNodeRequest) (response *models.GetContrailConfigDatabaseNodeResponse, err error) {
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
	listRequest := &models.ListContrailConfigDatabaseNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailConfigDatabaseNodeResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listContrailConfigDatabaseNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailConfigDatabaseNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailConfigDatabaseNodeResponse{
		ContrailConfigDatabaseNode: result.ContrailConfigDatabaseNodes[0],
	}
	return response, nil
}

//ListContrailConfigDatabaseNode handles a List service Request.
func (db *DB) ListContrailConfigDatabaseNode(
	ctx context.Context,
	request *models.ListContrailConfigDatabaseNodeRequest) (response *models.ListContrailConfigDatabaseNodeResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listContrailConfigDatabaseNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
