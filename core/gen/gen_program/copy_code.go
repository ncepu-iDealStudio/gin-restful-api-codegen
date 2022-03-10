// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	"LRYGoCodeGen/core/gen/gen_program/model"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

func CopyCodeFromTemplates() error {
	dirModel, err := model.GetDirModel(viper.GetString("genCode.templates_path"))
	if err != nil {
		return err
	}
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "keyword.json"))
	var replaceDict model.KeyWord
	err = json.Unmarshal(dictKeywordFile, &replaceDict)
	if err != nil {
		return err
	}
	err = dirModel.MakeDir(viper.GetString("genCode.result_path"), replaceDict)
	if err != nil {
		return err
	}
	return nil
}
