package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func New(addr string, maxOpenConn, maxIdleConn int, maxConnLifetime time.Duration) (*sql.DB, error) {
	db, err = sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
