package api

import (
	"fmt"
	"net/http"
)

type GetHealthOut struct {
	ContentType string `header:"Content-Type"`
	Body        []byte `example:"ok"`
}

func health(m *http.ServeMux) {
	m.HandleFunc(fmt.Sprintf("GET %s", "/api/health"), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
