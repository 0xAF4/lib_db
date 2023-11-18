package lib_db

import "time"

type DB_Mock struct{}

func NewMock() *DB_Mock {
	return &DB_Mock{}
}

func (d *DB_Mock) Open() error {
	return nil
}

func (d *DB_Mock) Close() {
	//Noop Close()
}

func (d *DB_Mock) Exec(txType int, query string, args ...interface{}) (*string, error) {
	s := "LastInsertId: 1; RowsAffected: 1;"
	return &s, nil
}

func (d *DB_Mock) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*string, error) {
	return d.Exec(txType, query, args...)
}

func (d *DB_Mock) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
	return &DBResult{
		map[string]interface{}{
			"username": "0xAF4",
			"password": "pass123456789",
			"role":     "admin",
		},
		map[string]interface{}{
			"username": "testUser",
			"password": "testPassword",
			"role":     "user",
		},
	}, nil
}

func (d *DB_Mock) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return d.QueryRow(txType, query, args...)
}
