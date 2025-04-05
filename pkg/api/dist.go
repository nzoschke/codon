package api

import (
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nzoschke/codon/build"
	"github.com/olekukonko/errors"
)

func dist(dev bool, mux *http.ServeMux) error {
	if dev {
		slog.Info("api", "dist", "proxy")

		url, err := url.Parse("http://localhost:3000")
		if err != nil {
			return errors.WithStack(err)
		}

		mux.Handle("/", httputil.NewSingleHostReverseProxy(url))

		return nil
	}

	slog.Info("api", "dist", "embed")

	dist, err := fs.Sub(build.Dist, "dist")
	if err != nil {
		return errors.WithStack(err)
	}

	mux.Handle("/", http.FileServerFS(dist))

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
