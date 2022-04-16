// coding: utf-8
// @Author : lryself
// @Software: GoLand

package permissions

import (
	"LRYGoCodeGen/internal/apis/api_1_0/roles/permissions"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitPermissionsRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("permissions")

	Api.GET("", permissions.RolePermissionsApi)
	Api.POST("", permissions.RolePermissionsApi)
	Api.PUT("", permissions.RolePermissionsApi)
	Api.DELETE("", permissions.RolePermissionsApi)
	Api.GET("list", permissions.GetListHandler)
	Api.GET("list/page", permissions.GetListByPage)
}
