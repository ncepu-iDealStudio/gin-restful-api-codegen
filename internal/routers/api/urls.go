package api

import (
	"LRYGoCodeGen/internal/apis/api_1_0"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.Any("version", api_1_0.GetVersion)
}
