package routers

import (
	"github.com/gin-gonic/gin"
	"tem_go_project/internal/routers/api_1_0"
)

func InitRouter(engine *gin.Engine) {
	api_1_0.InitAPI_1_0Router(engine)
}
