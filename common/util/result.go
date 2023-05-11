package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/logger"
)

// ResponseStruct
type ResponseStruct struct{}

// Response gloabl response fmt
var Response = ResponseStruct{}

// ResponsFmt
//
//	@receiver ResponseStruct
//	@param ctx
//	@param httpStatus
//	@param code
//	@param data
//	@param msg
func (ResponseStruct) ResponsFmt(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}

// Success
//
//	@receiver r
//	@param ctx
//	@param data
//	@param msg
func (r ResponseStruct) Success(ctx *gin.Context, data gin.H, msg string) {
	r.ResponsFmt(ctx, http.StatusOK, 200, data, msg)
}

// Error
//
//	@receiver r
//	@param ctx
//	@param data
//	@param msg
func (r ResponseStruct) Error(ctx *gin.Context, data gin.H, msg string) {
	logger.Logger.Error(msg)

	r.ResponsFmt(ctx, 200, 400, data, msg)
}
