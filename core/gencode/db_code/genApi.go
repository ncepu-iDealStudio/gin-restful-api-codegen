// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:58
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils"
	"LRYGoCodeGen/core/utils/str"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenApiCode(tableInfo *genmysql.DataBaseModel) error {
	var dbCode DBCode
	err := dbCode.Init("apis.json")
	//构建替换字典
	var kwargs map[string]string
	kwargs = dbCode.Code.StaticDict
	fileDir := replaceString(kwargs, dbCode.Code.Filepath)
	//判断结果文件夹存在
	if !utils.PathExists(fileDir) {
		err = os.Mkdir(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["PackageName"] = str.LineToLowCamel(table.TableName)
		kwargs["StructName"] = str.LineToUpCamel(table.TableName)
		var model GenCodeModel
		model.Header = replaceString(kwargs, dbCode.Code.FileHeader)
		model.Import = []string{replaceString(kwargs, dbCode.Code.Import.Header), replaceString(kwargs, dbCode.Code.Import.Footer)}
		model.Methods = []string{replaceString(kwargs, dbCode.Code.Methods["crud"])}
		if !utils.PathExists(filepath.Join(fileDir, str.LineToLowCamel(table.TableName))) {
			err = os.Mkdir(filepath.Join(fileDir, str.LineToLowCamel(table.TableName)), os.ModePerm)
			if err != nil {
				return err
			}
		}
		err = ioutil.WriteFile(
			filepath.Join(fileDir, str.LineToLowCamel(table.TableName), fmt.Sprintf("%sResource.go", getTableName(table.TableName))),
			[]byte(model.String()), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
