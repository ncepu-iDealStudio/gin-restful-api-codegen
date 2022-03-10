// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 17:35
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gen_mysql"
	"LRYGoCodeGen/core/gen_program"
	"fmt"
	"testing"
)

func TestDBModel(t *testing.T) {
	dbModel, err := gen_mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	println(dbModel)
}

func TestCopyCodeModel(t *testing.T) {
	err := gen_program.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
}
