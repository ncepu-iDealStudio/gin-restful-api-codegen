// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:13
// @Software: GoLand

package sys

import (
	"fmt"
)

type message struct {
	Context string
	Type    string
}

func Println(a ...interface{}) {
	if rwMutex.TryRLock() {
		defer rwMutex.RUnlock()
		outputChan.In <- message{
			Context: fmt.Sprintln(a...),
			Type:    "info",
		}
	}
}

func PrintWarn(a ...interface{}) {
	if rwMutex.TryRLock() {
		defer rwMutex.RUnlock()
		outputChan.In <- message{
			Context: fmt.Sprintln(a...),
			Type:    "warn",
		}
	}
}

func PrintErr(a ...interface{}) {
	if rwMutex.TryRLock() {
		defer rwMutex.RUnlock()
		outputChan.In <- message{
			Context: fmt.Sprintln(a...),
			Type:    "err",
		}
	}
}

func Printf(format string, a ...interface{}) {
	if rwMutex.TryRLock() {
		defer rwMutex.RUnlock()
		outputChan.In <- message{
			Context: fmt.Sprintf(format, a...),
			Type:    "msg",
		}
	}
}

func Exit() {
	if rwMutex.TryRLock() {
		defer rwMutex.RUnlock()
		outputChan.In <- message{
			Context: "",
			Type:    "exit",
		}
	}
}
