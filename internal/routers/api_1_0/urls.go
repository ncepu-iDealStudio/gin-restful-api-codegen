package api_1_0

import (
	"LRYGoCodeGen/internal/apis/api_1_0"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPI_1_0Router(engine *gin.Engine) {
	Api = engine.Group("api1_0")
	Api.Any("version", api_1_0.GetVersion)
}
