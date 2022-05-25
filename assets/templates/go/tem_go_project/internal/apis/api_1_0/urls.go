package api_1_0

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.Any("version", GetVersion)
}
