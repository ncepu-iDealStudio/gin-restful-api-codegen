package routers

import (
	"LRYGoCodeGen/internal/routers/api_1_0"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	api_1_0.InitAPI_1_0Router(engine)
}
