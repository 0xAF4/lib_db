package lib_db

import (
	"fmt"
	"time"
)

type DBConfig map[string]interface{}
type DBResult []map[string]interface{}

type DBInterface interface {
	Open() error
	Exec(query string, args ...interface{}) (*DBResult, error)
	ExecWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error)
	QueryRow(query string, args ...interface{}) (*DBResult, error)
	QueryRowWithTimeout(timeOut time.Duration, query string, args ...interface{}) (*DBResult, error)
}

type DB struct {
	DBIntr DBInterface
}

const (
	MSSQL = iota
	PostgreSQL
	SQLite
	Mock
)

func New(cfg *DBConfig) (*DB, error) {
	var dbIntr DBInterface

	driver, ok := (*cfg)["driver"].(int)
	if !ok {
		return nil, fmt.Errorf("Укажите driver")
	}
	connectionString, ok := (*cfg)["connectionString"].(string)
	if !ok {
		return nil, fmt.Errorf("Укажите строку подключения connectionString")
	}

	switch driver {
	case MSSQL:
		dbIntr = NewMSSQL(connectionString)
	case PostgreSQL:
		dbIntr = NewPostgreSQL(connectionString)
	case SQLite:
		dbIntr = NewSQLite(connectionString)		
	case Mock:
		dbIntr = NewMock(connectionString)
	default:
		return nil, fmt.Errorf("Укажите driver, тип БД")
	}

	if err := dbIntr.Open(); err != nil {
		return nil, err
	}

	return &DB{
		DBIntr: dbIntr,
	}, nil
}
