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

const insertInterfaceRouteTableQuery = "insert into `interface_route_table` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`route`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteInterfaceRouteTableQuery = "delete from `interface_route_table` where uuid = ?"

// InterfaceRouteTableFields is db columns for InterfaceRouteTable
var InterfaceRouteTableFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"route",
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

// InterfaceRouteTableRefFields is db reference fields for InterfaceRouteTable
var InterfaceRouteTableRefFields = map[string][]string{

	"service_instance": []string{
		// <schema.Schema Value>
		"interface_type",
	},
}

// InterfaceRouteTableBackRefFields is db back reference fields for InterfaceRouteTable
var InterfaceRouteTableBackRefFields = map[string][]string{}

// InterfaceRouteTableParentTypes is possible parents for InterfaceRouteTable
var InterfaceRouteTableParents = []string{

	"project",
}

const insertInterfaceRouteTableServiceInstanceQuery = "insert into `ref_interface_route_table_service_instance` (`from`, `to` ,`interface_type`) values (?, ?,?);"

// CreateInterfaceRouteTable inserts InterfaceRouteTable to DB
// nolint
func (db *DB) createInterfaceRouteTable(
	ctx context.Context,
	request *models.CreateInterfaceRouteTableRequest) error {
	tx := GetTransaction(ctx)
	model := request.InterfaceRouteTable
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInterfaceRouteTableQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertInterfaceRouteTableQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		common.MustJSON(model.GetInterfaceRouteTableRoutes().GetRoute()),
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

	stmtServiceInstanceRef, err := tx.Prepare(insertInterfaceRouteTableServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceInstanceRefs create statement failed")
	}
	defer stmtServiceInstanceRef.Close()
	for _, ref := range model.ServiceInstanceRefs {

		if ref.Attr == nil {
			ref.Attr = &models.ServiceInterfaceTag{}
		}

		_, err = stmtServiceInstanceRef.ExecContext(ctx, model.UUID, ref.UUID, string(ref.Attr.GetInterfaceType()))
		if err != nil {
			return errors.Wrap(err, "ServiceInstanceRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "interface_route_table",
		FQName: model.FQName,
	}
	err = CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = CreateSharing(tx, "interface_route_table", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

// nolint
func scanInterfaceRouteTable(values map[string]interface{}) (*models.InterfaceRouteTable, error) {
	m := models.MakeInterfaceRouteTable()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

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

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.InterfaceRouteTableRoutes.Route)

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

	if value, ok := values["display_name"]; ok {

		m.DisplayName = common.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_service_instance"]; ok {
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
			referenceModel := &models.InterfaceRouteTableServiceInstanceRef{}
			referenceModel.UUID = uuid
			m.ServiceInstanceRefs = append(m.ServiceInstanceRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListInterfaceRouteTable lists InterfaceRouteTable with list spec.
// nolint
func (db *DB) listInterfaceRouteTable(ctx context.Context, request *models.ListInterfaceRouteTableRequest) (response *models.ListInterfaceRouteTableResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)
	qb := &ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "interface_route_table"
	qb.Fields = InterfaceRouteTableFields
	qb.RefFields = InterfaceRouteTableRefFields
	qb.BackRefFields = InterfaceRouteTableBackRefFields
	result := []*models.InterfaceRouteTable{}

	if spec.ParentFQName != nil {
		parentMetaData, err := GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
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
		m, err := scanInterfaceRouteTable(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListInterfaceRouteTableResponse{
		InterfaceRouteTables: result,
	}
	return response, nil
}

// UpdateInterfaceRouteTable updates a resource
// nolint
func (db *DB) updateInterfaceRouteTable(
	ctx context.Context,
	request *models.UpdateInterfaceRouteTableRequest,
) error {
	//TODO
	return nil
}

// DeleteInterfaceRouteTable deletes a resource
// nolint
func (db *DB) deleteInterfaceRouteTable(
	ctx context.Context,
	request *models.DeleteInterfaceRouteTableRequest) error {
	deleteQuery := deleteInterfaceRouteTableQuery
	selectQuery := "select count(uuid) from interface_route_table where uuid = ?"
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

	err = DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}

//CreateInterfaceRouteTable handle a Create API
// nolint
func (db *DB) CreateInterfaceRouteTable(
	ctx context.Context,
	request *models.CreateInterfaceRouteTableRequest) (*models.CreateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createInterfaceRouteTable(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateInterfaceRouteTableResponse{
		InterfaceRouteTable: request.InterfaceRouteTable,
	}, nil
}

//UpdateInterfaceRouteTable handles a Update request.
func (db *DB) UpdateInterfaceRouteTable(
	ctx context.Context,
	request *models.UpdateInterfaceRouteTableRequest) (*models.UpdateInterfaceRouteTableResponse, error) {
	model := request.InterfaceRouteTable
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateInterfaceRouteTable(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "interface_route_table",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateInterfaceRouteTableResponse{
		InterfaceRouteTable: model,
	}, nil
}

//DeleteInterfaceRouteTable delete a resource.
func (db *DB) DeleteInterfaceRouteTable(ctx context.Context, request *models.DeleteInterfaceRouteTableRequest) (*models.DeleteInterfaceRouteTableResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteInterfaceRouteTable(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteInterfaceRouteTableResponse{
		ID: request.ID,
	}, nil
}

//GetInterfaceRouteTable a Get request.
// nolint
func (db *DB) GetInterfaceRouteTable(ctx context.Context, request *models.GetInterfaceRouteTableRequest) (response *models.GetInterfaceRouteTableResponse, err error) {
	spec := &models.ListSpec{
		Limit: 1,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListInterfaceRouteTableRequest{
		Spec: spec,
	}
	var result *models.ListInterfaceRouteTableResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listInterfaceRouteTable(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.InterfaceRouteTables) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetInterfaceRouteTableResponse{
		InterfaceRouteTable: result.InterfaceRouteTables[0],
	}
	return response, nil
}

//ListInterfaceRouteTable handles a List service Request.
// nolint
func (db *DB) ListInterfaceRouteTable(
	ctx context.Context,
	request *models.ListInterfaceRouteTableRequest) (response *models.ListInterfaceRouteTableResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listInterfaceRouteTable(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
