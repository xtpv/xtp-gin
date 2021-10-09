package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xtp2217866847/xtp-gin/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := controller.NewArticle()
	tag := controller.NewTag()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/tags", tag.Create)
		v1.DELETE("/tags/:id", tag.Delete)
		v1.PUT("/tags/:id", tag.Update)
		v1.PATCH("/tags/:id/state", tag.Update)
		v1.GET("/tags", tag.Get)

		v1.POST("/articles", article.Create)
		v1.DELETE("/articles/:id", article.Delete)
		v1.PUT("/articles/:id", article.Update)
		v1.PATCH("/articles/:id/state", article.Update)
		v1.GET("/articles/:id", article.Get)
		v1.GET("/articles", article.List)
	}

	return r
}
