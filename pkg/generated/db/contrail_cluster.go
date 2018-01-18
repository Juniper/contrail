package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertContrailClusterQuery = "insert into `contrail_cluster` (`uuid`,`statistics_ttl`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`flow_ttl`,`display_name`,`default_vrouter_bond_interface_members`,`default_vrouter_bond_interface`,`default_gateway`,`data_ttl`,`contrail_webui`,`config_audit_ttl`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteContrailClusterQuery = "delete from `contrail_cluster` where uuid = ?"

// ContrailClusterFields is db columns for ContrailCluster
var ContrailClusterFields = []string{
	"uuid",
	"statistics_ttl",
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
	"flow_ttl",
	"display_name",
	"default_vrouter_bond_interface_members",
	"default_vrouter_bond_interface",
	"default_gateway",
	"data_ttl",
	"contrail_webui",
	"config_audit_ttl",
	"key_value_pair",
}

// ContrailClusterRefFields is db reference fields for ContrailCluster
var ContrailClusterRefFields = map[string][]string{}

// ContrailClusterBackRefFields is db back reference fields for ContrailCluster
var ContrailClusterBackRefFields = map[string][]string{}

// ContrailClusterParentTypes is possible parents for ContrailCluster
var ContrailClusterParents = []string{}

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
	_, err = stmt.Exec(string(model.UUID),
		string(model.StatisticsTTL),
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
		string(model.FlowTTL),
		string(model.DisplayName),
		string(model.DefaultVrouterBondInterfaceMembers),
		string(model.DefaultVrouterBondInterface),
		string(model.DefaultGateway),
		string(model.DataTTL),
		string(model.ContrailWebui),
		string(model.ConfigAuditTTL),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "contrail_cluster",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "contrail_cluster", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanContrailCluster(values map[string]interface{}) (*models.ContrailCluster, error) {
	m := models.MakeContrailCluster()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["statistics_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.StatisticsTTL = castedValue

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

	if value, ok := values["flow_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.FlowTTL = castedValue

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["default_vrouter_bond_interface_members"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultVrouterBondInterfaceMembers = castedValue

	}

	if value, ok := values["default_vrouter_bond_interface"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultVrouterBondInterface = castedValue

	}

	if value, ok := values["default_gateway"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DefaultGateway = castedValue

	}

	if value, ok := values["data_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DataTTL = castedValue

	}

	if value, ok := values["contrail_webui"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ContrailWebui = castedValue

	}

	if value, ok := values["config_audit_ttl"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ConfigAuditTTL = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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
	spec.BackRefFields = ContrailClusterBackRefFields
	result := models.MakeContrailClusterSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanContrailCluster(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateContrailCluster updates a resource
func UpdateContrailCluster(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateContrailClusterQuery = "update `contrail_cluster` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateContrailClusterQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".StatisticsTTL", "."); ok {
		updateContrailClusterQuery += "`statistics_ttl` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateContrailClusterQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateContrailClusterQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateContrailClusterQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateContrailClusterQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateContrailClusterQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateContrailClusterQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateContrailClusterQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateContrailClusterQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateContrailClusterQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateContrailClusterQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateContrailClusterQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateContrailClusterQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateContrailClusterQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateContrailClusterQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateContrailClusterQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateContrailClusterQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateContrailClusterQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateContrailClusterQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FlowTTL", "."); ok {
		updateContrailClusterQuery += "`flow_ttl` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateContrailClusterQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DefaultVrouterBondInterfaceMembers", "."); ok {
		updateContrailClusterQuery += "`default_vrouter_bond_interface_members` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DefaultVrouterBondInterface", "."); ok {
		updateContrailClusterQuery += "`default_vrouter_bond_interface` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DefaultGateway", "."); ok {
		updateContrailClusterQuery += "`default_gateway` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DataTTL", "."); ok {
		updateContrailClusterQuery += "`data_ttl` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ContrailWebui", "."); ok {
		updateContrailClusterQuery += "`contrail_webui` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ConfigAuditTTL", "."); ok {
		updateContrailClusterQuery += "`config_audit_ttl` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateContrailClusterQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateContrailClusterQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateContrailClusterQuery += ","
	}

	updateContrailClusterQuery =
		updateContrailClusterQuery[:len(updateContrailClusterQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateContrailClusterQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateContrailClusterQuery,
	}).Debug("update query")
	_, err = stmt.Exec(updatedValues...)
	if err != nil {
		return errors.Wrap(err, "update failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("updated")
	return err
}

// DeleteContrailCluster deletes a resource
func DeleteContrailCluster(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteContrailClusterQuery
	selectQuery := "select count(uuid) from contrail_cluster where uuid = ?"
	var err error
	var count int

	if auth.IsAdmin() {
		row := tx.QueryRow(selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRow(selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.Exec(deleteQuery, uuid, auth.ProjectID())
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
