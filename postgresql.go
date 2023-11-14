package lib_db

import "time"

type DB_PostgreSQL struct {
	connStr string
}

func NewPostgreSQL(cStr string) *DB_PostgreSQL {
	return &DB_PostgreSQL{
		connStr: cStr,
	}
}

func (d *DB_PostgreSQL) Open() error {
	return nil
}

func (d *DB_PostgreSQL) Exec(query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_PostgreSQL) ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_PostgreSQL) QueryRow(query string, args ...interface{}) (*DBResult, error) {
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

func (d *DB_PostgreSQL) QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
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
