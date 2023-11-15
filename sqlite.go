package lib_db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB_SQLite struct {
	db      map[int]*sql.DB
	connStr string
}

func NewSQLite(cStr string) *DB_SQLite {
	return &DB_SQLite{
		connStr: cStr,
		db: map[int]*sql.DB{
			TxRead:  nil,
			TxWrite: nil,
		},
	}
}

func (d *DB_SQLite) Open() error {
	rdb, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", d.connStr))
	if err != nil {
		return err
	}
	d.db[TxRead] = rdb
	defer rdb.Close()

	wdb, err := sql.Open("sqlite3", d.connStr)
	if err != nil {
		return err
	}
	d.db[TxWrite] = wdb
	defer wdb.Close()

	return nil
}

func (d *DB_SQLite) Exec(txType int, query string, args ...interface{}) (*DBResult, error) {
	return nil, nil
}

func (d *DB_SQLite) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return nil, nil
}

func (d *DB_SQLite) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
	tx, err := d.db[txType].Begin()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows
	rows, err = tx.Query(query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var username, email string
		err := rows.Scan(&id, &username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
	}

	return nil, nil
}

func (d *DB_SQLite) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return nil, nil
}
