package jwtMgr

import (
	jwt "github.com/dgrijalva/jwt-go"
	"microshop/infrastructure/logger"
	"time"
)

type JwtMgr struct {

}

var (
	key = []byte("adfadf!@#2")
	expireTime = 24 * 365
	log = logger.GetLogger()
)

type CustomClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func (o JwtMgr) GetUserID(tokenstr string) string {
	token := o.Parse(tokenstr)
	if token == nil {
		return ""
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID
	}
	return ""
}

func (o JwtMgr) Parse(tokenstr string) *jwt.Token {
	token, err := jwt.ParseWithClaims(tokenstr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Errorln("parse token str err:", err)
		return nil
	}
	if token.Valid {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Errorln("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			log.Errorln("The token is expired or not valid.")
		} else {
			log.Errorln("Couldn't handle this token:", err)
		}
	} else {
		log.Errorln("Couldn't handle this token:", err)
	}
	return nil

}

func (o JwtMgr) Create(userId string) string {
	claims := CustomClaims{
		userId, jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expireTime)).Unix(),
			Issuer:    "custom",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(key)
	if err == nil {
		return tokenstr
	}
	return ""
}

func (o JwtMgr) Verify(tokenstr string) bool {
	token := o.Parse(tokenstr)
	return token != nil
}





