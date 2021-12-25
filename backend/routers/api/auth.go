package api

import (
	"fmt"
	"net/http"

	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/app"
	"github.com/csDeng/blog/pkg/e"
	"github.com/csDeng/blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func Refresh(c *gin.Context) {
	appG := app.Gin{c}
	code := e.SUCCESS
	data := make(map[string]interface{})
	token := c.Param("token")
	maps, err := util.ParseToken(token)
	fmt.Println(maps, err)
	if err != nil {
		code = e.ERROR
	} else {
		fmt.Println("token=>", maps)
	}
	appG.Response(http.StatusOK, code, data)

}

// 登录
func GetAuth(c *gin.Context) {
	appG := app.Gin{c}

	json := &models.Auth{}
	c.BindJSON(json)

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	validate = validator.New()

	err := validate.Struct(json)

	if err == nil {
		u := json.Username
		p := json.Password
		isExist := models.CheckAuth(u, p)
		if isExist {
			token, err := util.GenerateToken(u, p)
			user := models.CheckAdmin(u, p)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				data["user"] = user
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		data["error"] = fmt.Sprintf("%v", err)
	}

	appG.Response(http.StatusOK, code, data)
}

// 注册
func AddAuth(c *gin.Context) {
	appG := app.Gin{c}

	json := &models.Auth{}
	c.BindJSON(json)

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	validate = validator.New()

	err := validate.Struct(json)

	if err == nil {
		u := json.Username
		p := json.Password
		isExist := models.UsernameIsExisted(u)
		if isExist {
			code = e.ERROR_AUTH_USERNAMEISEXISTED
		} else {
			if models.AddUser(u, p) {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		}
	} else {
		data["error"] = fmt.Sprintf("%v", err)
	}

	appG.Response(http.StatusOK, code, data)
}
