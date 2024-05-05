package mysql

import (
	"GoofyLink/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("用户不存在")
	ErrorPasswordError = errors.New("密码错误")
)

const secret = "goofy.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入用户一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	if err != nil {
		return err
	}
	return
}

// encryptPassword 对密码进行加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// QueryUser 查询用户信息并返回
func QueryUser(user *models.User) (err error) {
	oldPassword := user.Password // 用户传过来的密码
	sqlStr := `select user_id , username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	// 没查到这个用户名
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库出错
		return err
	}
	// 密码验证
	password := encryptPassword(oldPassword)
	if password != user.Password {
		return ErrorPasswordError
	}
	return
}
