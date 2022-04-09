// coding: utf-8
// @Author : lryself
// @Date : 2022/4/5 20:28
// @Software: GoLand

package templates

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/globals/snowflake"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/services"
	"LRYGoCodeGen/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mime/multipart"
	"path/filepath"
)

func UploadTemplateZip(c *gin.Context) {
	var err error
	var Parser struct {
		TemplateName string                `form:"TemplateName" json:"TemplateName" binding:"required"`
		TemplateType int8                  `form:"TemplateType" json:"TemplateType" binding:"required"`
		IsPublic     bool                  `form:"IsPublic" json:"IsPublic"`
		File         *multipart.FileHeader `form:"File" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	file := Parser.File
	contentTypeExt := map[string]string{
		".zip": "application/zip",
		".tpl": "text/plain",
	}
	ext := filepath.Ext(file.Filename)
	contentType, ok := contentTypeExt[ext]
	if !ok {
		parser.JsonParameterIllegal(c, "请求参数不合法！", err)
		return
	}
	//上传文件
	TemplateID := snowflake.GetSnowflakeID()
	FileName := TemplateID + ext
	StoreType := viper.GetString("template.StoreType")
	StorePath := fmt.Sprintf(viper.GetString("template.StorePath"), viper.GetString("template.StoreType"), FileName)
	switch StoreType {
	case "2":
		minioClient := utils.GetMinioClient()
		f, err := file.Open()
		if err != nil {
			parser.JsonInternalError(c, "文件打开出错！", err)
		}
		err = minioClient.PutObject(StorePath, contentType, f, file.Size)
		if err != nil {
			parser.JsonInternalError(c, "文件上传出错！", err)
			return
		}
	}

	//添加数据库中信息
	user, err := ginModels.GetUser(c)
	var templateModel services.TemplatePoolService
	templateModel.Assign(Parser)
	templateModel.TemplateID = TemplateID
	templateModel.UserID = user.UserID
	templateModel.StoreType = int8(viper.GetInt("template.StoreType"))
	templateModel.StorePath = StorePath
	err = templateModel.Add()
	if err != nil {
		parser.JsonDBError(c, "", err)
	}
	parser.JsonOK(c, "", templateModel)
	return
}