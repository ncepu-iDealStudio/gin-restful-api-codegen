// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:21
// @Software: GoLand

package sys

import (
	"fmt"
	"gitee.com/lryself/go-utils/loggers"
	"github.com/fufuok/chanx"
	"os"
	"sync"
	"tem_go_project/globals"
)

var outputChan *chanx.UnboundedChan[message]
var msgMap map[string]func(...any) error
var rwMutex sync.RWMutex
var msgMapOnce sync.Once

func InitMsg() {
	rwMutex.Lock()
	defer globals.GetWatGroup().Done()
	var err error
	outputChan = chanx.NewUnboundedChan[message](10, 0)
	initMsgHandler()
	rwMutex.Unlock()
	for msg := range outputChan.Out {
		err = msgMap[msg.Type](msg.Context)
		if err != nil {
			panic("输出错误！")
		}
	}
	_, _ = fmt.Fprintf(os.Stdout, "\\033[1;37;40m%s\\033[0m\\n", "系统服务已结束")
	os.Exit(1)
}
func AddMsgHandler(msg string, f func(args ...any) error) {
	msgMapOnce.Do(func() {
		msgMap = map[string]func(...any) error{}
	})
	msgMap[msg] = f
}

func initMsgHandler() {
	AddMsgHandler("exit", func(args ...any) error {
		var log = loggers.GetLogger()
		rwMutex.Lock()
		log.Infoln("程序终止！")
		close(outputChan.In)
		return nil
	})
	AddMsgHandler("info", func(args ...any) error {
		var log = loggers.GetLogger()
		log.Infoln(args)
		var err error
		for _, arg := range args {
			_, err = fmt.Fprint(os.Stdout, arg)
			if err != nil {
				return err
			}
		}
		return nil
	})
	AddMsgHandler("err", func(args ...any) error {
		var log = loggers.GetLogger()
		log.Errorln(args)
		var err error
		for _, arg := range args {
			_, err = fmt.Fprint(os.Stderr, arg)
			if err != nil {
				return err
			}
		}
		return nil
	})
	AddMsgHandler("warn", func(args ...any) error {
		var log = loggers.GetLogger()
		log.Warnln(args)
		var err error
		for _, arg := range args {
			_, err = fmt.Fprintf(os.Stdout, "\\033[1;37;40m%s\\033[0m\\n", arg)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
