// coding: utf-8
// @Author : lryself
// @Software: GoLand

package member

import (
	"LRYGoCodeGen/internal/apis/api_1_0/projects/members"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitMemberRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("member")
	Api.GET("", members.ProjectMemberApi)
	Api.POST("", members.ProjectMemberApi)
	Api.PUT("", members.ProjectMemberApi)
	Api.DELETE("", members.ProjectMemberApi)
	Api.GET("list", members.GetListHandler)
	Api.GET("list/page", members.GetListByPage)
}
