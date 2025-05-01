package api

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func health(g huma.API) {
	huma.Register(g, huma.Operation{
		Description: "Returns status 200 if healthy, indeterminate response if not.",
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
		Summary: "Get API health",
	}, func(ctx context.Context, in *struct{}) (*HealthOut, error) {
		return &HealthOut{
			ContentType: "text/plain",
			Body:        []byte("ok"),
		}, nil
	})
}
