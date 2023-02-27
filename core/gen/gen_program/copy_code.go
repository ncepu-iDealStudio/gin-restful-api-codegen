// coding: utf-8
// @Author : lryself
// @Date : 2022/2/14 20:05
// @Software: GoLand

package gen_program

import (
	"GinCodeGen/core/gen/gen_program/model"
	initialization "GinCodeGen/init"
	"GinCodeGen/tools/errorPack"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GenProgramCodeFromTemplates() {
	codeGenViper := initialization.GetCodeGenViper()
	dirModel, err := model.GetDirModel(initialization.TemplatesPath)
	errorPack.ErrExit(err)
	dictKeywordFile, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "keyword.json"))
	errorPack.ErrExit(err)
	var replaceDict model.KeyWord
	errorPack.ErrExit(json.Unmarshal(dictKeywordFile, &replaceDict))
	replaceDict.Replace["tem_go_project"] = codeGenViper.GetGenViper().GetString("database.database")
	replaceDict.Replace["mysql_host"] = codeGenViper.GetGenViper().GetString("database.host")
	replaceDict.Replace["mysql_port"] = codeGenViper.GetGenViper().GetString("database.port")
	replaceDict.Replace["mysql_username"] = codeGenViper.GetGenViper().GetString("database.username")
	replaceDict.Replace["mysql_password"] = codeGenViper.GetGenViper().GetString("database.password")
	replaceDict.Replace["mysql_database"] = codeGenViper.GetGenViper().GetString("database.database")
	replaceDict.Replace["redis_host"] = codeGenViper.GetGenViper().GetString("redis.host")
	replaceDict.Replace["redis_password"] = codeGenViper.GetGenViper().GetString("redis.password")
	errorPack.ErrExit(dirModel.MakeDir(fmt.Sprintf("dist/%s", codeGenViper.GetGenViper().GetString("database.database")), replaceDict))
}
