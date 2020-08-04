package db

import (
	"bytes"
	"context"
	"database/sql"
	"net"
	"strings"

	"github.com/Juniper/asf/pkg/errutil"
	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/pkg/errors"
)

// EmptyIntOwner is useful for creating pool when owner is not relevant.
const EmptyIntOwner = ""

// IntPool represents the half-open integer range [Start, End) in the set of integers identified by Key.
type IntPool struct {
	Key   string
	Start int64
	End   int64
}

// CreateIntPool creates int pool.
func (db *DB) CreateIntPool(ctx context.Context, pool string, start int64, end int64) error {
	intPool := &IntPool{
		Key:   pool,
		Start: start,
		End:   end,
	}

	intPools, err := db.GetIntPools(ctx, intPool)
	if err != nil {
		return err
	}

	if len(intPools) > 0 {
		return errutil.ErrorConflictf("int pool %+v already in use", intPool)
	}

	return db.DeallocateIntRange(ctx, intPool)
}

// DeleteIntPool deletes int pool.
func (db *DB) DeleteIntPool(ctx context.Context, pool string) error {
	return db.deleteIntPools(ctx, &IntPool{
		Key: pool,
	})
}

// GetIntOwner returns owner of an allocated integer
func (db *DB) GetIntOwner(ctx context.Context, pool string, value int64) (string, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := GetTransaction(ctx)
	query.WriteString("select " + d.QuoteSep("owner") + "from int_owner")
	query.WriteString(" where " + d.Quote("pool") + " = " + d.Placeholder(1))
	query.WriteString(" and " + d.Quote("value") + " = " + d.Placeholder(2))

	var owner string
	err := tx.QueryRowContext(ctx, query.String(), pool, value).Scan(&owner)
	err = FormatDBError(err)
	if err != nil {
		if errutil.IsNotFound(err) {
			return "", err
		}
		return "", errors.Wrap(err, "failed to get int owner")
	}
	return owner, nil
}

// GetIntPools gets int pools overlaps in given the range.
// return all if target.End is zero.
func (db *DB) GetIntPools(ctx context.Context, target *IntPool) ([]*IntPool, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := GetTransaction(ctx)
	query.WriteString("select " + db.Dialect.QuoteSep("start", "end") + "from int_pool where ")
	query.WriteString(db.Dialect.Quote("key") + " = " + db.Dialect.Placeholder(1))
	var rows *sql.Rows
	var err error
	if target.End == 0 {
		query.WriteString(" order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String(), target.Key)
	} else {
		query.WriteString(" and " + d.Placeholder(2) + "<" + d.Quote("end") + " and " +
			d.Quote("start") + " < " + d.Placeholder(3) + " order by start for update")
		rows, err = tx.QueryContext(ctx, query.String(), target.Key, target.Start, target.End)
	}
	pools := []*IntPool{}
	err = FormatDBError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get int pool")
	}
	for rows.Next() {
		pool := &IntPool{
			Key: target.Key,
		}
		err := rows.Scan(&pool.Start, &pool.End)
		err = FormatDBError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get int pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// AllocateInt allocates integer.
func (db *DB) AllocateInt(ctx context.Context, key, owner string) (int64, error) {
	if key == "" {
		return 0, errors.New("empty int-pool key provided to allocate")
	}
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " +
		d.QuoteSep("start", "end") +
		" from int_pool where " +
		d.Quote("key") + " = " + d.Placeholder(1) +
		" order by " + d.Quote("start") +
		" limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)
	var start, end int64
	err := row.Scan(&start, &end)
	if err != nil {
		return 0, FormatDBError(err)
	}
	updatedStart := start + 1
	if updatedStart < end {
		_, err = tx.ExecContext(ctx,
			"update int_pool set "+d.Quote("start")+" = "+d.Placeholder(1)+
				" where "+d.Quote("key")+" = "+d.Placeholder(2)+" and "+
				d.Quote("start")+" = "+d.Placeholder(3),
			updatedStart,
			key,
			start,
		)
	} else {
		_, err = tx.ExecContext(ctx,
			"delete from int_pool where "+d.Quote("key")+" = "+d.Placeholder(1)+" and "+
				d.Quote("start")+" = "+d.Placeholder(2),
			key,
			start,
		)
	}
	if err != nil {
		return 0, FormatDBError(err)
	}

	if err := db.insertIntOwner(ctx, tx, start, key, owner); err != nil {
		return 0, FormatDBError(err)
	}

	return start, nil
}

// SetInt set a id for allocation pool.
func (db *DB) SetInt(ctx context.Context, key string, id int64, owner string) error {
	if key == "" {
		return errors.New("empty int-pool key provided to set")
	}

	storedOwner, err := db.GetIntOwner(ctx, key, id)
	if err != nil && !errutil.IsNotFound(err) {
		return err
	}

	if owner != "" && storedOwner == owner {
		return nil
	}

	tx := GetTransaction(ctx)

	if err := db.insertIntIntoPool(ctx, tx, key, id); err != nil {
		return FormatDBError(err)
	}

	if err := db.insertIntOwner(ctx, tx, id, key, owner); err != nil {
		return FormatDBError(err)
	}
	return nil
}

func (db *DB) insertIntIntoPool(ctx context.Context, tx *sql.Tx, key string, id int64) (err error) {
	d := db.Dialect
	rangePool := &IntPool{
		Key:   key,
		Start: id,
		End:   id + 1,
	}

	pools, err := db.GetIntPools(ctx, rangePool)
	if err != nil {
		return err
	}

	if len(pools) == 0 {
		return errutil.ErrorNotFound()
	}

	err = db.deleteIntPools(ctx, rangePool)
	if err != nil {
		return err
	}

	pool := pools[0]
	if pool.Start == id {
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.QuoteSep("key", "start", "end")+") values ("+
				d.Values("key", "start", "end")+");",
			key, pool.Start+1, pool.End)
	} else if pool.End-1 == id {
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.QuoteSep("key", "start", "end")+") values ("+
				d.Values("key", "start", "end")+");",
			key, pool.Start, pool.End-1)
	} else {
		// We need divide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.QuoteSep("key", "start", "end")+") values ("+
				d.Values("key", "start", "end")+");",
			key, pool.Start, id)
		if err != nil {
			return FormatDBError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.QuoteSep("key", "start", "end")+") values ("+
				d.Values("key", "start", "end")+");",
			key, id+1, pool.End)
	}

	return err
}

// DeallocateInt deallocate integer.
func (db *DB) DeallocateInt(ctx context.Context, key string, id int64) error {
	return db.DeallocateIntRange(ctx, &IntPool{
		Key:   key,
		Start: id,
		End:   id + 1,
	})
}

func (db *DB) insertIntOwner(ctx context.Context, tx *sql.Tx, value int64, pool, owner string) error {
	if owner == "" {
		// Inserting records with empty owner is pointless.
		return nil
	}

	d := db.Dialect
	_, err := tx.ExecContext(ctx,
		"insert into int_owner ("+d.QuoteSep("pool", "value", "owner")+") values ("+
			d.Values("pool", "value", "owner")+");", pool, value, owner,
	)
	return err
}

func (db *DB) deleteIntOwner(ctx context.Context, tx *sql.Tx, value int64, pool string) error {
	d := db.Dialect
	_, err := tx.ExecContext(ctx,
		"delete from int_owner where "+d.Quote("pool")+" = "+d.Placeholder(1)+" and "+
			d.Quote("value")+" = "+d.Placeholder(2), pool, value,
	)
	return err
}

func intMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func intMin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// DeallocateIntRange deallocate integer range
func (db *DB) DeallocateIntRange(ctx context.Context, target *IntPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	// range for pool we want to merge.
	// We need enlarge range so that we can merge pools on the next.
	mergePool := &IntPool{
		Key:   target.Key,
		Start: target.Start - 1,
		End:   target.End + 1,
	}
	pools, err := db.GetIntPools(ctx, mergePool)
	if err != nil && err != errutil.ErrorNotFound() {
		return err
	}

	start := target.Start
	end := target.End

	// Clear overlapping int pols
	if len(pools) > 0 {
		err = db.deleteIntPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = intMin(start, pools[0].Start)
		end = intMax(end, pools[len(pools)-1].End)
	}
	if err = db.deleteIntOwner(ctx, tx, target.Start, target.Key); err != nil {
		return FormatDBError(err)
	}
	_, err = tx.ExecContext(
		ctx,
		"insert into int_pool ("+d.QuoteSep("key", "start", "end")+") values ("+
			d.Values("key", "start", "end")+");",
		target.Key, start, end)
	return FormatDBError(err)
}

// SizeIntPool returns size of a int pool.
func (db *DB) SizeIntPool(ctx context.Context, key string) (int, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select sum( " +
		d.Quote("end") + " - " + d.Quote("start") + " ) as size" +
		" from int_pool where " +
		d.Quote("key") + " = " + d.Placeholder(1)
	row := tx.QueryRowContext(ctx, query, key)
	var size int
	err := row.Scan(&size)
	if err != nil {
		return 0, FormatDBError(err)
	}
	return size, nil
}

// deleteIntPools deletes int pool overlap with target range. delete all if target.End is zero.
func (db *DB) deleteIntPools(ctx context.Context, target *IntPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	var err error
	if target.End == 0 {
		_, err = tx.ExecContext(ctx, "delete from int_pool where "+
			db.Dialect.Quote("key")+" = "+db.Dialect.Placeholder(1)+";", target.Key)
	} else {
		_, err = tx.ExecContext(ctx, "delete from int_pool where "+
			db.Dialect.Quote("key")+" = "+db.Dialect.Placeholder(1)+" and "+
			d.Placeholder(2)+"<"+d.Quote("end")+" and "+
			d.Quote("start")+" < "+d.Placeholder(3),
			target.Key,
			target.Start,
			target.End,
		)
	}
	return errors.Wrap(FormatDBError(err), "failed to delete int pools")
}

// IPPool struct, represents a range of available ips.
type IPPool struct {
	Key   string
	Start net.IP
	End   net.IP
}

// CreateIPPool creates ip pool.
func (db *DB) CreateIPPool(ctx context.Context, target *IPPool) error {
	return db.DeallocateIPRange(ctx, target)
}

// GetIPPools gets ip pools overlapping given range.
// return all if target.End is zero.
func (db *DB) GetIPPools(ctx context.Context, target *IPPool) ([]*IPPool, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := GetTransaction(ctx)
	WriteStrings(
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
	if target.End.Equal(net.IP{}) {
		WriteStrings(&query, " order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String(), target.Key)
	} else {
		WriteStrings(
			&query,
			" and ",
			d.LiteralIP(target.Start),
			" < ",
			d.Quote("end"),
			" and ",
			d.Quote("start"),
			" < ",
			d.LiteralIP(target.End),
			" order by start for update ",
		)
		rows, err = tx.QueryContext(ctx, query.String(), target.Key)
	}
	pools := []*IPPool{}
	err = FormatDBError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ip pool")
	}
	for rows.Next() {
		pool := &IPPool{
			Key: target.Key,
		}
		var start, end string
		err := rows.Scan(&start, &end)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan ip pool")
		}
		pool.Start = stringToIP(start)
		pool.End = stringToIP(end)
		err = FormatDBError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse ip pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

// DeleteIPPools deletes ip pool that overlaps target range. delete all if target.End is zero.
// Whole pools are removed, not only overlapping parts.
func (db *DB) DeleteIPPools(ctx context.Context, target *IPPool) (int64, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := ""
	if target.End.Equal(net.IP{}) {
		query = "delete from ipaddress_pool where " + db.Dialect.Quote("key") + " = " + db.Dialect.Placeholder(1)
	} else {
		query = "delete from ipaddress_pool where " + db.Dialect.Quote("key") + " = " + db.Dialect.Placeholder(1) +
			" and " + d.LiteralIP(target.Start) + " < " + d.Quote("end") + " and " +
			d.Quote("start") + " < " + d.LiteralIP(target.End)
	}

	res, err := tx.ExecContext(ctx, query, target.Key)
	if err != nil {
		return 0, errors.Wrap(FormatDBError(err), "failed to delete ip pools")
	}

	return res.RowsAffected()
}

// AllocateIP allocates smallest available ip.
func (db *DB) AllocateIP(ctx context.Context, key string) (net.IP, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " + db.Dialect.SelectIP("start") + ", " + db.Dialect.SelectIP("end") +
		" from ipaddress_pool where " +
		d.Quote("key") + " = " + db.Dialect.Placeholder(1) + " limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)

	var start, end net.IP
	var startString, endString string
	err := row.Scan(&startString, &endString)
	if err != nil {
		return nil, FormatDBError(err)
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
		return nil, FormatDBError(err)
	}

	return start, nil
}

// SetIP allocates given ip, if it's available. Can split pools.
func (db *DB) SetIP(ctx context.Context, key string, ip net.IP) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	rangePool := &IPPool{
		Key:   key,
		Start: ip,
		End:   cidr.Inc(ip),
	}

	pools, err := db.GetIPPools(ctx, rangePool)
	if err != nil {
		return err
	}
	if len(pools) == 0 {
		return errors.Errorf("Cannot allocate address %s : pool containing this address not found", ip.String())
	}
	_, err = db.DeleteIPPools(ctx, rangePool)
	if err != nil {
		return err
	}
	pool := pools[0]

	if pool.Start.Equal(ip) {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+
				") values ( "+db.Dialect.Placeholder(1)+", "+d.LiteralIP(cidr.Inc(pool.Start))+", "+
				d.LiteralIP(pool.End)+")", key)
		if err != nil {
			return FormatDBError(err)
		}
	} else if cidr.Dec(pool.End).Equal(ip) {
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(pool.Start)+", "+d.LiteralIP(cidr.Dec(pool.End))+")", key)
		if err != nil {
			return FormatDBError(err)
		}
	} else {
		// We need divide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(pool.Start)+", "+d.LiteralIP(ip)+")", key)
		if err != nil {
			return FormatDBError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into ipaddress_pool ("+d.QuoteSep("key", "start", "end")+") values ( "+
				db.Dialect.Placeholder(1)+", "+d.LiteralIP(cidr.Inc(ip))+", "+d.LiteralIP(pool.End)+")", key)
		if err != nil {
			return FormatDBError(err)
		}
	}
	return nil
}

//DeallocateIP deallocates ip.
func (db *DB) DeallocateIP(ctx context.Context, key string, ip net.IP) error {
	return db.DeallocateIPRange(ctx, &IPPool{
		Key:   key,
		Start: ip,
		End:   cidr.Inc(ip),
	})
}

//DeallocateIPRange deallocates ip range.
func (db *DB) DeallocateIPRange(ctx context.Context, target *IPPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	// range for pool we want to merge.
	// We need enlarge range so that we can merge pools on the next.
	mergePool := &IPPool{
		Key:   target.Key,
		Start: cidr.Dec(target.Start),
		End:   cidr.Inc(target.End),
	}

	pools, err := db.GetIPPools(ctx, mergePool)
	if err != nil && err != errutil.ErrorNotFound() {
		return err
	}

	start := target.Start
	end := target.End

	// Clear overlapping ip pools
	if len(pools) > 0 {
		_, err = db.DeleteIPPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = ipMin(start, pools[0].Start)
		end = ipMax(end, pools[len(pools)-1].End)
	}
	q := "insert into ipaddress_pool (" + d.QuoteSep("key", "start", "end") + ") values ( " +
		db.Dialect.Placeholder(1) + ", " + d.LiteralIP(start) + ", " + d.LiteralIP(end) + ")"

	_, err = tx.ExecContext(ctx, q, target.Key)
	return FormatDBError(err)
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
