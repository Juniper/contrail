package db

import (
	"bytes"
	"context"

	"github.com/Juniper/contrail/pkg/types/ipam"
	"github.com/pkg/errors"
)

//CreateIntPool creates int pool.
func (db *DB) CreateIntPool(ctx context.Context, key string, start, end int) error {
	tx := GetTransaction(ctx)
	_, err := tx.Exec(
		"insert into int_pool ("+db.Dialect.quoteSep("key", "start", "end")+") values ("+
			db.Dialect.values("key", "start", "end")+");",
		key, start, end)
	err = handleError(err)
	return errors.Wrap(err, "failed to create int pool")
}

//GetIntPools gets int pools.
func (db *DB) GetIntPools(ctx context.Context, key string) ([]*ipam.IntPool, error) {
	var query bytes.Buffer
	tx := GetTransaction(ctx)
	query.WriteString("select " + db.Dialect.quoteSep("start", "end") + "from int_pool where ")
	query.WriteString(db.Dialect.quote("key") + " = " + db.Dialect.placeholder(1))
	pools := []*ipam.IntPool{}
	rows, err := tx.QueryContext(ctx, query.String(), key)
	err = handleError(err)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get int pool")
	}
	for rows.Next() {
		pool := &ipam.IntPool{
			Key: key,
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

//DeleteIntPools deletes int pool.
func (db *DB) DeleteIntPools(ctx context.Context, key string) error {
	tx := GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, "delete from int_pool where "+db.Dialect.quote("key")+" = "+db.Dialect.placeholder(1)+";", key)
	return errors.Wrap(handleError(err), "failed to delete int pools")
}

func (db *DB) CheckOverlap(ctx context.Context, key string, target ipam.IntPool) bool {
	return false
}

func (db *DB) AllocateInt(ctx context.Context, key string) (int, error) {
	tx := GetTransaction(ctx)
	d := db.Dialect
	query := "select " +
		d.quoteSep("start", "end") +
		" from int_pool where " +
		d.quote("key") + " = " + d.placeholder(1) +
		" order by " + d.quote("start") +
		" limit 1 for update"
	row := tx.QueryRowContext(ctx, query, key)
	var start, end int
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

func (db *DB) SetInt(ctx context.Context, key string, id int) error {
	return nil
}

func (db *DB) DeallocateInt(ctx context.Context, key string, id int) error {
	return nil
}

func (db *DB) SizeIntPool(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (db *DB) ClearIntPool(ctx context.Context, key string) error {
	return nil
}
