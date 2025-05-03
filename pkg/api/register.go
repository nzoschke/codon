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

type BodyIn[I any] struct {
	Body *I
}

type IDIn[I any] struct {
	ID I `path:"id"`
}

type AllIn[B any, I any] struct {
	Body *B
	ID   I `path:"id"`
}

type BodyOut[O any] struct {
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
	convenience(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *IDIn[I]) (*struct{}, error) {
		err := handler(ctx, in.ID)
		return nil, err
	})
}

func Get[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	convenience(g, "get", http.MethodGet, path, func(ctx context.Context, in *IDIn[I]) (*BodyOut[O], error) {
		out, err := handler(ctx, in.ID)
		return &BodyOut[O]{Body: &out}, err
	})
}

func List[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	convenience(g, "list", http.MethodGet, "/", func(ctx context.Context, in *I) (*BodyOut[O], error) {
		out, err := handler(ctx, *in)
		return &BodyOut[O]{Body: &out}, err
	})
}

func Post[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	convenience(g, "post", http.MethodPost, "/", func(ctx context.Context, in *BodyIn[I]) (*BodyOut[O], error) {
		out, err := handler(ctx, *in.Body)
		return &BodyOut[O]{Body: &out}, err
	})
}

func Put[I, B, O any](g Group, path string, handler func(context.Context, I, B) (O, error)) {
	convenience(g, "put", http.MethodPut, path, func(ctx context.Context, in *AllIn[B, I]) (*BodyOut[O], error) {
		out, err := handler(ctx, in.ID, *in.Body)
		return &BodyOut[O]{Body: &out}, err
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
