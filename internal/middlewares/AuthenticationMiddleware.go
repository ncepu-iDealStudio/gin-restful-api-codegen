// coding: utf-8
// @Author : lryself
// @Date : 2021/5/16 19:20
// @Software: GoLand

package middlewares

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/models/ginModels"
	"errors"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(allowRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		temp, ok := c.Get("user")
		if !ok {
			parser.JsonInternalError(c, "您可能没有登录！", errors.New("获取用户信息错误！"))
			c.Abort()
			return
		}
		user := temp.(ginModels.UserModel)
		// 验证权限
		if !user.Auth(allowRole...) {
			parser.JsonAccessDenied(c, "您无权访问！")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
