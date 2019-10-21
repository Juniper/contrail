package db

import (
	"bytes"
	"context"
	"database/sql"
	"net"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"

	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/errutil"
)

// ipPool struct, represents a range of available ips.
type ipPool struct {
	key   string
	start net.IP
	end   net.IP
}

// createIPPool creates ip pool.
func (db *Service) createIPPool(ctx context.Context, target *ipPool) error {
	return db.deallocateIPRange(ctx, target)
}

// getIPPools gets ip pools overlapping given range.
// return all if target.End is zero.
func (db *Service) getIPPools(ctx context.Context, target *ipPool) ([]*ipPool, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := basedb.GetTransaction(ctx)
	basedb.WriteStrings(
		&query,
		"select ",
		db.Dialect.SelectIP("start"),
		", ",
		db.Dialect.SelectIP("end"),
		" from ipaddress_pool where ",
		db.Dialect.Quote("key"),
		" = ",
		db.Dialect.Placeholder(1),
	)

	var rows *sql.Rows
	var err error
	if target.end.Equal(net.IP{}) {
		basedb.WriteStrings(&query, " order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String(), target.key)
	} else {
		basedb.WriteStrings(
			&query,
			" and ",
			d.LiteralIP(target.start),
			" < ",
			d.Quote("end"),
			" and ",
			d.Quote("start"),
			" < ",
			d.LiteralIP(target.end),
			" order by start for update ",
		)
		rows, err = tx.QueryContext(ctx, query.String(), target.key)
	}
	pools := []*ipPool{}
	err = basedb.FormatDBError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ip pool")
	}
	for rows.Next() {
		pool := &ipPool{
			key: target.key,
		}
		var start, end string
		err := rows.Scan(&start, &end)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan ip pool")
		}
		pool.start = stringToIP(start)
		pool.end = stringToIP(end)
		err = basedb.FormatDBError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse ip pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// deleteIPPools deletes ip pool that overlaps target range. delete all if target.End is zero.
// Whole pools are removed, not only overlapping parts.
func (db *Service) deleteIPPools(ctx context.Context, target *ipPool) (int64, error) {
	tx := basedb.GetTransaction(ctx)
	d := db.Dialect
	query := ""
	if target.end.Equal(net.IP{}) {
		query = "delete from ipaddress_pool where " + db.Dialect.Quote("key") + " = " + db.Dialect.Placeholder(1)
	} else {
		query = "delete from ipaddress_pool where " + db.Dialect.Quote("key") + " = " + db.Dialect.Placeholder(1) +
			" and " + d.LiteralIP(target.start) + " < " + d.Quote("end") + " and " +
			d.Quote("start") + " < " + d.LiteralIP(target.end)
	}

	res, err := tx.ExecContext(ctx, query, target.key)
	if err != nil {
		return 0, errors.Wrap(basedb.FormatDBError(err), "failed to delete ip pools")
	}

	return res.RowsAffected()
}

// allocateIP allocates smallest available ip.
func (db *Service) allocateIP(ctx context.Context, key string) (net.IP, error) {
	tx := basedb.GetTransaction(ctx)
	d := db.Dialect
	query := "select " + db.Dialect.SelectIP("start") + ", " + db.Dialect.SelectIP("end") +
		" from ipaddress_pool where " +
		d.Quote("key") + " = " + db.Dialect.Placeholder(1) + " limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)

	var start, end net.IP
	var startString, endString string
	err := row.Scan(&startString, &endString)
	if err != nil {
		return nil, basedb.FormatDBError(err)
	}

	start = stringToIP(startString)
	end = stringToIP(endString)
	updatedStart := cidr.Inc(start)

	if bytes.Compare(updatedStart.To16(), end.To16()) <= 0 {
		_, err = tx.ExecContext(ctx,
			"update ipaddress_pool set "+d.Quote("start")+" = "+d.LiteralIP(updatedStart)+
				" where "+d.Quote("key")+" = "+db.Dialect.Placeholder(1)+" and "+d.Quote("start")+
				" = "+d.LiteralIP(start), key)
	} else {
		_, err = tx.ExecContext(ctx,
			"delete from ipaddress_pool where "+d.Quote("key")+" = "+db.Dialect.Placeholder(1)+" and "+
				d.Quote("start")+" = "+d.LiteralIP(start), key)
	}
	if err != nil {
		return nil, basedb.FormatDBError(err)
	}

	return start, nil
}

//setIP allocates given ip, if it's available. Can split pools.
func (db *Service) setIP(ctx context.Context, key string, ip net.IP) error {
	tx := basedb.GetTransaction(ctx)
	d := db.Dialect
	rangePool := &ipPool{
		key:   key,
		start: ip,
		end:   cidr.Inc(ip),
	}

	pools, err := db.getIPPools(ctx, rangePool)
	if err != nil {
		return err
	}
	if len(pools) == 0 {
		return errors.Errorf("Cannot allocate address %s : pool containing this address not found", ip.String())
	}
	_, err = db.deleteIPPools(ctx, rangePool)
	if err != nil {
		return err
	}
	pool := pools[0]

	if pool.start.Equal(ip) {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+
				") values ( "+db.Dialect.Placeholder(1)+", "+d.LiteralIP(cidr.Inc(pool.start))+", "+
				d.LiteralIP(pool.end)+")", key)
		if err != nil {
			return basedb.FormatDBError(err)
		}
	} else if cidr.Dec(pool.end).Equal(ip) {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(pool.start)+", "+d.LiteralIP(cidr.Dec(pool.end))+")", key)
		if err != nil {
			return basedb.FormatDBError(err)
		}
	} else {
		// We need divide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(pool.start)+", "+d.LiteralIP(ip)+")", key)
		if err != nil {
			return basedb.FormatDBError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(cidr.Inc(ip))+", "+d.LiteralIP(pool.end)+")", key)
		if err != nil {
			return basedb.FormatDBError(err)
		}
	}
	return nil
}

//deallocateIP deallocates ip.
//nolint: unused
func (db *Service) deallocateIP(ctx context.Context, key string, ip net.IP) error {
	return db.deallocateIPRange(ctx, &ipPool{
		key:   key,
		start: ip,
		end:   cidr.Inc(ip),
	})
}

//deallocateIPRange deallocates ip range.
func (db *Service) deallocateIPRange(ctx context.Context, target *ipPool) error {
	tx := basedb.GetTransaction(ctx)
	d := db.Dialect
	// range for pool we want to merge.
	// We need enlarge range so that we can merge pools on the next.
	mergePool := &ipPool{
		key:   target.key,
		start: cidr.Dec(target.start),
		end:   cidr.Inc(target.end),
	}

	pools, err := db.getIPPools(ctx, mergePool)
	if err != nil && err != errutil.ErrorNotFound {
		return err
	}

	start := target.start
	end := target.end

	// Clear overlapping ip pools
	if len(pools) > 0 {
		_, err = db.deleteIPPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = ipMin(start, pools[0].start)
		end = ipMax(end, pools[len(pools)-1].end)
	}
	q := "insert into ipaddress_pool (" + d.QuoteSep("key", "start", "end") + ") values ( " +
		db.Dialect.Placeholder(1) + ", " + d.LiteralIP(start) + ", " + d.LiteralIP(end) + ")"

	_, err = tx.ExecContext(ctx, q, target.key)
	return basedb.FormatDBError(err)
}

// stringToIP translates string representation to IP, removing redundant '0' bytes from the end.
func stringToIP(s string) net.IP {
	s = strings.TrimRightFunc(s, func(r rune) bool { return r == 0 })
	return net.ParseIP(s).To16()
}

// ipMax returns bigger of two addresses
func ipMax(a, b net.IP) net.IP {
	if bytes.Compare(a.To16(), b.To16()) > 0 {
		return a.To16()
	}
	return b.To16()
}

// ipMin returns smaller of two addresses
func ipMin(a, b net.IP) net.IP {
	if bytes.Compare(a.To16(), b.To16()) < 0 {
		return a.To16()
	}
	return b.To16()
}
