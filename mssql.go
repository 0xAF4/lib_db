package lib_db

import "time"

type DB_MSSQL struct {
	connStr string
}

func NewMSSQL(cStr string) *DB_MSSQL {
	return &DB_MSSQL{
		connStr: cStr,
	}
}

func (d *DB_MSSQL) Open() error {
	return nil
}

func (d *DB_MSSQL) Exec(query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_MSSQL) ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_MSSQL) QueryRow(query string, args ...interface{}) (*DBResult, error) {
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

func (d *DB_MSSQL) QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
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
