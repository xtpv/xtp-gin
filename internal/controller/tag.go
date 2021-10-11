package controller

import (
	"github.com/gin-gonic/gin"
	G "github.com/xtpv/xtp-gin/global"
	"github.com/xtpv/xtp-gin/internal/service"
	"github.com/xtpv/xtp-gin/pkg/app"
	"github.com/xtpv/xtp-gin/pkg/ecode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} ecode.Error "请求错误"
// @Failure 500 {object} ecode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	req := service.TagListRequest{}
	G.Event.Trigger(G.Sync, req)
	resp := app.NewResponse(c)
	errs := app.Validate(c, &req)
	if errs != nil {
		resp.ToErrorResponse(ecode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(resp.Ctx)
	list, err := svc.GetTagList(req)
	if err != nil {
		resp.ToErrorResponse(ecode.ServerError.WithDetails(err.Error()))
		return
	}
	resp.ToResponse(list)
	return
}

func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
