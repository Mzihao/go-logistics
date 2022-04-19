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

		// 物流查询
		auth.GET("logistics/:barcode", v1.QueryExpress)
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
