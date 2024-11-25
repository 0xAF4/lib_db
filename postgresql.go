package lib_db

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB_PostgreSQL struct {
	db      *pgxpool.Pool
	connStr string
}

var mu sync.Mutex

func NewPostgreSQL(cStr string) *DB_PostgreSQL {
	return &DB_PostgreSQL{
		connStr: cStr,
		db:      nil,
	}
}

func (d *DB_PostgreSQL) Open() error {
	if d.db == nil {
		rdb, err := pgxpool.New(context.Background(), d.connStr)
		if err != nil {
			return err
		}
		d.db = rdb
	}

	return nil
}

func (d *DB_PostgreSQL) Close() {
	if d.db != nil {
		d.db.Close()
		d.db = nil
	}
}

func (d *DB_PostgreSQL) StartTx(txType int) (interface{}, error) {
	if err := d.ensureConnection(); err != nil {
		return nil, err
	}
	return d.db.Begin(context.Background())
}

func (d *DB_PostgreSQL) Exec(txType int, query string, args ...interface{}) (*string, error) {
	if err := d.ensureConnection(); err != nil {
		return nil, err
	}

	coma, err := d.db.Exec(context.Background(), query, args...)
	str := coma.String()
	return &str, err
}

func (d *DB_PostgreSQL) ExecWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*string, error) {
	if err := d.ensureConnection(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctxTime, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	tx, err := d.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var rows pgconn.CommandTag
	done := make(chan bool)

	go func() {
		rows, err = tx.Exec(ctxTime, query, args...)
		done <- true
	}()

	select {
	case <-ctxTime.Done():
		tx.Rollback(ctx)
		return nil, ctxTime.Err()
	case <-done:
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		} else {
			res := rows.String()
			if err != nil {
				tx.Rollback(ctx)
				return nil, err
			}
			tx.Commit(ctx)
			return &res, nil
		}
	}
}

func (d *DB_PostgreSQL) QueryRow(txType int, query string, args ...interface{}) (*DBResult, error) {
	// Проверяем соединение и восстанавливаем его при необходимости
	if err := d.ensureConnection(); err != nil {
		return nil, err
	}

	rows, err := d.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	return d.rowsToMap(rows)
}

func (d *DB_PostgreSQL) QueryRowWithTimeout(txType int, timeOut time.Duration, query string, args ...interface{}) (*DBResult, error) {
	// Проверяем соединение и восстанавливаем его при необходимости
	if err := d.ensureConnection(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctxTime, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	tx, err := d.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var rows pgx.Rows
	done := make(chan bool)

	go func() {
		rows, err = tx.Query(ctxTime, query, args...)
		done <- true
	}()

	select {
	case <-ctxTime.Done():
		rows.Close()
		tx.Rollback(ctx)
		return nil, ctxTime.Err()
	case <-done:
		if err != nil {
			rows.Close()
			tx.Rollback(ctx)
			return nil, err
		} else {
			res, err := d.rowsToMap(rows)
			defer rows.Close()
			if err != nil {
				tx.Rollback(ctx)
				return nil, err
			}
			tx.Commit(ctx)
			return res, nil
		}
	}
}

func (d *DB_PostgreSQL) rowsToMap(rows pgx.Rows) (*DBResult, error) {
	columns := rows.FieldDescriptions()
	var result DBResult

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, column := range columns {
			rowMap[string(column.Name)] = *(values[i].(*interface{}))
		}

		result = append(result, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}

// Новый метод для проверки соединения и реконнекта
func (d *DB_PostgreSQL) ensureConnection() error {
	mu.Lock()
	defer mu.Unlock()
	if d.db != nil {
		// Проверяем текущее состояние соединения
		if err := d.db.Ping(context.Background()); err == nil {
			return nil
		}
		// Закрываем старое невалидное соединение
		d.db.Close()
		d.db = nil
	}
	// Пытаемся установить новое соединение
	return d.Open()
}
