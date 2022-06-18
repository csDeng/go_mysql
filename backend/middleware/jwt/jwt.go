package jwt

import (
	"net/http"
	"time"

	"github.com/csDeng/blog/pkg/e"
	"github.com/csDeng/blog/pkg/util"
	"github.com/gin-gonic/gin"
)

var claims *util.Claims
var err error

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			code = e.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				code = e.ERROR_USER_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_USER_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		SetUserInfo(c)
		c.Next()
	}
}

func SetUserInfo(c *gin.Context) {
	c.Set("userinfo", claims)
}

func GetUserinfo(c *gin.Context) *util.Claims {
	return c.MustGet("userinfo").(*util.Claims)
}
