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

const insertBGPAsAServiceQuery = "insert into `bgp_as_a_service` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`bgpaas_suppress_route_advertisement`,`bgpaas_shared`,`bgpaas_session_attributes`,`bgpaas_ipv4_mapped_ipv6_nexthop`,`bgpaas_ip_address`,`autonomous_system`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteBGPAsAServiceQuery = "delete from `bgp_as_a_service` where uuid = ?"

// BGPAsAServiceFields is db columns for BGPAsAService
var BGPAsAServiceFields = []string{
	"uuid",
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
	"bgpaas_suppress_route_advertisement",
	"bgpaas_shared",
	"bgpaas_session_attributes",
	"bgpaas_ipv4_mapped_ipv6_nexthop",
	"bgpaas_ip_address",
	"autonomous_system",
	"key_value_pair",
}

// BGPAsAServiceRefFields is db reference fields for BGPAsAService
var BGPAsAServiceRefFields = map[string][]string{

	"service_health_check": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},
}

// BGPAsAServiceBackRefFields is db back reference fields for BGPAsAService
var BGPAsAServiceBackRefFields = map[string][]string{}

// BGPAsAServiceParentTypes is possible parents for BGPAsAService
var BGPAsAServiceParents = []string{

	"project",
}

const insertBGPAsAServiceVirtualMachineInterfaceQuery = "insert into `ref_bgp_as_a_service_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertBGPAsAServiceServiceHealthCheckQuery = "insert into `ref_bgp_as_a_service_service_health_check` (`from`, `to` ) values (?, ?);"

// CreateBGPAsAService inserts BGPAsAService to DB
func CreateBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.CreateBGPAsAServiceRequest) error {
	model := request.BGPAsAService
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertBGPAsAServiceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertBGPAsAServiceQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
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
		string(model.DisplayName),
		bool(model.BgpaasSuppressRouteAdvertisement),
		bool(model.BgpaasShared),
		string(model.BgpaasSessionAttributes),
		bool(model.BgpaasIpv4MappedIpv6Nexthop),
		string(model.BgpaasIPAddress),
		int(model.AutonomousSystem),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertBGPAsAServiceVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	stmtServiceHealthCheckRef, err := tx.Prepare(insertBGPAsAServiceServiceHealthCheckQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceHealthCheckRefs create statement failed")
	}
	defer stmtServiceHealthCheckRef.Close()
	for _, ref := range model.ServiceHealthCheckRefs {

		_, err = stmtServiceHealthCheckRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceHealthCheckRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "bgp_as_a_service",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "bgp_as_a_service", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanBGPAsAService(values map[string]interface{}) (*models.BGPAsAService, error) {
	m := models.MakeBGPAsAService()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["bgpaas_suppress_route_advertisement"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasSuppressRouteAdvertisement = castedValue

	}

	if value, ok := values["bgpaas_shared"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasShared = castedValue

	}

	if value, ok := values["bgpaas_session_attributes"]; ok {

		castedValue := common.InterfaceToString(value)

		m.BgpaasSessionAttributes = castedValue

	}

	if value, ok := values["bgpaas_ipv4_mapped_ipv6_nexthop"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.BgpaasIpv4MappedIpv6Nexthop = castedValue

	}

	if value, ok := values["bgpaas_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.BgpaasIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["autonomous_system"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.AutonomousSystem = models.AutonomousSystemType(castedValue)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
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
			referenceModel := &models.BGPAsAServiceVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["ref_service_health_check"]; ok {
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
			referenceModel := &models.BGPAsAServiceServiceHealthCheckRef{}
			referenceModel.UUID = uuid
			m.ServiceHealthCheckRefs = append(m.ServiceHealthCheckRefs, referenceModel)

		}
	}

	return m, nil
}

// ListBGPAsAService lists BGPAsAService with list spec.
func ListBGPAsAService(ctx context.Context, tx *sql.Tx, request *models.ListBGPAsAServiceRequest) (response *models.ListBGPAsAServiceResponse, err error) {
	var rows *sql.Rows
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "bgp_as_a_service"
	qb.Fields = BGPAsAServiceFields
	qb.RefFields = BGPAsAServiceRefFields
	qb.BackRefFields = BGPAsAServiceBackRefFields
	result := models.MakeBGPAsAServiceSlice()

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
		m, err := scanBGPAsAService(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListBGPAsAServiceResponse{
		BGPAsAServices: result,
	}
	return response, nil
}

// UpdateBGPAsAService updates a resource
func UpdateBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.UpdateBGPAsAServiceRequest,
) error {
	//TODO
	return nil
}

// DeleteBGPAsAService deletes a resource
func DeleteBGPAsAService(
	ctx context.Context,
	tx *sql.Tx,
	request *models.DeleteBGPAsAServiceRequest) error {
	deleteQuery := deleteBGPAsAServiceQuery
	selectQuery := "select count(uuid) from bgp_as_a_service where uuid = ?"
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
