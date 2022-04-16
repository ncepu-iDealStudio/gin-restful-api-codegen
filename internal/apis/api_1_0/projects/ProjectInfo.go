// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 16:47
// @Software: GoLand

package projects

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/snowflake"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"gitee.com/lryself/go-utils/structs"
	"github.com/gin-gonic/gin"
)

func ProjectPostHandler(c *gin.Context) {
	var err error

	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
		return
	}

	var Parser struct {
		ProjectName    string `json:"ProjectName" form:"ProjectName" binding:"required"`
		ProjectContext string `json:"ProjectContext" form:"ProjectContext" binding:"required"`
		OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
	}
	var Project services.ProjectPoolService
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	Project.Assign(Parser)
	Project.ProjectID = snowflake.GetSnowflakeID()
	Project.UserID = user.UserID

	err = Project.Add()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", Project)
}

func ProjectGetHandler(c *gin.Context) {
	var err error

	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
		return
	}

	var Parser struct {
		ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
	}
	var ProjectPool services.ProjectPoolService

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	ProjectPool.Assign(Parser)
	err = ProjectPool.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	if !user.AuthSelf(ProjectPool.UserID) {
		parser.JsonAccessDenied(c, "只允许操作自己的项目！")
		return
	}
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ProjectPool)
}

func ProjectPutHandler(c *gin.Context) {
	var err error

	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
		return
	}

	var Parser struct {
		ProjectID      string `json:"ProjectID" form:"ProjectID" binding:"required"`
		UserID         string `json:"UserID" form:"UserID"`
		ProjectName    string `json:"ProjectName" form:"ProjectName"`
		ProjectContext string `json:"ProjectContext" form:"ProjectContext"`
		OtherInfo      string `json:"OtherInfo" form:"OtherInfo"`
	}

	var ProjectPool services.ProjectPoolService
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	ProjectPool.ProjectID = Parser.ProjectID
	err = ProjectPool.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	if !user.AuthSelf(ProjectPool.UserID) {
		parser.JsonAccessDenied(c, "只能操作自己的项目！")
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
	parser.JsonOK(c, "", ProjectPool)
}

func ProjectDeleteHandler(c *gin.Context) {
	var err error

	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
		return
	}

	var Parser struct {
		ProjectID string `json:"ProjectID" form:"ProjectID" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ProjectPool services.ProjectPoolService
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	ProjectPool.ProjectID = Parser.ProjectID
	err = ProjectPool.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	if !user.AuthSelf(ProjectPool.UserID) {
		parser.JsonAccessDenied(c, "只能操作自己的项目！")
		return
	}

	ProjectPool.Assign(Parser)

	err = ProjectPool.Delete()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", ProjectPool)
}

func GetListHandler(c *gin.Context) {
	var err error
	var ProjectPoolService struct {
		ProjectID      string `json:"ProjectID" form:"ProjectID"`
		UserID         string `json:"UserID" form:"UserID"`
		ProjectName    string `json:"ProjectName" form:"ProjectName"`
		ProjectContext string `json:"ProjectContext" form:"ProjectContext"`
	}

	err = c.ShouldBind(&ProjectPoolService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	var ProjectMember services.VProjectMemberService
	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
	}
	ProjectMember.Assign(ProjectPoolService)
	if !user.IsAdmin() {
		ProjectMember.MemberID = user.UserID
	}
	results, err := ProjectMember.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}

func GetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		services.ProjectPoolService
		parser.ListParser
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ProjectMember services.VProjectMemberService
	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonGinUserError(c, err)
	}
	ProjectMember.Assign(Parser.ProjectPoolService)
	if !user.IsAdmin() {
		ProjectMember.MemberID = user.UserID
	}
	results, err := ProjectMember.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
