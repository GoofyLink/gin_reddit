package controller

import (
	"GoofyLink/logic"
	"GoofyLink/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	_ "github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// controller层
// 优化后的代码片段
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验 请求重定向
	signupParams := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(signupParams); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, TrimFieldPrefixes(errs.Translate(trans)))
		return
	}

	// 2. 处理业务 logic层
	if err := logic.SignUp(signupParams); err != nil {
		zap.L().Error("logic.SignUp is failed", zap.Error(err))
		ResponseError(c, CodeSignUpFailed)
		return
	}
	// 3. 返回数据
	ResponseSuccess(c, CodeSuccess)
}

// 用户登录
func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	loginParams := new(models.ParamLogin)
	if err := c.ShouldBindJSON(loginParams); err != nil {
		zap.L().Error("login is failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, TrimFieldPrefixes(errs.Translate(trans)))
		return
	}
	// 2. 处理业务 logic层
	if err := logic.Login(loginParams); err != nil {
		zap.L().Error("logic.Login is failed", zap.String("username", loginParams.Username), zap.Error(err))
		ResponseError(c, CodeUserPasswordError)
		return
	}
	// 3.返回数据
	ResponseError(c, CodeSuccess)
}
