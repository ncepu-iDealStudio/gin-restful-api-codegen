package main

import (
	"tem_go_project/cmd"
	"tem_go_project/utils/message"

	initialization "tem_go_project/init"
)

func main() {
	waitGroup := initialization.Init()

	// 执行cmd
	cmd.Execute()
	message.Println("系统已全部初始化完成！")

	// 退出系统
	waitGroup.Wait()
	message.Exit()
}
