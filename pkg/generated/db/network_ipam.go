package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"strings"
)

const insertNetworkIpamQuery = "insert into `network_ipam` (`ipam_dns_method`,`ip_address`,`virtual_dns_server_name`,`dhcp_option`,`route`,`ip_prefix`,`ip_prefix_len`,`ipam_method`,`ipam_subnet_method`,`fq_name`,`user_visible`,`last_modified`,`other_access`,`group`,`group_access`,`owner`,`owner_access`,`enable`,`description`,`created`,`creator`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`ipam_subnets`,`uuid`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateNetworkIpamQuery = "update `network_ipam` set `ipam_dns_method` = ?,`ip_address` = ?,`virtual_dns_server_name` = ?,`dhcp_option` = ?,`route` = ?,`ip_prefix` = ?,`ip_prefix_len` = ?,`ipam_method` = ?,`ipam_subnet_method` = ?,`fq_name` = ?,`user_visible` = ?,`last_modified` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`ipam_subnets` = ?,`uuid` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteNetworkIpamQuery = "delete from `network_ipam` where uuid = ?"
const listNetworkIpamQuery = "select `network_ipam`.`ipam_dns_method`,`network_ipam`.`ip_address`,`network_ipam`.`virtual_dns_server_name`,`network_ipam`.`dhcp_option`,`network_ipam`.`route`,`network_ipam`.`ip_prefix`,`network_ipam`.`ip_prefix_len`,`network_ipam`.`ipam_method`,`network_ipam`.`ipam_subnet_method`,`network_ipam`.`fq_name`,`network_ipam`.`user_visible`,`network_ipam`.`last_modified`,`network_ipam`.`other_access`,`network_ipam`.`group`,`network_ipam`.`group_access`,`network_ipam`.`owner`,`network_ipam`.`owner_access`,`network_ipam`.`enable`,`network_ipam`.`description`,`network_ipam`.`created`,`network_ipam`.`creator`,`network_ipam`.`perms2_owner`,`network_ipam`.`perms2_owner_access`,`network_ipam`.`global_access`,`network_ipam`.`share`,`network_ipam`.`ipam_subnets`,`network_ipam`.`uuid`,`network_ipam`.`display_name`,`network_ipam`.`key_value_pair` from `network_ipam`"
const showNetworkIpamQuery = "select `network_ipam`.`ipam_dns_method`,`network_ipam`.`ip_address`,`network_ipam`.`virtual_dns_server_name`,`network_ipam`.`dhcp_option`,`network_ipam`.`route`,`network_ipam`.`ip_prefix`,`network_ipam`.`ip_prefix_len`,`network_ipam`.`ipam_method`,`network_ipam`.`ipam_subnet_method`,`network_ipam`.`fq_name`,`network_ipam`.`user_visible`,`network_ipam`.`last_modified`,`network_ipam`.`other_access`,`network_ipam`.`group`,`network_ipam`.`group_access`,`network_ipam`.`owner`,`network_ipam`.`owner_access`,`network_ipam`.`enable`,`network_ipam`.`description`,`network_ipam`.`created`,`network_ipam`.`creator`,`network_ipam`.`perms2_owner`,`network_ipam`.`perms2_owner_access`,`network_ipam`.`global_access`,`network_ipam`.`share`,`network_ipam`.`ipam_subnets`,`network_ipam`.`uuid`,`network_ipam`.`display_name`,`network_ipam`.`key_value_pair` from `network_ipam` where uuid = ?"

const insertNetworkIpamVirtualDNSQuery = "insert into `ref_network_ipam_virtual_DNS` (`from`, `to` ) values (?, ?);"

func CreateNetworkIpam(tx *sql.Tx, model *models.NetworkIpam) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.NetworkIpamMGMT.IpamDNSMethod),
		string(model.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress),
		string(model.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName),
		utils.MustJSON(model.NetworkIpamMGMT.DHCPOptionList.DHCPOption),
		utils.MustJSON(model.NetworkIpamMGMT.HostRoutes.Route),
		string(model.NetworkIpamMGMT.CidrBlock.IPPrefix),
		int(model.NetworkIpamMGMT.CidrBlock.IPPrefixLen),
		string(model.NetworkIpamMGMT.IpamMethod),
		string(model.IpamSubnetMethod),
		utils.MustJSON(model.FQName),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.IpamSubnets),
		string(model.UUID),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair))

	stmtVirtualDNSRef, err := tx.Prepare(insertNetworkIpamVirtualDNSQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualDNSRef.Close()
	for _, ref := range model.VirtualDNSRefs {
		_, err = stmtVirtualDNSRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanNetworkIpam(rows *sql.Rows) (*models.NetworkIpam, error) {
	m := models.MakeNetworkIpam()

	var jsonNetworkIpamMGMTDHCPOptionListDHCPOption string

	var jsonNetworkIpamMGMTHostRoutesRoute string

	var jsonFQName string

	var jsonPerms2Share string

	var jsonIpamSubnets string

	var jsonAnnotationsKeyValuePair string

	if err := rows.Scan(&m.NetworkIpamMGMT.IpamDNSMethod,
		&m.NetworkIpamMGMT.IpamDNSServer.TenantDNSServerAddress.IPAddress,
		&m.NetworkIpamMGMT.IpamDNSServer.VirtualDNSServerName,
		&jsonNetworkIpamMGMTDHCPOptionListDHCPOption,
		&jsonNetworkIpamMGMTHostRoutesRoute,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefix,
		&m.NetworkIpamMGMT.CidrBlock.IPPrefixLen,
		&m.NetworkIpamMGMT.IpamMethod,
		&m.IpamSubnetMethod,
		&jsonFQName,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&jsonIpamSubnets,
		&m.UUID,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonNetworkIpamMGMTDHCPOptionListDHCPOption), &m.NetworkIpamMGMT.DHCPOptionList.DHCPOption)

	json.Unmarshal([]byte(jsonNetworkIpamMGMTHostRoutesRoute), &m.NetworkIpamMGMT.HostRoutes.Route)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonIpamSubnets), &m.IpamSubnets)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	return m, nil
}

func buildNetworkIpamWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["ipam_dns_method"]; ok {
		results = append(results, "ipam_dns_method = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_address"]; ok {
		results = append(results, "ip_address = ?")
		values = append(values, value)
	}

	if value, ok := where["virtual_dns_server_name"]; ok {
		results = append(results, "virtual_dns_server_name = ?")
		values = append(values, value)
	}

	if value, ok := where["ip_prefix"]; ok {
		results = append(results, "ip_prefix = ?")
		values = append(values, value)
	}

	if value, ok := where["ipam_method"]; ok {
		results = append(results, "ipam_method = ?")
		values = append(values, value)
	}

	if value, ok := where["ipam_subnet_method"]; ok {
		results = append(results, "ipam_subnet_method = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["description"]; ok {
		results = append(results, "description = ?")
		values = append(values, value)
	}

	if value, ok := where["created"]; ok {
		results = append(results, "created = ?")
		values = append(values, value)
	}

	if value, ok := where["creator"]; ok {
		results = append(results, "creator = ?")
		values = append(values, value)
	}

	if value, ok := where["perms2_owner"]; ok {
		results = append(results, "perms2_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	return "where " + strings.Join(results, " and "), values
}

func ListNetworkIpam(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.NetworkIpam, error) {
	result := models.MakeNetworkIpamSlice()
	whereQuery, values := buildNetworkIpamWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listNetworkIpamQuery)
	query.WriteRune(' ')
	query.WriteString(whereQuery)
	query.WriteRune(' ')
	query.WriteString(pagenationQuery)
	rows, err = tx.Query(query.String(), values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		m, _ := scanNetworkIpam(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowNetworkIpam(tx *sql.Tx, uuid string) (*models.NetworkIpam, error) {
	rows, err := tx.Query(showNetworkIpamQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanNetworkIpam(rows)
	}
	return nil, nil
}

func UpdateNetworkIpam(tx *sql.Tx, uuid string, model *models.NetworkIpam) error {
	return nil
}

func DeleteNetworkIpam(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteNetworkIpamQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
