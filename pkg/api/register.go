package api

import (
	"context"
	"log/slog"
	"net/http"
	"path"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/casing"
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

type OutCookie struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
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

func Delete(g Group, path string, handler func(context.Context) error) {
	register(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *struct{}) (*struct{}, error) {
		err := handler(ctx)
		return nil, err
	})
}

func DeleteID[I any](g Group, path string, handler func(context.Context, I) error) {
	register(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *InID[I]) (*struct{}, error) {
		err := handler(ctx, in.ID)
		return nil, err
	})
}

func DeleteIn[I any](g Group, path string, handler func(context.Context, I) error) {
	register(g, "delete", http.MethodDelete, path, func(ctx context.Context, in *I) (*struct{}, error) {
		err := handler(ctx, *in)
		return nil, err
	})
}

func Get[O any](g Group, path string, handler func(context.Context) (O, error)) {
	register(g, "get", http.MethodGet, path, func(ctx context.Context, in *struct{}) (*OutBody[O], error) {
		out, err := handler(ctx)
		return &OutBody[O]{Body: &out}, err
	})
}

func GetID[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	register(g, "get", http.MethodGet, path, func(ctx context.Context, in *InID[I]) (*OutBody[O], error) {
		out, err := handler(ctx, in.ID)
		return &OutBody[O]{Body: &out}, err
	})
}

func GetIn[I any, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	register(g, "get", http.MethodGet, path, func(ctx context.Context, in *I) (*OutBody[O], error) {
		out, err := handler(ctx, *in)
		return &OutBody[O]{Body: &out}, err
	})
}

func GetInOut[I any, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	register(g, "get", http.MethodGet, path, func(ctx context.Context, in *I) (*O, error) {
		out, err := handler(ctx, *in)
		return &out, err
	})
}

func List[I, O any](g Group, handler func(context.Context, I) (O, error)) {
	register(g, "list", http.MethodGet, "/", func(ctx context.Context, in *I) (*OutBody[O], error) {
		out, err := handler(ctx, *in)
		return &OutBody[O]{Body: &out}, err
	})
}

func Post[O any](g Group, path string, handler func(context.Context) (O, error)) {
	register(g, "create", http.MethodPost, path, func(ctx context.Context, in *struct{}) (*OutBody[O], error) {
		out, err := handler(ctx)
		return &OutBody[O]{Body: &out}, err
	})
}

func PostBody[I, O any](g Group, path string, handler func(context.Context, I) (O, error)) {
	register(g, "create", http.MethodPost, path, func(ctx context.Context, in *InBody[I]) (*OutBody[O], error) {
		out, err := handler(ctx, *in.Body)
		return &OutBody[O]{Body: &out}, err
	})
}

func PostCookie[I any](g Group, path string, handler func(context.Context, I) (http.Cookie, error)) {
	register(g, "create", http.MethodPost, path, func(ctx context.Context, in *InBody[I]) (*OutCookie, error) {
		out, err := handler(ctx, *in.Body)
		return &OutCookie{SetCookie: out}, err
	})
}

func Put[B, O any](g Group, path string, handler func(context.Context, B) (O, error)) {
	register(g, "update", http.MethodPut, path, func(ctx context.Context, in *InBody[B]) (*OutBody[O], error) {
		out, err := handler(ctx, *in.Body)
		return &OutBody[O]{Body: &out}, err
	})
}

func PutID[I, B, O any](g Group, path string, handler func(context.Context, I, B) (O, error)) {
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
