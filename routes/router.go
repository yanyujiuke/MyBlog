package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/admin/v1"
	"myblog/utils"
	"net/http"
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

		// 分类模块
		auth.GET("category/list", v1.GetCate)
		auth.POST("category/add", v1.CreateCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		// 文章模块
		auth.GET("article/info/:id", v1.GetArtInfo)
		auth.GET("article/list", v1.GetArts)
		auth.POST("article/add", v1.CreateArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 个人信息模块
		auth.GET("profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)

		auth.GET("test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"a": 1,
			})
		})
	}

	_ = r.Run(utils.HttpPort)
}
