package basedb

import (
	"database/sql"
	"strings"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/format"
)

//Sharable resource can share.
type Sharable interface {
	GetTenantAccess() int64
	GetTenant() string
}

//CreateSharing creates sharing information in DB.
func (db *BaseDB) CreateSharing(tx *sql.Tx, table string, uuid string, share Sharable) error {
	return db.createSharingEntry(tx, table, uuid, share.GetTenant(), int(share.GetTenantAccess()))
}

//UpdateSharing updates sharing data for a object by UUID.
func (db *BaseDB) UpdateSharing(tx *sql.Tx, table string, uuid string, shares []interface{}) error {
	if len(shares) == 0 {
		return nil
	}
	_, err := tx.Exec(
		"delete from "+db.Dialect.Quote("domain_share_"+table)+" where uuid = "+db.Dialect.Placeholder(1), uuid)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"delete from "+db.Dialect.Quote("tenant_share_"+table)+" where uuid = ?"+db.Dialect.Placeholder(1), uuid)
	if err != nil {
		return err
	}
	for _, share := range shares {
		err = db.createSharingEntry(tx, table, uuid, format.InterfaceToString(
			share.(map[string]interface{})["tenant"]), //nolint: errcheck
			format.InterfaceToInt(share.(map[string]interface{})["tenant_access"])) //nolint: errcheck
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *BaseDB) createSharingEntry(tx *sql.Tx, table string, uuid string, tenant string, tenantAccess int) error {
	shareParts := strings.Split(tenant, ":")
	if len(shareParts) < 2 {
		return errutil.ErrorBadRequest("invalid sharing entry")
	}

	shareType := shareParts[0]
	to := shareParts[1]

	_, err := tx.Exec(
		"insert into "+
			db.Dialect.Quote(shareType+"_share_"+table)+" (uuid, access, "+
			db.Dialect.Quote("to")+") values ("+db.Dialect.Values("uuid", "access", "to")+");",
		uuid, tenantAccess, to)
	return err
}
