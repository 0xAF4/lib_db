package lib_db

import "time"

type DB_SQLite struct {
	connStr string
}

func NewSQLite(cStr string) *DB_SQLite {
	return &DB_SQLite{
		connStr: cStr,
	}
}

func (d *DB_SQLite) Open() error {
	return nil
}

func (d *DB_SQLite) Exec(query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_SQLite) ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_SQLite) QueryRow(query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"username": "0xAF4",
		"password": "pass123456789",
		"role":     "admin",
	}

	rec2 := map[string]interface{}{
		"username": "testUser",
		"password": "testPassword",
		"role":     "user",
	}

	result = append(result, rec1, rec2)
	return &result, nil
}

func (d *DB_SQLite) QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"username": "0xAF4",
		"password": "pass123456789",
		"role":     "admin",
	}

	rec2 := map[string]interface{}{
		"username": "testUser",
		"password": "testPassword",
		"role":     "user",
	}

	result = append(result, rec1, rec2)
	return &result, nil
}
