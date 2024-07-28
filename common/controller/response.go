package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponseData(data interface{}) ResponseData {
	return ResponseData{
		Data: data,
	}
}

func newResponseCodeError(code string, err error) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  err.Error(),
	}
}

func newResponseCodeMsg(code, msg string) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
	}
}

// SendBadRequestBody
func SendBadRequestBody(ctx *gin.Context, err error) {
	if _, ok := err.(errorCode); ok {
		SendError(ctx, err)
	} else {
		ctx.JSON(
			http.StatusBadRequest,
			newResponseCodeMsg(errorBadRequestBody, err.Error()),
		)
	}
}

func SendBadRequestParam(ctx *gin.Context, err error) {
	if _, ok := err.(errorCode); ok {
		SendError(ctx, err)
	} else {
		ctx.JSON(
			http.StatusBadRequest,
			newResponseCodeMsg(errorBadRequestParam, err.Error()),
		)
	}
}

func SendRespOfPut(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(http.StatusAccepted, newResponseCodeMsg("", "success"))
	} else {
		ctx.JSON(http.StatusAccepted, newResponseData(data))
	}
}

func SendRespOfGet(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResponseData(data))
}

func SendRespOfPost(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(http.StatusCreated, newResponseCodeMsg("", "success"))
	} else {
		ctx.JSON(http.StatusCreated, newResponseData(data))
	}
}

func SendRespOfDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, newResponseCodeMsg("", "success"))
}

func SendError(ctx *gin.Context, err error) {
	sc, code := httpError(err)
	ctx.AbortWithError(sc, err)
	ctx.JSON(sc, newResponseCodeMsg(code, err.Error()))
}
