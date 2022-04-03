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
	Api.Any("", permissions.RolePermissionsApi)
	Api.GET("list", permissions.GetListHandler)
}
