// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:21
// @Software: GoLand

package sys

import (
	"fmt"
	"os"
	"sync"
)

func InitMsg(waitGroup *sync.WaitGroup) {
	var err error
	outputChan = make(chan message, 10)
	defer close(outputChan)
	for mes := range outputChan {
		if mes.Type == "exit" {
			break
		} else if mes.Type == "err" {
			_, err = fmt.Fprintln(os.Stderr, mes.Context)
		} else if mes.Type == "ln" {
			_, err = fmt.Fprintln(os.Stdout, mes.Context)
		}
		if err != nil {
			panic("输出错误！")
		}
	}
	waitGroup.Done()
}
