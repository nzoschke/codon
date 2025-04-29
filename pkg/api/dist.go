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

		proxy := httputil.NewSingleHostReverseProxy(url)

		// Handle all non-API routes with the proxy
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// Skip proxy for API routes
			if len(r.URL.Path) >= 4 && r.URL.Path[:4] == "/api" {
				return
			}
			proxy.ServeHTTP(w, r)
		})

		return nil
	}

	slog.Info("api", "dist", "embed")

	// Create a filesystem from the embedded dist directory
	distFS, err := fs.Sub(build.Dist, "dist")
	if err != nil {
		return errors.WithStack(err)
	}

	// Serve static files from the embedded filesystem
	fileServer := http.FileServer(http.FS(distFS))
	mux.Handle("/", fileServer)

	return nil
}
