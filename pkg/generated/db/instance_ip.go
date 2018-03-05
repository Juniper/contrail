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

const insertInstanceIPQuery = "insert into `instance_ip` (`uuid`,`subnet_uuid`,`service_instance_ip`,`service_health_check_ip`,`ip_prefix_len`,`ip_prefix`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`instance_ip_secondary`,`instance_ip_mode`,`instance_ip_local_ip`,`instance_ip_family`,`instance_ip_address`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteInstanceIPQuery = "delete from `instance_ip` where uuid = ?"

// InstanceIPFields is db columns for InstanceIP
var InstanceIPFields = []string{
	"uuid",
	"subnet_uuid",
	"service_instance_ip",
	"service_health_check_ip",
	"ip_prefix_len",
	"ip_prefix",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"instance_ip_secondary",
	"instance_ip_mode",
	"instance_ip_local_ip",
	"instance_ip_family",
	"instance_ip_address",
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

// InstanceIPRefFields is db reference fields for InstanceIP
var InstanceIPRefFields = map[string][]string{

	"virtual_machine_interface": []string{
	// <schema.Schema Value>

	},

	"physical_router": []string{
	// <schema.Schema Value>

	},

	"virtual_router": []string{
	// <schema.Schema Value>

	},

	"network_ipam": []string{
	// <schema.Schema Value>

	},

	"virtual_network": []string{
	// <schema.Schema Value>

	},
}

// InstanceIPBackRefFields is db back reference fields for InstanceIP
var InstanceIPBackRefFields = map[string][]string{

	"floating_ip": []string{
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
		"floating_ip_traffic_direction",
		"port_mappings",
		"floating_ip_port_mappings_enable",
		"floating_ip_is_virtual_ip",
		"floating_ip_fixed_ip_address",
		"floating_ip_address_family",
		"floating_ip_address",
		"display_name",
		"key_value_pair",
	},
}

// InstanceIPParentTypes is possible parents for InstanceIP
var InstanceIPParents = []string{}

const insertInstanceIPNetworkIpamQuery = "insert into `ref_instance_ip_network_ipam` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualNetworkQuery = "insert into `ref_instance_ip_virtual_network` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualMachineInterfaceQuery = "insert into `ref_instance_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertInstanceIPPhysicalRouterQuery = "insert into `ref_instance_ip_physical_router` (`from`, `to` ) values (?, ?);"

const insertInstanceIPVirtualRouterQuery = "insert into `ref_instance_ip_virtual_router` (`from`, `to` ) values (?, ?);"

// CreateInstanceIP inserts InstanceIP to DB
func (db *DB) createInstanceIP(
	ctx context.Context,
	request *models.CreateInstanceIPRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.InstanceIP
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertInstanceIPQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		string(model.GetSubnetUUID()),
		bool(model.GetServiceInstanceIP()),
		bool(model.GetServiceHealthCheckIP()),
		int(model.GetSecondaryIPTrackingIP().GetIPPrefixLen()),
		string(model.GetSecondaryIPTrackingIP().GetIPPrefix()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		bool(model.GetInstanceIPSecondary()),
		string(model.GetInstanceIPMode()),
		bool(model.GetInstanceIPLocalIP()),
		string(model.GetInstanceIPFamily()),
		string(model.GetInstanceIPAddress()),
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

	stmtVirtualRouterRef, err := tx.Prepare(insertInstanceIPVirtualRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualRouterRefs create statement failed")
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {

		_, err = stmtVirtualRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualRouterRefs create failed")
		}
	}

	stmtNetworkIpamRef, err := tx.Prepare(insertInstanceIPNetworkIpamQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkIpamRefs create statement failed")
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {

		_, err = stmtNetworkIpamRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	stmtVirtualNetworkRef, err := tx.Prepare(insertInstanceIPVirtualNetworkQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualNetworkRefs create statement failed")
	}
	defer stmtVirtualNetworkRef.Close()
	for _, ref := range model.VirtualNetworkRefs {

		_, err = stmtVirtualNetworkRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualNetworkRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertInstanceIPVirtualMachineInterfaceQuery)
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

	stmtPhysicalRouterRef, err := tx.Prepare(insertInstanceIPPhysicalRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing PhysicalRouterRefs create statement failed")
	}
	defer stmtPhysicalRouterRef.Close()
	for _, ref := range model.PhysicalRouterRefs {

		_, err = stmtPhysicalRouterRef.ExecContext(ctx, model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "PhysicalRouterRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "instance_ip",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "instance_ip", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanInstanceIP(values map[string]interface{}) (*models.InstanceIP, error) {
	m := models.MakeInstanceIP()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["subnet_uuid"]; ok {

		m.SubnetUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["service_instance_ip"]; ok {

		m.ServiceInstanceIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["service_health_check_ip"]; ok {

		m.ServiceHealthCheckIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["ip_prefix_len"]; ok {

		m.SecondaryIPTrackingIP.IPPrefixLen = schema.InterfaceToInt64(value)

	}

	if value, ok := values["ip_prefix"]; ok {

		m.SecondaryIPTrackingIP.IPPrefix = schema.InterfaceToString(value)

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

	if value, ok := values["instance_ip_secondary"]; ok {

		m.InstanceIPSecondary = schema.InterfaceToBool(value)

	}

	if value, ok := values["instance_ip_mode"]; ok {

		m.InstanceIPMode = schema.InterfaceToString(value)

	}

	if value, ok := values["instance_ip_local_ip"]; ok {

		m.InstanceIPLocalIP = schema.InterfaceToBool(value)

	}

	if value, ok := values["instance_ip_family"]; ok {

		m.InstanceIPFamily = schema.InterfaceToString(value)

	}

	if value, ok := values["instance_ip_address"]; ok {

		m.InstanceIPAddress = schema.InterfaceToString(value)

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

	if value, ok := values["ref_physical_router"]; ok {
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
			referenceModel := &models.InstanceIPPhysicalRouterRef{}
			referenceModel.UUID = uuid
			m.PhysicalRouterRefs = append(m.PhysicalRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_router"]; ok {
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
			referenceModel := &models.InstanceIPVirtualRouterRef{}
			referenceModel.UUID = uuid
			m.VirtualRouterRefs = append(m.VirtualRouterRefs, referenceModel)

		}
	}

	if value, ok := values["ref_network_ipam"]; ok {
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
			referenceModel := &models.InstanceIPNetworkIpamRef{}
			referenceModel.UUID = uuid
			m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_network"]; ok {
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
			referenceModel := &models.InstanceIPVirtualNetworkRef{}
			referenceModel.UUID = uuid
			m.VirtualNetworkRefs = append(m.VirtualNetworkRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
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
			referenceModel := &models.InstanceIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	if value, ok := values["backref_floating_ip"]; ok {
		var childResources []interface{}
		stringValue := schema.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &childResources)
		for _, childResource := range childResources {
			childResourceMap, ok := childResource.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := schema.InterfaceToString(childResourceMap["uuid"])
			if uuid == "" {
				continue
			}
			childModel := models.MakeFloatingIP()
			m.FloatingIPs = append(m.FloatingIPs, childModel)

			if propertyValue, ok := childResourceMap["uuid"]; ok && propertyValue != nil {

				childModel.UUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["share"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Perms2.Share)

			}

			if propertyValue, ok := childResourceMap["owner_access"]; ok && propertyValue != nil {

				childModel.Perms2.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["owner"]; ok && propertyValue != nil {

				childModel.Perms2.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["global_access"]; ok && propertyValue != nil {

				childModel.Perms2.GlobalAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_uuid"]; ok && propertyValue != nil {

				childModel.ParentUUID = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["parent_type"]; ok && propertyValue != nil {

				childModel.ParentType = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["user_visible"]; ok && propertyValue != nil {

				childModel.IDPerms.UserVisible = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["permissions_owner"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Owner = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["other_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group_access"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(propertyValue)

			}

			if propertyValue, ok := childResourceMap["group"]; ok && propertyValue != nil {

				childModel.IDPerms.Permissions.Group = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["last_modified"]; ok && propertyValue != nil {

				childModel.IDPerms.LastModified = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["enable"]; ok && propertyValue != nil {

				childModel.IDPerms.Enable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["description"]; ok && propertyValue != nil {

				childModel.IDPerms.Description = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["creator"]; ok && propertyValue != nil {

				childModel.IDPerms.Creator = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["created"]; ok && propertyValue != nil {

				childModel.IDPerms.Created = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["fq_name"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FQName)

			}

			if propertyValue, ok := childResourceMap["floating_ip_traffic_direction"]; ok && propertyValue != nil {

				childModel.FloatingIPTrafficDirection = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["port_mappings"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.FloatingIPPortMappings.PortMappings)

			}

			if propertyValue, ok := childResourceMap["floating_ip_port_mappings_enable"]; ok && propertyValue != nil {

				childModel.FloatingIPPortMappingsEnable = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_is_virtual_ip"]; ok && propertyValue != nil {

				childModel.FloatingIPIsVirtualIP = schema.InterfaceToBool(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_fixed_ip_address"]; ok && propertyValue != nil {

				childModel.FloatingIPFixedIPAddress = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_address_family"]; ok && propertyValue != nil {

				childModel.FloatingIPAddressFamily = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["floating_ip_address"]; ok && propertyValue != nil {

				childModel.FloatingIPAddress = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["display_name"]; ok && propertyValue != nil {

				childModel.DisplayName = schema.InterfaceToString(propertyValue)

			}

			if propertyValue, ok := childResourceMap["key_value_pair"]; ok && propertyValue != nil {

				json.Unmarshal(schema.InterfaceToBytes(propertyValue), &childModel.Annotations.KeyValuePair)

			}

		}
	}

	return m, nil
}

// ListInstanceIP lists InstanceIP with list spec.
func (db *DB) listInstanceIP(ctx context.Context, request *models.ListInstanceIPRequest) (response *models.ListInstanceIPResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "instance_ip"
	qb.Fields = InstanceIPFields
	qb.RefFields = InstanceIPRefFields
	qb.BackRefFields = InstanceIPBackRefFields
	result := []*models.InstanceIP{}

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
		m, err := scanInstanceIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListInstanceIPResponse{
		InstanceIPs: result,
	}
	return response, nil
}

// UpdateInstanceIP updates a resource
func (db *DB) updateInstanceIP(
	ctx context.Context,
	request *models.UpdateInstanceIPRequest,
) error {
	//TODO
	return nil
}

// DeleteInstanceIP deletes a resource
func (db *DB) deleteInstanceIP(
	ctx context.Context,
	request *models.DeleteInstanceIPRequest) error {
	deleteQuery := deleteInstanceIPQuery
	selectQuery := "select count(uuid) from instance_ip where uuid = ?"
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

//CreateInstanceIP handle a Create API
func (db *DB) CreateInstanceIP(
	ctx context.Context,
	request *models.CreateInstanceIPRequest) (*models.CreateInstanceIPResponse, error) {
	model := request.InstanceIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createInstanceIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateInstanceIPResponse{
		InstanceIP: request.InstanceIP,
	}, nil
}

//UpdateInstanceIP handles a Update request.
func (db *DB) UpdateInstanceIP(
	ctx context.Context,
	request *models.UpdateInstanceIPRequest) (*models.UpdateInstanceIPResponse, error) {
	model := request.InstanceIP
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateInstanceIP(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "instance_ip",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateInstanceIPResponse{
		InstanceIP: model,
	}, nil
}

//DeleteInstanceIP delete a resource.
func (db *DB) DeleteInstanceIP(ctx context.Context, request *models.DeleteInstanceIPRequest) (*models.DeleteInstanceIPResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteInstanceIP(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteInstanceIPResponse{
		ID: request.ID,
	}, nil
}

//GetInstanceIP a Get request.
func (db *DB) GetInstanceIP(ctx context.Context, request *models.GetInstanceIPRequest) (response *models.GetInstanceIPResponse, err error) {
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
	listRequest := &models.ListInstanceIPRequest{
		Spec: spec,
	}
	var result *models.ListInstanceIPResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listInstanceIP(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.InstanceIPs) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetInstanceIPResponse{
		InstanceIP: result.InstanceIPs[0],
	}
	return response, nil
}

//ListInstanceIP handles a List service Request.
func (db *DB) ListInstanceIP(
	ctx context.Context,
	request *models.ListInstanceIPRequest) (response *models.ListInstanceIPResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listInstanceIP(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
