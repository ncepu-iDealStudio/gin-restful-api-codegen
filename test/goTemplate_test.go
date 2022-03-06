// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 15:50
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gencode/templete_code"
	"LRYGoCodeGen/core/go_template/db_code"
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	var err error
	//生成项目代码
	err = templete_code.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dbModel, err := genmysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	//生成model层
	err = db_code.GenCode(dbModel,
		"assert/templates/go/tmpl/model.go.tpl",
		"result/usercenter-service/internal/models/mysqlModel",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成dao层
	err = db_code.GenCode(dbModel,
		"assert/templates/go/tmpl/dao.go.tpl",
		"result/usercenter-service/internal/dao",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成service层
	err = db_code.GenCode(dbModel,
		"assert/templates/go/tmpl/service.go.tpl",
		"result/usercenter-service/internal/services",
		false,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成api层
	err = db_code.GenCode(dbModel,
		"assert/templates/go/tmpl/api.go.tpl",
		"result/usercenter-service/internal/apis/api1_0",
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成router层
	err = db_code.GenCode(dbModel,
		"assert/templates/go/tmpl/router.go.tpl",
		"result/usercenter-service/internal/routers/api1_0",
		true,
	)
	if err != nil {
		fmt.Println(err)
	}
	//生成url文件
	err = db_code.GenUrlCode(dbModel,
		"assert/templates/go/tmpl/urls.go.tpl",
		"result/usercenter-service/internal/routers/api1_0",
	)
	if err != nil {
		fmt.Println(err)
	}
}
