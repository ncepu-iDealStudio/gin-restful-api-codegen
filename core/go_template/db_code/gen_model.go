// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 14:19
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/gencode/db_code/models"
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils/str"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type CodeTemplate struct {
	TypeDict  models.TypeDict
	TableInfo tableModel
	Dict      map[string]string
}

type tableModel struct {
	TableName   string
	StructName  string
	ProjectName string
	Columns     []columnModel
}

type columnModel struct {
	Field      string
	Type       string
	GoType     string
	Collation  *string
	Null       string
	Key        string
	Default    string
	Extra      string
	Privileges []string
	Comment    string
}

func (d *CodeTemplate) Init(table genmysql.TableModel) error {
	dictTypeDict, err := ioutil.ReadFile("configs/type_dict.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictTypeDict, &d.TypeDict)
	staticDict, err := ioutil.ReadFile("configs/static_dict.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(staticDict, &d.Dict)
	if err != nil {
		return err
	}
	d.TableInfo.TableName = table.TableName
	d.TableInfo.StructName = str.LineToUpCamel(table.TableName)
	d.TableInfo.ProjectName = str.LineToLowCamel(table.TableName)

	for _, column := range table.Columns {
		var column1 columnModel
		column1.Field = column.Field
		column1.Type = tagColumnType(column.Type)
		column1.GoType = d.TypeDict.GetGoType(column.Type)
		column1.Collation = column.Collation
		column1.Null = tagNull(column.Null)
		column1.Key = tagKey(column.Field)
		column1.Default = tagDefault(column.Default)
		column1.Extra = column.Extra
		column1.Privileges = column.Privileges
		column1.Comment = column.Comment
		d.TableInfo.Columns = append(d.TableInfo.Columns, column1)
	}
	return nil
}

func GenModelCode(tableInfo *genmysql.DataBaseModel) error {
	for _, table := range tableInfo.Tables {
		var codeTemplate CodeTemplate
		err := codeTemplate.Init(table)
		if err != nil {
			return err
		}
		tmpl, err := template.ParseFiles("./model.go.tpl")
		if err != nil {
			return err
		}
		file, err := os.Create(filepath.Join("result", fmt.Sprintf("%s.go", table.TableName)))
		if err != nil {
			return err
		}
		err = tmpl.Execute(file, codeTemplate)
		if err != nil {
			return err
		}
	}
	return nil
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
	return fmt.Sprintf("default:%s;", *s)
}
