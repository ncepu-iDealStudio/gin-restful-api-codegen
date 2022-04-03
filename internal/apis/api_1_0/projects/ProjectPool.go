// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 16:47
// @Software: GoLand

package projects

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/snowflake"
	"LRYGoCodeGen/internal/services"
	"LRYGoCodeGen/internal/utils"
	"gitee.com/lryself/go-utils/structs"
	"github.com/gin-gonic/gin"
)

type projectPoolGetParser struct {
	ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
	UserID    string `json:"UserID" form:"UserID"`
}
type projectPoolPostParser struct {
	UserID         string `json:"UserID" form:"UserID" binding:"required"`
	ProjectName    string `json:"ProjectName" form:"ProjectName" binding:"required"`
	ProjectContext string `json:"ProjectContext" form:"ProjectContext" binding:"required"`
	OtherInfo      string `json:"OtherInfo" form:"OtherInfo" binding:"required"`
}
type projectPoolPutParser struct {
	ProjectID      string `json:"ProjectID" form:"ProjectID" binding:"required"`
	UserID         string `json:"UserID" form:"UserID"`
	ProjectName    string `json:"ProjectName" form:"ProjectName"`
	ProjectContext string `json:"ProjectContext" form:"ProjectContext"`
	OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
}
type projectPoolDeleteParser struct {
	ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
}

func ProjectPoolApi(c *gin.Context) {
	var err error
	var ProjectPool services.ProjectPoolService

	if c.Request.Method == "GET" {
		var Parser projectPoolGetParser
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		utils.StructAssign(ProjectPool, Parser, "json")
		err = ProjectPool.Get()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		var Parser projectPoolPostParser
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		utils.StructAssign(ProjectPool, Parser, "json")

		ProjectPool.ProjectID = snowflake.GetSnowflakeID()
		err = ProjectPool.Add()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		var Parser projectPoolPutParser
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		args, err := structs.StructToMap(Parser, "json")
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		//不能修改业务主键
		delete(args, "ProjectID")

		err = ProjectPool.Update(args)
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		var Parser projectPoolDeleteParser
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		utils.StructAssign(ProjectPool, Parser, "json")

		err = ProjectPool.Delete()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	}

	parser.JsonOK(c, "", ProjectPool)
}

func GetListHandler(c *gin.Context) {
	var err error
	var ProjectPoolService services.ProjectPoolService

	err = c.ShouldBind(&ProjectPoolService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := ProjectPoolService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
