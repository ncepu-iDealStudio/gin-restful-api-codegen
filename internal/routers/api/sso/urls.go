// coding: utf-8
// @Author : lryself
// @Date : 2022/3/30 18:42
// @Software: GoLand

package sso

import (
	"LRYGoCodeGen/internal/apis/api_1_0/sso"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitSsoRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("sso")
	Api.POST("register", sso.Register)
	Api.POST("login", sso.Login)
	Api.POST("password", sso.LoginByPassword)
	Api.PUT("password", sso.ChangePassword)
	Api.POST("verifyCode/email", sso.LoginByEmailVerifyCode)
	Api.POST("makeVerifyCode/email", sso.MakeEmailVerifyCode)
	Api.POST("logout", sso.Logout)
}
