// coding: utf-8
// @Author : lryself
// @Software: GoLand

package project

import (
	"LRYGoCodeGen/internal/apis/api_1_0/projects"
	"LRYGoCodeGen/internal/middlewares"
	"LRYGoCodeGen/internal/routers/api/project/member"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitProjectRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("project")
	Api.Use(middlewares.TokenRequire())
	Api.GET("", projects.ProjectGetHandler)
	Api.POST("", projects.ProjectPostHandler)
	Api.PUT("", projects.ProjectPutHandler)
	Api.DELETE("", projects.ProjectDeleteHandler)
	Api.GET("list", projects.GetListHandler)
	Api.GET("list/page", projects.GetListByPage)
	member.InitMemberRouterGroup(Api)
}
