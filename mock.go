package lib_db

import "time"

type DB_Mock struct{}

func NewMock() *DB_Mock {
	return &DB_Mock{}
}

func (d *DB_Mock) Open() error {
	return nil
}

func (d *DB_Mock) Exec(query string, args ...interface{}) (*DBResult, error) {
	return &DBResult{
		map[string]interface{}{
			"exec_result": true,
		},
	}, nil
}

func (d *DB_Mock) ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return d.Exec(query, args...)
}

func (d *DB_Mock) QueryRow(query string, args ...interface{}) (*DBResult, error) {
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

func (d *DB_Mock) QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	return d.QueryRow(query, args...)
}
