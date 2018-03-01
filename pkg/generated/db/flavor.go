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

const insertFlavorQuery = "insert into `flavor` (`vcpus`,`uuid`,`swap`,`rxtx_factor`,`ram`,`property`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`name`,`type`,`rel`,`href`,`is_public`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`id`,`fq_name`,`ephemeral`,`display_name`,`disk`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteFlavorQuery = "delete from `flavor` where uuid = ?"

// FlavorFields is db columns for Flavor
var FlavorFields = []string{
	"vcpus",
	"uuid",
	"swap",
	"rxtx_factor",
	"ram",
	"property",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"name",
	"type",
	"rel",
	"href",
	"is_public",
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
	"id",
	"fq_name",
	"ephemeral",
	"display_name",
	"disk",
	"key_value_pair",
}

// FlavorRefFields is db reference fields for Flavor
var FlavorRefFields = map[string][]string{}

// FlavorBackRefFields is db back reference fields for Flavor
var FlavorBackRefFields = map[string][]string{}

// FlavorParentTypes is possible parents for Flavor
var FlavorParents = []string{}

// CreateFlavor inserts Flavor to DB
func (db *DB) createFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.Flavor
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertFlavorQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertFlavorQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, int(model.GetVcpus()),
		string(model.GetUUID()),
		int(model.GetSwap()),
		int(model.GetRXTXFactor()),
		int(model.GetRAM()),
		string(model.GetProperty()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetName()),
		string(model.GetLinks().GetType()),
		string(model.GetLinks().GetRel()),
		string(model.GetLinks().GetHref()),
		bool(model.GetIsPublic()),
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
		string(model.GetID()),
		common.MustJSON(model.GetFQName()),
		int(model.GetEphemeral()),
		string(model.GetDisplayName()),
		int(model.GetDisk()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "flavor",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "flavor", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanFlavor(values map[string]interface{}) (*models.Flavor, error) {
	m := models.MakeFlavor()

	if value, ok := values["vcpus"]; ok {

		m.Vcpus = schema.InterfaceToInt64(value)

	}

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["swap"]; ok {

		m.Swap = schema.InterfaceToInt64(value)

	}

	if value, ok := values["rxtx_factor"]; ok {

		m.RXTXFactor = schema.InterfaceToInt64(value)

	}

	if value, ok := values["ram"]; ok {

		m.RAM = schema.InterfaceToInt64(value)

	}

	if value, ok := values["property"]; ok {

		m.Property = schema.InterfaceToString(value)

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

	if value, ok := values["name"]; ok {

		m.Name = schema.InterfaceToString(value)

	}

	if value, ok := values["type"]; ok {

		m.Links.Type = schema.InterfaceToString(value)

	}

	if value, ok := values["rel"]; ok {

		m.Links.Rel = schema.InterfaceToString(value)

	}

	if value, ok := values["href"]; ok {

		m.Links.Href = schema.InterfaceToString(value)

	}

	if value, ok := values["is_public"]; ok {

		m.IsPublic = schema.InterfaceToBool(value)

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

	if value, ok := values["id"]; ok {

		m.ID = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["ephemeral"]; ok {

		m.Ephemeral = schema.InterfaceToInt64(value)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["disk"]; ok {

		m.Disk = schema.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListFlavor lists Flavor with list spec.
func (db *DB) listFlavor(ctx context.Context, request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "flavor"
	qb.Fields = FlavorFields
	qb.RefFields = FlavorRefFields
	qb.BackRefFields = FlavorBackRefFields
	result := []*models.Flavor{}

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
		m, err := scanFlavor(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListFlavorResponse{
		Flavors: result,
	}
	return response, nil
}

// UpdateFlavor updates a resource
func (db *DB) updateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest,
) error {
	//TODO
	return nil
}

// DeleteFlavor deletes a resource
func (db *DB) deleteFlavor(
	ctx context.Context,
	request *models.DeleteFlavorRequest) error {
	deleteQuery := deleteFlavorQuery
	selectQuery := "select count(uuid) from flavor where uuid = ?"
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

//CreateFlavor handle a Create API
func (db *DB) CreateFlavor(
	ctx context.Context,
	request *models.CreateFlavorRequest) (*models.CreateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createFlavor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateFlavorResponse{
		Flavor: request.Flavor,
	}, nil
}

//UpdateFlavor handles a Update request.
func (db *DB) UpdateFlavor(
	ctx context.Context,
	request *models.UpdateFlavorRequest) (*models.UpdateFlavorResponse, error) {
	model := request.Flavor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateFlavor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "flavor",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateFlavorResponse{
		Flavor: model,
	}, nil
}

//DeleteFlavor delete a resource.
func (db *DB) DeleteFlavor(ctx context.Context, request *models.DeleteFlavorRequest) (*models.DeleteFlavorResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteFlavor(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteFlavorResponse{
		ID: request.ID,
	}, nil
}

//GetFlavor a Get request.
func (db *DB) GetFlavor(ctx context.Context, request *models.GetFlavorRequest) (response *models.GetFlavorResponse, err error) {
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
	listRequest := &models.ListFlavorRequest{
		Spec: spec,
	}
	var result *models.ListFlavorResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listFlavor(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.Flavors) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetFlavorResponse{
		Flavor: result.Flavors[0],
	}
	return response, nil
}

//ListFlavor handles a List service Request.
func (db *DB) ListFlavor(
	ctx context.Context,
	request *models.ListFlavorRequest) (response *models.ListFlavorResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listFlavor(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
