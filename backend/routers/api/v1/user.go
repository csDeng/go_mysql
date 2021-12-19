package v1

import (
	"net/http"

	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/app"
	"github.com/csDeng/blog/pkg/e"
	"github.com/csDeng/blog/pkg/setting"
	"github.com/csDeng/blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Getusers(c *gin.Context) {
	appG := app.Gin{c}

	code := e.SUCCESS

	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	data["lists"] = models.GetUsers(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetUserTotal(maps)

	appG.Response(http.StatusOK, code, data)
}

func EditUser(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_ID
	data := make(map[string]interface{})

	if id < 1 {
		appG.Response(http.StatusOK, code, data)
	}

	json := &models.Auth{}

	c.BindJSON(json)

	if models.UserIsExistedBYId(id) {
		code = e.SUCCESS
		models.EditUser(id, json)
		data["user"] = models.GetUserById(id)
	} else {
		code = e.ERROR_AUTH_USERISEISTEDBYID
	}

	appG.Response(http.StatusOK, code, data)
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_ID

	data := make(map[string]interface{})
	if id > 1 {
		if models.UserIsExistedBYId(id) {
			if models.DeleteUser(id) {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		} else {
			code = e.ERROR_AUTH_USERISEISTEDBYID
		}

	}

	appG.Response(http.StatusOK, code, data)
}
