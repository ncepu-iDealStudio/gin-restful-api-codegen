// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"GinCodeGen/cmd"
	"GinCodeGen/tools/logger"
	"fmt"
	"os"
)

func main() {
	// 初始化日志模块
	logger.InitLogger(logger.FileType)

	// 命令行执行
	cmd.Execute()

	// 程序完成
	fmt.Println("执行完毕，程序退出!")
	os.Exit(0)
}
