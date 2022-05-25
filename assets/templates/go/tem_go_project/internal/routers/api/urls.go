package api

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/apis/api_1_0"
)

var (
	Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.Any("version", api_1_0.GetVersion)
}
