// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	"GinCodeGen/core/gen/gen_program/model"
	"GinCodeGen/globals/vipers"
	"GinCodeGen/utils/errHelper"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func GenProgramCodeFromTemplates() {
	codeGenViper := vipers.GetCodeGenViper()
	dirModel, err := model.GetDirModel(codeGenViper.GetSysViper().GetString("genCode.templates_path"))
	errHelper.ErrExit(err)
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(codeGenViper.GetSysViper().GetString("genCode.dict_path"), "keyword.json"))
	errHelper.ErrExit(err)
	var replaceDict model.KeyWord
	errHelper.ErrExit(json.Unmarshal(dictKeywordFile, &replaceDict))
	errHelper.ErrExit(dirModel.MakeDir(codeGenViper.GetGenViper().GetString("genCode.result_path"), replaceDict))
}
