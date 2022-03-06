// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 15:50
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"fmt"
	"testing"
)

func TestGoTemplate(t *testing.T) {
	var err error
	dbModel, err := genmysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}
	err = GenModelCode(dbModel)
	if err != nil {
		fmt.Println(err)
	}
}
