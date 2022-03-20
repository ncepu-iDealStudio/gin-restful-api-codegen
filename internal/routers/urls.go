package routers

import (
	"LRYGoCodeGen/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	api.InitAPIRouter(engine)
}
