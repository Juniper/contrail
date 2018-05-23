package db

import (
	"database/sql"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
)

//CreateSharing creates sharing information in DB.
func (db *Service) CreateSharing(tx *sql.Tx, table string, uuid string, shares []*models.ShareType) error {
	for _, share := range shares {
		err := db.createSharingEntry(tx, table, uuid, share.Tenant, int(share.TenantAccess))
		if err != nil {
			return err
		}
	}
	return nil
}

//UpdateSharing updates sharing data for a object by UUID.
func (db *Service) UpdateSharing(tx *sql.Tx, table string, uuid string, shares []interface{}) error {
	if len(shares) == 0 {
		return nil
	}
	_, err := tx.Exec(
		"delete from "+db.Dialect.quote("domain_share_"+table)+" where uuid = "+db.Dialect.placeholder(1), uuid)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"delete from "+db.Dialect.quote("tenant_share_"+table)+" where uuid = ?"+db.Dialect.placeholder(1), uuid)
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

func (db *Service) createSharingEntry(tx *sql.Tx, table string, uuid string, tenant string, tenantAccess int) error {
	shareParts := strings.Split(tenant, ":")
	if len(shareParts) < 2 {
		return common.ErrorBadRequest("invalid sharing entry")
	}

	shareType := shareParts[0]
	to := shareParts[1]

	_, err := tx.Exec(
		"insert into "+
			db.Dialect.quote(shareType+"_share_"+table)+" (uuid, access, "+
			db.Dialect.quote("to")+") values ("+db.Dialect.values("uuid", "access", "to")+");", // nolint
		uuid, tenantAccess, to)
	return err
}
