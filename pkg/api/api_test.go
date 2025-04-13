package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/run"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/stretchr/testify/assert"
)

var reISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)

func TestUser(t *testing.T) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "11234"
	go run.Run(ctx, []string{"test", "-db", "file::memory:?mode=memory&cache=shared", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)
	err := run.Health(ctx, 100*time.Millisecond, port)
	a.NoError(err)

	tests := []struct {
		in     any
		want   any
		method string
		path   string
	}{
		{
			in: q.ContactCreateParams{
				Email: "a@example.com",
				Name:  "Ann",
			},
			want: q.ContactCreateRes{
				CreatedAt: timeAny(),
				Email:     "a@example.com",
				Id:        1,
				Name:      "Ann",
			},
			method: http.MethodPost,
			path:   "/api/contacts",
		},
		{
			in: nil,
			want: q.ContactCreateRes{
				CreatedAt: timeAny(),
				Email:     "a@example.com",
				Id:        1,
				Name:      "Ann",
			},
			method: http.MethodGet,
			path:   "/api/contacts/1",
		},
		{
			in: q.ContactUpdateParams{
				Email: "a@new.com",
				Name:  "Ann",
			},
			want: q.ContactReadRes{
				CreatedAt: timeAny(),
				Email:     "a@new.com",
				Id:        1,
				Name:      "Ann",
			},
			method: http.MethodPut,
			path:   "/api/contacts/1",
		},
		{
			in:     nil,
			want:   nil,
			method: http.MethodDelete,
			path:   "/api/contacts/1",
		},
	}

	for _, test := range tests {
		bs, err := json.Marshal(test.in)
		a.NoError(err)

		req, err := http.NewRequestWithContext(ctx, test.method, fmt.Sprintf("http://localhost:%s%s", port, test.path), bytes.NewBuffer(bs))
		a.NoError(err)
		req.Header.Add("Content-Type", "application/json")

		res, err := http.DefaultClient.Do(req)
		a.NoError(err)

		got := test.want
		err = json.NewDecoder(res.Body).Decode(&got)
		a.NoError(err)
		defer res.Body.Close()

		JSONEq(a, test.want, got)
	}
}

func timeAny() time.Time {
	return time.Unix(0, 0).UTC()
}

func JSONEq(a *assert.Assertions, expected any, actual any) {
	be, err := json.Marshal(expected)
	a.NoError(err)

	ba, err := json.Marshal(actual)
	a.NoError(err)

	// replace all ISO 8601 UTC strings (eg "2025-04-12T16:25:32Z") with "1970-01-01T00:00:00Z"
	t := timeAny()
	be = reISO8601.ReplaceAll(be, []byte(t.Format("2006-01-02T15:04:05.999Z")))
	ba = reISO8601.ReplaceAll(ba, []byte(t.Format("2006-01-02T15:04:05.999Z")))

	a.JSONEq(string(be), string(ba))
}
