// coding: utf-8
// @Author : lryself
// @Date : 2022/2/25 16:38
// @Software: GoLand

package crud

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitCRUDRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("crud")
	
}
