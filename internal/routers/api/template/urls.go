// coding: utf-8
// @Author : lryself
// @Date : 2022/4/5 22:33
// @Software: GoLand

package template

import (
	"LRYGoCodeGen/internal/apis/api_1_0/templates"
	"LRYGoCodeGen/internal/middlewares"
	"LRYGoCodeGen/internal/models/ginModels"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitTemplateRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("template")
	Api.GET("list/public", templates.GetPublicList)
	Api.Use(middlewares.TokenRequire())
	Api.GET("", templates.TemplatePoolApi)
	Api.PUT("", templates.TemplatePoolApi)
	Api.DELETE("", templates.TemplatePoolApi)
	Api.POST("upload", templates.UploadTemplateZip)
	Api.GET("list/private", templates.GetPrivateList)
	Api.GET("list/page", middlewares.AuthMiddleware(ginModels.Platform), templates.GetListByPage)
}
