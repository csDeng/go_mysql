package v1

import (
	"fmt"
	"net/http"

	"github.com/csDeng/go-gin-example/models"
	"github.com/csDeng/go-gin-example/pkg/e"
	"github.com/csDeng/go-gin-example/pkg/setting"
	"github.com/csDeng/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/unknwon/com"
)

type Tag struct {
	Name string `json:"name" validate:"required"`
	// State     int    `json:"state" validate:"required,min=0,max=1"`
	// tag不能有空格
	// 0不符合required
	State      int    `json:"state" validate:"min=0,max=1"`
	CreatedBy  string `json:"created_by" `
	ModifiedBy string `json:"modified_by"`
}

var validate *validator.Validate

// 获取所有标签，或者根据query name进行查询
func GetTags(c *gin.Context) {

	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	// 配置参数校验
	validate = validator.New()
	json := &Tag{}
	c.BindJSON(json)

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	// fmt.Println("tag json=>", json)
	// fmt.Printf("%v, %T\n", json, json)
	// fmt.Println(json.Name, json.State, json.CreatedBy)
	err := validate.Struct(json)

	if err != nil {
		fmt.Println("err=>", err)
		data["error"] = fmt.Sprint(err)
	} else {
		if !models.ExistTagByName(json.Name) {
			code = e.SUCCESS
			models.AddTag(json.Name, int(json.State), json.CreatedBy)
			data["success"] = fmt.Sprintf("%s created successfully", json.Name)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	validate = validator.New()
	json := &Tag{}
	c.BindJSON(json)
	err := validate.Struct(json)
	code := e.INVALID_PARAMS

	if err == nil {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = json.ModifiedBy
			if json.Name != "" {
				data["name"] = json.Name
			}
			if json.State != -1 {
				data["state"] = json.State
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var code int
	var data = make(map[string]string)
	if id < 1 {
		code = e.INVALID_PARAMS
	} else {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			// 判断tag存在再删除
			models.DeleteTag(id)
			data["success"] = fmt.Sprintf("%d tag deleted successfully", id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
