// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:58
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenServicesCode(tableInfo *genmysql.DataBaseModel) error {
	var dbCode DBCode
	err := dbCode.Init("service.json")
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
	for _, table := range tableInfo.Tables {
		kwargs["StructName"] = getTableName(table.TableName)
		var model GenCodeModel
		model.Header = replaceString(kwargs, dbCode.Code.FileHeader)
		model.Import = []string{replaceString(kwargs, dbCode.Code.Import.Header), replaceString(kwargs, dbCode.Code.Import.Footer)}
		model.Struct = []string{replaceString(kwargs, dbCode.Code.Struct.Header), replaceString(kwargs, dbCode.Code.Struct.Footer)}
		err = ioutil.WriteFile(filepath.Join(fileDir, fmt.Sprintf("%s.go", getTableName(table.TableName))), []byte(model.String()), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
