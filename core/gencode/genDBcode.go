// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 18:09
// @Software: GoLand

package gencode

import (
	"LRYGoCodeGen/core/globals/config"
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type dbCode struct {
	StaticDict map[string]string `json:"static_dict"`
	Models     struct {
		Filepath   string `json:"filepath,omitempty"`
		FileHeader string `json:"file_header,omitempty"`
		Struct     struct {
			TableHeader string `json:"table_header,omitempty"`
			Column      string `json:"column,omitempty"`
			TableFooter string `json:"table_footer,omitempty"`
		} `json:"struct"`
		Methods struct {
			TableName string `json:"table_name,omitempty"`
			GetSet    string `json:"get_set,omitempty"`
		} `json:"methods"`
	} `json:"models"`
	Dao struct {
		Filepath   string `json:"filepath,omitempty"`
		FileHeader string `json:"file_header,omitempty"`
		Struct     string `json:"struct"`
		Methods    struct {
			CRUD string `json:"CRUD"`
		} `json:"methods"`
	} `json:"dao"`
	Service struct {
		Filepath   string `json:"filepath,omitempty"`
		FileHeader string `json:"file_header,omitempty"`
		Struct     string `json:"struct"`
	} `json:"services"`
}

type typeDict struct {
	Accurate map[string]string `json:"accurate"`
	Fuzzy    map[string]string `json:"fuzzy"`
}

func GenModelsCode(tableInfo *genmysql.DataBaseModel) error {
	// 读取配置文件
	userConfig, err := config.GetUserConfig()
	if err != nil {
		return err
	}
	dictDBCode, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "dbCode.json"))
	if err != nil {
		return err
	}
	var dbCodeModel dbCode
	err = json.Unmarshal(dictDBCode, &dbCodeModel)

	dictTypeDict, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "typeDict.json"))
	if err != nil {
		return err
	}
	var typeDictModel typeDict
	err = json.Unmarshal(dictTypeDict, &typeDictModel)
	if err != nil {
		return err
	}
	//构建替换字典
	var kwargs map[string]string
	kwargs = dbCodeModel.StaticDict

	//判断结果文件夹存在
	if !utils.PathExists(dbCodeModel.Models.Filepath) {
		err = os.Mkdir(dbCodeModel.Models.Filepath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	var timeFlag bool
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["StructName"] = getTableName(table.TableName)
		timeFlag = false
		s := ""

		s = s + replaceString(kwargs, dbCodeModel.Models.Struct.TableHeader)

		for _, column := range table.Columns {
			goType := getGoType(column.Type, typeDictModel)
			if ok, _ := regexp.MatchString("^time\\..*", goType); ok {
				timeFlag = true
			}
			kwargs["ColumnName"] = column.Field
			kwargs["GoType"] = goType
			kwargs["ColumnKey"] = tagKey(column.Key)
			kwargs["ColumnType"] = tagColumnType(column.Type)
			kwargs["ColumnNull"] = tagNull(column.Null)
			kwargs["ColumnDefault"] = tagDefault(column.Default)
			s = s + replaceString(kwargs, dbCodeModel.Models.Struct.Column)
		}

		s = s + dbCodeModel.Models.Struct.TableFooter

		kwargs["TableName"] = table.TableName
		s = s + replaceString(kwargs, dbCodeModel.Models.Methods.TableName)
		for _, column := range table.Columns {
			kwargs["ColumnName"] = column.Field
			kwargs["GoType"] = getGoType(column.Type, typeDictModel)
			s = s + replaceString(kwargs, dbCodeModel.Models.Methods.GetSet)
		}
		if timeFlag {
			s = dbCodeModel.Models.FileHeader + "import \"time\"\n\n" + s
		} else {
			s = dbCodeModel.Models.FileHeader + s
		}
		err = ioutil.WriteFile(filepath.Join(dbCodeModel.Models.Filepath, fmt.Sprintf("%s.go", getTableName(table.TableName))), []byte(s), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenDaoCode(tableInfo *genmysql.DataBaseModel) error {
	// 读取配置文件
	userConfig, err := config.GetUserConfig()
	if err != nil {
		return err
	}
	dictDBCode, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "dbCode.json"))
	if err != nil {
		return err
	}
	var dbCodeModel dbCode
	err = json.Unmarshal(dictDBCode, &dbCodeModel)

	dictTypeDict, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "typeDict.json"))
	if err != nil {
		return err
	}
	var typeDictModel typeDict
	err = json.Unmarshal(dictTypeDict, &typeDictModel)
	if err != nil {
		return err
	}
	//构建替换字典
	var kwargs map[string]string
	kwargs = dbCodeModel.StaticDict

	//判断结果文件夹存在
	if !utils.PathExists(dbCodeModel.Dao.Filepath) {
		err = os.Mkdir(dbCodeModel.Dao.Filepath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["StructName"] = getTableName(table.TableName)
		s := ""

		s = s + replaceString(kwargs, dbCodeModel.Dao.Struct)

		s = s + replaceString(kwargs, dbCodeModel.Dao.Methods.CRUD)

		s = replaceString(kwargs, dbCodeModel.Dao.FileHeader) + s
		err = ioutil.WriteFile(filepath.Join(dbCodeModel.Dao.Filepath, fmt.Sprintf("%s.go", getTableName(table.TableName))), []byte(s), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenServicesCode(tableInfo *genmysql.DataBaseModel) error {
	// 读取配置文件
	userConfig, err := config.GetUserConfig()
	if err != nil {
		return err
	}
	dictDBCode, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "dbCode.json"))
	if err != nil {
		return err
	}
	var dbCodeModel dbCode
	err = json.Unmarshal(dictDBCode, &dbCodeModel)

	dictTypeDict, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "typeDict.json"))
	if err != nil {
		return err
	}
	var typeDictModel typeDict
	err = json.Unmarshal(dictTypeDict, &typeDictModel)
	if err != nil {
		return err
	}
	//构建替换字典
	var kwargs map[string]string
	kwargs = dbCodeModel.StaticDict

	//判断结果文件夹存在
	if !utils.PathExists(dbCodeModel.Dao.Filepath) {
		err = os.Mkdir(dbCodeModel.Dao.Filepath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	//生成代码
	for _, table := range tableInfo.Tables {
		kwargs["StructName"] = getTableName(table.TableName)
		s := ""

		s = s + replaceString(kwargs, dbCodeModel.Service.Struct)

		s = replaceString(kwargs, dbCodeModel.Service.FileHeader) + s
		err = ioutil.WriteFile(filepath.Join(dbCodeModel.Service.Filepath, fmt.Sprintf("%s.go", getTableName(table.TableName))), []byte(s), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func replaceString(args map[string]string, s string) string {
	for k, v := range args {
		s = strings.ReplaceAll(s, "{{"+k+"}}", v)
	}
	return s
}

func getGoType(name string, typeDictModel typeDict) string {
	// Precise matching first.先精确匹配
	if v, ok := typeDictModel.Accurate[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range typeDictModel.Fuzzy {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way", name))
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
