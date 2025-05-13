package bun

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/run"
	"github.com/stretchr/testify/assert"
)

func TestBun(t *testing.T) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "21234"
	go run.Run(ctx, []string{"test", "-db", "file::memory:?mode=memory&cache=shared", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)
	err := run.Health(ctx, 1000*time.Millisecond, port)
	a.NoError(err)

	_, b, _, _ := runtime.Caller(0)

	cmd := exec.Command("bun", "test")
	cmd.Dir = filepath.Dir(filepath.Dir(filepath.Dir(b)))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	a.NoError(cmd.Run())
}
