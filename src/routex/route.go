package routex

import (
	"github.com/Natchalit/account/src/actions"
	"github.com/Natchalit/account/src/services/ginx"
	"github.com/Natchalit/account/src/services/migration"
	"github.com/gin-gonic/gin"
)

var MODULE_S = ginx.RGX{
	MODULE: `finance`,
	PORT:   8080,
}

func Route() {
	MODULE_S.Connect(func(r *ginx.RX, rg *gin.RouterGroup) {
		rGruest := rg.Group(``)
		{
			r.GET(rGruest, `actions/get-data`, actions.XGetData)
			r.POST(rGruest, `migration`, migration.XAutoMigrate)
		}
	})
}
