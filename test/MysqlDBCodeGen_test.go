// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 19:34
// @Software: GoLand

package test

import (
	"LRYGoCodeGen/core/model/mysql"
	db_code2 "LRYGoCodeGen/core/old_code/db_code"
	"fmt"
	"testing"
)

func TestGenDBCode(t *testing.T) {
	var err error
	dbModel, err := mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	err = db_code2.GenModelsCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
	err = db_code2.GenDaoCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
	err = db_code2.GenServicesCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
}
