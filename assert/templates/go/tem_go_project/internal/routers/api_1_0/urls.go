package api_1_0

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/apis/api_1_0"
)

var (
	Api *gin.RouterGroup
)

func InitAPI_1_0Router(engine *gin.Engine) {
	Api = engine.Group("api_1_0")
	Api.Any("version", api_1_0.GetVersion)
}
