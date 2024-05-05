package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应信息
type Response struct {
	Status  ResCode     `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// 错误响应状态码
func ResponseError(c *gin.Context, status ResCode) {
	c.JSON(http.StatusOK, &Response{
		Status:  status,
		Message: status.Msg(),
		Data:    nil,
	})
}

// 成功响应状态码
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Status:  CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}

// 自定义错误信息响应状态码
func ResponseErrorWithMsg(c *gin.Context, status ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Status:  status,
		Message: msg,
		Data:    nil,
	})
}
