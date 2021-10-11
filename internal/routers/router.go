package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/xtpv/xtp-gin/docs"
	"github.com/xtpv/xtp-gin/internal/controller"
	"github.com/xtpv/xtp-gin/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	// 添加中间件
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Translations(),
	)

	// 绑定 swagger
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 添加路由
	addRoutes(r)

	return r
}

func addRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		// 标签
		tag := controller.NewTag()
		v1.POST("/tags", tag.Create)
		v1.DELETE("/tags/:id", tag.Delete)
		v1.PUT("/tags/:id", tag.Update)
		v1.PATCH("/tags/:id/state", tag.Update)
		v1.GET("/tags/:id", tag.Get)
		v1.GET("/tags", tag.List)

		// 文章
		article := controller.NewArticle()
		v1.POST("/articles", article.Create)
		v1.DELETE("/articles/:id", article.Delete)
		v1.PUT("/articles/:id", article.Update)
		v1.PATCH("/articles/:id/state", article.Update)
		v1.GET("/articles/:id", article.Get)
		v1.GET("/articles", article.List)
	}
}
