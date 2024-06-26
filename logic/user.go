package logic

import (
	"GoofyLink/dao/mysql"
	"GoofyLink/models"
	"GoofyLink/pkg/jwt"
	"GoofyLink/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	// 1.获取用户名和密码
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 这里存取token 传递的是指针可以获取到UserID
	// 2.查询数据库，是否有这个用户并返回对应账号和密码比较如果相等则成功登录
	if err := mysql.QueryUser(user); err != nil {
		return " ", err
	}

	// 可以拿到userID 生成token
	return jwt.GenToken(user.UserID, user.Username)
}
