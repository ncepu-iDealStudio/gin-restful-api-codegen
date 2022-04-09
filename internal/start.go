// coding: utf-8
// @Author : lryself
// @Date : 2022/3/31 21:28
// @Software: GoLand

package internal

import (
	"LRYGoCodeGen/globals/sys"
	"LRYGoCodeGen/internal/models/ginModels"
	"LRYGoCodeGen/internal/settings"
	"encoding/gob"
	"fmt"
	"gitee.com/lryself/go-utils/loggers"
	"github.com/spf13/viper"
	"time"
)

func StartHttp() {
	gob.Register(time.Time{})
	gob.Register(ginModels.UserModel{})
	var err error
	//初始化viper
	err = settings.InitViper()
	if err != nil {
		sys.PrintErr("配置文件加载出错！", err)
		return
	}
	loggers.InitLogger(viper.GetString("log.type"))
	var log = loggers.GetLogger()

	//初始化数据库（mysql、redis）
	err = settings.InitDatabase()
	if err != nil {
		log.Errorln(err)
		return
	}

	//初始化gin引擎
	engine, err := settings.InitGinEngine()
	if err != nil {
		log.Errorln(err)
		return
	}
	//开始运行
	err = engine.Run(fmt.Sprintf("%s:%s", viper.GetString("system.SysIP"), viper.GetString("system.SysPort")))
	if err != nil {
		log.Errorln(err)
		return
	}
}
