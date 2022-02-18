package api1_0

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/apis/api1_0"
)

var (
	Api *gin.RouterGroup
)

func InitAPI1_0Router(engine *gin.Engine) {
	Api = engine.Group("api1_0")
	Api.Any("version", api1_0.GetVersion)

}
