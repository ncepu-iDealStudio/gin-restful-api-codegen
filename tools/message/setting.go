// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:21
// @Software: GoLand

package message

import (
	"fmt"
	"os"
	"sync"
)

func InitMsg(waitGroup *sync.WaitGroup) {
	var err error
	outputChan = make(chan message, 10)
	defer close(outputChan)
	for msg := range outputChan {
		switch msg.Type {
		case "exit":
			os.Exit(1)
		case "err":
			_, _ = fmt.Fprint(os.Stderr, msg.Context)
		case "msg":
			_, _ = fmt.Fprint(os.Stdout, msg.Context)
		}
		if err != nil {
			//panic("输出错误！")
			fmt.Printf("输出错误！")
		}

	}
	waitGroup.Done()
}
