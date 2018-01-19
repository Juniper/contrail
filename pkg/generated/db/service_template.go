package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertServiceTemplateQuery = "insert into `service_template` (`uuid`,`vrouter_instance_type`,`version`,`service_virtualization_type`,`service_type`,`service_scaling`,`service_mode`,`ordered_interfaces`,`interface_type`,`instance_data`,`image_name`,`flavor`,`availability_zone_enable`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteServiceTemplateQuery = "delete from `service_template` where uuid = ?"

// ServiceTemplateFields is db columns for ServiceTemplate
var ServiceTemplateFields = []string{
	"uuid",
	"vrouter_instance_type",
	"version",
	"service_virtualization_type",
	"service_type",
	"service_scaling",
	"service_mode",
	"ordered_interfaces",
	"interface_type",
	"instance_data",
	"image_name",
	"flavor",
	"availability_zone_enable",
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
	"key_value_pair",
}

// ServiceTemplateRefFields is db reference fields for ServiceTemplate
var ServiceTemplateRefFields = map[string][]string{

	"service_appliance_set": {
	// <common.Schema Value>

	},
}

// ServiceTemplateBackRefFields is db back reference fields for ServiceTemplate
var ServiceTemplateBackRefFields = map[string][]string{}

// ServiceTemplateParentTypes is possible parents for ServiceTemplate
var ServiceTemplateParents = []string{

	"domain",
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
		string(model.ServiceTemplateProperties.VrouterInstanceType),
		int(model.ServiceTemplateProperties.Version),
		string(model.ServiceTemplateProperties.ServiceVirtualizationType),
		string(model.ServiceTemplateProperties.ServiceType),
		bool(model.ServiceTemplateProperties.ServiceScaling),
		string(model.ServiceTemplateProperties.ServiceMode),
		bool(model.ServiceTemplateProperties.OrderedInterfaces),
		common.MustJSON(model.ServiceTemplateProperties.InterfaceType),
		string(model.ServiceTemplateProperties.InstanceData),
		string(model.ServiceTemplateProperties.ImageName),
		string(model.ServiceTemplateProperties.Flavor),
		bool(model.ServiceTemplateProperties.AvailabilityZoneEnable),
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
		common.MustJSON(model.Annotations.KeyValuePair))
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

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "service_template",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "service_template", model.UUID, model.Perms2.Share)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanServiceTemplate(values map[string]interface{}) (*models.ServiceTemplate, error) {
	m := models.MakeServiceTemplate()

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["vrouter_instance_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.VrouterInstanceType = models.VRouterInstanceType(castedValue)

	}

	if value, ok := values["version"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.ServiceTemplateProperties.Version = castedValue

	}

	if value, ok := values["service_virtualization_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceVirtualizationType = models.ServiceVirtualizationType(castedValue)

	}

	if value, ok := values["service_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceType = models.ServiceType(castedValue)

	}

	if value, ok := values["service_scaling"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.ServiceScaling = castedValue

	}

	if value, ok := values["service_mode"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ServiceMode = models.ServiceModeType(castedValue)

	}

	if value, ok := values["ordered_interfaces"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.OrderedInterfaces = castedValue

	}

	if value, ok := values["interface_type"]; ok {

		json.Unmarshal(value.([]byte), &m.ServiceTemplateProperties.InterfaceType)

	}

	if value, ok := values["instance_data"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.InstanceData = castedValue

	}

	if value, ok := values["image_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.ImageName = castedValue

	}

	if value, ok := values["flavor"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ServiceTemplateProperties.Flavor = castedValue

	}

	if value, ok := values["availability_zone_enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.ServiceTemplateProperties.AvailabilityZoneEnable = castedValue

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

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

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
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.ServiceTemplateServiceApplianceSetRef{}
			referenceModel.UUID = uuid
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
	spec.BackRefFields = ServiceTemplateBackRefFields
	result := models.MakeServiceTemplateSlice()

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
		m, err := scanServiceTemplate(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateServiceTemplate updates a resource
func UpdateServiceTemplate(tx *sql.Tx, uuid string, model map[string]interface{}) error {
	//TODO (handle references)
	// Prepare statement for updating data
	var updateServiceTemplateQuery = "update `service_template` set "

	updatedValues := make([]interface{}, 0)

	if value, ok := common.GetValueByPath(model, ".UUID", "."); ok {
		updateServiceTemplateQuery += "`uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.VrouterInstanceType", "."); ok {
		updateServiceTemplateQuery += "`vrouter_instance_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.Version", "."); ok {
		updateServiceTemplateQuery += "`version` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.ServiceVirtualizationType", "."); ok {
		updateServiceTemplateQuery += "`service_virtualization_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.ServiceType", "."); ok {
		updateServiceTemplateQuery += "`service_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.ServiceScaling", "."); ok {
		updateServiceTemplateQuery += "`service_scaling` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.ServiceMode", "."); ok {
		updateServiceTemplateQuery += "`service_mode` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.OrderedInterfaces", "."); ok {
		updateServiceTemplateQuery += "`ordered_interfaces` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.InterfaceType", "."); ok {
		updateServiceTemplateQuery += "`interface_type` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.InstanceData", "."); ok {
		updateServiceTemplateQuery += "`instance_data` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.ImageName", "."); ok {
		updateServiceTemplateQuery += "`image_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.Flavor", "."); ok {
		updateServiceTemplateQuery += "`flavor` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ServiceTemplateProperties.AvailabilityZoneEnable", "."); ok {
		updateServiceTemplateQuery += "`availability_zone_enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Share", "."); ok {
		updateServiceTemplateQuery += "`share` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.OwnerAccess", "."); ok {
		updateServiceTemplateQuery += "`owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.Owner", "."); ok {
		updateServiceTemplateQuery += "`owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Perms2.GlobalAccess", "."); ok {
		updateServiceTemplateQuery += "`global_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentUUID", "."); ok {
		updateServiceTemplateQuery += "`parent_uuid` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".ParentType", "."); ok {
		updateServiceTemplateQuery += "`parent_type` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.UserVisible", "."); ok {
		updateServiceTemplateQuery += "`user_visible` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OwnerAccess", "."); ok {
		updateServiceTemplateQuery += "`permissions_owner_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Owner", "."); ok {
		updateServiceTemplateQuery += "`permissions_owner` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.OtherAccess", "."); ok {
		updateServiceTemplateQuery += "`other_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.GroupAccess", "."); ok {
		updateServiceTemplateQuery += "`group_access` = ?"

		updatedValues = append(updatedValues, common.InterfaceToInt(value.(float64)))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Permissions.Group", "."); ok {
		updateServiceTemplateQuery += "`group` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.LastModified", "."); ok {
		updateServiceTemplateQuery += "`last_modified` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Enable", "."); ok {
		updateServiceTemplateQuery += "`enable` = ?"

		updatedValues = append(updatedValues, common.InterfaceToBool(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Description", "."); ok {
		updateServiceTemplateQuery += "`description` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Creator", "."); ok {
		updateServiceTemplateQuery += "`creator` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".IDPerms.Created", "."); ok {
		updateServiceTemplateQuery += "`created` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".FQName", "."); ok {
		updateServiceTemplateQuery += "`fq_name` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".DisplayName", "."); ok {
		updateServiceTemplateQuery += "`display_name` = ?"

		updatedValues = append(updatedValues, common.InterfaceToString(value))

		updateServiceTemplateQuery += ","
	}

	if value, ok := common.GetValueByPath(model, ".Annotations.KeyValuePair", "."); ok {
		updateServiceTemplateQuery += "`key_value_pair` = ?"

		updatedValues = append(updatedValues, common.MustJSON(value))

		updateServiceTemplateQuery += ","
	}

	updateServiceTemplateQuery =
		updateServiceTemplateQuery[:len(updateServiceTemplateQuery)-1] + " where `uuid` = ? ;"
	updatedValues = append(updatedValues, string(uuid))
	stmt, err := tx.Prepare(updateServiceTemplateQuery)
	if err != nil {
		return errors.Wrap(err, "preparing update statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": updateServiceTemplateQuery,
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

// DeleteServiceTemplate deletes a resource
func DeleteServiceTemplate(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	deleteQuery := deleteServiceTemplateQuery
	selectQuery := "select count(uuid) from service_template where uuid = ?"
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
