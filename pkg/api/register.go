package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"path"
	pth "path"
	"strconv"
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

func DeleteID[I any](g Group, path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) error) {
	p := pth.Join("/api", g.prefix, path)

	r.Delete(p).
		HasPathParameter("id", rest.PathParam{
			Description: "id",
			Regexp:      `\d+`,
		}).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("DELETE %s", p), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			in, err := convertID[I](r.PathValue("id"))
			if err != nil {
				return errors.WithStack(err)
			}

			if err := handler(r.Context(), in); err != nil {
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

func GetID[I, O any](g Group, path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	p := pth.Join("/api", g.prefix, path)

	r.Get(p).
		HasPathParameter("id", rest.PathParam{
			Description: "id",
			Regexp:      `\d+`,
		}).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("GET %s", p), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			in, err := convertID[I](r.PathValue("id"))
			if err != nil {
				return errors.WithStack(err)
			}

			out, err := handler(r.Context(), in)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := json.NewEncoder(w).Encode(out); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			slog.Error("register.GetID", "path", p, "error", err)

			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(respond.Error{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}); err != nil {
				slog.Error("register.GetID", "path", p, "error", err)
			}
		}
	})
}

func List[I, O any](g Group, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	p := pth.Join("/api", g.prefix)

	r.Get(p).
		HasQueryParameter("limit", rest.QueryParam{ // FIXME
			Regexp: `\d+`,
		}).
		HasQueryParameter("offset", rest.QueryParam{
			Regexp: `\d+`,
		}).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("GET %s", p), func(w http.ResponseWriter, r *http.Request) {
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
			slog.Error("register.List", "path", p, "error", err)

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

func Post[I, O any](g Group, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	p := pth.Join("/api", g.prefix)

	r.Post(p).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("POST %s", p), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			var in I
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				return errors.WithStack(err)
			}

			out, err := handler(r.Context(), in)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := json.NewEncoder(w).Encode(out); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			slog.Error("register.Post", "path", p, "error", err)

			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode(respond.Error{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}); err != nil {
				slog.Error("register.Post", "path", p, "error", err)
			}
		}
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

func convertID[I any](id string) (I, error) {
	var result I
	switch any(result).(type) {
	case int64:
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return result, errors.WithStack(err)
		}
		return any(val).(I), nil
	case int:
		val, err := strconv.Atoi(id)
		if err != nil {
			return result, errors.WithStack(err)
		}
		return any(val).(I), nil
	case string:
		return any(id).(I), nil
	default:
		return result, fmt.Errorf("unsupported ID type: %T", result)
	}
}
