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

const insertContrailStorageNodeQuery = "insert into `contrail_storage_node` (`uuid`,`storage_backend_bond_interface_members`,`storage_access_bond_interface_members`,`provisioning_state`,`provisioning_start_time`,`provisioning_progress_stage`,`provisioning_progress`,`provisioning_log`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`osd_drives`,`journal_drives`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailStorageNodeQuery = "delete from `contrail_storage_node` where uuid = ?"

// ContrailStorageNodeFields is db columns for ContrailStorageNode
var ContrailStorageNodeFields = []string{
	"uuid",
	"storage_backend_bond_interface_members",
	"storage_access_bond_interface_members",
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
	"osd_drives",
	"journal_drives",
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

// ContrailStorageNodeRefFields is db reference fields for ContrailStorageNode
var ContrailStorageNodeRefFields = map[string][]string{

	"node": []string{
	// <schema.Schema Value>

	},
}

// ContrailStorageNodeBackRefFields is db back reference fields for ContrailStorageNode
var ContrailStorageNodeBackRefFields = map[string][]string{}

// ContrailStorageNodeParentTypes is possible parents for ContrailStorageNode
var ContrailStorageNodeParents = []string{

	"contrail_cluster",
}

const insertContrailStorageNodeNodeQuery = "insert into `ref_contrail_storage_node_node` (`from`, `to` ) values (?, ?);"

// CreateContrailStorageNode inserts ContrailStorageNode to DB
func (db *DB) createContrailStorageNode(
	ctx context.Context,
	request *models.CreateContrailStorageNodeRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.ContrailStorageNode
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailStorageNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailStorageNodeQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetStorageBackendBondInterfaceMembers()),
		string(model.GetStorageAccessBondInterfaceMembers()),
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
		string(model.GetOsdDrives()),
		string(model.GetJournalDrives()),
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

	stmtNodeRef, err := tx.Prepare(insertContrailStorageNodeNodeQuery)
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
		Type:   "contrail_storage_node",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "contrail_storage_node", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanContrailStorageNode(values map[string]interface{}) (*models.ContrailStorageNode, error) {
	m := models.MakeContrailStorageNode()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["storage_backend_bond_interface_members"]; ok {

		m.StorageBackendBondInterfaceMembers = schema.InterfaceToString(value)

	}

	if value, ok := values["storage_access_bond_interface_members"]; ok {

		m.StorageAccessBondInterfaceMembers = schema.InterfaceToString(value)

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

	if value, ok := values["osd_drives"]; ok {

		m.OsdDrives = schema.InterfaceToString(value)

	}

	if value, ok := values["journal_drives"]; ok {

		m.JournalDrives = schema.InterfaceToString(value)

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
			referenceModel := &models.ContrailStorageNodeNodeRef{}
			referenceModel.UUID = uuid
			m.NodeRefs = append(m.NodeRefs, referenceModel)

		}
	}

	return m, nil
}

// ListContrailStorageNode lists ContrailStorageNode with list spec.
func (db *DB) listContrailStorageNode(ctx context.Context, request *models.ListContrailStorageNodeRequest) (response *models.ListContrailStorageNodeResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "contrail_storage_node"
	qb.Fields = ContrailStorageNodeFields
	qb.RefFields = ContrailStorageNodeRefFields
	qb.BackRefFields = ContrailStorageNodeBackRefFields
	result := []*models.ContrailStorageNode{}

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
		m, err := scanContrailStorageNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListContrailStorageNodeResponse{
		ContrailStorageNodes: result,
	}
	return response, nil
}

// UpdateContrailStorageNode updates a resource
func (db *DB) updateContrailStorageNode(
	ctx context.Context,
	request *models.UpdateContrailStorageNodeRequest,
) error {
	//TODO
	return nil
}

// DeleteContrailStorageNode deletes a resource
func (db *DB) deleteContrailStorageNode(
	ctx context.Context,
	request *models.DeleteContrailStorageNodeRequest) error {
	deleteQuery := deleteContrailStorageNodeQuery
	selectQuery := "select count(uuid) from contrail_storage_node where uuid = ?"
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

//CreateContrailStorageNode handle a Create API
func (db *DB) CreateContrailStorageNode(
	ctx context.Context,
	request *models.CreateContrailStorageNodeRequest) (*models.CreateContrailStorageNodeResponse, error) {
	model := request.ContrailStorageNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createContrailStorageNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_storage_node",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateContrailStorageNodeResponse{
		ContrailStorageNode: request.ContrailStorageNode,
	}, nil
}

//UpdateContrailStorageNode handles a Update request.
func (db *DB) UpdateContrailStorageNode(
	ctx context.Context,
	request *models.UpdateContrailStorageNodeRequest) (*models.UpdateContrailStorageNodeResponse, error) {
	model := request.ContrailStorageNode
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateContrailStorageNode(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "contrail_storage_node",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateContrailStorageNodeResponse{
		ContrailStorageNode: model,
	}, nil
}

//DeleteContrailStorageNode delete a resource.
func (db *DB) DeleteContrailStorageNode(ctx context.Context, request *models.DeleteContrailStorageNodeRequest) (*models.DeleteContrailStorageNodeResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteContrailStorageNode(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteContrailStorageNodeResponse{
		ID: request.ID,
	}, nil
}

//GetContrailStorageNode a Get request.
func (db *DB) GetContrailStorageNode(ctx context.Context, request *models.GetContrailStorageNodeRequest) (response *models.GetContrailStorageNodeResponse, err error) {
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
	listRequest := &models.ListContrailStorageNodeRequest{
		Spec: spec,
	}
	var result *models.ListContrailStorageNodeResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listContrailStorageNode(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.ContrailStorageNodes) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetContrailStorageNodeResponse{
		ContrailStorageNode: result.ContrailStorageNodes[0],
	}
	return response, nil
}

//ListContrailStorageNode handles a List service Request.
func (db *DB) ListContrailStorageNode(
	ctx context.Context,
	request *models.ListContrailStorageNodeRequest) (response *models.ListContrailStorageNodeResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listContrailStorageNode(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
