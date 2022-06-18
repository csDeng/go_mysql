package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ISADMIN(c *gin.Context) {
	fmt.Println("hha")
	c.JSON(http.StatusOK, gin.H{
		"msg": "I'm admin",
	})
}
