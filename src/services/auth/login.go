package auth

import (
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

func Login(c *gin.Context) {
	// login := LoginDto{}
	// if ex := c.ShouldBindJSON(&login); ex != nil {
	// 	c.JSON(http.StatusBadRequest, ex.Error())
	// }

	// username := login.Username
	// pass := login.Password
	// data := make([]byte, base64.StdEncoding.EncodedLen(len(pass)))
	// password, ex := base64.StdEncoding.Decode(data, pass)
	// if ex != nil {
	// 	c.JSON(http.StatusInternalServerError, ex.Error())
	// }

	// px, ex := xpg.Account()
	// if ex != nil {
	// 	c.JSON(http.StatusInternalServerError, ex.Error())
	// }

	// c.JSON(http.StatusOK, rows)

}
