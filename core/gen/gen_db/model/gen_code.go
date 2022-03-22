// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 14:19
// @Software: GoLand

package model

import (
	"LRYGoCodeGen/core/model/mysql"
	"LRYGoCodeGen/utils"
	"LRYGoCodeGen/utils/str"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenTableCode(tableInfo *mysql.DataBaseModel, tmplPath string, outPath string, divideDir bool) error {
	for _, table := range tableInfo.Tables {
		var codeTemplate tableCodeDict
		err := codeTemplate.Init(&table)
		if err != nil {
			return err
		}
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			return err
		}
		var file *os.File
		if divideDir {
			if !utils.PathExists(filepath.Join(outPath, str.LineToLowCamel(table.TableName))) {
				err = os.Mkdir(filepath.Join(outPath, str.LineToLowCamel(table.TableName)), os.ModePerm)
				if err != nil {
					return err
				}
			}
			file, err = os.Create(filepath.Join(outPath, str.LineToLowCamel(table.TableName), fmt.Sprintf("%s.go", str.LineToUpCamel(table.TableName))))
		} else {
			file, err = os.Create(filepath.Join(outPath, fmt.Sprintf("%s.go", str.LineToUpCamel(table.TableName))))
		}
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
func GenTablesCode(tablesInfo *mysql.DataBaseModel, tmplPath string, outPath string) error {
	var codeDict tablesCodeDict
	err := codeDict.Init(tablesInfo)
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, codeDict)
	if err != nil {
		return err
	}
	return nil
}
