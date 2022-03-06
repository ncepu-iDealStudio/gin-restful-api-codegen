// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 19:39
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/model/genmysql"
	"LRYGoCodeGen/core/utils/str"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type UrlCodeDict struct {
	ProjectName string
	Tables      []map[string]string
}

func GenUrlCode(tableInfo *genmysql.DataBaseModel, tmplPath string, outPath string) error {
	var tmp map[string]string
	staticDict, err := ioutil.ReadFile("assert/templates/go/dict/static_dict.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(staticDict, &tmp)
	if err != nil {
		return err
	}

	var codeDict UrlCodeDict
	codeDict.ProjectName = tmp["ProjectName"]

	for _, table := range tableInfo.Tables {
		d := map[string]string{}
		d["TableName"] = table.TableName
		d["PackageName"] = str.LineToLowCamel(table.TableName)
		d["StructName"] = str.LineToUpCamel(table.TableName)
		codeDict.Tables = append(codeDict.Tables, d)
	}
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(outPath, "urls.go"))
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, codeDict)
	if err != nil {
		return err
	}
	return nil
}
