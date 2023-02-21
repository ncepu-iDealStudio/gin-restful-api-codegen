// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 23:24
// @Software: GoLand

package gen_db

import (
	"GinCodeGen/core/database/mysql"
	"GinCodeGen/core/gen/gen_db/model"
	initialization "GinCodeGen/init"
	"GinCodeGen/tools/errorPack"
	"fmt"

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
	errorPack.ErrExit(err)
	codeGenViper := initialization.GetCodeGenViper()
	tmplPath := initialization.TemplatesPath
	resultPath := fmt.Sprintf("dist/%s", codeGenViper.GetGenViper().GetString("database.database"))

	var makefiles []makeFileDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(initialization.DictPath, "makefile.json"))
	errorPack.ErrExit(err)
	errorPack.Error(json.Unmarshal(dictTypeDict, &makefiles))

	for _, d := range makefiles {
		if d.IsTables {
			errorPack.Error(model.GenTablesCode(dbModel,
				filepath.Join(tmplPath, d.TmplPath),
				filepath.Join(resultPath, d.OutPath),
			))
		} else {
			errorPack.Error(model.GenTableCode(dbModel,
				filepath.Join(tmplPath, d.TmplPath),
				filepath.Join(resultPath, d.OutPath),
				d.DivideDir,
				d.GetFileName(),
			))
		}
	}
}
