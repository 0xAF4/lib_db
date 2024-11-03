package lib_db

import (
	fb "lib_db/fb"
	"time"

	_ "github.com/nakagami/firebirdsql"
)

type DB_FireBird struct {
	db   *fb.Database
	conn *fb.Connection
}

func NewFireBird(cStr string) *DB_FireBird {
	db, _ := fb.New("database=" + cStr)
	return &DB_FireBird{
		db: db,
	}
}

func (d *DB_FireBird) Open() error {
	conn, err := d.db.Connect()
	if err != nil {
		return err
	}
	d.conn = conn
	return nil
}

func (d *DB_FireBird) Close() {
	d.conn.Close()
}
func (d *DB_FireBird) Exec(txType int, query string, args ...interface{}) (*string, error) {
	_, err := d.conn.Execute(query, args...)
	return nil, err
}

func (d *DB_FireBird) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*string, error) {
	return d.Exec(txType, query, args...)
}

func (d *DB_FireBird) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
	row, err := d.conn.QueryRowMaps(query)
	return (*DBResult)(&row), err
}

func (d *DB_FireBird) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return d.QueryRow(txType, query, args...)
}

func (d *DB_FireBird) StartTx(txType int) (interface{}, error) {
	return nil, nil
}
