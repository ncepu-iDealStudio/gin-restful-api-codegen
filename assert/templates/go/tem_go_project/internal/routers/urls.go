package routers

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/routers/api1_0"
)

func InitRouter(engine *gin.Engine) {
	api1_0.InitAPI1_0Router(engine)
}
