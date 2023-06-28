package xgorm

import (
	"fmt"

	"github.com/Natchalit/account/src/services/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm(dbName string) (*gorm.DB, error) {
	PG_HOST := env.GetEnv(`PG_HOST`)
	// PG_PORT := env.GetEnv(`PG_PORT`)
	PG_USER := env.GetEnv(`PG_USER`)
	PG_PASS := env.GetEnv(`PG_PASS`)

	dsn := fmt.Sprintf(
		`postgres://%s:%s@%s/%s`,
		PG_USER, PG_PASS, PG_HOST, dbName,
	)

	// dsn = `postgres://natchalit:zZ25vnagqyBXnEmBkPRzwwuOiLAEGy7k@dpg-cg5d08ndvk4pls4ds66g-a.singapore-postgres.render.com/account_test`

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// defer func() {
	// 	dbInstance, _ := db.DB()
	// 	_ = dbInstance.Close()
	// }()
	return db, nil
}
