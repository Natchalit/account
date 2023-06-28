package migration

import "github.com/Natchalit/account/src/services/migration/xmigrator"

var Migration = map[string]func(string){
	`account_test`: xmigrator.MAccountTest,
}
