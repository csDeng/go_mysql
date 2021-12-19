package v1

import (
	"fmt"
	"net/http"

	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/app"
	"github.com/csDeng/blog/pkg/e"
	"github.com/csDeng/blog/pkg/logging"
	"github.com/csDeng/blog/pkg/setting"
	"github.com/csDeng/blog/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/unknwon/com"
)

// 这里的文章类型，用于接收绑定的json数据，与models的不一样哦
type Article struct {
	TagID         int    `json:"tag_id" validate:"min=1" `
	Title         string `json:"title" validate:"required,max=100"`
	Desc          string `json:"desc" validate:"max=255"`
	Content       string `json:"content" validate:"required"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state" validate:"min=0,max=1"`
	CoverImageUrl string `json:"cover_image_url"`
}

// 获取首页的文章
func GetArticleIndex(c *gin.Context) {
	appG := app.Gin{c}
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["lists"] = models.GetArticleIndex()
	appG.Response(http.StatusOK, code, data)
}

//获取单个文章
// params : id
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	code := 0
	var data interface{}
	if id < 1 {
		code = e.INVALID_PARAMS
		data = "id必须大于0"
		logging.Info(code, data)
	}

	if code == 0 {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
			data = fmt.Sprintf("%d 文章不存在", id)
		}
	}
	appG.Response(http.StatusOK, code, data)
}

//获取多个文章
/*
query:
state:
tag_id
*/
func GetArticles(c *gin.Context) {
	appG := app.Gin{c}
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		if state < 0 || state > 1 {
			data["state"] = "state 只能是0或1哦"
		}
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		if tagId < 1 {
			data["tag_id"] = "标签ID必须大于0"
		}
	}
	code := e.INVALID_PARAMS
	if len(data) == 0 {
		code = e.SUCCESS
		offset := util.GetPage(c)
		data["total"] = models.GetArticleTotal(maps)
		data["lists"] = models.GetArticles(offset, setting.AppSetting.PageSize, maps)
	} else {
		code = e.INVALID_PARAMS
	}
	appG.Response(http.StatusOK, code, data)
}

//新增文章
func AddArticle(c *gin.Context) {
	appG := app.Gin{c}
	json := &Article{}
	c.BindJSON(json)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	validate = validator.New()
	err := validate.Struct(json)

	if err != nil {
		data["error"] = fmt.Sprint(err)
	} else {
		data["tag_id"] = json.TagID
		data["title"] = json.Title
		data["desc"] = json.Desc
		data["content"] = json.Content
		data["created_by"] = json.CreatedBy
		data["state"] = json.State
		data["cover_image_url"] = json.CoverImageUrl
		fmt.Println("添加文章=>", data)
		if models.AddArticle(data) {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	}

	appG.Response(http.StatusOK, code, data)
}

//修改文章
func EditArticle(c *gin.Context) {
	appG := app.Gin{c}

	validate = validator.New()
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	id := com.StrTo(c.Param("id")).MustInt()
	if id < 1 {
		appG.Response(http.StatusBadRequest, code, data)
	}
	json := &Article{}
	c.BindJSON(json)
	err := validate.Struct(json)
	if err == nil {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(json.TagID) {
				data["tag_id"] = json.TagID
				if json.Title != "" {
					data["title"] = json.Title
				}
				if json.Desc != "" {
					data["desc"] = json.Desc
				}
				if json.Content != "" {
					data["content"] = json.Content
				}
				if json.CoverImageUrl != "" {
					data["cover_image_url"] = json.CoverImageUrl
				}
				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		data["error"] = fmt.Sprint(err)
	}
	appG.Response(http.StatusOK, code, data)

}

//删除文章
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	data := make(map[string]interface{})
	code := e.INVALID_ID
	if id > 0 {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}

	}
	appG.Response(http.StatusOK, code, data)
}
