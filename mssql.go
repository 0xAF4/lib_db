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

func (d *DB_MSSQL) Close() {
	//Noop Close()
}

func (d *DB_MSSQL) Exec(txType int, query string, args ...interface{}) (*string, error) {
	s := "LastInsertId: 1; RowsAffected: 1;"
	return &s, nil
}

func (d *DB_MSSQL) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*string, error) {
	s := "LastInsertId: 1; RowsAffected: 1;"
	return &s, nil
}

func (d *DB_MSSQL) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
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

func (d *DB_MSSQL) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
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
