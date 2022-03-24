// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"LRYGoCodeGen/internal/settings"
	"encoding/gob"
	"fmt"
	logs "gitee.com/lryself/go-utils/loggers"
	"github.com/spf13/viper"
	"time"
)

func main() {
	gob.Register(time.Time{})
	var err error
	//初始化viper
	err = settings.InitViper()
	if err != nil {
		fmt.Println("配置文件加载出错！", err)
		return
	}
	logs.InitLogger(viper.GetString("log.type"))
	var log = logs.GetLogger()

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
