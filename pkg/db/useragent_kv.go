package db

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

func (db *Service) StoreKV(ctx context.Context, key string, value string) error {
	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		return db.storeKV(ctx, key, value)
	}); err != nil {
		return err
	}

	return nil
}

func (db *Service) RetrieveValue(ctx context.Context, key string) (val string, err error) {
	if err := db.DoInTransaction(ctx, func(ctx context.Context) error {
		val, err = db.retrieveValue(ctx, key)
		return err
	}); err != nil {
		return "", err
	}

	return val, nil
}

func (db *Service) RetrieveValues(ctx context.Context, keys []string) (val []string, err error) {
	return []string{}, nil
}

func (db *Service) RetrieveKVPs(ctx context.Context) (kvps []models.KeyValuePair, err error) {
	return []models.KeyValuePair{}, nil
}

func (db *Service) storeKV(ctx context.Context, key string, value string) error {
	d := db.Dialect
	query := fmt.Sprintf(
		"insert into kv_store (%s) values (%s, %s)",
		d.quoteSep("key", "value"), d.placeholder(1), d.placeholder(2))

	tx := GetTransaction(ctx)
	_, err := tx.ExecContext(ctx, query, key, value)
	if err != nil {
		err = handleError(err)
		return errors.Wrap(err, "failed to store KV")
	}

	return nil
}

func (db *Service) retrieveValue(ctx context.Context, key string) (val string, err error) {
	d := db.Dialect
	query := fmt.Sprintf(
		"select %s from kv_store where %s = %s",
		d.quote("value"), d.quote("key"), d.placeholder(1))

	row := tx.QueryRowContext(ctx, query, key)
	if err = row.Scan(&val); err != nil {
		return "", handleError(err)
	}

	return val, nil
}
