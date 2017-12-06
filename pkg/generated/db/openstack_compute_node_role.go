package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/utils"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertOpenstackComputeNodeRoleQuery = "insert into `openstack_compute_node_role` (`fq_name`,`uuid`,`key_value_pair`,`provisioning_start_time`,`provisioning_state`,`default_gateway`,`vrouter_bond_interface_members`,`vrouter_type`,`owner`,`owner_access`,`global_access`,`share`,`provisioning_progress`,`provisioning_progress_stage`,`vrouter_bond_interface`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`display_name`,`provisioning_log`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateOpenstackComputeNodeRoleQuery = "update `openstack_compute_node_role` set `fq_name` = ?,`uuid` = ?,`key_value_pair` = ?,`provisioning_start_time` = ?,`provisioning_state` = ?,`default_gateway` = ?,`vrouter_bond_interface_members` = ?,`vrouter_type` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`provisioning_progress` = ?,`provisioning_progress_stage` = ?,`vrouter_bond_interface` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`display_name` = ?,`provisioning_log` = ?;"
const deleteOpenstackComputeNodeRoleQuery = "delete from `openstack_compute_node_role` where uuid = ?"

// OpenstackComputeNodeRoleFields is db columns for OpenstackComputeNodeRole
var OpenstackComputeNodeRoleFields = []string{
	"fq_name",
	"uuid",
	"key_value_pair",
	"provisioning_start_time",
	"provisioning_state",
	"default_gateway",
	"vrouter_bond_interface_members",
	"vrouter_type",
	"owner",
	"owner_access",
	"global_access",
	"share",
	"provisioning_progress",
	"provisioning_progress_stage",
	"vrouter_bond_interface",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
	"group_access",
	"display_name",
	"provisioning_log",
}

// OpenstackComputeNodeRoleRefFields is db reference fields for OpenstackComputeNodeRole
var OpenstackComputeNodeRoleRefFields = map[string][]string{}

// CreateOpenstackComputeNodeRole inserts OpenstackComputeNodeRole to DB
func CreateOpenstackComputeNodeRole(tx *sql.Tx, model *models.OpenstackComputeNodeRole) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertOpenstackComputeNodeRoleQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertOpenstackComputeNodeRoleQuery,
	}).Debug("create query")
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
		string(model.UUID),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.ProvisioningStartTime),
		string(model.ProvisioningState),
		string(model.DefaultGateway),
		string(model.VrouterBondInterfaceMembers),
		string(model.VrouterType),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		int(model.ProvisioningProgress),
		string(model.ProvisioningProgressStage),
		string(model.VrouterBondInterface),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.DisplayName),
		string(model.ProvisioningLog))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanOpenstackComputeNodeRole(values map[string]interface{}) (*models.OpenstackComputeNodeRole, error) {
	m := models.MakeOpenstackComputeNodeRole()

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["provisioning_start_time"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningStartTime = castedValue

	}

	if value, ok := values["provisioning_state"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningState = castedValue

	}

	if value, ok := values["default_gateway"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DefaultGateway = castedValue

	}

	if value, ok := values["vrouter_bond_interface_members"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VrouterBondInterfaceMembers = castedValue

	}

	if value, ok := values["vrouter_type"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VrouterType = castedValue

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["provisioning_progress"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ProvisioningProgress = castedValue

	}

	if value, ok := values["provisioning_progress_stage"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningProgressStage = castedValue

	}

	if value, ok := values["vrouter_bond_interface"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.VrouterBondInterface = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["provisioning_log"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ProvisioningLog = castedValue

	}

	return m, nil
}

// ListOpenstackComputeNodeRole lists OpenstackComputeNodeRole with list spec.
func ListOpenstackComputeNodeRole(tx *sql.Tx, spec *db.ListSpec) ([]*models.OpenstackComputeNodeRole, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "openstack_compute_node_role"
	spec.Fields = OpenstackComputeNodeRoleFields
	spec.RefFields = OpenstackComputeNodeRoleRefFields
	result := models.MakeOpenstackComputeNodeRoleSlice()
	query, columns, values := db.BuildListQuery(spec)
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
		m, err := scanOpenstackComputeNodeRole(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowOpenstackComputeNodeRole shows OpenstackComputeNodeRole resource
func ShowOpenstackComputeNodeRole(tx *sql.Tx, uuid string) (*models.OpenstackComputeNodeRole, error) {
	list, err := ListOpenstackComputeNodeRole(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateOpenstackComputeNodeRole updates a resource
func UpdateOpenstackComputeNodeRole(tx *sql.Tx, uuid string, model *models.OpenstackComputeNodeRole) error {
	//TODO(nati) support update
	return nil
}

// DeleteOpenstackComputeNodeRole deletes a resource
func DeleteOpenstackComputeNodeRole(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteOpenstackComputeNodeRoleQuery)
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
