package ginx

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (rx *RX) _Router(onWorkerX func(*GinX) (any, error)) []gin.HandlerFunc {
	handlerX := []gin.HandlerFunc{}

	handFunc := func(c *gin.Context) {
		hl := rx._Handler(c, onWorkerX)
		if hl.StatusCode == http.StatusNoContent {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			if hl.ResponseBody == nil {
				//c.AbortWithStatus(http.StatusNoContent)
			} else if vx, ok := hl.ResponseBody.(string); ok {
				if vs := strings.TrimSpace(vx); vs == `` || vs == `{}` {
					// c.AbortWithStatus(http.StatusNoContent)
				} else {
					if json.Valid([]byte(vx)) {
						c.JSON(hl.StatusCode, hl.ResponseBody)
					} else {
						c.String(hl.StatusCode, vx)
					}
				}
			} else {
				c.JSON(hl.StatusCode, hl.ResponseBody)
			}
		}
	}
	handlerX = append(handlerX, handFunc)

	return handlerX
}
