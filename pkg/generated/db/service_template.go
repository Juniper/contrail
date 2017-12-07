package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceTemplateQuery = "insert into `service_template` (`uuid`,`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`owner`,`owner_access`,`other_access`,`group`,`enable`,`display_name`,`key_value_pair`,`instance_data`,`vrouter_instance_type`,`interface_type`,`service_mode`,`availability_zone_enable`,`service_type`,`ordered_interfaces`,`service_virtualization_type`,`service_scaling`,`version`,`flavor`,`image_name`,`perms2_owner`,`perms2_owner_access`,`global_access`,`share`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateServiceTemplateQuery = "update `service_template` set `uuid` = ?,`fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`instance_data` = ?,`vrouter_instance_type` = ?,`interface_type` = ?,`service_mode` = ?,`availability_zone_enable` = ?,`service_type` = ?,`ordered_interfaces` = ?,`service_virtualization_type` = ?,`service_scaling` = ?,`version` = ?,`flavor` = ?,`image_name` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`share` = ?;"
const deleteServiceTemplateQuery = "delete from `service_template` where uuid = ?"

// ServiceTemplateFields is db columns for ServiceTemplate
var ServiceTemplateFields = []string{
	"uuid",
	"fq_name",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"owner",
	"owner_access",
	"other_access",
	"group",
	"enable",
	"display_name",
	"key_value_pair",
	"instance_data",
	"vrouter_instance_type",
	"interface_type",
	"service_mode",
	"availability_zone_enable",
	"service_type",
	"ordered_interfaces",
	"service_virtualization_type",
	"service_scaling",
	"version",
	"flavor",
	"image_name",
	"perms2_owner",
	"perms2_owner_access",
	"global_access",
	"share",
}

// ServiceTemplateRefFields is db reference fields for ServiceTemplate
var ServiceTemplateRefFields = map[string][]string{

	"service_appliance_set": {
	// <common.Schema Value>

	},
}

const insertServiceTemplateServiceApplianceSetQuery = "insert into `ref_service_template_service_appliance_set` (`from`, `to` ) values (?, ?);"

// CreateServiceTemplate inserts ServiceTemplate to DB
func CreateServiceTemplate(tx *sql.Tx, model *models.ServiceTemplate) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertServiceTemplateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertServiceTemplateQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.UUID),
		common.MustJSON(model.FQName),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		string(model.ServiceTemplateProperties.InstanceData),
		string(model.ServiceTemplateProperties.VrouterInstanceType),
		common.MustJSON(model.ServiceTemplateProperties.InterfaceType),
		string(model.ServiceTemplateProperties.ServiceMode),
		bool(model.ServiceTemplateProperties.AvailabilityZoneEnable),
		string(model.ServiceTemplateProperties.ServiceType),
		bool(model.ServiceTemplateProperties.OrderedInterfaces),
		string(model.ServiceTemplateProperties.ServiceVirtualizationType),
		bool(model.ServiceTemplateProperties.ServiceScaling),
		int(model.ServiceTemplateProperties.Version),
		string(model.ServiceTemplateProperties.Flavor),
		string(model.ServiceTemplateProperties.ImageName),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		common.MustJSON(model.Perms2.Share))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtServiceApplianceSetRef, err := tx.Prepare(insertServiceTemplateServiceApplianceSetQuery)
	if err != nil {
		return errors.Wrap(err, "preparing ServiceApplianceSetRefs create statement failed")
	}
	defer stmtServiceApplianceSetRef.Close()
	for _, ref := range model.ServiceApplianceSetRefs {

		_, err = stmtServiceApplianceSetRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "ServiceApplianceSetRefs create failed")
		}
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanServiceTemplate(values map[string]interface{}) (*models.ServiceTemplate, error) {
	m := models.MakeServiceTemplate()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

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

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

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

	if value, ok := values["instance_data"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.InstanceData = castedValue

	}

	if value, ok := values["vrouter_instance_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.VrouterInstanceType = models.VRouterInstanceType(castedValue)

	}

	if value, ok := values["interface_type"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceTemplateProperties.InterfaceType)

	}

	if value, ok := values["service_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceMode = models.ServiceModeType(castedValue)

	}

	if value, ok := values["availability_zone_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.AvailabilityZoneEnable = castedValue

	}

	if value, ok := values["service_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceType = models.ServiceType(castedValue)

	}

	if value, ok := values["ordered_interfaces"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.OrderedInterfaces = castedValue

	}

	if value, ok := values["service_virtualization_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceVirtualizationType = models.ServiceVirtualizationType(castedValue)

	}

	if value, ok := values["service_scaling"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.ServiceScaling = castedValue

	}

	if value, ok := values["version"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceTemplateProperties.Version = castedValue

	}

	if value, ok := values["flavor"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.Flavor = castedValue

	}

	if value, ok := values["image_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ImageName = castedValue

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

	if value, ok := values["ref_service_appliance_set"]; ok {
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
			referenceModel := &models.ServiceTemplateServiceApplianceSetRef{}
			referenceModel.UUID = common.InterfaceToString(referenceMap["to"])
			m.ServiceApplianceSetRefs = append(m.ServiceApplianceSetRefs, referenceModel)

		}
	}

	return m, nil
}

// ListServiceTemplate lists ServiceTemplate with list spec.
func ListServiceTemplate(tx *sql.Tx, spec *common.ListSpec) ([]*models.ServiceTemplate, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "service_template"
	spec.Fields = ServiceTemplateFields
	spec.RefFields = ServiceTemplateRefFields
	result := models.MakeServiceTemplateSlice()
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
		m, err := scanServiceTemplate(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowServiceTemplate shows ServiceTemplate resource
func ShowServiceTemplate(tx *sql.Tx, uuid string) (*models.ServiceTemplate, error) {
	list, err := ListServiceTemplate(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateServiceTemplate updates a resource
func UpdateServiceTemplate(tx *sql.Tx, uuid string, model *models.ServiceTemplate) error {
	//TODO(nati) support update
	return nil
}

// DeleteServiceTemplate deletes a resource
func DeleteServiceTemplate(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteServiceTemplateQuery)
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
