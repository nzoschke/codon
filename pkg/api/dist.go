package api

import (
	"io/fs"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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

		// proxy all non-api routes to bun
		p := httputil.NewSingleHostReverseProxy(url)
		paths := []string{"/api", "/spec", "/swagger"}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			for _, p := range paths {
				if strings.HasPrefix(r.URL.Path, p) {
					return
				}
			}

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
