package lib_db

import "time"

type DB_Mock struct {
	connStr string
}

func NewMock(cStr string) *DB_Mock {
	return &DB_Mock{
		connStr: cStr,
	}
}

func (d *DB_Mock) Open() error {
	return nil
}

func (d *DB_Mock) Exec(query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_Mock) ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	var result DBResult
	rec1 := map[string]interface{}{
		"exec_result": true,
	}

	result = append(result, rec1)
	return &result, nil
}

func (d *DB_Mock) QueryRow(query string, args ...interface{}) (*DBResult, error) {
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

func (d *DB_Mock) QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
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
