package routers

import (
	"net/http"

	"github.com/csDeng/blog/pkg/upload"
	"github.com/csDeng/blog/routers/api"
	v1 "github.com/csDeng/blog/routers/api/v1"
)

func globalRoutes() {
	// 开放静态资源
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.POST("/login", api.GetUser)
	r.POST("/register", api.AddUser)
	r.GET("refresh/:token", api.Refresh)

	r.POST("/upload", api.UploadImage)

	r.GET("/index", v1.GetArticleIndex)

}
