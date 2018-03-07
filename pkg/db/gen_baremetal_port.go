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

// BaremetalPortFields is db columns for BaremetalPort
var BaremetalPortFields = []string{
	"uuid",
	"updated_at",
	"pxe_enabled",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"node",
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
	"created_at",
	"configuration_version",
	"key_value_pair",
}

// BaremetalPortRefFields is db reference fields for BaremetalPort
var BaremetalPortRefFields = map[string][]string{}

// BaremetalPortBackRefFields is db back reference fields for BaremetalPort
var BaremetalPortBackRefFields = map[string][]string{}

// BaremetalPortParentTypes is possible parents for BaremetalPort
var BaremetalPortParents = []string{}

// CreateBaremetalPort inserts BaremetalPort to DB
// nolint
func (db *DB) createBaremetalPort(
	ctx context.Context,
	request *models.CreateBaremetalPortRequest) error {
	qb := db.queryBuilders["baremetal_port"]
	tx := GetTransaction(ctx)
	model := request.BaremetalPort
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		string(model.GetUpdatedAt()),
		bool(model.GetPxeEnabled()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetNode()),
		string(model.GetMacAddress()),
		string(model.GetLocalLinkConnection().GetSwitchInfo()),
		string(model.GetLocalLinkConnection().GetSwitchID()),
		string(model.GetLocalLinkConnection().GetPortID()),
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
		string(model.GetCreatedAt()),
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		log.WithFields(log.Fields{
			"model": model,
			"err":   err}).Debug("create failed")
		return errors.Wrap(err, "create failed")
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "baremetal_port",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "baremetal_port", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBaremetalPort(values map[string]interface{}) (*models.BaremetalPort, error) {
	m := models.MakeBaremetalPort()

	if value, ok := values["uuid"]; ok {

		m.UUID = common.InterfaceToString(value)

	}

	if value, ok := values["updated_at"]; ok {

		m.UpdatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["pxe_enabled"]; ok {

		m.PxeEnabled = common.InterfaceToBool(value)

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

	if value, ok := values["node"]; ok {

		m.Node = common.InterfaceToString(value)

	}

	if value, ok := values["mac_address"]; ok {

		m.MacAddress = common.InterfaceToString(value)

	}

	if value, ok := values["switch_info"]; ok {

		m.LocalLinkConnection.SwitchInfo = common.InterfaceToString(value)

	}

	if value, ok := values["switch_id"]; ok {

		m.LocalLinkConnection.SwitchID = common.InterfaceToString(value)

	}

	if value, ok := values["port_id"]; ok {

		m.LocalLinkConnection.PortID = common.InterfaceToString(value)

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

	if value, ok := values["created_at"]; ok {

		m.CreatedAt = common.InterfaceToString(value)

	}

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListBaremetalPort lists BaremetalPort with list spec.
func (db *DB) listBaremetalPort(ctx context.Context, request *models.ListBaremetalPortRequest) (response *models.ListBaremetalPortResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["baremetal_port"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.BaremetalPort{}

	if spec.ParentFQName != nil {
		parentMetaData, err := db.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = models.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}
	query, columns, values := qb.ListQuery(auth, spec)
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
		m, err := scanBaremetalPort(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBaremetalPortResponse{
		BaremetalPorts: result,
	}
	return response, nil
}

// UpdateBaremetalPort updates a resource
func (db *DB) updateBaremetalPort(
	ctx context.Context,
	request *models.UpdateBaremetalPortRequest,
) error {
	//TODO
	return nil
}

// DeleteBaremetalPort deletes a resource
func (db *DB) deleteBaremetalPort(
	ctx context.Context,
	request *models.DeleteBaremetalPortRequest) error {
	qb := db.queryBuilders["baremetal_port"]

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

//CreateBaremetalPort handle a Create API
// nolint
func (db *DB) CreateBaremetalPort(
	ctx context.Context,
	request *models.CreateBaremetalPortRequest) (*models.CreateBaremetalPortResponse, error) {
	model := request.BaremetalPort
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createBaremetalPort(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateBaremetalPortResponse{
		BaremetalPort: request.BaremetalPort,
	}, nil
}

//UpdateBaremetalPort handles a Update request.
func (db *DB) UpdateBaremetalPort(
	ctx context.Context,
	request *models.UpdateBaremetalPortRequest) (*models.UpdateBaremetalPortResponse, error) {
	model := request.BaremetalPort
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateBaremetalPort(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "baremetal_port",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateBaremetalPortResponse{
		BaremetalPort: model,
	}, nil
}

//DeleteBaremetalPort delete a resource.
func (db *DB) DeleteBaremetalPort(ctx context.Context, request *models.DeleteBaremetalPortRequest) (*models.DeleteBaremetalPortResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteBaremetalPort(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteBaremetalPortResponse{
		ID: request.ID,
	}, nil
}

//GetBaremetalPort a Get request.
func (db *DB) GetBaremetalPort(ctx context.Context, request *models.GetBaremetalPortRequest) (response *models.GetBaremetalPortResponse, err error) {
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
	listRequest := &models.ListBaremetalPortRequest{
		Spec: spec,
	}
	var result *models.ListBaremetalPortResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listBaremetalPort(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.BaremetalPorts) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetBaremetalPortResponse{
		BaremetalPort: result.BaremetalPorts[0],
	}
	return response, nil
}

//ListBaremetalPort handles a List service Request.
// nolint
func (db *DB) ListBaremetalPort(
	ctx context.Context,
	request *models.ListBaremetalPortRequest) (response *models.ListBaremetalPortResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listBaremetalPort(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
