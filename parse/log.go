package parse

import (
	"fmt"

	"github.com/ije/esbuild-internal/logger"
)

type LogMap struct {
	Verbose []string
	Debug   []string
	Info    []string
	Err     []string
	Warning []string
}

func NewLogMap(options logger.OutputOptions, logMap *LogMap) logger.Log {
	hasErrors := false
	var msgs []logger.Msg

	return logger.Log{
		Level: options.LogLevel,
		AddMsg: func(msg logger.Msg) {
			msgs = append(msgs, msg)

			switch msg.Kind {
			case logger.Verbose:
				if options.LogLevel <= logger.LevelVerbose {
					logMap.Verbose = append(logMap.Verbose, msgString(&msg))
				}

			case logger.Debug:
				if options.LogLevel <= logger.LevelDebug {
					logMap.Debug = append(logMap.Debug, msgString(&msg))
				}

			case logger.Info:
				if options.LogLevel <= logger.LevelInfo {
					logMap.Info = append(logMap.Info, msgString(&msg))
				}

			case logger.Error:
				if options.LogLevel <= logger.LevelError {
					hasErrors = true
					logMap.Err = append(logMap.Err, msgString(&msg))
				}

			case logger.Warning:
				if options.LogLevel <= logger.LevelWarning {
					logMap.Warning = append(logMap.Warning, msgString(&msg))
				}
			}
		},

		HasErrors: func() bool {
			return hasErrors
		},

		AlmostDone: func() {
			// noop
		},

		Done: func() []logger.Msg {
			return msgs
		},
	}
}

func msgString(msg *logger.Msg) string {
	if loc := msg.Data.Location; loc != nil {
		return fmt.Sprintf("%s: %s\n", loc.File, msg.Data.Text)
	}
	return fmt.Sprintf("%s: \n", msg.Data.Text)
}
