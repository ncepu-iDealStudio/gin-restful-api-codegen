// coding: utf-8
// @Author : lryself
// @Date : 2022/2/18 18:16
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gencode/db_code"
	"LRYGoCodeGen/core/gencode/templete_code"
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestGenCode(t *testing.T) {
	err := templete_code.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	dbModel, err := genmysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	err = db_code.GenModelsCode(dbModel)
	if err != nil {
		t.Fatal("model生成报错：", err)
	}
	err = db_code.GenDaoCode(dbModel)
	if err != nil {
		t.Fatal("dao生成报错：", err)
	}
	err = db_code.GenServicesCode(dbModel)
	if err != nil {
		t.Fatal("service生成报错：", err)
	}
	err = db_code.GenApiCode(dbModel)
	if err != nil {
		t.Fatal("api生成报错：", err)
	}
	err = db_code.GenRouterCode(dbModel)
	if err != nil {
		t.Fatal("router生成报错：", err)
	}
}
