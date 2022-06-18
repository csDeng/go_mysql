package util

import (
	"time"

	"github.com/csDeng/blog/models"
	"github.com/csDeng/blog/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	*models.User
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
