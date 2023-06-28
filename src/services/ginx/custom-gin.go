package ginx

import (
	"net/http"

	"github.com/Natchalit/account/src/services/errory"
	"github.com/gin-gonic/gin"
)

type GinX struct {
	c *gin.Context
}

func (cx *GinX) Context() *gin.Context {
	if cx == nil {
		return nil
	}
	return cx.c
}

func (cx *GinX) Error(err error) error {
	if err == nil {
		return nil
	}

	if ex, ok := err.(*errory.ErrorX); ok {
		return ex
	}

	return errory.ErrorGinX(cx.Context(), http.StatusInternalServerError, err.Error())
}

func (cx *GinX) ShouldBindJSON(obj any) error {
	return cx.Context().ShouldBindJSON(obj)
}
