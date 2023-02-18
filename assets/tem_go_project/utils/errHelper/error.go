// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 22:04
// @Software: GoLand

package errHelper

import (
	"GinCodeGen/globals/sys"
)

func Error(err error) {
	if err != nil {
		sys.PrintErr(err)
	}
}

func ErrExit(err error) {
	if err != nil {
		sys.PrintErr(err)
		sys.Exit()
	}
}

func ErrorFunc(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
