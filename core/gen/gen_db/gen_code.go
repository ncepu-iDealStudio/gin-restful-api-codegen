// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 23:24
// @Software: GoLand

package gen_db

import (
	"LRYGoCodeGen/core/gen/gen_db/model"
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
	IsTables  bool   `json:"isTables"`
}

func GenDBCodeFromTemplate() {
	dbModel, err := mysql.GetMysqlDBModel()
	errHelper.ErrExit(err)

	tmplPath := viper.GetString("genCode.tmplPath")
	resultPath := viper.GetString("genCode.result_path")

	var makefiles []makeFileDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "makefile.json"))
	errHelper.ErrExit(err)
	errHelper.Error(json.Unmarshal(dictTypeDict, &makefiles))
	for _, d := range makefiles {
		if d.IsTables {
			errHelper.Error(model.GenTablesCode(dbModel,
				filepath.Join(tmplPath, d.TmplPath),
				filepath.Join(resultPath, d.OutPath),
			))
		} else {
			errHelper.Error(model.GenTableCode(dbModel,
				filepath.Join(tmplPath, d.TmplPath),
				filepath.Join(resultPath, d.OutPath),
				d.DivideDir,
			))
		}
	}
}
