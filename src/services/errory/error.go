package errory

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorGinX(c *gin.Context, statusCode int, f string, a ...any) error {
	// create new error
	ex := ErrorX{
		Message: fmt.Sprintf(f, a...),
	}

	if c != nil {
		ex.StatusCode = c.Writer.Status()
		ex.ClientIP = c.ClientIP()
		ex.Method = c.Request.Method
		ex.URL = c.Request.Host + c.Request.URL.Path
	}
	if statusCode > 0 {
		ex.StatusCode = statusCode
	}
	return Error(&ex)
}

func Error(err error) error {
	ex, ok := err.(*ErrorX)
	if !ok {
		ex = &ErrorX{
			Message: ex.Error(),
		}
	}

	if ex.StatusCode == 0 {
		ex.StatusCode = http.StatusInternalServerError
	}

	return ex
}
