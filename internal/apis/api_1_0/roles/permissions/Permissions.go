// coding: utf-8
// @Author : lryself
// @Software: GoLand

package permissions

import (
	"LRYGoCodeGen/internal/globals/codes"
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RolePermissionsApi(c *gin.Context) {
	var err error
	var RolePermissionsService services.RolePermissionsService

	err = c.ShouldBind(&RolePermissionsService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePremissions services.RolePermissionsService
	//针对业务主键处理
	RolePremissions.PermissionID = RolePermissionsService.PermissionID

	if c.Request.Method == "GET" {
		err = RolePermissionsService.Get()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		err = RolePremissions.Get()
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

		err = RolePermissionsService.Add()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		args, err := RolePermissionsService.GetModelMap()
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		//不能修改业务主键
		delete(args, "PermissionID")

		err = RolePremissions.Update(args)
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
		RolePermissionsService = RolePremissions
	} else if c.Request.Method == "DELETE" {
		err = RolePremissions.Delete()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
		RolePermissionsService.IsDeleted = true
	}

	parser.JsonOK(c, "", RolePermissionsService)
}

func GetListHandler(c *gin.Context) {
	var err error
	var RolePermissionsService services.RolePermissionsService

	err = c.ShouldBind(&RolePermissionsService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := RolePermissionsService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
