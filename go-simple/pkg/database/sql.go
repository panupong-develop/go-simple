package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// Resources:
// https://hackernoon.com/comparing-optimistic-and-pessimistic-locking-with-go-and-postgresql

type IDatabaseConfig interface {
	GetDatabaseURL() string
	GetMaxOpenConnections() int
}

func ConnectSQLDatabase(dbConf IDatabaseConfig) (*sqlx.DB, error) {
	const driverName = "pgx"
	db, err := sqlx.Connect(driverName, dbConf.GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	db.DB.SetMaxOpenConns(dbConf.GetMaxOpenConnections())
	return db, nil
}
