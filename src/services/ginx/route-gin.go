package ginx

import "github.com/gin-gonic/gin"

type RGX struct {
	PORT   int
	MODULE string
}

type RX struct {
	SRV    string
	Router *gin.Engine
}

func (rx *RX) GET(g *gin.RouterGroup, relativePath string, onWorkerX func(*GinX) (any, error), handlers ...gin.HandlerFunc) {
	path := ``

	handlerx := []gin.HandlerFunc{}
	handlerx = append(handlerx, rx._Router(onWorkerX)...)
	handlerx = append(handlerx, Guard(path))
	handlerx = append(handlerx, handlers...)
	g.GET(path, handlerx...)
}

func (rx *RX) POST(g *gin.RouterGroup, relativePath string, onWorkerX func(*GinX) (any, error), handlers ...gin.HandlerFunc) {
	path := ``

	handlerx := []gin.HandlerFunc{}
	handlerx = append(handlerx, rx._Router(onWorkerX)...)
	handlerx = append(handlerx, Guard(path))
	handlerx = append(handlerx, handlers...)
	g.POST(path, handlerx...)
}
