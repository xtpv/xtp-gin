package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xtpv/xtp-gin/pkg/app"
	"github.com/xtpv/xtp-gin/pkg/ecode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(ecode.ServerError)
	return
}
func (a Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(ecode.ServerError)
	return
}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
