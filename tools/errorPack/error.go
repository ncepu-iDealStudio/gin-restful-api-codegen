package errorPack

import (
	"fmt"
	"os"
)

func Error(err error) {
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Sprintln(err))
	}
}

func ErrExit(err error) {
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Sprintln(err))
		os.Exit(1)
	}
}

func ErrorFunc(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
