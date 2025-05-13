package api

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/a-h/rest"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nzoschke/codon/pkg/db"
	"github.com/olekukonko/errors"
)

//go:generate ./oapi.sh

func New(ctx context.Context, addr string, db db.DB, dev bool) error {
	m := http.NewServeMux()
	if _, err := NewAPI(m, db, dev); err != nil {
		return errors.WithStack(err)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
			code:           http.StatusOK,
		}

		m.ServeHTTP(rw, r)

		slog.Info("api",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.code,
			"remote_addr", r.RemoteAddr,
		)
	})

	s := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func() {
		slog.Info("api", "serve", addr)
		if err := s.ListenAndServe(); err != nil {
			slog.Error("api", "err", err)
		}
	}()

	<-ctx.Done()

	sctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.Shutdown(sctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewAPI(m *http.ServeMux, db db.DB, dev bool) (*rest.API, error) {
	r := rest.NewAPI("Codon")
	r.StripPkgPaths = []string{"github.com/a-h/respond", "github.com/nzoschke/codon"}
	r.RegisterModel(rest.ModelOf[Error](), rest.WithDescription("Standard JSON error"), func(s *openapi3.Schema) {
		status := s.Properties["statusCode"]
		status.Value.WithMin(100).WithMax(600)
	})

	if err := dist(m, dev); err != nil {
		return nil, errors.WithStack(err)
	}
	contacts(db, m, r)
	health(m)

	// generate /openapi.json last
	if err := oapi(m, r); err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}
