// coding: utf-8
// @Author : lryself
// @Date : 2022/2/12 17:35
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gencode"
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestDBModel(t *testing.T) {
	dbModel, err := genmysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	println(dbModel)
}

func TestCopyCodeModel(t *testing.T) {
	err := gencode.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
}
