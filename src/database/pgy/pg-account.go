package pgy

func Account() (*CusSql, error) {
	return ConnectPostgres(`account_test`)
}
