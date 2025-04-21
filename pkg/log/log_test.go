package log_test

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/nzoschke/codon/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	a := assert.New(t)

	// default level is INFO
	bs := bytes.Buffer{}
	log.SetDefault(func(string) string {
		return ""
	}, &bs)
	slog.Debug("debug")
	slog.Info("info")

	a.Equal("level=INFO msg=info\n", bs.String())

	bs = bytes.Buffer{}
	log.SetDefault(func(string) string {
		return "debug"
	}, &bs)
	slog.Debug("debug")
	slog.Info("info")

	a.Equal("level=DEBUG msg=debug\nlevel=INFO msg=info\n", bs.String())
}
