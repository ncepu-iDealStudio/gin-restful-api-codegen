// coding: utf-8
// @Author : lryself
// @Software: GoLand

package role

import (
	"LRYGoCodeGen/internal/apis/api_1_0/roles"
	"LRYGoCodeGen/internal/middlewares"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/routers/api/role/permissions"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitRoleRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("role")
	Api.GET("list", roles.GetListHandler)

	Api.Use(middlewares.TokenRequire())
	Api.Use(middlewares.AuthMiddleware(ginModels.Platform))

	Api.Any("", roles.RolePoolApi)
	Api.GET("list/page", roles.GetListByPage)

	permissions.InitPermissionsRouterGroup(Api)
}
