package api

import (
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nzoschke/codon/build"
	"github.com/pkg/errors"
)

func dist(devAddr string, mux *http.ServeMux) error {
	if !dial(devAddr) {
		slog.Info("static", "fs", "dist")

		dist, err := fs.Sub(build.Dist, "dist")
		if err != nil {
			return errors.WithStack(err)
		}

		mux.Handle("/", http.FileServerFS(dist))

		return nil
	}

	slog.Info("static", "proxy", devAddr)

	url, err := url.Parse("http://" + devAddr)
	if err != nil {
		return errors.WithStack(err)
	}

	mux.Handle("/", httputil.NewSingleHostReverseProxy(url))

	return nil
}

func dial(addr string) bool {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
