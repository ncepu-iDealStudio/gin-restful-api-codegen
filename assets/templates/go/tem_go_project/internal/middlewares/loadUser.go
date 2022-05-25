// coding: utf-8
// @Author : lryself
// @Date : 2022/3/30 23:20
// @Software: GoLand

package middlewares

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/globals/extensions/currentUser"
	"tem_go_project/internal/globals/parser"
	"tem_go_project/internal/remote/rpcReq"
)

func LoadUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		if token == "" {
			c.Next()
			return
		}

		res, err := rpcReq.VerifyToken(token)
		if err != nil {
			c.Next()
			return
		}

		//加载用户信息到上下文
		User, err := currentUser.NewUser(res.UserID, currentUser.UnKnown)
		defer currentUser.Release(User)
		if err != nil {
			parser.JsonDBError(c, "用户信息未找到", err)
			c.Abort()
			return
		}

		c.Set("user", User)
		c.Next()
	}
}
