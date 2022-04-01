// coding: utf-8
// @Author : lryself
// @Date : 2022/3/30 23:20
// @Software: GoLand

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"tem_go_project/internal/globals/jwt"
	"tem_go_project/internal/globals/parser"
	"tem_go_project/internal/models/ginModels"
)

func LoadUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		jwtChaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
		if err != nil {
			c.Next()
			return
		}

		//加载用户信息到上下文
		User, err := ginModels.NewUser(jwtChaim.UserID, jwtChaim.UserType)
		if err != nil {
			parser.JsonDBError(c, "用户信息未找到", err)
			c.Abort()
			return
		}

		c.Set("user", User)
		c.Next()
	}
}
