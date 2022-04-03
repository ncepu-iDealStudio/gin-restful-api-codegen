// coding: utf-8
// @Author : lryself
// @Date : 2022/4/3 15:37
// @Software: GoLand

package program

import (
	"LRYGoCodeGen/internal/globals/parser"
	"LRYGoCodeGen/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type uploadTemplateZipParser struct {
	Account     string `form:"Account" json:"Account" binding:"required"`
	Password    string `form:"Password" json:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" json:"NewPassword" binding:"required"`
}

func UploadTemplateZip(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		parser.JsonParameterIllegal(c, "获取上传的文件有误", err)
		return
	}
	files := form.File["file"]
	minioClient := utils.GetMinioClient()
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			parser.JsonInternalError(c, "文件打开出错！", err)
		}
		err = minioClient.PutObject(file.Filename, f, file.Size)
		if err != nil {
			parser.JsonInternalError(c, "文件上传出错！", err)
			return
		}

		fmt.Println(file.Size)
	}
	return
}
