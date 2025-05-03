package api

import (
	"context"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/casing"
)

type InBody[O any] struct {
	Body *O
}

type OutBody[O any] struct {
	Body *O
}

type Group struct {
	g      *huma.Group
	prefix string
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
	convenience(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *I) (*struct{}, error) {
		err := handler(ctx, *in)
		return nil, err
	})
}

func Get[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	convenience(g, "get", http.MethodGet, path, func(ctx context.Context, in *I) (*OutBody[O], error) {
		out, err := handler(ctx, *in)
		return &OutBody[O]{Body: &out}, err
	})
}

func List[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	convenience(g, "list", http.MethodGet, "/", func(ctx context.Context, in *I) (*OutBody[O], error) {
		out, err := handler(ctx, *in)
		return &OutBody[O]{Body: &out}, err
	})
}

func Post[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	convenience(g, "post", http.MethodPost, "/", func(ctx context.Context, in *InBody[I]) (*OutBody[O], error) {
		out, err := handler(ctx, *in.Body)
		return &OutBody[O]{Body: &out}, err
	})
}

func Put[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	convenience(g, "put", http.MethodPut, path, func(ctx context.Context, in *I) (*OutBody[O], error) {
		out, err := handler(ctx, *in)
		return &OutBody[O]{Body: &out}, err
	})
}

func convenience[I, O any](g Group, action, method, path string, handler func(context.Context, *I) (*O, error)) {
	var o *O
	operation := huma.Operation{
		OperationID: huma.GenerateOperationID(action, filepath.Join(g.prefix, path), o),
		Summary:     huma.GenerateSummary(action, filepath.Join(g.prefix, path), o),
		Method:      method,
		Path:        path,
		Metadata:    map[string]any{},
	}

	huma.Register(g.g, operation, func(ctx context.Context, in *I) (*O, error) {
		out, err := handler(ctx, in)
		if err != nil {
			slog.Error("handler", "err", err)
		}

		return out, err
	})
}
