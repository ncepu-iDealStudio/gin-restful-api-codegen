// coding: utf-8
// @Author : lryself
// @Date : 2021/5/16 19:20
// @Software: GoLand

package middlewares

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/globals/parser"
	"tem_go_project/internal/models/ginModels"
)

// AuthUserType 垂直鉴权
func AuthUserType(allowRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息

		user, err := ginModels.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			return
		}
		// 验证权限
		var f bool
		for _, role := range allowRole {
			if user.UserType == role {
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

// AuthUserID todo 水平鉴权
func AuthUserID(allowRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		user, err := ginModels.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			return
		}
		// 验证权限
		var f bool
		for _, role := range allowRole {
			if user.UserType == role {
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
