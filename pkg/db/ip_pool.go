package db

import (
	"bytes"
	"context"
	"database/sql"
	"net"
	"strings"

	"fmt"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"
)

// ipPool struct.
type ipPool struct {
	key   string
	start net.IP
	end   net.IP
}

// createIPPool creates ip pool.
func (db *Service) createIPPool(ctx context.Context, target *ipPool) error {
	return db.deallocateIPRange(ctx, target)
}

// getIPPools gets ip pools overlaps in given the range.
// return all if target.End is zero.
func (db *Service) getIPPools(ctx context.Context, target *ipPool) ([]*ipPool, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := GetTransaction(ctx)
	query.WriteString("select " + db.Dialect.wrapSelectIp("start") + ", " + db.Dialect.wrapSelectIp("end") + " from ipaddress_pool where ")
	query.WriteString(db.Dialect.quote("key") + " = '" + target.key + "'")
	var rows *sql.Rows
	var err error

	if target.end.Equal(net.IP{}) {
		query.WriteString(" order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String())
	} else {
		query.WriteString(" and " + d.comparableIp(target.start) + " < " + d.quote("end") + " and " +
			d.quote("start") + " < " + d.comparableIp(target.end) + " order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String())
	}
	pools := []*ipPool{}
	err = handleError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ip pool")
	}
	for rows.Next() {
		pool := &ipPool{
			key: target.key,
		}
		var start, end string
		err := rows.Scan(&start, &end)
		pool.start = stringToIP(start)
		pool.end = stringToIP(end)
		err = handleError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get ip pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// deleteIPPools deletes ip pool overlap with target range. delete all if target.End is zero. Whole pools are removed, not only overlapping parts.
func (db *Service) deleteIPPools(ctx context.Context, target *ipPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	var err error

	if target.end.Equal(net.IP{}) {
		_, err = tx.ExecContext(ctx, "delete from ipaddress_pool where "+
			db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+";", target.key)
	} else {
		query := "delete from ipaddress_pool where " +
			db.Dialect.quote("key") + " = '" + target.key + "' and " +
			d.comparableIp(target.start) + " < " + d.quote("end") + " and " +
			d.quote("start") + " < " + d.comparableIp(target.end)
		_, err = tx.ExecContext(ctx, query)
	}
	return errors.Wrap(handleError(err), "failed to delete ip pools")
}

//allocateIP allocates integer.
func (db *Service) allocateIP(ctx context.Context, key string) (net.IP, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " + db.Dialect.wrapSelectIp("start") + ", " + db.Dialect.wrapSelectIp("end") +
		" from ipaddress_pool where " +
		d.quote("key") + " = " + d.placeholder(1) +
		" limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)

	var start, end net.IP
	var startString, endString string
	err := row.Scan(&startString, &endString)

	start = stringToIP(startString)
	end = stringToIP(endString)

	if err != nil {
		return nil, handleError(err)
	}
	updatedStart := cidr.Inc(start)

	if bytes.Compare(updatedStart.To16(), end.To16()) < 0 {
		_, err = tx.ExecContext(ctx,
			"update ipaddress_pool set "+d.quote("start")+" = "+d.insertIp(updatedStart)+
				" where "+d.quote("key")+" = "+d.placeholder(1)+" and "+
				d.quote("start")+" = "+d.insertIp(start),
			key,
		)
	} else {
		_, err = tx.ExecContext(ctx,
			"delete from ipaddress_pool where "+d.quote("key")+" = "+d.placeholder(1)+" and "+
				d.quote("start")+" = "+d.comparableIp(start),
			key,
		)
	}
	if err != nil {
		return nil, handleError(err)
	}

	return start, nil
}

//setIP set a id for allocation pool.
func (db *Service) setIP(ctx context.Context, key string, ip net.IP) error {
	tx := GetTransaction(ctx)
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
		return common.ErrorNotFound
	}
	err = db.deleteIPPools(ctx, rangePool)
	if err != nil {
		return err
	}
	pool := pools[0]

	if pool.start.Equal(ip) {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ( '"+key+"', "+d.insertIp(cidr.Inc(pool.start))+", "+d.insertIp(pool.end)+");")
		if err != nil {
			return handleError(err)
		}
	} else if cidr.Dec(pool.end).Equal(ip) {

		fmt.Printf("MYK!\n")
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ( '"+key+"', "+d.insertIp(pool.start)+", "+d.insertIp(cidr.Dec(pool.end))+");")
		if err != nil {
			return handleError(err)
		}
	} else {
		// We need divide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ( '"+key+"', "+d.insertIp(pool.start)+", "+d.insertIp(ip)+");")
		if err != nil {
			return handleError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ( '"+key+"', "+d.insertIp(cidr.Inc(ip))+", "+d.insertIp(pool.end)+");")
		if err != nil {
			return handleError(err)
		}
	}
	return nil
}

//deallocateIP deallocate integer.
func (db *Service) deallocateIP(ctx context.Context, key string, ip net.IP) error {
	return db.deallocateIPRange(ctx, &ipPool{
		key:   key,
		start: ip,
		end:   cidr.Inc(ip),
	})
}

//deallocateIPRange deallocate integer range
func (db *Service) deallocateIPRange(ctx context.Context, target *ipPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	// range for pool we want to merge.
	// We need enlarge range so that we can merge pools on the next.
	mergePool := &ipPool{
		key:   target.key,
		start: cidr.Dec(target.start),
		end:   cidr.Inc(target.end),
	}

	pools, err := db.getIPPools(ctx, mergePool)
	if err != nil && err != common.ErrorNotFound {
		return err
	}

	start := target.start
	end := target.end

	// Clear overlapping ip pools
	if len(pools) > 0 {
		err := db.deleteIPPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = ipMin(start, pools[0].start)
		end = ipMax(end, pools[len(pools)-1].end)
	}
	q := "insert into ipaddress_pool (" + d.quoteSep("key", "start", "end") + ") values ( '" + target.key + "', " + d.insertIp(start) + ", " + d.insertIp(end) + ");"

	_, err = tx.ExecContext(
		ctx, q)
	return handleError(err)
}

// stringToIP translates string representation to IP, removing redundant '0' bytes from the end
func stringToIP(s string) net.IP {
	s = strings.TrimRightFunc(s, func(r rune) bool { return r == 0 })
	return net.ParseIP(s).To16()
}

// ipMax returns bigger of two addresses
func ipMax(a, b net.IP) net.IP {
	if bytes.Compare(a.To16(), b.To16()) > 0 {
		return a
	}
	return b
}

// ipMin returns smaller of two addresses
func ipMin(a, b net.IP) net.IP {
	if bytes.Compare(a.To16(), b.To16()) < 0 {
		return a.To16()
	}
	return b.To16()
}
