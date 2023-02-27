// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 19:01
// @Software: GoLand

package model

import (
	"GinCodeGen/core/database/mysql"
	initialization "GinCodeGen/init"
	"GinCodeGen/tools/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type TypeDict struct {
	Accurate map[string]string `json:"accurate"`
	Fuzzy    map[string]string `json:"fuzzy"`
}

func (t *TypeDict) GetDataType(name string) string {
	// Precise matching first.先精确匹配
	if v, ok := t.Accurate[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range t.Fuzzy {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way", name))
}

type tableModel struct {
	TableName    string
	StructName   string
	PackageName  string
	Columns      []columnModel
	NaturalKey   []string
	HasTimeField bool
	TableType    string
}

type columnModel struct {
	FieldName  string
	Field      string
	Type       string
	DataType   string
	Collation  *string
	Null       bool
	Key        string
	Default    *string
	Extra      string
	Privileges []string
	Comment    string
	NaturalKey bool
}

func (c columnModel) TagNull() string {
	if !c.Null {
		return "not null;"
	}
	return ""
}

func (c columnModel) TagKey() string {
	if c.Key == "" {
		return ""
	}
	if c.Key == "PRI" {
		return "primaryKey;"
	}
	return fmt.Sprintf("%s;", c.Key)
}

func (c columnModel) TagColumnType() string {
	if c.Type == "" {
		return ""
	}
	return fmt.Sprintf("type:%s;", c.Type)
}

func (c columnModel) TagDefault() string {
	if c.Default == nil {
		return ""
	}
	return fmt.Sprintf("default:%s;", *c.Default)
}

type tableCodeDict struct {
	TableInfo tableModel
	Dict      map[string]string
}

func (d *tableCodeDict) Init(table *mysql.TableModel) error {
	var typeDict TypeDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "type_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictTypeDict, &typeDict)
	staticDict, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "static_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(staticDict, &d.Dict)
	d.Dict["ProjectName"] = initialization.GetCodeGenViper().GetGenViper().GetString("database.database")
	if err != nil {
		return err
	}
	d.TableInfo.TableName = table.TableName
	d.TableInfo.StructName = common.LineToUpCamel(table.TableName)
	d.TableInfo.PackageName = common.LineToLowCamel(table.TableName)

	for _, column := range table.Columns {
		var column1 columnModel
		column1.FieldName = common.LineToUpCamel(column.Field)
		column1.Field = column.Field
		column1.Type = column.Type
		column1.DataType = typeDict.GetDataType(column.Type)
		column1.Collation = column.Collation
		column1.Null = column.Null
		column1.Key = column.Key
		column1.Default = column.Default
		column1.Extra = column.Extra
		column1.Privileges = column.Privileges
		column1.Comment = column.Comment

		switch column1.Type {
		case "timestamp", "time", "smalldatetime", "datetime", "date":
			d.TableInfo.HasTimeField = true
		}

		//if strings.ToLower(column.Field) != "autoid" && column.Key == "PRI" {
		if column.Key == "PRI" {
			column1.NaturalKey = true
			d.TableInfo.NaturalKey = append(d.TableInfo.NaturalKey, column.Field)
		}
		d.TableInfo.Columns = append(d.TableInfo.Columns, column1)
	}
	return nil
}

type tablesCodeDict struct {
	TablesInfo []tableModel
	Dict       map[string]string
}

func (d *tablesCodeDict) Init(tables *mysql.DataBaseModel) error {
	var typeDict TypeDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "type_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictTypeDict, &typeDict)
	staticDict, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "static_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(staticDict, &d.Dict)
	d.Dict["ProjectName"] = initialization.GetCodeGenViper().GetGenViper().GetString("database.database")
	if err != nil {
		return err
	}

	for _, table := range tables.Tables {
		var table1 tableModel
		table1.TableName = table.TableName
		table1.StructName = common.LineToUpCamel(table.TableName)
		table1.PackageName = common.LineToLowCamel(table.TableName)
		table1.TableType = table.TableType

		for _, column := range table.Columns {
			var column1 columnModel
			column1.Field = column.Field
			column1.Type = column.Type
			column1.DataType = typeDict.GetDataType(column.Type)
			column1.Collation = column.Collation
			column1.Null = column.Null
			column1.Key = column.Key
			column1.Default = column.Default
			column1.Extra = column.Extra
			column1.Privileges = column.Privileges
			column1.Comment = column.Comment

			switch column1.Type {
			case "timestamp", "time", "smalldatetime", "datetime", "date":
				table1.HasTimeField = true
			}

			if strings.ToLower(column.Field) != "autoid" && column.Key == "PRI" {
				table1.NaturalKey = append(table1.NaturalKey, column.Field)
			}
			table1.Columns = append(table1.Columns, column1)
		}
		d.TablesInfo = append(d.TablesInfo, table1)
	}
	return nil
}
