package admin

import (
	"net/http"

	"github.com/csDeng/blog/middleware/jwt"
	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/e"
	"github.com/gin-gonic/gin"
)

func ISADMIN() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.ERROR_TOKEN_ISNOTADMIN
		var data interface{}
		v := jwt.GetUserinfo(c)
		username, password := v.Username, v.Password
		if models.CheckAdmin(username, password) {
			c.Next()
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		c.Abort()
	}
}
