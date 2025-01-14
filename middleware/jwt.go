package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/flipped-aurora/gin-vue-admin/server/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()

			return
		}
		tokenStr = tokenStr[len(service.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.DEMO_CONFIG.Jwt.Secret), nil
		})
		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
