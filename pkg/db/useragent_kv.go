package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"
)

// StoreKV stores a value under given key.
// Updates the value if key is already present.
func (db *Service) StoreKV(ctx context.Context, key string, value string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		return db.storeKV(ctx, key, value)
	})
}

// RetrieveValue retrieves the value stored under the given key.
// Returns an error if key is not present.
func (db *Service) RetrieveValue(ctx context.Context, key string) (val string, err error) {
	if err = db.DoInTransaction(ctx, func(ctx context.Context) error {
		val, err = db.retrieveValue(ctx, key)
		return err
	}); err != nil {
		return "", err
	}

	return val, nil
}

// RetrieveValues retrieves values corresponding to the given list of keys.
// The values are returned in an arbitrary order. Keys not present in the store are ignored.
func (db *Service) RetrieveValues(ctx context.Context, keys []string) (vals []string, err error) {
	if err = db.DoInTransaction(ctx, func(ctx context.Context) error {
		vals, err = db.retrieveValues(ctx, keys)
		return err
	}); err != nil {
		return nil, err
	}

	return vals, nil
}

// DeleteKey deletes the value under the given key.
// Nothing happens if the key is not present.
func (db *Service) DeleteKey(ctx context.Context, key string) error {
	return db.DoInTransaction(ctx, func(ctx context.Context) error {
		return db.deleteKey(ctx, key)
	})
}

// RetrieveKVPs returns the entire store as a list of (key, value) pairs.
func (db *Service) RetrieveKVPs(ctx context.Context) (kvps []*models.KeyValuePair, err error) {
	if err = db.DoInTransaction(ctx, func(ctx context.Context) error {
		kvps, err = db.retrieveKVPs(ctx)
		return err
	}); err != nil {
		return nil, err
	}

	return kvps, nil
}

func (db *Service) storeKV(ctx context.Context, key string, value string) error {
	d := db.Dialect

	// Try to update first. Insert if the key is not present.
	queryUpdate := fmt.Sprintf(
		"update kv_store set %s = %s where %s = %s",
		d.quote("value"), d.placeholder(1), d.quote("key"), d.placeholder(2))
	queryInsert := fmt.Sprintf(
		"insert into kv_store (%s) select %s where not exists (select 1 from kv_store where %s = %s)",
		d.quoteSep("key", "value"), d.values(d.placeholder(1), d.placeholder(2)), d.quote("key"), d.placeholder(3))

	tx := GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, queryUpdate, value, key)
	if err != nil {
		return errors.Wrap(handleError(err), "failed to update KV")
	}
	_, err = tx.ExecContext(ctx, queryInsert, key, value, key)
	if err != nil {
		return errors.Wrap(handleError(err), "failed to insert KV")
	}

	return nil
}

func (db *Service) retrieveValue(ctx context.Context, key string) (val string, err error) {
	d := db.Dialect
	query := fmt.Sprintf(
		"select %s from kv_store where %s = %s",
		d.quote("value"), d.quote("key"), d.placeholder(1))

	tx := GetTransaction(ctx)
	row := tx.QueryRowContext(ctx, query, key)
	if err = row.Scan(&val); err != nil {
		return "", handleError(err)
	}

	return val, nil
}

func (db *Service) retrieveValues(ctx context.Context, keys []string) (vals []string, err error) {
	d := db.Dialect

	var or []string
	for i := range keys {
		or = append(or, fmt.Sprintf("%s = %s", d.quote("key"), d.placeholder(i+1)))
	}
	query := fmt.Sprintf(
		"select %s from kv_store where %s", d.quote("value"), strings.Join(or, " or "))

	fmt.Println("retrieveValues query:\n", query)

	var keyInterfaces []interface{}
	for _, k := range keys {
		keyInterfaces = append(keyInterfaces, k)
	}

	tx := GetTransaction(ctx)
	rows, err := tx.QueryContext(ctx, query, keyInterfaces...)
	if err != nil {
		return nil, errors.Wrap(handleError(err), "failed to get values")
	}

	vals = []string{}
	for rows.Next() {
		var val string
		err = rows.Scan(&val)
		if err != nil {
			return nil, errors.Wrap(handleError(err), "failed to retrieve value")
		}
		vals = append(vals, val)
	}

	return vals, nil
}

func (db *Service) deleteKey(ctx context.Context, key string) error {
	d := db.Dialect
	query := fmt.Sprintf(
		"delete from kv_store where %s = %s",
		d.quote("key"), d.placeholder(1))

	tx := GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, query, key)
	if err != nil {
		return errors.Wrap(handleError(err), "failed to delete KV")
	}

	return nil
}

func (db *Service) retrieveKVPs(ctx context.Context) (kvps []*models.KeyValuePair, err error) {
	d := db.Dialect
	query := fmt.Sprintf(
		"select %s from kv_store",
		d.quoteSep("key", "value"))

	tx := GetTransaction(ctx)
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(handleError(err), "failed to get store")
	}

	kvps = []*models.KeyValuePair{}
	for rows.Next() {
		var kvp models.KeyValuePair
		err = rows.Scan(&kvp.Key, &kvp.Value)
		if err != nil {
			return nil, errors.Wrap(handleError(err), "failed to retrieve key value pair")
		}
		kvps = append(kvps, &kvp)
	}

	return kvps, nil
}
