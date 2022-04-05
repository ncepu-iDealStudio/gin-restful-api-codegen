// coding: utf-8
// @Author : lryself
// @Software: GoLand

package templates

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"gitee.com/lryself/go-utils/structs"
	"github.com/gin-gonic/gin"
)

func TemplatePoolApi(c *gin.Context) {
	var err error

	var Parser struct {
		TemplateID   string `json:"TemplateID" form:"TemplateID" binding:"required"`
		TemplateName string `json:"TemplateName" form:"TemplateName"`
		TemplateType int8   `json:"TemplateType" form:"TemplateType"`
		IsPublic     bool   `json:"IsPublic" form:"IsPublic"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var TemplatePool services.TemplatePoolService
	TemplatePool.TemplateID = Parser.TemplateID
	err = TemplatePool.Get()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	user, err := ginModels.GetUser(c)
	if !user.AuthSelf(TemplatePool.UserID) {
		if TemplatePool.IsPublic && c.Request.Method == "GET" {
			parser.JsonOK(c, "", TemplatePool)
		}
		parser.JsonAccessDenied(c, "您没有权限获取此项目")
		return
	}

	if c.Request.Method == "PUT" {
		args, err := structs.StructToMap(Parser, "json")
		delete(args, "TemplateID")

		err = TemplatePool.Update(args)
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	} else if c.Request.Method == "DELETE" {
		err = TemplatePool.Delete()
		if err != nil {
			parser.JsonDBError(c, "", err)
			return
		}
	}

	parser.JsonOK(c, "", TemplatePool)
}
