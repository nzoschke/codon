package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"github.com/olekukonko/errors"
)

func writeErr(w http.ResponseWriter, err error) {
	slog.Error("api", "err", err)

	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(respond.Error{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}); err != nil {
		slog.Error("writeErr", "error", err)
	}
}

func DeleteID[I any](path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) error) {
	r.Delete(path).
		HasPathParameter("id", rest.PathParam{
			Description: "id",
		}).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("DELETE %s", path), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			id, err := convertID[I](r.PathValue("id"))
			if err != nil {
				return errors.WithStack(err)
			}

			if err := handler(r.Context(), id); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			writeErr(w, err)
		}
	})
}

func GetID[I, O any](path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	r.Get(path).
		HasPathParameter("id", rest.PathParam{
			Description: "id",
			Regexp:      `\d+`,
		}).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("GET %s", path), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			id, err := convertID[I](r.PathValue("id"))
			if err != nil {
				return errors.WithStack(err)
			}

			out, err := handler(r.Context(), id)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := json.NewEncoder(w).Encode(out); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			writeErr(w, err)
		}
	})
}

func List[I, O any](path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	r.Get(path).
		HasQueryParameter("limit", rest.QueryParam{ // FIXME
			Regexp: `\d+`,
		}).
		HasQueryParameter("offset", rest.QueryParam{
			Regexp: `\d+`,
		}).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("GET %s", path), func(w http.ResponseWriter, r *http.Request) {
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
			writeErr(w, err)
		}
	})
}

func Post[I, O any](path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I) (O, error)) {
	r.Post(path).
		HasRequestModel(rest.ModelOf[I]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("POST %s", path), func(w http.ResponseWriter, r *http.Request) {
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
			writeErr(w, err)
		}
	})
}

func Put[I, B, O any](path string, m *http.ServeMux, r *rest.API, handler func(context.Context, I, B) (O, error)) {
	r.Put(path).
		HasPathParameter("id", rest.PathParam{
			Description: "id",
			Regexp:      `\d+`,
		}).
		HasRequestModel(rest.ModelOf[B]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[O]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]())

	m.HandleFunc(fmt.Sprintf("PUT %s", path), func(w http.ResponseWriter, r *http.Request) {
		if err := func() error {
			id, err := convertID[I](r.PathValue("id"))
			if err != nil {
				return errors.WithStack(err)
			}

			var in B
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				return errors.WithStack(err)
			}

			out, err := handler(r.Context(), id, in)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := json.NewEncoder(w).Encode(out); err != nil {
				return errors.WithStack(err)
			}

			return nil
		}(); err != nil {
			writeErr(w, err)
		}
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
