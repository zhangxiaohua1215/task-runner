package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

// Success 成功的响应
//
//	①使用默认成功消息：操作成功，且无响应数据时 api.Success(c, "")
//	②使用自定义成功消息，且无响应数据时 api.Success(c, "OK")
//	③使用自定义成功消息，且有响应数据时 api.Success(c, "OK", data)
func Success(c *gin.Context, msg string, data ...interface{}) {
	if msg == "" {
		msg = "操作成功"
	}

	ResponseWithCode(c, SUCCESS, msg, data...)
}

// Fail 失败的响应
//
//	①使用默认失败消息：操作失败，且无响应数据时 api.Fail(c, "")
//	②使用自定义失败消息，且无响应数据时 api.Fail(c, "Fail")
//	③使用自定义失败消息，且有响应数据时 api.Fail(c, "Fail", data)
func Fail(c *gin.Context, msg string, data ...interface{}) {
	if msg == "" {
		msg = "操作失败"
	}

	// 设置请求状态，HttpRequestLog 中间件中要使用
	c.Error(errors.New(msg))

	ResponseWithCode(c, ERROR, msg, data...)
}

// ResponseWithCode 自定义 code 的响应（请优先使用 Success 和 Fail 方法）
func ResponseWithCode(c *gin.Context, code int, msg string, data ...interface{}) {
	var resp interface{}
	if len(data) > 0 {
		resp = data[0]
	}

	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: resp,
	})
}
