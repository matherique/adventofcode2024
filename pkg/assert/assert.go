package assert

import (
	"log/slog"
	"os"
)

func NotNill(err error, message string, args ...any) {
	if err != nil {
		slog.Error(message, "err", err, "args", args)
		os.Exit(1)
	}
}

func True(cond bool, message string, args ...any) {
	if cond {
		slog.Error(message, "args", args)
		os.Exit(1)
	}
}
