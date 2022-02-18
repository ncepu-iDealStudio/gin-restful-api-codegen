// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 19:34
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/gencode"
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestGenDBCode(t *testing.T) {
	var err error
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
