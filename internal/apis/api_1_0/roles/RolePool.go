// coding: utf-8
// @Author : lryself
// @Software: GoLand

package roles

import (
	"LRYGoCodeGen/internal/globals/codes"
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RolePoolApi(c *gin.Context) {
	var err error
	var RolePoolService services.RolePoolService

	err = c.ShouldBind(&RolePoolService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePool services.RolePoolService
	//针对业务主键处理
	RolePool.RoleID = RolePoolService.RoleID

	if c.Request.Method == "GET" {
		err = RolePoolService.Get()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = RolePool.Get()
		if err != nil {
			if err.Error() != "record not found" {
				parser.JsonDBError(c, "", err)
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.DataExist,
				"message": "数据已存在！",
			})
			return
		}

		err = RolePoolService.Add()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := RolePoolService.GetModelMap()
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		//不能修改业务主键
		delete(args, "RoleID")

		err = RolePool.Update(args)
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
		RolePoolService = RolePool
	} else if c.Request.Method == "DELETE" {
		err = RolePool.Delete()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
		RolePoolService.IsDeleted = true
	}

	parser.JsonOK(c, "", RolePoolService)
}

func GetListHandler(c *gin.Context) {
	var err error
	var RolePoolService services.RolePoolService

	err = c.ShouldBind(&RolePoolService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := RolePoolService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
