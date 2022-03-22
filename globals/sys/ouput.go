// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:13
// @Software: GoLand

package sys

import "fmt"

var outputChan chan message

type message struct {
	Context string
	Type    string
}

func Println(a ...interface{}) {
	outputChan <- message{
		Context: fmt.Sprintln(a...),
		Type:    "msg",
	}
}

func PrintErr(a ...interface{}) {
	outputChan <- message{
		Context: fmt.Sprintln(a...),
		Type:    "err",
	}
}

func Printf(format string, a ...interface{}) {
	outputChan <- message{
		Context: fmt.Sprintf(format, a...),
		Type:    "msg",
	}
}

func Exit() {
	outputChan <- message{
		Context: "",
		Type:    "exit",
	}
}
