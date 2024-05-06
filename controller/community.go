package controller

import (
	"GoofyLink/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CommunityHandler(c *gin.Context) {
	// 获取社区信息
	// 1.查询数据库
	data, err := logic.QueryCommunity()
	if err != nil {
		zap.L().Error("查询社区信息失败", zap.Any("err", err))
		ResponseError(c, CodeServerBusy)
	}
	// 2.返回数据
	ResponseSuccess(c, data)
}
