package ginx

import (
	"net/http"

	"github.com/Natchalit/account/src/services/errory"
	"github.com/Natchalit/account/src/services/helper"
	"github.com/gin-gonic/gin"
)

type HandlerRT struct {
	StatusCode     int
	ResponseHeader map[string]string
	ResponseBody   any
}

func (r *RX) _Handler(c *gin.Context, onWorkerX func(*GinX) (any, error)) (hl *HandlerRT) {

	hl = &HandlerRT{
		ResponseHeader: map[string]string{},
	}
	// context
	ctx := &GinX{c: c}

	// worker
	rx, ex := onWorkerX(ctx)
	if ex != nil {
		// if ex.StatusCode == 0 {
		// 	ex.StatusCode = http.StatusInternalServerError
		// }
		// if ex.Message == `` {
		// 	ex.Message = `Unknown`
		// }
		// respError = ex.Message
		// hl.StatusCode = ex.StatusCode

		isNilResponse := helper.IsNil(rx)
		if !isNilResponse {
			hl.ResponseBody = rx
		}
		if v, ok := ex.(*errory.ErrorX); ok {
			if v != nil {
				if isNilResponse {
					hl.ResponseBody = v
				}
				if v.StatusCode == 0 {
					hl.StatusCode = http.StatusInternalServerError
				} else {
					hl.StatusCode = v.StatusCode
				}
			}
		} else {
			if isNilResponse {
				hl.ResponseBody = errory.ErrorGinX(c, http.StatusInternalServerError, ex.Error())
			}
			if hl.StatusCode == 0 {
				hl.StatusCode = http.StatusInternalServerError
			} else {
				hl.StatusCode = v.StatusCode
			}
		}
		// if rx != nil {
		// 	hl.ResponseBody = rx
		// } else {
		// 	hl.ResponseBody = ex
		// }
	} else if rx != nil {
		if rc, ok := rx.(*_RC); ok {
			hl.StatusCode = rc.StatusCode
			hl.ResponseBody = rc.ResponseBody
		} else {
			hl.StatusCode = http.StatusOK
			hl.ResponseBody = rx
		}
	}

	return hl
}

type _RC struct {
	StatusCode   int
	ResponseBody any
}
