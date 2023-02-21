// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:14
// @Software: GoLand

package core

import (
	"GinCodeGen/core/gen/gen_db"
	"GinCodeGen/core/gen/gen_program"
	"GinCodeGen/tools/logger"
)

func Execute() {
	var log = logger.GetLogger()
	//生成项目代码
	log.Infoln("开始生成项目代码")
	gen_program.GenProgramCodeFromTemplates()
	log.Infoln("项目代码生成成功")

	//生成数据库代码
	log.Infoln("生成数据库代码")
	gen_db.GenDBCodeFromTemplate()
	log.Infoln("数据库代码生成成功")
}
