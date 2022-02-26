// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:59
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils"
	"LRYGoCodeGen/core/utils/str"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenRouterCode(tableInfo *genmysql.DataBaseModel) error {
	var dbCode DBCode
	err := dbCode.Init("router.json")
	//构建替换字典
	kwargs := dbCode.Code.StaticDict
	fileDir := replaceString(kwargs, dbCode.Code.Filepath)

	//判断结果文件夹存在
	if !utils.PathExists(fileDir) {
		err = os.Mkdir(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	//生成代码
	var model GenCodeModel
	model.Header = replaceString(kwargs, dbCode.Code.FileHeader)
	model.Import = []string{replaceString(kwargs, dbCode.Code.Import.Header), replaceString(kwargs, dbCode.Code.Import.Footer)}
	model.Struct = []string{replaceString(kwargs, dbCode.Code.Struct.Header)}
	for _, table := range tableInfo.Tables {
		kwargs["PackageName"] = str.LineToLowCamel(table.TableName)
		kwargs["StructName"] = str.LineToUpCamel(table.TableName)
		model.Struct = append(model.Struct, replaceString(kwargs, dbCode.Code.Struct.Row))
	}
	model.Struct = append(model.Struct, replaceString(kwargs, dbCode.Code.Struct.Footer))
	err = ioutil.WriteFile(filepath.Join(fileDir, "urls.go"), []byte(model.String()), os.ModePerm)
	return nil
}
