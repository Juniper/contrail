package db

import (
	"bytes"
	"context"
	"database/sql"
	"net"

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
	query.WriteString("select " + db.Dialect.quoteSep("start", "end") + "from ipaddress_pool where ")
	query.WriteString(db.Dialect.quote("key") + " = " + db.Dialect.placeholder(1))
	var rows *sql.Rows
	var err error

	if bytes.Compare(target.end, net.IP{}) == 0 {
		query.WriteString(" order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String(), target.key)
	} else {
		query.WriteString(" and " + d.placeholder(2) + "<" + d.quote("end") + " and " +
			d.quote("start") + " < " + d.placeholder(3) + " order by start for update")
		rows, err = tx.QueryContext(ctx, query.String(), target.key, target.start.String(), target.end.String())
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
		pool.start = net.ParseIP(start)
		pool.end = net.ParseIP(end)
		err = handleError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get ip pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// deleteIPPools deletes ip pool overlap with target range. delete all if target.End is zero.
func (db *Service) deleteIPPools(ctx context.Context, target *ipPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	var err error

	//TODO: Not sure if we should handle this that way
	if bytes.Compare(target.end, net.IP{}) == 0 {
		_, err = tx.ExecContext(ctx, "delete from ipaddress_pool where "+
			db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+";", target.key)
	} else {
		_, err = tx.ExecContext(ctx, "delete from ipaddress_pool where "+
			db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+" and "+
			d.placeholder(2)+"<"+d.quote("end")+" and "+
			d.quote("start")+" < "+d.placeholder(3),
			target.key,
			target.start.String(),
			target.end.String(),
		)
	}
	return errors.Wrap(handleError(err), "failed to delete ip pools")
}

//allocateIP allocates integer.
func (db *Service) allocateIP(ctx context.Context, key string) (net.IP, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " +
		d.quoteSep("start", "end") +
		" from ipaddress_pool where " +
		d.quote("key") + " = " + d.placeholder(1) +
		" order by " + d.quote("start") +
		" limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)
	var start, end net.IP
	err := row.Scan(&start, &end)
	if err != nil {
		return nil, handleError(err)
	}
	updatedStart := cidr.Inc(start)
	if bytes.Compare(updatedStart, end) < 0 {
		_, err = tx.ExecContext(ctx,
			"update ipaddress_pool set "+d.quote("start")+" = "+d.placeholder(1)+
				" where "+d.quote("key")+" = "+d.placeholder(2)+" and "+
				d.quote("start")+" = "+d.placeholder(3),
			updatedStart,
			key,
			start.String(),
		)
	} else {
		_, err = tx.ExecContext(ctx,
			"delete from ipaddress_pool where "+d.quote("key")+" = "+d.placeholder(1)+" and "+
				d.quote("start")+" = "+d.placeholder(2),
			key,
			start.String(),
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
	if bytes.Compare(pool.start, ip) == 0 {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, cidr.Inc(pool.start).String(), pool.end.String())
		if err != nil {
			return handleError(err)
		}
	} else if bytes.Compare(cidr.Dec(pool.end), ip) == 0 {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, pool.start.String(), cidr.Dec(pool.end).String())
		if err != nil {
			return handleError(err)
		}
	} else {
		// We need divide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, pool.start.String(), ip.String())
		if err != nil {
			return handleError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, cidr.Inc(ip).String(), pool.end.String())
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

func ipMax(a, b net.IP) net.IP {
	if bytes.Compare(a, b) > 0 {
		return a
	}
	return b
}

func ipMin(a, b net.IP) net.IP {
	if bytes.Compare(a, b) < 0 {
		return a
	}
	return b
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

	// Clear overlapping int pols
	if len(pools) > 0 {
		err := db.deleteIPPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = ipMin(start, pools[0].start)
		end = ipMax(end, pools[len(pools)-1].end)
	}

	_, err = tx.ExecContext(
		ctx,
		"insert into ipaddress_pool ("+d.quoteSep("key", "start", "end")+") values ("+
			d.values("key", "start", "end")+");",
		target.key, start.String(), end.String())
	return handleError(err)
}

//sizeIPPool returns size of a ip pool.
func (db *Service) sizeIPPool(ctx context.Context, key string) (int, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select sum( inet " +
		d.quote("end") + " - inet " + d.quote("start") + " ) as size" +
		" from ipaddress_pool where " +
		d.quote("key") + " = " + d.placeholder(1)
	row := tx.QueryRowContext(ctx, query, key)
	var size int
	err := row.Scan(&size)
	if err != nil {
		return 0, handleError(err)
	}
	return size, nil
}
