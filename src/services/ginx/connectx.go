package ginx

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (rgx *RGX) Connect(routerx func(*RX, *gin.RouterGroup)) {
	rx := RX{}
	// c := gin.Default()
	gin.ForceConsoleColor()
	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))

	router.Use(gin.Recovery())

	rx.Router = router
	rx.SRV = rgx.MODULE

	routerx(&rx, router.Group(rgx.MODULE))

	// rg := router.Group(`/`)
	// {
	// 	rg.POST(`/migrate`, migration.XAutoMigrate)
	// }
	// gr := router.Group(`/finance`)
	// {
	// 	gr.GET(`/currency`, actions.XGetCurrency)
	// 	gr.GET(`/data`, actions.XGetData)

	// }
	_ = router.Run(fmt.Sprintf(":%v", rgx.PORT))
}
