// coding: utf-8
// @Author : lryself
// @Date : 2021/5/16 19:20
// @Software: GoLand

package middlewares

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/globals/extensions/currentUser"
	"tem_go_project/internal/globals/parser"
)

// AuthUserType 垂直鉴权
func AuthUserType(allowType ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}
		// 验证权限
		if !user.AuthType(allowType...) {
			parser.JsonAccessDenied(c, "您无权访问！")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// AuthUserID todo 水平鉴权
func AuthUserID(allowUserID ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}
		// 验证权限
		var f bool
		for _, userID := range allowUserID {
			if user.UserID == userID {
				f = true
				break
			}
		}
		if !f {
			parser.JsonAccessDenied(c, "您无权访问！")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
