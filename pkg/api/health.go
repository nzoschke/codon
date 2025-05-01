package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetHealthOut struct {
	ContentType string `header:"Content-Type"`
	Body        []byte `example:"ok"`
}

func health(g huma.API) {
	huma.Register(g, huma.Operation{
		Description: "Returns 200 if healthy, indeterminate response if not.",
		Method:      http.MethodGet,
		Path:        "/health",
		Responses: map[string]*huma.Response{
			"200": {
				Content: map[string]*huma.MediaType{
					"text/plain": {
						Example: "ok",
					},
				},
			},
		},
		Summary: "Get health",
	}, func(ctx context.Context, in *struct{}) (*GetHealthOut, error) {
		return &GetHealthOut{
			ContentType: "text/plain",
			Body:        []byte("ok"),
		}, nil
	})
}
