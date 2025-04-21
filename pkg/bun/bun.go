package bun

import (
	"context"
	"os"
	"os/exec"

	"github.com/olekukonko/errors"
)

func Dev(ctx context.Context) error {
	if err := run("bun", "install"); err != nil {
		return errors.WithStack(err)
	}

	cmd := exec.Command("bun", "run", "dev", "--silent")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
