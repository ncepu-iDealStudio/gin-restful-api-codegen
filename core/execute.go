// coding: utf-8
// @Author : lryself
// @Date : 2022/3/20 23:14
// @Software: GoLand

package core

import (
	"LRYGoCodeGen/core/gen/gen_db"
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/core/model/mysql"
	"LRYGoCodeGen/utils/errHelper"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

type makeFileDict struct {
	TmplPath  string `json:"tmplPath"`
	OutPath   string `json:"outPath"`
	DivideDir bool   `json:"divideDir"`
}

func Execute() {
	var err error

	//生成项目代码
	errHelper.ErrExit(gen_program.CopyCodeFromTemplates())
	dbModel, err := mysql.GetMysqlDBModel()
	errHelper.ErrExit(err)

	tmplPath := viper.GetString("genCode.tmplPath")
	resultPath := viper.GetString("genCode.result_path")

	var makefiles []makeFileDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "makefile.json"))
	errHelper.ErrExit(err)
	errHelper.Error(json.Unmarshal(dictTypeDict, &makefiles))
	for _, d := range makefiles {
		errHelper.Error(gen_db.GenCode(dbModel,
			filepath.Join(tmplPath, d.TmplPath),
			filepath.Join(resultPath, d.OutPath),
			d.DivideDir,
		))
	}

	//生成url文件
	errHelper.Error(gen_db.GenUrlCode(dbModel,
		filepath.Join(tmplPath, "urls.go.tpl"),
		filepath.Join(resultPath, "internal/routers/api1_0"),
	))
}
