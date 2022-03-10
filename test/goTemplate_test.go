// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 15:50
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gen_db"
	"LRYGoCodeGen/core/gen_mysql"
	"LRYGoCodeGen/core/gen_program"
	"fmt"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	var err error
	//生成项目代码
	err = gen_program.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dbModel, err := gen_mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	//生成model层
	err = gen_db.GenCode(dbModel,
		"assert/templates/go/tmpl/model.go.tpl",
		"result/project-center-service/internal/models/mysqlModel",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成dao层
	err = gen_db.GenCode(dbModel,
		"assert/templates/go/tmpl/dao.go.tpl",
		"result/project-center-service/internal/dao",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成service层
	err = gen_db.GenCode(dbModel,
		"assert/templates/go/tmpl/service.go.tpl",
		"result/project-center-service/internal/services",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成api层
	err = gen_db.GenCode(dbModel,
		"assert/templates/go/tmpl/api.go.tpl",
		"result/project-center-service/internal/apis/api1_0",
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成router层
	err = gen_db.GenCode(dbModel,
		"assert/templates/go/tmpl/router.go.tpl",
		"result/project-center-service/internal/routers/api1_0",
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成url文件
	err = gen_db.GenUrlCode(dbModel,
		"assert/templates/go/tmpl/urls.go.tpl",
		"result/project-center-service/internal/routers/api1_0",
	)
	if err != nil {
		fmt.Println(err)
	}
}
