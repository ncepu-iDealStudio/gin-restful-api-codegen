// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	"LRYGoCodeGen/core/gen/gen_program/model"
	"LRYGoCodeGen/utils/errHelper"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

func GenProgramCodeFromTemplates() {
	dirModel, err := model.GetDirModel(viper.GetString("genCode.templates_path"))
	errHelper.ErrExit(err)
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "keyword.json"))
	errHelper.ErrExit(err)
	var replaceDict model.KeyWord
	errHelper.ErrExit(json.Unmarshal(dictKeywordFile, &replaceDict))
	errHelper.ErrExit(dirModel.MakeDir(viper.GetString("genCode.result_path"), replaceDict))
}
