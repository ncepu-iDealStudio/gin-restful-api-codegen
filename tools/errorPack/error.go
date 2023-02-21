package errorPack

import (
	"GinCodeGen/tools/message"
)

func Error(err error) {
	if err != nil {
		message.PrintErr(err)
	}
}

func ErrExit(err error) {
	if err != nil {
		message.PrintErr(err)
		message.Exit()
	}
}

func ErrorFunc(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
