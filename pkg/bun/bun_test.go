package bun

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBun(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)

	cmd := exec.Command("bun", "test")
	cmd.Dir = filepath.Dir(filepath.Dir(filepath.Dir(b)))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	assert.NoError(t, cmd.Run())
}
