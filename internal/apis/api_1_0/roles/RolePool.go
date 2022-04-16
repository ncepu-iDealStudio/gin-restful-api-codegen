// coding: utf-8
// @Author : lryself
// @Software: GoLand

package roles

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/snowflake"
	"LRYGoCodeGen/internal/services"
	"gitee.com/lryself/go-utils/structs"
	"github.com/gin-gonic/gin"
)

func RolePoolPostHandler(c *gin.Context) {
	var err error

	var Parser struct {
		RoleName       string `json:"RoleName" form:"RoleName" binding:"required"`
		PermissionList string `json:"PermissionList" form:"PermissionList" binding:"required"`
		OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePoolService services.RolePoolService
	RolePoolService.Assign(Parser)
	
	RolePoolService.RoleID = snowflake.GetSnowflakeID()
	err = RolePoolService.Add()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", RolePoolService)
}

func RolePoolGetHandler(c *gin.Context) {
	var err error

	var Parser struct {
		RoleID         string `json:"RoleID" form:"RoleID" binding:"required"`
		RoleName       string `json:"RoleName" form:"RoleName"`
		PermissionList string `json:"PermissionList" form:"PermissionList"`
		OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
		IsDeleted      bool   `json:"IsDeleted" form:"IsDeleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePoolService services.RolePoolService
	RolePoolService.Assign(Parser)
	err = RolePoolService.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", RolePoolService)
}

func RolePoolPutHandler(c *gin.Context) {
	var err error

	var Parser struct {
		RoleID         string `json:"RoleID" form:"RoleID" binding:"required"`
		RoleName       string `json:"RoleName" form:"RoleName"`
		PermissionList string `json:"PermissionList" form:"PermissionList"`
		OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePoolService services.RolePoolService

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	//不能修改业务主键
	delete(args, "RoleID")

	err = RolePoolService.Update(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", RolePoolService)
}

func RolePoolDeleteHandler(c *gin.Context) {
	var err error

	var Parser struct {
		RoleID string `json:"RoleID" form:"RoleID" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var RolePoolService services.RolePoolService
	RolePoolService.Assign(Parser)

	err = RolePoolService.Delete()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", RolePoolService)
}

// 获取列表
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

// 获取列表（分页）
func GetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		services.RolePoolService
		parser.ListParser
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := Parser.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
