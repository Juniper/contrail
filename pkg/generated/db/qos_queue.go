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

const insertQosQueueQuery = "insert into `qos_queue` (`uuid`,`qos_queue_identifier`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`min_bandwidth`,`max_bandwidth`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteQosQueueQuery = "delete from `qos_queue` where uuid = ?"

// QosQueueFields is db columns for QosQueue
var QosQueueFields = []string{
	"uuid",
	"qos_queue_identifier",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"min_bandwidth",
	"max_bandwidth",
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

// QosQueueRefFields is db reference fields for QosQueue
var QosQueueRefFields = map[string][]string{}

// QosQueueBackRefFields is db back reference fields for QosQueue
var QosQueueBackRefFields = map[string][]string{}

// QosQueueParentTypes is possible parents for QosQueue
var QosQueueParents = []string{

	"global_qos_config",
}

// CreateQosQueue inserts QosQueue to DB
func (db *DB) createQosQueue(
	ctx context.Context,
	request *models.CreateQosQueueRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.QosQueue
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertQosQueueQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertQosQueueQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		int(model.GetQosQueueIdentifier()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		int(model.GetMinBandwidth()),
		int(model.GetMaxBandwidth()),
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "qos_queue",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "qos_queue", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanQosQueue(values map[string]interface{}) (*models.QosQueue, error) {
	m := models.MakeQosQueue()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["qos_queue_identifier"]; ok {

		m.QosQueueIdentifier = schema.InterfaceToInt64(value)

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

	if value, ok := values["min_bandwidth"]; ok {

		m.MinBandwidth = schema.InterfaceToInt64(value)

	}

	if value, ok := values["max_bandwidth"]; ok {

		m.MaxBandwidth = schema.InterfaceToInt64(value)

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

	return m, nil
}

// ListQosQueue lists QosQueue with list spec.
func (db *DB) listQosQueue(ctx context.Context, request *models.ListQosQueueRequest) (response *models.ListQosQueueResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "qos_queue"
	qb.Fields = QosQueueFields
	qb.RefFields = QosQueueRefFields
	qb.BackRefFields = QosQueueBackRefFields
	result := []*models.QosQueue{}

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
		m, err := scanQosQueue(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListQosQueueResponse{
		QosQueues: result,
	}
	return response, nil
}

// UpdateQosQueue updates a resource
func (db *DB) updateQosQueue(
	ctx context.Context,
	request *models.UpdateQosQueueRequest,
) error {
	//TODO
	return nil
}

// DeleteQosQueue deletes a resource
func (db *DB) deleteQosQueue(
	ctx context.Context,
	request *models.DeleteQosQueueRequest) error {
	deleteQuery := deleteQosQueueQuery
	selectQuery := "select count(uuid) from qos_queue where uuid = ?"
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

//CreateQosQueue handle a Create API
func (db *DB) CreateQosQueue(
	ctx context.Context,
	request *models.CreateQosQueueRequest) (*models.CreateQosQueueResponse, error) {
	model := request.QosQueue
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createQosQueue(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateQosQueueResponse{
		QosQueue: request.QosQueue,
	}, nil
}

//UpdateQosQueue handles a Update request.
func (db *DB) UpdateQosQueue(
	ctx context.Context,
	request *models.UpdateQosQueueRequest) (*models.UpdateQosQueueResponse, error) {
	model := request.QosQueue
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateQosQueue(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "qos_queue",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateQosQueueResponse{
		QosQueue: model,
	}, nil
}

//DeleteQosQueue delete a resource.
func (db *DB) DeleteQosQueue(ctx context.Context, request *models.DeleteQosQueueRequest) (*models.DeleteQosQueueResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteQosQueue(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteQosQueueResponse{
		ID: request.ID,
	}, nil
}

//GetQosQueue a Get request.
func (db *DB) GetQosQueue(ctx context.Context, request *models.GetQosQueueRequest) (response *models.GetQosQueueResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListQosQueueRequest{
		Spec: spec,
	}
	var result *models.ListQosQueueResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listQosQueue(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.QosQueues) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetQosQueueResponse{
		QosQueue: result.QosQueues[0],
	}
	return response, nil
}

//ListQosQueue handles a List service Request.
func (db *DB) ListQosQueue(
	ctx context.Context,
	request *models.ListQosQueueRequest) (response *models.ListQosQueueResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listQosQueue(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
