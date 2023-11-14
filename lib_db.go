package libdb

import (

)

type DBConfig map[string]interface{}

type DBInterface interface {
	
}


type DB struct {
	DBIntr DBInterface
}

const (
	MSSQL = iota
	PostgreSQL
	SQLite
)

func New(cfg *DBConfig) *DB {
	return &DB{
		
	}
}