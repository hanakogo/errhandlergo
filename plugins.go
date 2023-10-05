package errhandlergo

import (
	"github.com/gookit/slog"
	"runtime"
)

type HandlerPlugin string

const (
	Logging HandlerPlugin = "logging"
	Panic   HandlerPlugin = "panic"
)

var plugins = map[HandlerPlugin]func(err error){
	Logging: func(err error) {
		if err != nil {
			_, file, line, _ := runtime.Caller(2)
			slog.Errorf("error encountered[%s:%d]", file, line)
			slog.Errorf("errors: %s", err.Error())
		}
	},
	Panic: func(err error) {
		if err != nil {
			panic(err)
		}
	},
}
