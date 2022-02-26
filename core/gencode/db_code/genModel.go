// coding: utf-8
// @Author : lryself
// @Date : 2022/2/26 11:56
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GenModelsCode(tableInfo *genmysql.DataBaseModel) error {
	var dbCode DBCode
	err := dbCode.Init("models.json")
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

	var timeFlag bool
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["StructName"] = getTableName(table.TableName)
		timeFlag = false
		var model GenCodeModel
		model.Header = replaceString(kwargs, dbCode.Code.FileHeader)
		model.Import = []string{replaceString(kwargs, dbCode.Code.Import.Header)}
		model.Struct = []string{replaceString(kwargs, dbCode.Code.Struct.Header)}

		for _, column := range table.Columns {
			goType := dbCode.Type.GetGoType(column.Type)
			if ok, _ := regexp.MatchString("^time\\..*", goType); ok {
				timeFlag = true
			}
			kwargs["ColumnName"] = column.Field
			kwargs["GoType"] = goType
			kwargs["ColumnKey"] = tagKey(column.Key)
			kwargs["ColumnType"] = tagColumnType(column.Type)
			kwargs["ColumnNull"] = tagNull(column.Null)
			kwargs["ColumnDefault"] = tagDefault(column.Default)
			model.Struct = append(model.Struct, replaceString(kwargs, dbCode.Code.Struct.Row))
		}
		model.Struct = append(model.Struct, replaceString(kwargs, dbCode.Code.Struct.Footer))

		kwargs["TableName"] = table.TableName
		model.Methods = []string{replaceString(kwargs, dbCode.Code.Methods["table_name"])}
		for _, column := range table.Columns {
			kwargs["ColumnName"] = column.Field
			kwargs["GoType"] = dbCode.Type.GetGoType(column.Type)
			model.Methods = append(model.Methods, replaceString(kwargs, dbCode.Code.Methods["get_set"]))
		}
		if timeFlag {
			model.Import = append(model.Import, "import \"time\"")
		}
		model.Import = append(model.Import, replaceString(kwargs, dbCode.Code.Import.Footer))

		err = ioutil.WriteFile(filepath.Join(fileDir, fmt.Sprintf("%s.go", getTableName(table.TableName))), []byte(model.String()), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTableName(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.Title(name)
	return strings.ReplaceAll(name, " ", "")
}

func getKey(name string) string {
	if name == "PRI" {
		return "primaryKey"
	}
	return ""
}

func tagKey(s string) string {
	if s == "" {
		return ""
	}
	return fmt.Sprintf("%s;", getKey(s))
}

func tagColumnName(s string) string {
	if s == "" {
		return ""
	}
	return fmt.Sprintf("column:%s;", s)
}

func tagColumnType(s string) string {
	if s == "" {
		return ""
	}
	return fmt.Sprintf("type:%s;", s)
}

func tagNull(s bool) string {
	if !s {
		return "not null;"
	}
	return ""
}

func tagDefault(s *string) string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("default:%s;", s)
}
