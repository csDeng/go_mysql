package routers

import (
	"github.com/csDeng/blog/middleware/admin"
	"github.com/csDeng/blog/middleware/jwt"
	adminHandler "github.com/csDeng/blog/routers/api/admin"
)

func adminRoutes() {
	root := r.Group("/api/admin")
	root.Use(jwt.JWT(), admin.ISADMIN())
	{
		root.GET("/", adminHandler.ISADMIN)

		// 用户管理
		u := root.Group("users")
		{
			// 获取用户列表
			u.GET("/", adminHandler.Getusers)

			// 获取指定用户

			// 更新指定用户
			u.PUT("/:id", adminHandler.EditUser)
			// 删除指定用户
			u.DELETE("/:id", adminHandler.DeleteUser)
		}
	}
}
