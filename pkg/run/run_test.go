package run_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/run"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "11234"
	go run.Run(ctx, []string{"test", "-db", "", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)
	err := run.Health(ctx, 100*time.Millisecond, port)
	a.NoError(err)
}
