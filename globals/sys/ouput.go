// coding: utf-8
// @Author : lryself
// @Date : 2022/3/21 22:13
// @Software: GoLand

package sys

var outputChan chan message

type message struct {
	Context string
	Type    string
}

func Println(s string) {
	outputChan <- message{
		Context: s,
		Type:    "ln",
	}
}

func PrintErr(s string) {
	outputChan <- message{
		Context: s,
		Type:    "err",
	}
}

func Exit() {
	outputChan <- message{
		Context: "",
		Type:    "exit",
	}
}
