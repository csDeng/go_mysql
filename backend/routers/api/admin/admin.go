package admin

import (
	"net/http"

	"github.com/csDeng/blog/pkg/app"
	"github.com/csDeng/blog/pkg/e"
	"github.com/csDeng/blog/pkg/logging"
	"github.com/gin-gonic/gin"
)

func ISADMIN(c *gin.Context) {
	logging.Info("admin======")
	appG := app.Gin{C: c}
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["msg"] = "I.m admin"
	appG.Response(http.StatusOK, code, data)
}
