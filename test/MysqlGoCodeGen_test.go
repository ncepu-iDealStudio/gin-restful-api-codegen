// coding: utf-8
// @Author : lryself
// @Date : 2022/2/18 18:16
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gencode"
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestGenCode(t *testing.T) {
	err := gencode.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
	dbModel, err := genmysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	err = gencode.GenModelsCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
	err = gencode.GenDaoCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
	err = gencode.GenServicesCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
}
