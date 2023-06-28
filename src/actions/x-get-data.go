package actions

import (
	"fmt"

	"github.com/Natchalit/account/src/database/pgy"
	"github.com/Natchalit/account/src/services/ginx"
	"github.com/Natchalit/account/src/services/migration/xtable"
)

func XGetData(c *ginx.GinX) (any, error) {

	px, ex := pgy.Account()
	if ex != nil {
		return nil, ex
	}

	query := fmt.Sprintf(`select * from %s`, xtable.Users{}.TableName())
	res, ex := px.QueryData(query)
	if ex != nil {
		return nil, ex
	}

	return res.Rows, nil
}
