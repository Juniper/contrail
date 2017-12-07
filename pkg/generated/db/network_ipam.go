package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertNetworkIpamQuery = "insert into `network_ipam` (`ipam_method`,`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ip_prefix`,`ip_prefix_len`,`ipam_subnets`,`ipam_subnet_method`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group`,`group_access`,`owner`,`owner_access`,`other_access`,`enable`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`uuid`,`fq_name`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkIpamQuery = "update `network_ipam` set `ipam_method` = ?,`ipam_dns_method` = ?,`ip_address` = ?,`virtual_dns_server_name` = ?,`dhcp_option` = ?,`route` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`ipam_subnets` = ?,`ipam_subnet_method` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`uuid` = ?,`fq_name` = ?;"
const deleteNetworkIpamQuery = "delete from `network_ipam` where uuid = ?"

// NetworkIpamFields is db columns for NetworkIpam
var NetworkIpamFields = []string{
	"ipam_method",
	"ipam_dns_method",
	"ip_address",
	"virtual_dns_server_name",
	"dhcp_option",
	"route",
	"ip_prefix",
	"ip_prefix_len",
	"ipam_subnets",
	"ipam_subnet_method",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"enable",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"uuid",
	"fq_name",
}

// NetworkIpamRefFields is db reference fields for NetworkIpam
var NetworkIpamRefFields = map[string][]string{

	"virtual_DNS": {
	// <common.Schema Value>

	},
}

const insertNetworkIpamVirtualDNSQuery = "insert into `ref_network_ipam_virtual_DNS` (`from`, `to` ) values (?, ?);"

// CreateNetworkIpam inserts NetworkIpam to DB
func CreateNetworkIpam(tx *sql.Tx, model *models.NetworkIpam) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkIpamQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertNetworkIpamQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.NetworkIpamMGMT.IpamMethod),
		string(model.NetworkIpamMGMT.IpamDNSMethod),
		string(model.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress),
		string(model.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName),
		common.MustJSON(model.NetworkIpamMGMT.DHCPOptionList.DHCPOption),
		common.MustJSON(model.NetworkIpamMGMT.HostRoutes.Route),
		string(model.NetworkIpamMGMT.CidrBlock.IPPrefix),
		int(model.NetworkIpamMGMT.CidrBlock.IPPrefixLen),
		common.MustJSON(model.IpamSubnets),
		string(model.IpamSubnetMethod),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.UUID),
		common.MustJSON(model.FQName))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtVirtualDNSRef, err := tx.Prepare(insertNetworkIpamVirtualDNSQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualDNSRefs create statement failed")
	}
	defer stmtVirtualDNSRef.Close()
	for _, ref := range model.VirtualDNSRefs {

		_, err = stmtVirtualDNSRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualDNSRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanNetworkIpam(values map[string]interface{}) (*models.NetworkIpam, error) {
	m := models.MakeNetworkIpam()

	if value, ok := values["ipam_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.NetworkIpamMGMT.IpamMethod = models.IpamMethodType(castedValue)

	}

	if value, ok := values["ipam_dns_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.NetworkIpamMGMT.IpamDNSMethod = models.IpamDnsMethodType(castedValue)

	}

	if value, ok := values["ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["virtual_dns_server_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName = castedValue

	}

	if value, ok := values["dhcp_option"]; ok {

		json.Unmarshal(value.([]byte), &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption)

	}

	if value, ok := values["route"]; ok {

		json.Unmarshal(value.([]byte), &m.NetworkIpamMGMT.HostRoutes.Route)

	}

	if value, ok := values["ip_prefix"]; ok {

		castedValue := common.InterfaceToString(value)

		m.NetworkIpamMGMT.CidrBlock.IPPrefix = castedValue

	}

	if value, ok := values["ip_prefix_len"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.NetworkIpamMGMT.CidrBlock.IPPrefixLen = castedValue

	}

	if value, ok := values["ipam_subnets"]; ok {

		json.Unmarshal(value.([]byte), &m.IpamSubnets)

	}

	if value, ok := values["ipam_subnet_method"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IpamSubnetMethod = models.SubnetMethodType(castedValue)

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

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

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

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

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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
			if referenceMap["to"] == "" {
				continue
			}
			referenceModel := &models.NetworkIpamVirtualDNSRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.VirtualDNSRefs = append(m.VirtualDNSRefs, referenceModel)

		}
	}

	return m, nil
}

// ListNetworkIpam lists NetworkIpam with list spec.
func ListNetworkIpam(tx *sql.Tx, spec *common.ListSpec) ([]*models.NetworkIpam, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "network_ipam"
	spec.Fields = NetworkIpamFields
	spec.RefFields = NetworkIpamRefFields
	result := models.MakeNetworkIpamSlice()
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
		m, err := scanNetworkIpam(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowNetworkIpam shows NetworkIpam resource
func ShowNetworkIpam(tx *sql.Tx, uuid string) (*models.NetworkIpam, error) {
	list, err := ListNetworkIpam(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateNetworkIpam updates a resource
func UpdateNetworkIpam(tx *sql.Tx, uuid string, model *models.NetworkIpam) error {
	//TODO(nati) support update
	return nil
}

// DeleteNetworkIpam deletes a resource
func DeleteNetworkIpam(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkIpamQuery)
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
