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
	replaceDict.Replace["tem_go_project"] = codeGenViper.GetGenViper().GetString("database.database")
	replaceDict.Replace["mysql_host"] = codeGenViper.GetGenViper().GetString("database.host")
	replaceDict.Replace["mysql_port"] = codeGenViper.GetGenViper().GetString("database.port")
	replaceDict.Replace["mysql_username"] = codeGenViper.GetGenViper().GetString("database.username")
	replaceDict.Replace["mysql_password"] = codeGenViper.GetGenViper().GetString("database.password")
	replaceDict.Replace["mysql_database"] = codeGenViper.GetGenViper().GetString("database.database")
	replaceDict.Replace["redis_host"] = codeGenViper.GetGenViper().GetString("redis.host")
	replaceDict.Replace["redis_password"] = codeGenViper.GetGenViper().GetString("redis.password")
	errHelper.ErrExit(dirModel.MakeDir(codeGenViper.GetGenViper().GetString("genCode.result_path"), replaceDict))
}
