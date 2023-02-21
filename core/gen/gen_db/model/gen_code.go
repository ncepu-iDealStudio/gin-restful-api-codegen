// coding: utf-8
// @Author : lryself
// @Date : 2022/3/6 14:19
// @Software: GoLand

package model

import (
	"GinCodeGen/core/database/mysql"
	"GinCodeGen/tools/common"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenTableCode(tableInfo *mysql.DataBaseModel, tmplPath string, outPath string, divideDir bool, filename string) error {
	exts := strings.Split(tmplPath, ".")
	ext := exts[len(exts)-2]
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
		var codeFilename string
		if strings.Index(filename, "%s") != -1 {
			codeFilename = fmt.Sprintf(filename, common.LineToUpCamel(table.TableName)) + fmt.Sprintf(".%s", ext)
		} else {
			codeFilename = filename + fmt.Sprintf(".%s", ext)
		}
		if divideDir {
			err = os.MkdirAll(filepath.Join(outPath, common.LineToLowCamel(table.TableName)), os.ModePerm)
			if err != nil {
				return err
			}
			file, err = os.Create(filepath.Join(outPath, common.LineToLowCamel(table.TableName), codeFilename))
		} else {
			err = os.MkdirAll(outPath, os.ModePerm)
			if err != nil {
				return err
			}
			file, err = os.Create(filepath.Join(outPath, codeFilename))
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
	err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
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
