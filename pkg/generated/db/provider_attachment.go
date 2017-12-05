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

const insertProviderAttachmentQuery = "insert into `provider_attachment` (`fq_name`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`owner`,`owner_access`,`other_access`,`group`,`group_access`,`enable`,`display_name`,`key_value_pair`,`share`,`perms2_owner`,`perms2_owner_access`,`global_access`,`uuid`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateProviderAttachmentQuery = "update `provider_attachment` set `fq_name` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`owner` = ?,`owner_access` = ?,`other_access` = ?,`group` = ?,`group_access` = ?,`enable` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`perms2_owner` = ?,`perms2_owner_access` = ?,`global_access` = ?,`uuid` = ?;"
const deleteProviderAttachmentQuery = "delete from `provider_attachment` where uuid = ?"
const listProviderAttachmentQuery = "select `provider_attachment`.`fq_name`,`provider_attachment`.`description`,`provider_attachment`.`created`,`provider_attachment`.`creator`,`provider_attachment`.`user_visible`,`provider_attachment`.`last_modified`,`provider_attachment`.`owner`,`provider_attachment`.`owner_access`,`provider_attachment`.`other_access`,`provider_attachment`.`group`,`provider_attachment`.`group_access`,`provider_attachment`.`enable`,`provider_attachment`.`display_name`,`provider_attachment`.`key_value_pair`,`provider_attachment`.`share`,`provider_attachment`.`perms2_owner`,`provider_attachment`.`perms2_owner_access`,`provider_attachment`.`global_access`,`provider_attachment`.`uuid` from `provider_attachment`"
const showProviderAttachmentQuery = "select `provider_attachment`.`fq_name`,`provider_attachment`.`description`,`provider_attachment`.`created`,`provider_attachment`.`creator`,`provider_attachment`.`user_visible`,`provider_attachment`.`last_modified`,`provider_attachment`.`owner`,`provider_attachment`.`owner_access`,`provider_attachment`.`other_access`,`provider_attachment`.`group`,`provider_attachment`.`group_access`,`provider_attachment`.`enable`,`provider_attachment`.`display_name`,`provider_attachment`.`key_value_pair`,`provider_attachment`.`share`,`provider_attachment`.`perms2_owner`,`provider_attachment`.`perms2_owner_access`,`provider_attachment`.`global_access`,`provider_attachment`.`uuid` from `provider_attachment` where uuid = ?"

const insertProviderAttachmentVirtualRouterQuery = "insert into `ref_provider_attachment_virtual_router` (`from`, `to` ) values (?, ?);"

func CreateProviderAttachment(tx *sql.Tx, model *models.ProviderAttachment) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertProviderAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(utils.MustJSON(model.FQName),
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
		bool(model.IDPerms.Enable),
		string(model.DisplayName),
		utils.MustJSON(model.Annotations.KeyValuePair),
		utils.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID))

	stmtVirtualRouterRef, err := tx.Prepare(insertProviderAttachmentVirtualRouterQuery)
	if err != nil {
		return err
	}
	defer stmtVirtualRouterRef.Close()
	for _, ref := range model.VirtualRouterRefs {
		_, err = stmtVirtualRouterRef.Exec(model.UUID, ref.UUID)
	}

	return err
}

func scanProviderAttachment(rows *sql.Rows) (*models.ProviderAttachment, error) {
	m := models.MakeProviderAttachment()

	var jsonFQName string

	var jsonAnnotationsKeyValuePair string

	var jsonPerms2Share string

	if err := rows.Scan(&jsonFQName,
		&m.IDPerms.Description,
		&m.IDPerms.Created,
		&m.IDPerms.Creator,
		&m.IDPerms.UserVisible,
		&m.IDPerms.LastModified,
		&m.IDPerms.Permissions.Owner,
		&m.IDPerms.Permissions.OwnerAccess,
		&m.IDPerms.Permissions.OtherAccess,
		&m.IDPerms.Permissions.Group,
		&m.IDPerms.Permissions.GroupAccess,
		&m.IDPerms.Enable,
		&m.DisplayName,
		&jsonAnnotationsKeyValuePair,
		&jsonPerms2Share,
		&m.Perms2.Owner,
		&m.Perms2.OwnerAccess,
		&m.Perms2.GlobalAccess,
		&m.UUID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(jsonFQName), &m.FQName)

	json.Unmarshal([]byte(jsonAnnotationsKeyValuePair), &m.Annotations.KeyValuePair)

	json.Unmarshal([]byte(jsonPerms2Share), &m.Perms2.Share)

	return m, nil
}

func buildProviderAttachmentWhereQuery(where map[string]interface{}) (string, []interface{}) {
	if where == nil {
		return "", nil
	}
	results := []string{}
	values := []interface{}{}

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

	if value, ok := where["last_modified"]; ok {
		results = append(results, "last_modified = ?")
		values = append(values, value)
	}

	if value, ok := where["owner"]; ok {
		results = append(results, "owner = ?")
		values = append(values, value)
	}

	if value, ok := where["group"]; ok {
		results = append(results, "group = ?")
		values = append(values, value)
	}

	if value, ok := where["display_name"]; ok {
		results = append(results, "display_name = ?")
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

	return "where " + strings.Join(results, " and "), values
}

func ListProviderAttachment(tx *sql.Tx, where map[string]interface{}, offset int, limit int) ([]*models.ProviderAttachment, error) {
	result := models.MakeProviderAttachmentSlice()
	whereQuery, values := buildProviderAttachmentWhereQuery(where)
	var rows *sql.Rows
	var err error
	var query bytes.Buffer
	pagenationQuery := fmt.Sprintf("limit %d offset %d", limit, offset)
	query.WriteString(listProviderAttachmentQuery)
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
		m, _ := scanProviderAttachment(rows)
		result = append(result, m)
	}
	return result, nil
}

func ShowProviderAttachment(tx *sql.Tx, uuid string) (*models.ProviderAttachment, error) {
	rows, err := tx.Query(showProviderAttachmentQuery, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanProviderAttachment(rows)
	}
	return nil, nil
}

func UpdateProviderAttachment(tx *sql.Tx, uuid string, model *models.ProviderAttachment) error {
	return nil
}

func DeleteProviderAttachment(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteProviderAttachmentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return err
}
