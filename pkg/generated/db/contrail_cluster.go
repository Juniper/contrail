package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertContrailClusterQuery = "insert into `contrail_cluster` (`fq_name`,`flow_ttl`,`global_access`,`share`,`owner`,`owner_access`,`default_gateway`,`statistics_ttl`,`config_audit_ttl`,`data_ttl`,`default_vrouter_bond_interface_members`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`permissions_owner`,`display_name`,`key_value_pair`,`uuid`,`contrail_webui`,`default_vrouter_bond_interface`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateContrailClusterQuery = "update `contrail_cluster` set `fq_name` = ?,`flow_ttl` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`default_gateway` = ?,`statistics_ttl` = ?,`config_audit_ttl` = ?,`data_ttl` = ?,`default_vrouter_bond_interface_members` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`permissions_owner` = ?,`display_name` = ?,`key_value_pair` = ?,`uuid` = ?,`contrail_webui` = ?,`default_vrouter_bond_interface` = ?;"
const deleteContrailClusterQuery = "delete from `contrail_cluster` where uuid = ?"

// ContrailClusterFields is db columns for ContrailCluster
var ContrailClusterFields = []string{
	"fq_name",
	"flow_ttl",
	"global_access",
	"share",
	"owner",
	"owner_access",
	"default_gateway",
	"statistics_ttl",
	"config_audit_ttl",
	"data_ttl",
	"default_vrouter_bond_interface_members",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"permissions_owner",
	"display_name",
	"key_value_pair",
	"uuid",
	"contrail_webui",
	"default_vrouter_bond_interface",
}

// ContrailClusterRefFields is db reference fields for ContrailCluster
var ContrailClusterRefFields = map[string][]string{}

// CreateContrailCluster inserts ContrailCluster to DB
func CreateContrailCluster(tx *sql.Tx, model *models.ContrailCluster) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertContrailClusterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertContrailClusterQuery,
	}).Debug("create query")
	_, err = stmt.Exec(common.MustJSON(model.FQName),
		string(model.FlowTTL),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		string(model.DefaultGateway),
		string(model.StatisticsTTL),
		string(model.ConfigAuditTTL),
		string(model.DataTTL),
		string(model.DefaultVrouterBondInterfaceMembers),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.UUID),
		string(model.ContrailWebui),
		string(model.DefaultVrouterBondInterface))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanContrailCluster(values map[string]interface{}) (*models.ContrailCluster, error) {
	m := models.MakeContrailCluster()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["flow_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FlowTTL = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["default_gateway"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultGateway = castedValue

	}

	if value, ok := values["statistics_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.StatisticsTTL = castedValue

	}

	if value, ok := values["config_audit_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ConfigAuditTTL = castedValue

	}

	if value, ok := values["data_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DataTTL = castedValue

	}

	if value, ok := values["default_vrouter_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultVrouterBondInterfaceMembers = castedValue

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

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["contrail_webui"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ContrailWebui = castedValue

	}

	if value, ok := values["default_vrouter_bond_interface"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultVrouterBondInterface = castedValue

	}

	return m, nil
}

// ListContrailCluster lists ContrailCluster with list spec.
func ListContrailCluster(tx *sql.Tx, spec *common.ListSpec) ([]*models.ContrailCluster, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "contrail_cluster"
	spec.Fields = ContrailClusterFields
	spec.RefFields = ContrailClusterRefFields
	result := models.MakeContrailClusterSlice()
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
		m, err := scanContrailCluster(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowContrailCluster shows ContrailCluster resource
func ShowContrailCluster(tx *sql.Tx, uuid string) (*models.ContrailCluster, error) {
	list, err := ListContrailCluster(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateContrailCluster updates a resource
func UpdateContrailCluster(tx *sql.Tx, uuid string, model *models.ContrailCluster) error {
	//TODO(nati) support update
	return nil
}

// DeleteContrailCluster deletes a resource
func DeleteContrailCluster(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteContrailClusterQuery)
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
