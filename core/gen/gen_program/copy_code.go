// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	"LRYGoCodeGen/core/gen/gen_program/model"
	"LRYGoCodeGen/globals/vipers"
	"LRYGoCodeGen/utils/errHelper"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func GenProgramCodeFromTemplates() {
	genViper := vipers.GetGenViper()
	dirModel, err := model.GetDirModel(genViper.GetString("genCode.templates_path"))
	errHelper.ErrExit(err)
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(genViper.GetString("genCode.dict_path"), "keyword.json"))
	errHelper.ErrExit(err)
	var replaceDict model.KeyWord
	errHelper.ErrExit(json.Unmarshal(dictKeywordFile, &replaceDict))
	errHelper.ErrExit(dirModel.MakeDir(genViper.GetString("genCode.result_path"), replaceDict))
}
