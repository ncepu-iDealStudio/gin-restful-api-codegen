// coding: utf-8
// @Author : lryself
// @Date : 2022/3/22 22:04
// @Software: GoLand

package errHelper

import (
<<<<<<< HEAD
	"GinCodeGen/globals/sys"
=======
	"tem_go_project/utils/message"
>>>>>>> develop
)

func Error(err error) {
	if err != nil {
<<<<<<< HEAD
		sys.PrintErr(err)
=======
		message.PrintErr(err)
>>>>>>> develop
	}
}

func ErrExit(err error) {
	if err != nil {
<<<<<<< HEAD
		sys.PrintErr(err)
		sys.Exit()
=======
		message.PrintErr(err)
		message.Exit()
>>>>>>> develop
	}
}

func ErrorFunc(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
