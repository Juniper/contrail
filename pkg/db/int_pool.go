package db

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/pkg/errors"
)

//IntPool for int pool.
type IntPool struct {
	Key   string
	Start int64
	End   int64
}

//CreateIntPool creates int pool.
func (db *Service) CreateIntPool(ctx context.Context, target *IntPool) error {
	return db.DeallocateIntRange(ctx, target)
}

//GetIntPools gets int pools overlaps in given the range.
//return all if target.End is zero.
func (db *Service) GetIntPools(ctx context.Context, target *IntPool) ([]*IntPool, error) {
	var query bytes.Buffer
	d := db.Dialect
	tx := GetTransaction(ctx)
	query.WriteString("select " + db.Dialect.quoteSep("start", "end") + "from int_pool where ")
	query.WriteString(db.Dialect.quote("key") + " = " + db.Dialect.placeholder(1))
	var rows *sql.Rows
	var err error
	if target.End == 0 {
		query.WriteString(" order by start for update ")
		rows, err = tx.QueryContext(ctx, query.String(), target.Key)
	} else {
		query.WriteString(" and " + d.placeholder(2) + "<" + d.quote("end") + " and " +
			d.quote("start") + " < " + d.placeholder(3) + " order by start for update")
		rows, err = tx.QueryContext(ctx, query.String(), target.Key, target.Start, target.End)
	}
	pools := []*IntPool{}
	err = handleError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get int pool")
	}
	for rows.Next() {
		pool := &IntPool{
			Key: target.Key,
		}
		err := rows.Scan(&pool.Start, &pool.End)
		err = handleError(err)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get int pool")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

//DeleteIntPools deletes int pool overlap with target range. delete all if target.End is zero.
func (db *Service) DeleteIntPools(ctx context.Context, target *IntPool) error {
	tx := GetTransaction(ctx)
	d := db.Dialect
	var err error
	if target.End == 0 {
		_, err = tx.ExecContext(ctx, "delete from int_pool where "+
			db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+";", target.Key)
	} else {
		_, err = tx.ExecContext(ctx, "delete from int_pool where "+
			db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+" and "+
			d.placeholder(2)+"<"+d.quote("end")+" and "+
			d.quote("start")+" < "+d.placeholder(3),
			target.Key,
			target.Start,
			target.End,
		)
	}
	return errors.Wrap(handleError(err), "failed to delete int pools")
}

//AllocateInt allocates integer.
func (db *Service) AllocateInt(ctx context.Context, key string) (int64, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " +
		d.quoteSep("start", "end") +
		" from int_pool where " +
		d.quote("key") + " = " + d.placeholder(1) +
		" order by " + d.quote("start") +
		" limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)
	var start, end int64
	err := row.Scan(&start, &end)
	if err != nil {
		return 0, handleError(err)
	}
	updatedStart := start + 1
	if updatedStart < end {
		_, err = tx.ExecContext(ctx,
			"update int_pool set "+d.quote("start")+" = "+d.placeholder(1)+
				" where "+d.quote("key")+" = "+d.placeholder(2)+" and "+
				d.quote("start")+" = "+d.placeholder(3),
			updatedStart,
			key,
			start,
		)
	} else {
		_, err = tx.ExecContext(ctx,
			"delete from int_pool where "+d.quote("key")+" = "+d.placeholder(1)+" and "+
				d.quote("start")+" = "+d.placeholder(2),
			key,
			start,
		)
	}
	if err != nil {
		return 0, handleError(err)
	}
	return start, nil
}

//SetInt set a id for allocation pool.
func (db *Service) SetInt(ctx context.Context, key string, id int64) error {
	tx := GetTransaction(ctx)
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
		return common.ErrorNotFound
	}
	err = db.DeleteIntPools(ctx, rangePool)
	if err != nil {
		return err
	}
	pool := pools[0]
	if pool.Start == id {
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, pool.Start+1, pool.End)
		if err != nil {
			return handleError(err)
		}
	} else if pool.End-1 == id {
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, pool.Start, pool.End-1)
		if err != nil {
			return handleError(err)
		}
	} else {
		// We need devide one pool to two.
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, pool.Start, id)
		if err != nil {
			return handleError(err)
		}
		_, err = tx.ExecContext(
			ctx,
			"insert into int_pool ("+d.quoteSep("key", "start", "end")+") values ("+
				d.values("key", "start", "end")+");",
			key, id+1, pool.End)
		if err != nil {
			return handleError(err)
		}
	}
	return nil
}

//DeallocateInt deallocate integer.
func (db *Service) DeallocateInt(ctx context.Context, key string, id int64) error {
	return db.DeallocateIntRange(ctx, &IntPool{
		Key:   key,
		Start: id,
		End:   id + 1,
	})
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

//DeallocateIntRange deallocate integer range
func (db *Service) DeallocateIntRange(ctx context.Context, target *IntPool) error {
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
	if err != nil && err != common.ErrorNotFound {
		return err
	}

	start := target.Start
	end := target.End

	// Clear overlapping int pols
	if len(pools) > 0 {
		err := db.DeleteIntPools(ctx, mergePool)
		if err != nil {
			return err
		}
		// Extend range based on existing pools.
		start = intMin(start, pools[0].Start)
		end = intMax(end, pools[len(pools)-1].End)
	}
	_, err = tx.ExecContext(
		ctx,
		"insert into int_pool ("+d.quoteSep("key", "start", "end")+") values ("+
			d.values("key", "start", "end")+");",
		target.Key, start, end)
	return handleError(err)
}

//SizeIntPool returns size of a int pool.
func (db *Service) SizeIntPool(ctx context.Context, key string) (int, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select sum( " +
		d.quote("end") + " - " + d.quote("start") + " ) as size" +
		" from int_pool where " +
		d.quote("key") + " = " + d.placeholder(1)
	row := tx.QueryRowContext(ctx, query, key)
	var size int
	err := row.Scan(&size)
	if err != nil {
		return 0, handleError(err)
	}
	return size, nil
}
