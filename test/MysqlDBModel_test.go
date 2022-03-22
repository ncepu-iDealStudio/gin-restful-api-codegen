// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 17:35
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/core/model/mysql"
	"fmt"
	"testing"
)

func TestDBModel(t *testing.T) {
	dbModel, err := mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	println(dbModel)
}

func TestCopyCodeModel(t *testing.T) {
	err := gen_program.GenProgramCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
}
