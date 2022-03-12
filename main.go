// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"LRYGoCodeGen/core/gen/gen_db"
	"LRYGoCodeGen/core/gen/gen_program"
	"LRYGoCodeGen/core/globals/vipers"
	"LRYGoCodeGen/core/model/mysql"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
)

type makeFileDict struct {
	TmplPath  string `json:"tmplPath"`
	OutPath   string `json:"outPath"`
	DivideDir bool   `json:"divideDir"`
}

func main() {
	var err error

	err = vipers.InitViper()
	if err != nil {
		fmt.Println(err)
	}
	//生成项目代码
	err = gen_program.CopyCodeFromTemplates()
	if err != nil {
		fmt.Println(err)
	}
	dbModel, err := mysql.GetMysqlDBModel()
	if err != nil {
		fmt.Println(err)
	}

	tmplPath := viper.GetString("genCode.tmplPath")
	resultPath := viper.GetString("genCode.result_path")

	var makefiles []makeFileDict
	dictTypeDict, err := ioutil.ReadFile(filepath.Join(viper.GetString("genCode.dict_path"), "makefile.json"))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dictTypeDict, &makefiles)
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range makefiles {
		err = gen_db.GenCode(dbModel,
			filepath.Join(tmplPath, d.TmplPath),
			filepath.Join(resultPath, d.OutPath),
			d.DivideDir,
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	//生成url文件
	err = gen_db.GenUrlCode(dbModel,
		filepath.Join(tmplPath, "urls.go.tpl"),
		filepath.Join(resultPath, "internal/routers/api1_0"),
	)
	if err != nil {
		fmt.Println(err)
	}
}
