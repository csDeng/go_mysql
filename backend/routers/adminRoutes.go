package routers

import (
	"github.com/csDeng/blog/middleware/admin"
	"github.com/csDeng/blog/middleware/jwt"
	adminController "github.com/csDeng/blog/routers/api/admin"
)

func adminRoutes() {
	root := r.Group("/api/admin")
	root.Use(jwt.JWT(), admin.ISADMIN())
	{
		root.GET("/a", adminController.ISADMIN)

		// 用户管理
		u := root.Group("/users")
		{
			// 增加用户
			u.POST("/", adminController.AddUsers)

			// 获取用户列表
			u.GET("/", adminController.GetUsers)

			// 获取指定用户
			u.GET("/:id", adminController.Get)
			// 更新指定用户
			u.PUT("/:id", adminController.EditUser)
			// 删除指定用户
			u.DELETE("/:id", adminController.DeleteUser)
		}
	}
}
