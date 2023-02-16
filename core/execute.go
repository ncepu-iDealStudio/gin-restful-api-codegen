// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:14
// @Software: GoLand

package core

import (
	"GinCodeGen/core/gen/gen_db"
	"GinCodeGen/core/gen/gen_program"
)

func Execute() {
	//生成项目代码
	gen_program.GenProgramCodeFromTemplates()

	//生成数据库代码
	gen_db.GenDBCodeFromTemplate()
}
