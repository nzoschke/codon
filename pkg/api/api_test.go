package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nzoschke/codon/pkg/api"
	"github.com/nzoschke/codon/pkg/run"
	"github.com/nzoschke/codon/pkg/sql/models"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
)

var reISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)

type Test struct {
	in     any
	method string
	path   string
	want   any
}

func TestContact(t *testing.T) {
	test(t, []Test{
		{
			in: q.ContactCreateIn{
				Email: "a@example.com",
				Info: models.ContactInfo{
					Age: 21,
				},
				Name: "Ann",
			},
			method: http.MethodPost,
			path:   "/api/contacts",
			want: q.Contact{
				Email: "a@example.com",
				ID:    1,
				Info: models.ContactInfo{
					Age: 21,
				},
				Name: "Ann",
			},
		},
		{
			in:     nil,
			method: http.MethodGet,
			path:   "/api/contacts/1",
			want: q.Contact{
				Email: "a@example.com",
				ID:    1,
				Info: models.ContactInfo{
					Age: 21,
				},
				Name: "Ann",
			},
		},
		{
			in: api.ContactUpdateIn{
				Email: "a@new.com",
				Info: models.ContactInfo{
					Age: 22,
				},
				Name: "Ann",
			},
			method: http.MethodPut,
			path:   "/api/contacts/1",
			want: q.Contact{
				Email: "a@new.com",
				ID:    1,
				Info: models.ContactInfo{
					Age: 22,
				},
				Name: "Ann",
			},
		},
		{
			in:     nil,
			method: http.MethodDelete,
			path:   "/api/contacts/1",
			want:   "",
		},
	})
}

func TestUser(t *testing.T) {
	test(t, []Test{
		{
			in: api.UserCreateIn{
				Email:    "a@example.com",
				Password: "password",
			},
			method: http.MethodPost,
			path:   "/api/users",
			want: q.UserGetOut{
				Email: "a@example.com",
				ID:    1,
			},
		},
		{
			in: api.UserCreateIn{
				Email:    "a@example.com",
				Password: "password",
			},
			method: http.MethodPost,
			path:   "/api/users/session",
			want:   nil,
		},
		{
			method: http.MethodGet,
			path:   "/api/users/session",
			want: q.UserGetOut{
				Email: "a@example.com",
				ID:    1,
			},
		},
		{
			method: http.MethodDelete,
			path:   "/api/users/session",
			want:   nil,
		},
		{
			method: http.MethodGet,
			path:   "/api/users/session",
			want:   huma.Error401Unauthorized("invalid session"),
		},
	})
}

func test(t *testing.T, tests []Test) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "11234"
	go run.Run(ctx, []string{"test", "-db", "file::memory:?mode=memory&cache=shared", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)
	err := run.Health(ctx, 100*time.Millisecond, port)
	a.NoError(err)

	cookie := ""
	for _, test := range tests {
		name := fmt.Sprintf("%s %s", test.method, test.path)

		bs, err := json.Marshal(test.in)
		a.NoError(err)

		req, err := http.NewRequestWithContext(ctx, test.method, fmt.Sprintf("http://localhost:%s%s", port, test.path), bytes.NewBuffer(bs))
		a.NoError(err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Cookie", cookie)

		res, err := http.DefaultClient.Do(req)
		a.NoError(err)

		cookie = res.Header.Get("Set-Cookie")

		if test.want == nil {
			continue
		} else if _, ok := test.want.(string); ok {
			bs, _ = io.ReadAll(res.Body)
			a.Equal(test.want, string(bs))
		} else {
			got := test.want
			err = json.NewDecoder(res.Body).Decode(&got)
			a.NoError(err)
			defer res.Body.Close()

			JSONEq(a, test.want, got, name)
		}
	}
}

func timeAny() time.Time {
	return time.Unix(0, 0).UTC()
}

func JSONEq(a *assert.Assertions, expected any, actual any, msgAndArgs ...any) {
	be, err := json.Marshal(expected)
	a.NoError(err)

	// remove $schema
	if m, ok := actual.(map[string]any); ok {
		delete(m, "$schema")
		actual = m
	}

	ba, err := json.Marshal(actual)
	a.NoError(err)

	// replace all ISO 8601 UTC strings (eg "2025-04-12T16:25:32Z") with "1970-01-01T00:00:00Z"
	t := timeAny()
	be = reISO8601.ReplaceAll(be, []byte(t.Format("2006-01-02T15:04:05.999Z")))
	ba = reISO8601.ReplaceAll(ba, []byte(t.Format("2006-01-02T15:04:05.999Z")))

	a.JSONEq(string(be), string(ba), msgAndArgs...)
}
