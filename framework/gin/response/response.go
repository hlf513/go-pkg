package response

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any, opts ...Option) {
	opt := newOptions(opts...)
	ctx.JSON(opt.HttpCode, map[string]any{
		opt.CodeName:    opt.SuccessCode,
		opt.MessageName: opt.SuccessMessage,
		opt.DataName:    data,
	})
}

func Error(ctx *gin.Context, message string, opts ...Option) {
	opt := newOptions(opts...)
	ctx.JSON(opt.HttpCode, map[string]any{
		opt.CodeName:    opt.ErrorCode,
		opt.MessageName: message,
	})
}
