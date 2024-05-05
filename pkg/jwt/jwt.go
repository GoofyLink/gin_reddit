package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

var mySecret = []byte("夏天夏天悄悄过去")

// 创建一个jwt对象
type MyClaims struct {
	UserID   int64  `json:"username"`
	Username string `json:"password"`
	jwt.StandardClaims
}

// 生成jwt
func GenToken(userID int64, username string) (string, error) {
	// 创建一个MyClaims对象
	myclaims := &MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				// 过期时间
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "reddit", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaims)
	return token.SignedString(mySecret)
}

// 解析jwt
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
