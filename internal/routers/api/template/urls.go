// coding: utf-8
// @Author : lryself
// @Date : 2022/4/5 22:33
// @Software: GoLand

package template

import (
	"LRYGoCodeGen/internal/apis/api_1_0/templates"
	"LRYGoCodeGen/internal/routers/api/project/member"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitTemplateRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("template")
	Api.POST("", templates.TemplatePoolApi)
	Api.POST("update", templates.UpdateTemplateZip)
	Api.GET("list/private", templates.GetPrivateList)
	Api.GET("list/public", templates.GetPublicList)
	Api.GET("list/page", templates.GetListByPage)
	member.InitMemberRouterGroup(Api)
}
