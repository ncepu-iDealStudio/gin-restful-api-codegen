// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 11:36
// @Software: GoLand

package middlewares

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/globals/extensions/currentUser"
	"tem_go_project/internal/globals/parser"
)

func LoginRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取当前用户
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}

		//查询用户信息
		//var userInfoService services.UserInfoService
		//userInfoService.SetUserID(user.UserID)
		//err = userInfoService.Get()
		//if err != nil {
		//	parser.JsonDBError(c, "", err)
		//	c.Abort()
		//	return
		//}
		//user.UserType = userInfoService.UserType
		c.Set("user", user)
		c.Next()
	}
}
