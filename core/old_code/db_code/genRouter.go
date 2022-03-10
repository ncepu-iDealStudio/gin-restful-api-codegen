// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:59
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/gen_mysql"
	"LRYGoCodeGen/core/utils"
	"LRYGoCodeGen/core/utils/str"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenRouterCode(tableInfo *gen_mysql.DataBaseModel) error {
	var dbcodeRouter DBCode
	var dbcodeUrl DBCode
	err := dbcodeRouter.Init("router.json")
	if err != nil {
		return err
	}
	err = dbcodeUrl.Init("urls.json")
	if err != nil {
		return err
	}
	//构建替换字典
	kwargs := dbcodeRouter.Code.StaticDict
	fileDir := replaceString(kwargs, dbcodeRouter.Code.Filepath)

	//判断结果文件夹存在
	if !utils.PathExists(fileDir) {
		err = os.Mkdir(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	var modelUrl GenCodeModel
	modelUrl.Header = replaceString(kwargs, dbcodeUrl.Code.FileHeader)
	modelUrl.Import = []string{replaceString(kwargs, dbcodeUrl.Code.Import.Header)}
	modelUrl.Struct = []string{replaceString(kwargs, dbcodeUrl.Code.Struct.Header)}
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["PackageName"] = str.LineToLowCamel(table.TableName)
		kwargs["StructName"] = str.LineToUpCamel(table.TableName)
		var modelRouter GenCodeModel
		modelRouter.Header = replaceString(kwargs, dbcodeRouter.Code.FileHeader)

		modelRouter.Import = []string{replaceString(kwargs, dbcodeRouter.Code.Import.Header)}
		modelRouter.Import = append(modelRouter.Import, replaceString(kwargs, dbcodeRouter.Code.Import.Row))
		modelRouter.Import = append(modelRouter.Import, replaceString(kwargs, dbcodeRouter.Code.Import.Footer))

		modelRouter.Struct = []string{replaceString(kwargs, dbcodeRouter.Code.Struct.Header)}
		modelRouter.Struct = append(modelRouter.Struct, replaceString(kwargs, dbcodeRouter.Code.Struct.Row))
		modelRouter.Struct = append(modelRouter.Struct, replaceString(kwargs, dbcodeRouter.Code.Struct.Footer))

		if !utils.PathExists(filepath.Join(fileDir, str.LineToLowCamel(table.TableName))) {
			err = os.Mkdir(filepath.Join(fileDir, str.LineToLowCamel(table.TableName)), os.ModePerm)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(fileDir, str.LineToLowCamel(table.TableName), "urls.go"), []byte(modelRouter.String()), os.ModePerm)
		if err != nil {
			return err
		}

		modelUrl.Import = append(modelUrl.Import, replaceString(kwargs, dbcodeUrl.Code.Import.Row))
		modelUrl.Struct = append(modelUrl.Struct, replaceString(kwargs, dbcodeUrl.Code.Struct.Row))
	}

	modelUrl.Import = append(modelUrl.Import, replaceString(kwargs, dbcodeUrl.Code.Import.Footer))
	modelUrl.Struct = append(modelUrl.Struct, replaceString(kwargs, dbcodeUrl.Code.Struct.Footer))
	err = ioutil.WriteFile(filepath.Join(fileDir, "urls.go"), []byte(modelUrl.String()), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
