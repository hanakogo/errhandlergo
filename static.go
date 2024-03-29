package errhandlergo

import (
	"github.com/gookit/slog"
	"runtime"
)

func HandleRecover(do func(err any)) {
	if r := recover(); r != nil {
		do(r)
	}
}

func HandleErrorIsNil(err error, postprocess func()) {
	if err == nil {
		if postprocess != nil {
			postprocess()
		}
	}
}

func HandleError(err error, postprocess func()) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		slog.Errorf("error encountered[%s:%d]", file, line)
		slog.Errorf("errors: %s", err.Error())
		if postprocess != nil {
			postprocess()
		}
	}
}
