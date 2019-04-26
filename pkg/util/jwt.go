package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/selinplus/goweb/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username,password string) (string,error){
	nowTime := time.Now()
	expireTime := nowTime.Add(3*time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"daerdo-blog",
		},
	}
	tokeClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token,err := tokeClaims.SignedString(jwtSecret)

	return token,err
}
func ParseToken(token string) (*Claims,error){
	tokenClaims,err := jwt.ParseWithClaims(token, &Claims{},func(token *jwt.Token) (interface{}, error){
		return jwtSecret,nil
	})
	if tokenClaims != nil{

		if claims,ok:= tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err
}