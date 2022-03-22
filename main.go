// coding: utf-8
// @Author : lryself
// @Date : 2022/3/10 22:18
// @Software: GoLand

package main

import (
	"LRYGoCodeGen/cmd"
	"LRYGoCodeGen/globals/sys"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go sys.InitMsg(&waitGroup)

	cmd.Execute()

	sys.Exit()
	waitGroup.Wait()
}
