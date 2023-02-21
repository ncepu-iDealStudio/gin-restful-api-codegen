// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"GinCodeGen/cmd"
	"GinCodeGen/tools/message"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go message.InitMsg(&waitGroup)

	cmd.Execute()

	message.Exit()
	waitGroup.Wait()
}
