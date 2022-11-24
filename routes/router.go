package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/admin/v1"
	"myblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	// 后台管理路由接口
	auth := r.Group("admin/v1")
	{
		// 用户模块
		auth.GET("users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.PUT("changepw/:id", v1.ChangePassword)

	}

	_ = r.Run(utils.HttpPort)
}
