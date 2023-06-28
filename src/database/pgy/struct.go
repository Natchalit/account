package pgy

import "database/sql"

type CusSql struct {
	DB     sql.DB
	Driver string
}
