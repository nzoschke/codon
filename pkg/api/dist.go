package api

import (
	"io/fs"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nzoschke/codon/build"
	"github.com/olekukonko/errors"
)

func dist(mux *http.ServeMux, dev bool) error {
	if dev {
		slog.Info("api", "dist", "proxy")

		url, err := url.Parse("http://localhost:3000")
		if err != nil {
			return errors.WithStack(err)
		}

		p := httputil.NewSingleHostReverseProxy(url)
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			p.ServeHTTP(w, r)
		})

		return nil
	}

	slog.Info("api", "dist", "embed")

	// serve all non-api routes from build
	f, err := fs.Sub(build.Dist, "dist")
	if err != nil {
		return errors.WithStack(err)
	}

	s := http.FileServer(http.FS(f))
	mux.Handle("/", s)

	return nil
}
