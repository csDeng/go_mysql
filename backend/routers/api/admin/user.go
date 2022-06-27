package admin

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

// 获取某一用户的信息
func Get(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	code := e.INVALID_ID
	if id < 1 {
		data["id"] = id
		appG.Response(http.StatusOK, code, data)
		return
	}
	user := models.GetUserById(id)

	// 用户不存在
	if data["user"] != nil {
		code = e.SUCCESS
		data["user"] = user
	} else {
		code = e.USER_NOT_FOUND
	}
	appG.Response(http.StatusOK, code, data)
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	code := e.SUCCESS

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	data["lists"] = models.GetUsers(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetUserTotal(maps)

	appG.Response(http.StatusOK, code, data)
}

// 添加用户
func AddUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.INVALID_PARAMS
	json := &models.User{}

	data := make(map[string]interface{})

	if err := c.ShouldBindJSON(json); err != nil {
		data["error"] = err.Error()
		appG.Response(http.StatusBadRequest, code, data)
		return
	}
	// 校验通过
	if models.ADDUserWithLevel(json) {
		code = e.SUCCESS
	} else {
		code = e.USER_ADD_FAIL
	}
	appG.Response(http.StatusOK, code, data)

}

// 修改用户信息
func EditUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_ID
	data := make(map[string]interface{})
	if id < 1 {
		data["id"] = id
		appG.Response(http.StatusOK, code, data)
		return
	}

	json := &models.User{}

	if err := c.ShouldBindJSON(json); err != nil {
		data["errors"] = err.Error()
		code = e.INVALID_PARAMS
		appG.Response(http.StatusBadRequest, code, data)
		return
	}

	if models.UserIsExistedBYId(id) {
		if models.EditUser(id, json) {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
		data["user"] = models.GetUserById(id)
	} else {
		code = e.ERROR_USER_USERISEISTEDBYID
	}

	appG.Response(http.StatusOK, code, data)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}

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
			code = e.ERROR_USER_USERISEISTEDBYID
		}
	}

	appG.Response(http.StatusOK, code, data)
}
