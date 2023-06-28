package xmigrator

import (
	"log"

	"github.com/Natchalit/account/src/services/logx"
	"github.com/Natchalit/account/src/services/xgorm"
)

func AutoMigrate(dbName string, dst ...interface{}) {
	db, ex := xgorm.ConnectGorm(dbName)
	if ex != nil {
		log.Fatalf(ex.Error())
	}

	if ex = db.AutoMigrate(dst...); ex != nil {
		logx.Errorf("AutoMigration [%v], %v", dbName, ex)
	} else {
		logx.Infof("AutoMigration success. [%v]", dbName)
	}
}
