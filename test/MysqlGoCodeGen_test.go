// coding: utf-8
// @Author : lryself
// @Date : 2022/2/18 18:16
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/core/model/mysql"
	db_code2 "LRYGoCodeGen/core/old_code/db_code"
	"fmt"
	"testing"
)

func TestGenCode(t *testing.T) {
	err := gen_program.GenProgramCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	dbModel, err := mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	err = db_code2.GenModelsCode(dbModel)
	if err != nil {
		t.Fatal("model生成报错：", err)
	}
	err = db_code2.GenDaoCode(dbModel)
	if err != nil {
		t.Fatal("dao生成报错：", err)
	}
	err = db_code2.GenServicesCode(dbModel)
	if err != nil {
		t.Fatal("service生成报错：", err)
	}
	err = db_code2.GenApiCode(dbModel)
	if err != nil {
		t.Fatal("api生成报错：", err)
	}
	err = db_code2.GenRouterCode(dbModel)
	if err != nil {
		t.Fatal("router生成报错：", err)
	}
}
