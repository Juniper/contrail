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

// NetworkIpamFields is db columns for NetworkIpam
var NetworkIpamFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"ipam_method",
	"virtual_dns_server_name",
	"ip_address",
	"ipam_dns_method",
	"route",
	"dhcp_option",
	"ip_prefix_len",
	"ip_prefix",
	"subnets",
	"ipam_subnet_method",
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
}

// NetworkIpamRefFields is db reference fields for NetworkIpam
var NetworkIpamRefFields = map[string][]string{

	"virtual_DNS": []string{
	// <schema.Schema Value>

	},
}

// NetworkIpamBackRefFields is db back reference fields for NetworkIpam
var NetworkIpamBackRefFields = map[string][]string{}

// NetworkIpamParentTypes is possible parents for NetworkIpam
var NetworkIpamParents = []string{

	"project",
}

// CreateNetworkIpam inserts NetworkIpam to DB
// nolint
func (db *DB) createNetworkIpam(
	ctx context.Context,
	request *models.CreateNetworkIpamRequest) error {
	qb := db.queryBuilders["network_ipam"]
	tx := GetTransaction(ctx)
	model := request.NetworkIpam
	_, err := tx.ExecContext(ctx, qb.CreateQuery(), string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetNetworkIpamMGMT().GetIpamMethod()),
		string(model.GetNetworkIpamMGMT().GetIpamDNSServer().GetVirtualDNSServerName()),
		string(model.GetNetworkIpamMGMT().GetIpamDNSServer().GetTenantDNSServerAddress().GetIPAddress()),
		string(model.GetNetworkIpamMGMT().GetIpamDNSMethod()),
		common.MustJSON(model.GetNetworkIpamMGMT().GetHostRoutes().GetRoute()),
		common.MustJSON(model.GetNetworkIpamMGMT().GetDHCPOptionList().GetDHCPOption()),
		int(model.GetNetworkIpamMGMT().GetCidrBlock().GetIPPrefixLen()),
		string(model.GetNetworkIpamMGMT().GetCidrBlock().GetIPPrefix()),
		common.MustJSON(model.GetIpamSubnets().GetSubnets()),
		string(model.GetIpamSubnetMethod()),
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
		int(model.GetConfigurationVersion()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	for _, ref := range model.VirtualDNSRefs {

		_, err = tx.ExecContext(ctx, qb.CreateRefQuery("virtual_DNS"), model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualDNSRefs create failed")
		}
	}

	metaData := &MetaData{
		UUID:   model.UUID,
		Type:   "network_ipam",
		FQName: model.FQName,
	}
	err = db.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = db.CreateSharing(tx, "network_ipam", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanNetworkIpam(values map[string]interface{}) (*models.NetworkIpam, error) {
	m := models.MakeNetworkIpam()

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

	if value, ok := values["ipam_method"]; ok {

		m.NetworkIpamMGMT.IpamMethod = common.InterfaceToString(value)

	}

	if value, ok := values["virtual_dns_server_name"]; ok {

		m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName = common.InterfaceToString(value)

	}

	if value, ok := values["ip_address"]; ok {

		m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress = common.InterfaceToString(value)

	}

	if value, ok := values["ipam_dns_method"]; ok {

		m.NetworkIpamMGMT.IpamDNSMethod = common.InterfaceToString(value)

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.NetworkIpamMGMT.HostRoutes.Route)

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		m.NetworkIpamMGMT.CidrBlock.IPPrefixLen = common.InterfaceToInt64(value)

	}

	if value, ok := values["ip_prefix"]; ok {

		m.NetworkIpamMGMT.CidrBlock.IPPrefix = common.InterfaceToString(value)

	}

	if value, ok := values["subnets"]; ok {

		json.Unmarshal(value.([]byte), &m.IpamSubnets.Subnets)

	}

	if value, ok := values["ipam_subnet_method"]; ok {

		m.IpamSubnetMethod = common.InterfaceToString(value)

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

	if value, ok := values["configuration_version"]; ok {

		m.ConfigurationVersion = common.InterfaceToInt64(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_virtual_DNS"]; ok {
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
			referenceModel := &models.NetworkIpamVirtualDNSRef{}
			referenceModel.UUID = uuid
			m.VirtualDNSRefs = append(m.VirtualDNSRefs, referenceModel)

		}
	}

	return m, nil
}

// ListNetworkIpam lists NetworkIpam with list spec.
func (db *DB) listNetworkIpam(ctx context.Context, request *models.ListNetworkIpamRequest) (response *models.ListNetworkIpamResponse, err error) {
	var rows *sql.Rows
	tx := GetTransaction(ctx)

	qb := db.queryBuilders["network_ipam"]

	auth := common.GetAuthCTX(ctx)
	spec := request.Spec
	result := []*models.NetworkIpam{}

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
		m, err := scanNetworkIpam(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListNetworkIpamResponse{
		NetworkIpams: result,
	}
	return response, nil
}

// UpdateNetworkIpam updates a resource
func (db *DB) updateNetworkIpam(
	ctx context.Context,
	request *models.UpdateNetworkIpamRequest,
) error {
	//TODO
	return nil
}

// DeleteNetworkIpam deletes a resource
func (db *DB) deleteNetworkIpam(
	ctx context.Context,
	request *models.DeleteNetworkIpamRequest) error {
	qb := db.queryBuilders["network_ipam"]

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

//CreateNetworkIpam handle a Create API
// nolint
func (db *DB) CreateNetworkIpam(
	ctx context.Context,
	request *models.CreateNetworkIpamRequest) (*models.CreateNetworkIpamResponse, error) {
	model := request.NetworkIpam
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createNetworkIpam(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateNetworkIpamResponse{
		NetworkIpam: request.NetworkIpam,
	}, nil
}

//UpdateNetworkIpam handles a Update request.
func (db *DB) UpdateNetworkIpam(
	ctx context.Context,
	request *models.UpdateNetworkIpamRequest) (*models.UpdateNetworkIpamResponse, error) {
	model := request.NetworkIpam
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateNetworkIpam(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "network_ipam",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateNetworkIpamResponse{
		NetworkIpam: model,
	}, nil
}

//DeleteNetworkIpam delete a resource.
func (db *DB) DeleteNetworkIpam(ctx context.Context, request *models.DeleteNetworkIpamRequest) (*models.DeleteNetworkIpamResponse, error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteNetworkIpam(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteNetworkIpamResponse{
		ID: request.ID,
	}, nil
}

//GetNetworkIpam a Get request.
func (db *DB) GetNetworkIpam(ctx context.Context, request *models.GetNetworkIpamRequest) (response *models.GetNetworkIpamResponse, err error) {
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
	listRequest := &models.ListNetworkIpamRequest{
		Spec: spec,
	}
	var result *models.ListNetworkIpamResponse
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listNetworkIpam(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.NetworkIpams) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetNetworkIpamResponse{
		NetworkIpam: result.NetworkIpams[0],
	}
	return response, nil
}

//ListNetworkIpam handles a List service Request.
// nolint
func (db *DB) ListNetworkIpam(
	ctx context.Context,
	request *models.ListNetworkIpamRequest) (response *models.ListNetworkIpamResponse, err error) {
	if err := DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listNetworkIpam(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
