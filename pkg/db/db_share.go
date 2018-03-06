package db

import (
	"database/sql"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//CreateSharing creates sharing information in DB.
func (db *DB) CreateSharing(tx *sql.Tx, table string, uuid string, shares []*models.ShareType) error {
	for _, share := range shares {
		err := db.createSharingEntry(tx, table, uuid, share.Tenant, int(share.TenantAccess))
		if err != nil {
			return err
		}
	}
	return nil
}

//UpdateSharing updates sharing data for a object by UUID.
func (db *DB) UpdateSharing(tx *sql.Tx, table string, uuid string, shares []interface{}) error {
	if len(shares) == 0 {
		return nil
	}
	_, err := tx.Exec(
		"delete from "+db.Dialect.quote("domain_share_"+table)+" where uuid = ?;", uuid)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"delete from "+db.Dialect.quote("tenant_share_"+table)+" where uuid = ?;", uuid)
	if err != nil {
		return err
	}
	for _, share := range shares {
		err = db.createSharingEntry(tx, table, uuid, common.InterfaceToString(share.(map[string]interface{})["tenant"]),
			common.InterfaceToInt(share.(map[string]interface{})["tenant_access"]))
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) createSharingEntry(tx *sql.Tx, table string, uuid string, tenant string, tenantAccess int) error {
	shareParts := strings.Split(tenant, ":")
	shareType := "domain"
	if len(shareParts) > 1 {
		shareType = "tenant"
	}
	resourceMetaData, err := db.GetMetaData(tx, "", shareParts)
	if err != nil {
		return errors.Wrap(err, "can't find resource")
	}
	to := resourceMetaData.UUID
	_, err = tx.Exec(
		"insert into "+db.Dialect.quote(shareType+"_share_"+table)+" (uuid, access, "+db.Dialect.quote("to")+") values (?,?,?);", // nolint
		uuid, tenantAccess, to)
	return err
}
