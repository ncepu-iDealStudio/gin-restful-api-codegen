package errorPack

import (
	"GinCodeGen/tools/logger"
	"fmt"
	"os"
)

func Error(err error) {
	if err != nil {
		log := logger.GetLogger()
		_, _ = fmt.Fprint(os.Stderr, fmt.Sprintln(err))
		log.Error(err)
	}
}

func ErrExit(err error) {
	log := logger.GetLogger()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, fmt.Sprintln(err))
		log.Error(err)
		os.Exit(1)
	}
}

func ErrorFunc(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}
