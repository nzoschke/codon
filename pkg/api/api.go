package api

import (
	"log/slog"
	"net/http"

	"github.com/pkg/errors"
)

func New(addr string, devAddr string) error {
	mux := http.NewServeMux()

	dist(devAddr, mux)

	slog.Info("app", "serve", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
