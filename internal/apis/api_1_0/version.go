// coding: utf-8
// @Author : lryself
// @Date : 2021/12/28 21:32
// @Software: GoLand

package api_1_0

import (
	"LRYGoCodeGen/internal/globals/codes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "当前接口版本信息。",
		"data": gin.H{
			"version": "1.0",
		},
	})
	return
}
