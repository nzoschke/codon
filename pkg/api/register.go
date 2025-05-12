package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"path"
	"strings"

	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/casing"
	"github.com/olekukonko/errors"
)

type Group struct {
	g      *huma.Group
	prefix string
}

type InBody[B any] struct {
	Body *B
}

type InID[I any] struct {
	ID I `path:"id"`
}

type InBodyID[B any, I any] struct {
	Body *B
	ID   I `path:"id"`
}

type OutBody[B any] struct {
	Body *B
}

func NewGroup(a huma.API, prefix string) Group {
	g := huma.NewGroup(a, prefix)

	g.UseModifier(func(op *huma.Operation, next func(*huma.Operation)) {
		op.Path = strings.TrimSuffix(op.Path, "/")
		op.Tags = []string{casing.Camel(prefix)}
		next(op)
	})

	return Group{
		g:      g,
		prefix: prefix,
	}
}

func Delete[I any](g Group, path string, handler func(context.Context, I) error) {
	register(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *InID[I]) (*struct{}, error) {
		err := handler(ctx, in.ID)
		return nil, err
	})
}

func Get[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	register(g, "get", http.MethodGet, path, func(ctx context.Context, in *InID[I]) (*OutBody[O], error) {
		out, err := handler(ctx, in.ID)
		return &OutBody[O]{Body: &out}, err
	})
}

func List[I, O any](g Group, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	p := path.Join("/api", g.prefix)

	r.Get(p).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			var in I
			out, err := handler(r.Context(), in)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := json.NewEncoder(w).Encode(out); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(respond.Error{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}); err != nil {
				slog.Error("register.List", "path", p, "error", err)
			}
		}
	})
}

func Post[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	register(g, "create", http.MethodPost, "/", func(ctx context.Context, in *InBody[I]) (*OutBody[O], error) {
		out, err := handler(ctx, *in.Body)
		return &OutBody[O]{Body: &out}, err
	})
}

func Put[I, B, O any](g Group, path string, handler func(context.Context, I, B) (O, error)) {
	register(g, "update", http.MethodPut, path, func(ctx context.Context, in *InBodyID[B, I]) (*OutBody[O], error) {
		out, err := handler(ctx, in.ID, *in.Body)
		return &OutBody[O]{Body: &out}, err
	})
}

func register[I, O any](g Group, action, method, p string, handler func(context.Context, *I) (*O, error)) {
	var o *O
	operation := huma.Operation{
		OperationID: huma.GenerateOperationID(action, path.Join(g.prefix, p), o),
		Summary:     huma.GenerateSummary(action, path.Join(g.prefix, p), o),
		Method:      method,
		Path:        p,
	}

	huma.Register(g.g, operation, func(ctx context.Context, in *I) (*O, error) {
		out, err := handler(ctx, in)
		if err != nil {
			slog.Error("handler", "err", err)
		}

		return out, err
	})
}
