package app

import (
	"github.com/gin-gonic/gin"
	G "github.com/xtpv/xtp-gin/global"
	"github.com/xtpv/xtp-gin/pkg/ecode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = 0
	SuccessMsg  = "success"
)

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	response := result{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: data,
	}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ToErrorResponse(err *ecode.Error) {
	response := result{
		Code: err.Code(),
		Msg:  err.Msg(),
		Data: err.Details(),
	}

	G.Log.Errorf("resp err: %+v", response)
	r.Ctx.JSON(err.StatusCode(), response)
}
