// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"LRYGoCodeGen/core/gen/gen_db"
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/core/globals/vipers"
	"LRYGoCodeGen/core/model/mysql"
	"fmt"
	"path/filepath"
)

func main() {
	var err error

	err = vipers.InitViper()
	if err != nil {
		fmt.Println(err)
	}
	//生成项目代码
	err = gen_program.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
	dbModel, err := mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	//生成model层
	tmplPath := "assert/templates/go/tmpl"
	resultPath := "result/project-center-service"
	err = gen_db.GenCode(dbModel,
		filepath.Join(tmplPath, "model.go.tpl"),
		filepath.Join(resultPath, "internal/models/mysqlModel"),
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成dao层
	err = gen_db.GenCode(dbModel,
		filepath.Join(tmplPath, "dao.go.tpl"),
		filepath.Join(resultPath, "internal/dao"),
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成service层
	err = gen_db.GenCode(dbModel,
		filepath.Join(tmplPath, "service.go.tpl"),
		filepath.Join(resultPath, "internal/services"),
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成api层
	err = gen_db.GenCode(dbModel,
		filepath.Join(tmplPath, "api.go.tpl"),
		filepath.Join(resultPath, "internal/apis/api1_0"),
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成router层
	err = gen_db.GenCode(dbModel,
		filepath.Join(tmplPath, "router.go.tpl"),
		filepath.Join(resultPath, "internal/routers/api1_0"),
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成url文件
	err = gen_db.GenUrlCode(dbModel,
		filepath.Join(tmplPath, "urls.go.tpl"),
		filepath.Join(resultPath, "internal/routers/api1_0"),
	)
	if err != nil {
		fmt.Println(err)
	}
}
