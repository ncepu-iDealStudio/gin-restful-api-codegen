// coding: utf-8
// @Author : lryself
// @Date : 2022/4/5 21:45
// @Software: GoLand

package templates

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"LRYGoCodeGen/internal/utils"
	"github.com/gin-gonic/gin"
)

type listParser struct {
	TemplateID   string `json:"TemplateID" form:"TemplateID"`
	UserID       string `json:"UserID" form:"UserID"`
	TemplateName string `json:"TemplateName" form:"TemplateName"`
	TemplateType int8   `json:"TemplateType" form:"TemplateType"`
	parser.ListParser
}

func GetPublicList(c *gin.Context) {
	var err error
	var Parser listParser

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	var TemplatePoolService services.TemplatePoolService
	utils.StructAssign(TemplatePoolService, Parser, "json")
	TemplatePoolService.IsPublic = true
	results, err := TemplatePoolService.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}
func GetPrivateList(c *gin.Context) {
	var err error
	var Parser listParser

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	user, err := ginModels.GetUser(c)
	if err != nil {
		parser.JsonAccessDenied(c, "请先登录！")
	}

	var TemplatePoolService services.TemplatePoolService
	utils.StructAssign(TemplatePoolService, Parser, "json")
	TemplatePoolService.IsPublic = false
	TemplatePoolService.UserID = user.UserID
	results, err := TemplatePoolService.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}

func GetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		services.TemplatePoolService
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
