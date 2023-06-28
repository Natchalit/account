package pgy

import (
	"database/sql"
	"fmt"

	"github.com/Natchalit/account/src/services/env"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dbName string) (*CusSql, error) {

	PG_HOST := env.GetEnv(`PG_HOST`)
	PG_PORT := env.GetEnv(`PG_PORT`)
	PG_USER := env.GetEnv(`PG_USER`)
	PG_PASS := env.GetEnv(`PG_PASS`)
	PG_NAME := env.GetEnv(`PG_NAME`)
	if dbName == `` {
		dbName = PG_NAME
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", PG_USER, PG_PASS, PG_HOST, PG_PORT, PG_NAME)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf(`connect db %s`, err.Error())
	}

	res := CusSql{}
	res.DB = *db

	return &res, err
}
