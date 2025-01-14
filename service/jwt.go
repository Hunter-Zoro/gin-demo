package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type jwtService struct{}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "Bearer"
	AppGuardName = "app"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.DEMO_CONFIG.Jwt.JwtTtl,
			Id:        user.GetUid(),
			Issuer:    GuardName,
			NotBefore: time.Now().Unix() - 1000,
		},
	})

	tokenStr, err := token.SignedString([]byte(global.DEMO_CONFIG.Jwt.Secret))

	tokenData = TokenOutPut{
		tokenStr,
		int(global.DEMO_CONFIG.Jwt.JwtTtl),
		TokenType,
	}
	return

}
