// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	model2 "LRYGoCodeGen/core/gen_program/model"
	"LRYGoCodeGen/core/globals/config"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func CopyCodeFromTemplates() error {
	userConfig, err := config.GetUserConfig()
	if err != nil {
		return err
	}
	dirModel, err := model2.GetDirModel(userConfig.GenCodeConfig.TemplatesPath)
	if err != nil {
		return err
	}
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(userConfig.GenCodeConfig.DictPath, "keyword.json"))
	var replaceDict model2.KeyWord
	err = json.Unmarshal(dictKeywordFile, &replaceDict)
	if err != nil {
		return err
	}
	err = dirModel.MakeDir(userConfig.GenCodeConfig.ResultPath, replaceDict)
	if err != nil {
		return err
	}
	return nil
}
