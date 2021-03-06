package routers

import (
	"net/http"

	"github.com/csDeng/blog/middleware/jwt"
	"github.com/csDeng/blog/pkg/setting"
	"github.com/csDeng/blog/pkg/upload"
	"github.com/csDeng/blog/routers/api"
	v1 "github.com/csDeng/blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 开放静态资源
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.POST("/login", api.GetAuth)
	r.POST("/register", api.AddAuth)
	r.GET("refresh/:token", api.Refresh)

	r.POST("/upload", api.UploadImage)

	r.GET("/index", v1.GetArticleIndex)
	apiv1 := r.Group("/api/v1")
	//获取文章列表
	apiv1.GET("/articles", v1.GetArticles)
	//获取标签列表
	apiv1.GET("/tags", v1.GetTags)
	//获取指定文章
	apiv1.GET("/articles/:id", v1.GetArticle)
	// 获取用户列表
	apiv1.GET("/users", v1.Getusers)
	apiv1.Use(jwt.JWT())
	{

		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		// 获取指定用户

		// 更新指定用户
		apiv1.PUT("/users/:id", v1.EditUser)
		// 删除指定用户
		apiv1.DELETE("/users/:id", v1.DeleteUser)
	}

	return r
}
