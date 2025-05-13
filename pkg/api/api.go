package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	m := http.NewServeMux()
	if err := NewAPI(m, db, dev); err != nil {
		return errors.WithStack(err)
	}

	s := &http.Server{
		Addr:    addr,
		Handler: m,
	}

	go func() {
		slog.Info("api", "serve", addr)
		if err := s.ListenAndServe(); err != nil {
			slog.Error("api", "err", err)
		}
	}()

	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(sctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewAPI(m *http.ServeMux, db db.DB, dev bool) error {
	r := rest.NewAPI("Codon")
	r.StripPkgPaths = []string{"github.com/a-h/rest/example", "github.com/a-h/respond", "github.com/nzoschke/codon"}
	r.RegisterModel(rest.ModelOf[respond.Error](), rest.WithDescription("Standard JSON error"), func(s *openapi3.Schema) {
		status := s.Properties["statusCode"]
		status.Value.WithMin(100).WithMax(600)
	})

	dist(m, dev)
	contacts(db, m, r)
	health(m)

	spec, err := r.Spec()
	if err != nil {
		return errors.WithStack(err)
	}

	m.Handle("/openapi.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(spec)
	}))

	m.Handle("/spec", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		specURL := "/openapi.json"

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!doctype html>
		<html lang="en">
		<head>
			<meta charset="utf-8" />
			<meta name="referrer" content="same-origin" />
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<link rel="icon" type="image/svg+xml" href="https://go-fuego.dev/img/logo.svg">
			<title>OpenAPI specification</title>
			<script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
			<link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css" />
		</head>
		<body style="height: 100vh;">
			<elements-api
				apiDescriptionUrl="` + specURL + `"
				layout="responsive"
				router="hash"
				logo="https://go-fuego.dev/img/logo.svg"
				tryItCredentialsPolicy="same-origin"
			/>
		</body>
		</html>`))
	}))

	return nil
}
