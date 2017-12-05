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

const insertCustomerAttachmentQuery = "insert into `customer_attachment` (`display_name`,`key_value_pair`,`owner_access`,`global_access`,`share`,`owner`,`uuid`,`fq_name`,`last_modified`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`,`group_access`,`enable`,`description`,`created`,`creator`,`user_visible`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateCustomerAttachmentQuery = "update `customer_attachment` set `display_name` = ?,`key_value_pair` = ?,`owner_access` = ?,`global_access` = ?,`share` = ?,`owner` = ?,`uuid` = ?,`fq_name` = ?,`last_modified` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?;"
const deleteCustomerAttachmentQuery = "delete from `customer_attachment` where uuid = ?"
const listCustomerAttachmentQuery = "select `customer_attachment`.`display_name`,`customer_attachment`.`key_value_pair`,`customer_attachment`.`owner_access`,`customer_attachment`.`global_access`,`customer_attachment`.`share`,`customer_attachment`.`owner`,`customer_attachment`.`uuid`,`customer_attachment`.`fq_name`,`customer_attachment`.`last_modified`,`customer_attachment`.`permissions_owner`,`customer_attachment`.`permissions_owner_access`,`customer_attachment`.`other_access`,`customer_attachment`.`group`,`customer_attachment`.`group_access`,`customer_attachment`.`enable`,`customer_attachment`.`description`,`customer_attachment`.`created`,`customer_attachment`.`creator`,`customer_attachment`.`user_visible` from `customer_attachment`"
const showCustomerAttachmentQuery = "select `customer_attachment`.`display_name`,`customer_attachment`.`key_value_pair`,`customer_attachment`.`owner_access`,`customer_attachment`.`global_access`,`customer_attachment`.`share`,`customer_attachment`.`owner`,`customer_attachment`.`uuid`,`customer_attachment`.`fq_name`,`customer_attachment`.`last_modified`,`customer_attachment`.`permissions_owner`,`customer_attachment`.`permissions_owner_access`,`customer_attachment`.`other_access`,`customer_attachment`.`group`,`customer_attachment`.`group_access`,`customer_attachment`.`enable`,`customer_attachment`.`description`,`customer_attachment`.`created`,`customer_attachment`.`creator`,`customer_attachment`.`user_visible` from `customer_attachment` where uuid = ?"

const insertCustomerAttachmentVirtualMachineInterfaceQuery = "insert into `ref_customer_attachment_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

const insertCustomerAttachmentFloatingIPQuery = "insert into `ref_customer_attachment_floating_ip` (`from`, `to` ) values (?, ?);"

func CreateCustomerAttachment(tx *sql.Tx, model *models.CustomerAttachment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertCustomerAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		string(model.UUID),
		utils.MustJSON(model.FQName),
		string(model.IDPerms.LastModified),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group),
		int(model.IDPerms.Permissions.GroupAccess),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible))

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertCustomerAttachmentVirtualMachineInterfaceQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {
		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
	}

	stmtFloatingIPRef, err := tx.Prepare(insertCustomerAttachmentFloatingIPQuery)
	if err != nil {
		return err
	}
	defer stmtFloatingIPRef.Close()
	for _, ref := range model.FloatingIPRefs {
		_, err = stmtFloatingIPRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanCustomerAttachment(rows *sql.Rows) (*models.CustomerAttachment, error) {
	m := models.MakeCustomerAttachment()

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	var jsonFQName string

	if err := rows.Scan(&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.UUID,
		&jsonFQName,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	return m, nil
}

func buildCustomerAttachmentWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["uuid"]; ok {
		results = append(results, "uuid = ?")
		values = append(values, value)
	}

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["permissions_owner"]; ok {
		results = append(results, "permissions_owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListCustomerAttachment(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.CustomerAttachment, error) {
	result := models.MakeCustomerAttachmentSlice()
	whereQuery, values := buildCustomerAttachmentWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listCustomerAttachmentQuery)
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
		m, _ := scanCustomerAttachment(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowCustomerAttachment(tx *sql.Tx, uuid string) (*models.CustomerAttachment, error) {
	rows, err := tx.Query(showCustomerAttachmentQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanCustomerAttachment(rows)
	}
	return nil, nil
}

func UpdateCustomerAttachment(tx *sql.Tx, uuid string, model *models.CustomerAttachment) error {
	return nil
}

func DeleteCustomerAttachment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteCustomerAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
