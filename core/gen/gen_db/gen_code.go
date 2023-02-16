// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 23:24
// @Software: GoLand

package gen_db

import (
	"GinCodeGen/core/database/mysql"
	"GinCodeGen/core/gen/gen_db/model"
	"GinCodeGen/globals/vipers"
	"GinCodeGen/utils/errHelper"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type makeFileDict struct {
	TmplPath  string `json:"tmplPath"`
	OutPath   string `json:"outPath"`
	DivideDir bool   `json:"divideDir"`
	IsTables  bool   `json:"isTables"`
	Filename  string `json:"filename"`
}

func (m makeFileDict) GetFileName() string {
	if m.Filename == "" {
		return "%s"
	}
	return m.Filename
}

func GenDBCodeFromTemplate() {
	dbModel, err := mysql.GetMysqlDBModel()
	errHelper.ErrExit(err)
	codeGenViper := vipers.GetCodeGenViper()
	tmplPath := codeGenViper.GetSysViper().GetString("genCode.tmplPath")
	resultPath := codeGenViper.GetGenViper().GetString("genCode.result_path")

	var makefiles []makeFileDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(codeGenViper.GetSysViper().GetString("genCode.dict_path"), "makefile.json"))
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
				d.GetFileName(),
			))
		}
	}
}
