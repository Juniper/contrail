package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualRouterQuery = "insert into `virtual_router` (`fq_name`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`description`,`created`,`display_name`,`perms2_owner_access`,`global_access`,`share`,`perms2_owner`,`virtual_router_dpdk_enabled`,`virtual_router_ip_address`,`uuid`,`key_value_pair`,`virtual_router_type`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualRouterQuery = "update `virtual_router` set `fq_name` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`display_name` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`perms2_owner` = ?,`virtual_router_dpdk_enabled` = ?,`virtual_router_ip_address` = ?,`uuid` = ?,`key_value_pair` = ?,`virtual_router_type` = ?;"
const deleteVirtualRouterQuery = "delete from `virtual_router` where uuid = ?"

// VirtualRouterFields is db columns for VirtualRouter
var VirtualRouterFields = []string{
	"fq_name",
	"creator",
	"user_visible",
	"last_modified",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"enable",
	"description",
	"created",
	"display_name",
	"perms2_owner_access",
	"global_access",
	"share",
	"perms2_owner",
	"virtual_router_dpdk_enabled",
	"virtual_router_ip_address",
	"uuid",
	"key_value_pair",
	"virtual_router_type",
}

// VirtualRouterRefFields is db reference fields for VirtualRouter
var VirtualRouterRefFields = map[string][]string{

	"network_ipam": {
		// <common.Schema Value>
		"subnet",
		"allocation_pools",
	},

	"virtual_machine": {
	// <common.Schema Value>

	},
}

const insertVirtualRouterNetworkIpamQuery = "insert into `ref_virtual_router_network_ipam` (`from`, `to` ,`subnet`,`allocation_pools`) values (?, ?,?,?);"

const insertVirtualRouterVirtualMachineQuery = "insert into `ref_virtual_router_virtual_machine` (`from`, `to` ) values (?, ?);"

// CreateVirtualRouter inserts VirtualRouter to DB
func CreateVirtualRouter(tx *sql.Tx, model *models.VirtualRouter) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualRouterQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.FQName),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.DisplayName),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		bool(model.VirtualRouterDPDKEnabled),
		string(model.VirtualRouterIPAddress),
		string(model.UUID),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.VirtualRouterType))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualMachineRef, err := tx.Prepare(insertVirtualRouterVirtualMachineQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineRefs create statement failed")
	}
	defer stmtVirtualMachineRef.Close()
	for _, ref := range model.VirtualMachineRefs {

		_, err = stmtVirtualMachineRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineRefs create failed")
		}
	}

	stmtNetworkIpamRef, err := tx.Prepare(insertVirtualRouterNetworkIpamQuery)
	if err != nil {
		return errors.Wrap(err, "preparing NetworkIpamRefs create statement failed")
	}
	defer stmtNetworkIpamRef.Close()
	for _, ref := range model.NetworkIpamRefs {

		if ref.Attr == nil {
			ref.Attr = models.MakeVirtualRouterNetworkIpamType()
		}

		_, err = stmtNetworkIpamRef.Exec(model.UUID, ref.UUID, common.MustJSON(ref.Attr.Subnet),
			common.MustJSON(ref.Attr.AllocationPools))
		if err != nil {
			return errors.Wrap(err, "NetworkIpamRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualRouter(values map[string]interface{}) (*models.VirtualRouter, error) {
	m := models.MakeVirtualRouter()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["virtual_router_dpdk_enabled"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualRouterDPDKEnabled = castedValue

	}

	if value, ok := values["virtual_router_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualRouterIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["virtual_router_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualRouterType = models.VirtualRouterType(castedValue)

	}

	if value, ok := values["ref_network_ipam"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualRouterNetworkIpamRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.NetworkIpamRefs = append(m.NetworkIpamRefs, referenceModel)

			attr := models.MakeVirtualRouterNetworkIpamType()
			referenceModel.Attr = attr

		}
	}

	if value, ok := values["ref_virtual_machine"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.VirtualRouterVirtualMachineRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualMachineRefs = append(m.VirtualMachineRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualRouter lists VirtualRouter with list spec.
func ListVirtualRouter(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualRouter, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_router"
	spec.Fields = VirtualRouterFields
	spec.RefFields = VirtualRouterRefFields
	result := models.MakeVirtualRouterSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
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
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanVirtualRouter(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowVirtualRouter shows VirtualRouter resource
func ShowVirtualRouter(tx *sql.Tx, uuid string) (*models.VirtualRouter, error) {
	list, err := ListVirtualRouter(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateVirtualRouter updates a resource
func UpdateVirtualRouter(tx *sql.Tx, uuid string, model *models.VirtualRouter) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualRouter deletes a resource
func DeleteVirtualRouter(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteVirtualRouterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
