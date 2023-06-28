package xmigrator

import "github.com/Natchalit/account/src/services/migration/xtable"

func MAccountTest(dbname string) {
	AutoMigrate(dbname, &xtable.Users{})
}
