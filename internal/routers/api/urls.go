package api

import (
	"LRYGoCodeGen/internal/apis/api_1_0"
	"LRYGoCodeGen/internal/routers/api/project"
	"LRYGoCodeGen/internal/routers/api/role"
	"LRYGoCodeGen/internal/routers/api/sso"
	"LRYGoCodeGen/internal/routers/api/template"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.Any("version", api_1_0.GetVersion)
	sso.InitSsoRouterGroup(Api)
	project.InitProjectRouterGroup(Api)
	role.InitRoleRouterGroup(Api)
	template.InitTemplateRouterGroup(Api)
}
