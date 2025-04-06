package log

import (
	"io"
	"log/slog"
)

func SetDefault(getenv func(string) string, stdout io.Writer) {
	level := slog.LevelInfo
	levels := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}

	if l, ok := levels[getenv("LEVEL")]; ok {
		level = l
	}

	l := slog.New(slog.NewTextHandler(stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	slog.SetDefault(l)
}
