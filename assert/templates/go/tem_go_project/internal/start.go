// coding: utf-8
// @Author : lryself
// @Date : 2022/3/31 21:28
// @Software: GoLand

package internal

import (
	"encoding/gob"
	"fmt"
	"github.com/spf13/viper"
	"tem_go_project/globals"
	"tem_go_project/globals/sys"
	"tem_go_project/internal/globals/extensions/currentUser"
	"tem_go_project/internal/rpcServer"
	"tem_go_project/internal/settings"
	"time"
)

func StartHttp() {
	defer globals.GetWatGroup().Done()
	gob.Register(time.Time{})
	gob.Register(currentUser.CurrentUser{})
	var err error

	//初始化数据库（mysql、redis）
	err = settings.InitDatabase()
	if err != nil {
		sys.PrintErr(err)
		return
	}

	//初始化gin引擎
	engine, err := settings.InitGinEngine()
	if err != nil {
		sys.PrintErr(err)
		return
	}

	//开启rpc服务
	go rpcServer.StartRPCEngine()

	//开始运行
	err = engine.Run(fmt.Sprintf("%s:%s", viper.GetString("system.SysIP"), viper.GetString("system.SysPort")))
	if err != nil {
		sys.PrintErr(err)
		return
	}
}
