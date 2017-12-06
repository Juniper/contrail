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

const insertServiceInstanceQuery = "insert into `service_instance` (`auto_policy`,`left_ip_address`,`right_virtual_network`,`max_instances`,`auto_scale`,`interface_list`,`left_virtual_network`,`availability_zone`,`right_ip_address`,`management_virtual_network`,`ha_mode`,`virtual_router_id`,`uuid`,`fq_name`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`display_name`,`key_value_pair`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`,`service_instance_bindings`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceInstanceQuery = "update `service_instance` set `auto_policy` = ?,`left_ip_address` = ?,`right_virtual_network` = ?,`max_instances` = ?,`auto_scale` = ?,`interface_list` = ?,`left_virtual_network` = ?,`availability_zone` = ?,`right_ip_address` = ?,`management_virtual_network` = ?,`ha_mode` = ?,`virtual_router_id` = ?,`uuid` = ?,`fq_name` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`display_name` = ?,`key_value_pair` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?,`service_instance_bindings` = ?;"
const deleteServiceInstanceQuery = "delete from `service_instance` where uuid = ?"

// ServiceInstanceFields is db columns for ServiceInstance
var ServiceInstanceFields = []string{
	"auto_policy",
	"left_ip_address",
	"right_virtual_network",
	"max_instances",
	"auto_scale",
	"interface_list",
	"left_virtual_network",
	"availability_zone",
	"right_ip_address",
	"management_virtual_network",
	"ha_mode",
	"virtual_router_id",
	"uuid",
	"fq_name",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"group_access",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"display_name",
	"key_value_pair",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
	"service_instance_bindings",
}

// ServiceInstanceRefFields is db reference fields for ServiceInstance
var ServiceInstanceRefFields = map[string][]string{

	"service_template": {
	// <utils.Schema Value>

	},

	"instance_ip": {
	// <utils.Schema Value>

	},
}

const insertServiceInstanceServiceTemplateQuery = "insert into `ref_service_instance_service_template` (`from`, `to` ) values (?, ?);"

const insertServiceInstanceInstanceIPQuery = "insert into `ref_service_instance_instance_ip` (`from`, `to` ) values (?, ?);"

// CreateServiceInstance inserts ServiceInstance to DB
func CreateServiceInstance(tx *sql.Tx, model *models.ServiceInstance) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceInstanceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceInstanceQuery,
	}).Debug("create query")
	_, err = stmt.Exec(bool(model.ServiceInstanceProperties.AutoPolicy),
		string(model.ServiceInstanceProperties.LeftIPAddress),
		string(model.ServiceInstanceProperties.RightVirtualNetwork),
		int(model.ServiceInstanceProperties.ScaleOut.MaxInstances),
		bool(model.ServiceInstanceProperties.ScaleOut.AutoScale),
		utils.MustJSON(model.ServiceInstanceProperties.InterfaceList),
		string(model.ServiceInstanceProperties.LeftVirtualNetwork),
		string(model.ServiceInstanceProperties.AvailabilityZone),
		string(model.ServiceInstanceProperties.RightIPAddress),
		string(model.ServiceInstanceProperties.ManagementVirtualNetwork),
		string(model.ServiceInstanceProperties.HaMode),
		string(model.ServiceInstanceProperties.VirtualRouterID),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		utils.MustJSON(model.ServiceInstanceBindings))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtInstanceIPRef, err := tx.Prepare(insertServiceInstanceInstanceIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing InstanceIPRefs create statement failed")
	}
	defer stmtInstanceIPRef.Close()
	for _, ref := range model.InstanceIPRefs {
		_, err = stmtInstanceIPRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "InstanceIPRefs create failed")
		}
	}

	stmtServiceTemplateRef, err := tx.Prepare(insertServiceInstanceServiceTemplateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceTemplateRefs create statement failed")
	}
	defer stmtServiceTemplateRef.Close()
	for _, ref := range model.ServiceTemplateRefs {
		_, err = stmtServiceTemplateRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceTemplateRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceInstance(values map[string]interface{}) (*models.ServiceInstance, error) {
	m := models.MakeServiceInstance()

	if value, ok := values["auto_policy"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.ServiceInstanceProperties.AutoPolicy = castedValue

	}

	if value, ok := values["left_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.LeftIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["right_virtual_network"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.RightVirtualNetwork = castedValue

	}

	if value, ok := values["max_instances"]; ok {

		castedValue := utils.InterfaceToInt(value)

		m.ServiceInstanceProperties.ScaleOut.MaxInstances = castedValue

	}

	if value, ok := values["auto_scale"]; ok {

		castedValue := utils.InterfaceToBool(value)

		m.ServiceInstanceProperties.ScaleOut.AutoScale = castedValue

	}

	if value, ok := values["interface_list"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceProperties.InterfaceList)

	}

	if value, ok := values["left_virtual_network"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.LeftVirtualNetwork = castedValue

	}

	if value, ok := values["availability_zone"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.AvailabilityZone = castedValue

	}

	if value, ok := values["right_ip_address"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.RightIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["management_virtual_network"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.ManagementVirtualNetwork = castedValue

	}

	if value, ok := values["ha_mode"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.HaMode = models.AddressMode(castedValue)

	}

	if value, ok := values["virtual_router_id"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.ServiceInstanceProperties.VirtualRouterID = castedValue

	}

	if value, ok := values["uuid"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

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

	if value, ok := values["display_name"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["perms2_owner"]; ok {

		castedValue := utils.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["perms2_owner_access"]; ok {

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

	if value, ok := values["service_instance_bindings"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceInstanceBindings)

	}

	if value, ok := values["ref_service_template"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ServiceInstanceServiceTemplateRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.ServiceTemplateRefs = append(m.ServiceTemplateRefs, referenceModel)

		}
	}

	if value, ok := values["ref_instance_ip"]; ok {
		var references []interface{}
		stringValue := utils.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap := reference.(map[string]interface{})
			referenceModel := &models.ServiceInstanceInstanceIPRef{}
			referenceModel.UUID = utils.InterfaceToString(referenceMap["uuid"])
			m.InstanceIPRefs = append(m.InstanceIPRefs, referenceModel)

			attr := models.MakeServiceInterfaceTag()
			referenceModel.Attr = attr

		}
	}

	return m, nil
}

// ListServiceInstance lists ServiceInstance with list spec.
func ListServiceInstance(tx *sql.Tx, spec *db.ListSpec) ([]*models.ServiceInstance, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_instance"
	spec.Fields = ServiceInstanceFields
	spec.RefFields = ServiceInstanceRefFields
	result := models.MakeServiceInstanceSlice()
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
		m, err := scanServiceInstance(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceInstance shows ServiceInstance resource
func ShowServiceInstance(tx *sql.Tx, uuid string) (*models.ServiceInstance, error) {
	list, err := ListServiceInstance(tx, &db.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceInstance updates a resource
func UpdateServiceInstance(tx *sql.Tx, uuid string, model *models.ServiceInstance) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceInstance deletes a resource
func DeleteServiceInstance(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceInstanceQuery)
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
