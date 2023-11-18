package lib_db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type DB_PostgreSQL struct {
	db      map[int]*pgx.Conn
	connStr string
}

func NewPostgreSQL(cStr string) *DB_PostgreSQL {
	return &DB_PostgreSQL{
		connStr: cStr,
		db: map[int]*pgx.Conn{
			TxRead:  nil,
			TxWrite: nil,
		},
	}
}

func (d *DB_PostgreSQL) Open() error {
	if d.db[TxRead] == nil {
		rdb, err := pgx.Connect(context.Background(), d.connStr)
		if err != nil {
			return err
		}
		d.db[TxRead] = rdb
	}

	if d.db[TxWrite] == nil {
		wdb, err := pgx.Connect(context.Background(), d.connStr)
		if err != nil {
			return err
		}
		d.db[TxWrite] = wdb
	}

	return nil
}

func (d *DB_PostgreSQL) Close() {
	if d.db[TxRead] == nil {
		d.db[TxRead].Close(context.Background())
	}

	if d.db[TxWrite] == nil {
		d.db[TxWrite].Close(context.Background())
	}
}

func (d *DB_PostgreSQL) Exec(txType int, query string, args ...interface{}) (*string, error) {
	return nil, nil
}

func (d *DB_PostgreSQL) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*string, error) {
	return nil, nil
}

func (d *DB_PostgreSQL) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
	return nil, nil
}

func (d *DB_PostgreSQL) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return nil, nil
}
