package log

import (
	"io"
	"log/slog"
	"strings"
)

func SetDefault(getenv func(string) string, stdout io.Writer) {
	level := slog.LevelInfo
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	if l, ok := levels[strings.ToLower(getenv("LEVEL"))]; ok {
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
