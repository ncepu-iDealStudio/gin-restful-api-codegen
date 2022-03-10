// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 19:01
// @Software: GoLand

package gen_db

import (
	"LRYGoCodeGen/core/model/mysql"
	"LRYGoCodeGen/core/utils/str"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type TypeDict struct {
	Accurate map[string]string `json:"accurate"`
	Fuzzy    map[string]string `json:"fuzzy"`
}

func (t *TypeDict) GetGoType(name string) string {
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

type CodeDict struct {
	TableInfo tableModel
	Dict      map[string]string
}

func (d *CodeDict) Init(table mysql.TableModel) error {
	var typeDict TypeDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "type_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictTypeDict, &typeDict)
	staticDict, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "static_dict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(staticDict, &d.Dict)
	if err != nil {
		return err
	}
	d.TableInfo.TableName = table.TableName
	d.TableInfo.StructName = str.LineToUpCamel(table.TableName)
	d.TableInfo.PackageName = str.LineToLowCamel(table.TableName)

	for _, column := range table.Columns {
		var column1 columnModel
		column1.Field = column.Field
		column1.Type = column.Type
		column1.GoType = typeDict.GetGoType(column.Type)
		column1.Collation = column.Collation
		column1.Null = column.Null
		column1.Key = column.Key
		column1.Default = column.Default
		column1.Extra = column.Extra
		column1.Privileges = column.Privileges
		column1.Comment = column.Comment
		if strings.ToLower(column.Field) != "autoid" && column.Key == "PRI" {
			d.TableInfo.NaturalKey = append(d.TableInfo.NaturalKey, column.Field)
		}
		d.TableInfo.Columns = append(d.TableInfo.Columns, column1)
	}
	return nil
}

type tableModel struct {
	TableName   string
	StructName  string
	PackageName string
	Columns     []columnModel
	NaturalKey  []string
}

type columnModel struct {
	Field      string
	Type       string
	GoType     string
	Collation  *string
	Null       bool
	Key        string
	Default    *string
	Extra      string
	Privileges []string
	Comment    string
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
