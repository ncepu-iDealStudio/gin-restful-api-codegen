// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 20:40
// @Software: GoLand

package members

import (
	"LRYGoCodeGen/internal/globals/codes"
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/services"
	"gitee.com/lryself/go-utils/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProjectMemberApi(c *gin.Context) {
	var err error

	var ProjectMember services.ProjectMemberService

	if c.Request.Method == "GET" {
		err = c.ShouldBind(&ProjectMember)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		err = ProjectMember.Get()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "POST" {
		var Parser struct {
			ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
			UserID    string `json:"UserID" form:"UserID" binding:"required"`
			RoleID    string `json:"RoleID" form:"RoleID" binding:"required"`
			OtherInfo string `json:"OtherInfo" form:"OtherInfo"`
		}
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}

		ProjectMember.ProjectID = Parser.ProjectID
		ProjectMember.UserID = Parser.UserID
		err = ProjectMember.Get()
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

		err = ProjectMember.Add()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "PUT" {
		var Parser struct {
			ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
			UserID    string `json:"UserID" form:"UserID" binding:"required"`
			RoleID    string `json:"RoleID" form:"RoleID"`
			OtherInfo string `json:"OtherInfo" form:"OtherInfo"`
		}
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		ProjectMember.ProjectID = Parser.ProjectID
		ProjectMember.UserID = Parser.UserID

		args, err := structs.StructToMap(Parser, "json")
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		//不能修改业务主键
		delete(args, "ProjectID")
		delete(args, "UserID")

		err = ProjectMember.Update(args)
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		var Parser struct {
			ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
			UserID    string `json:"UserID" form:"UserID" binding:"required"`
		}
		err = c.ShouldBind(&Parser)
		if err != nil {
			parser.JsonParameterIllegal(c, "", err)
			return
		}
		ProjectMember.ProjectID = Parser.ProjectID
		ProjectMember.UserID = Parser.UserID
		err = ProjectMember.Delete()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	}

	parser.JsonOK(c, "", ProjectMember)
}

func GetListHandler(c *gin.Context) {
	var err error
	var ProjectMemberService services.ProjectMemberService
	var Parser struct {
		ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
		UserID    string `json:"UserID" form:"UserID"`
		RoleID    string `json:"RoleID" form:"RoleID"`
		OtherInfo string `json:"OtherInfo" form:"OtherInfo"`
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	ProjectMemberService.Assign(Parser)

	results, err := ProjectMemberService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
func GetListByPage(c *gin.Context) {
	var err error
	var ProjectMemberService services.ProjectMemberService
	var Parser struct {
		ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
		UserID    string `json:"UserID" form:"UserID"`
		RoleID    string `json:"RoleID" form:"RoleID"`
		OtherInfo string `json:"OtherInfo" form:"OtherInfo"`
		parser.ListParser
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	ProjectMemberService.Assign(Parser)

	results, err := ProjectMemberService.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
