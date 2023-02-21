// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"GinCodeGen/cmd"
	"fmt"
	"os"
)

func main() {
	// 命令行执行
	cmd.Execute()

	// 程序完成
	fmt.Println("执行完毕，程序推出!")
	os.Exit(0)
}
