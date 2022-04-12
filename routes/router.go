package routes

import (
	v1 "go-logistics/api/v1"
	"go-logistics/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-logistics/docs"
)

func InitRouter() {
	//gin.SetMode(viper.GetString("server.AppMode"))
	r := gin.Default()

	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.Use(middleware.Log())
	//r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("users", v1.GetUsers)
		auth.GET("user/:id", v1.GetUserInfo)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//修改密码
		auth.PUT("admin/password/:id", v1.ChangeUserPassword)
		//// 分类模块的路由接口
		//auth.GET("admin/category", v1.GetCate)
		//auth.POST("category/add", v1.AddCategory)
		//auth.PUT("category/:id", v1.EditCate)
		//auth.DELETE("category/:id", v1.DeleteCate)
		//// 文章模块的路由接口
		//auth.GET("admin/article/info/:id", v1.GetArtInfo)
		//auth.GET("admin/article", v1.GetArt)
		//auth.POST("article/add", v1.AddArticle)
		//auth.PUT("article/:id", v1.EditArt)
		//auth.DELETE("article/:id", v1.DeleteArt)
		//// 上传文件
		//auth.POST("upload", v1.UpLoad)
		//// 更新个人设置
		//auth.GET("admin/profile/:id", v1.GetProfile)
		//auth.PUT("profile/:id", v1.UpdateProfile)
		//// 评论模块
		//auth.GET("comment/list", v1.GetCommentList)
		//auth.DELETE("delcomment/:id", v1.DeleteComment)
		//auth.PUT("checkcomment/:id", v1.CheckComment)
		//auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	router := r.Group("api/v1")
	{

		// 新增用户
		router.POST("user/add", v1.AddUser)

		// 用户登录
		router.POST("login", v1.Login)
	}

	//监听端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}
