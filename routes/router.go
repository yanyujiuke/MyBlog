package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "myblog/admin/v1"
	"myblog/middleware"
	"myblog/utils"
	"net/http"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

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

		// 评论模块
		auth.GET("comment/list", v1.GetComments)
		auth.DELETE("comment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)

		// 上传
		auth.POST("upload", v1.Upload)

		auth.GET("test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"a": 1,
			})
		})
	}

	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/find/:id", v1.FindUser)
		router.GET("users", v1.GetUsers)

		// 文章分类
		router.GET("cate/get", v1.GetCate)
		router.GET("cate/find/:id", v1.FindCate)

		// 文章模块
		router.GET("art/get", v1.GetArts)           // 获取文章列表
		router.GET("art/get/:cid", v1.GetArtsByCid) // 获取耨个分类下的所有文章
		router.GET("art/find/:id", v1.GetArtInfo)   // 获取耨个分类下的所有文章

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router.POST("comment/add", v1.AddComment)
		router.GET("comment/find/:id", v1.FindComment)
		router.GET("comment/get/:art_id", v1.GetCommentByArtId)
		router.GET("comment/count/:art_id", v1.GetCommentCount)

	}

	_ = r.Run(utils.HttpPort)
}
