// coding: utf-8
// @Author : lryself
// @Date : 2022/2/15 18:09
// @Software: GoLand

package db_code

import (
	"LRYGoCodeGen/core/globals/config"
	"LRYGoCodeGen/core/old_code/db_code/models"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DBCode struct {
	Code models.CodeModel
	Type models.TypeDict
}

func replaceString(args map[string]string, s string) string {
	for k, v := range args {
		s = strings.ReplaceAll(s, "{{"+k+"}}", v)
	}
	return s
}

func (d *DBCode) Init(fileName string) error {
	// 读取配置文件
	userConfig, err := config.GetUserConfig()
	if err != nil {
		return err
	}
	dictDBCode, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.CodeModelPath, fileName))
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictDBCode, &d.Code)

	dictTypeDict, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "typeDict.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(dictTypeDict, &d.Type)
	if err != nil {
		return err
	}
	return nil
}

type GenCodeModel struct {
	Header  string
	Import  []string
	Struct  []string
	Methods []string
}

func (g GenCodeModel) String() string {
	s := g.Header
	s += "\n"
	for _, s1 := range g.Import {
		s += s1
		s += "\n"
	}
	s += "\n"
	for _, s1 := range g.Struct {
		s += s1
		s += "\n"
	}
	s += "\n"
	for _, s1 := range g.Methods {
		s += s1
		s += "\n"
	}
	s += "\n"
	return s
}
