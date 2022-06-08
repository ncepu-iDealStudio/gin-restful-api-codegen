package main

import (
	"gitee.com/lryself/go-utils/loggers"
	"github.com/spf13/viper"
	"tem_go_project/cmd"
	"tem_go_project/globals"
	"tem_go_project/globals/errHelper"
	"tem_go_project/globals/sys"
	"tem_go_project/globals/vipers"
)

func main() {
	waitGroup := globals.GetWaitGroup()
	waitGroup.Add(1)
	go sys.InitMsg()

	//初始化viper
	errHelper.ErrExit(vipers.InitViper())
	sys.Println("配置文件加载完成")

	//初始化日志
	loggers.InitLogger(viper.GetString("log.type"))
	sys.Println("日志组件初始化完成")

	//执行cmd
	cmd.Execute()

	sys.Println("系统已全部初始化完成！")
	//退出系统
	waitGroup.Wait()
	sys.Exit()
}
