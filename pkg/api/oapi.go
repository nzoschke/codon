package api

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/rest"
	"github.com/olekukonko/errors"
)

func oapi(m *http.ServeMux, r *rest.API) error {
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
