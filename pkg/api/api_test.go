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
	"github.com/olekukonko/errors"
	"github.com/stretchr/testify/assert"
)

var reISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z`)

func TestHealth(t *testing.T) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "11234"
	go run.Run(ctx, []string{"test", "-db", "file::memory:?mode=memory&cache=shared", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)

	err := waitForHealth(ctx, 100*time.Millisecond, port)
	a.NoError(err)
}

func TestUser(t *testing.T) {
	ctx := t.Context()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	a := assert.New(t)

	port := "11234"
	go run.Run(ctx, []string{"test", "-db", "file::memory:?mode=memory&cache=shared", "-port", port}, func(string) string { return "DEBUG" }, os.Stdout)

	err := waitForHealth(ctx, 100*time.Millisecond, port)
	a.NoError(err)

	tests := []struct {
		in     any
		want   any
		method string
		path   string
	}{
		{
			in: q.UserCreateParams{
				Email: "user@example.com",
				Name:  "user",
			},
			want: q.UserCreateRes{
				CreatedAt: epoch(),
				Email:     "user@example.com",
				Id:        1,
				Name:      "user",
			},
			method: http.MethodPost,
			path:   "/api/users",
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

func epoch() *time.Time {
	t := time.Unix(0, 0).UTC()
	return &t
}

func JSONEq(a *assert.Assertions, expected any, actual any) {
	be, err := json.Marshal(expected)
	a.NoError(err)

	ba, err := json.Marshal(actual)
	a.NoError(err)

	// replace all ISO 8601 UTC strings (eg "2025-04-12T16:25:32Z") with "1970-01-01T00:00:00Z"
	e := *epoch()
	be = reISO8601.ReplaceAll(be, []byte(e.Format("2006-01-02T15:04:05.999Z")))
	ba = reISO8601.ReplaceAll(ba, []byte(e.Format("2006-01-02T15:04:05.999Z")))

	a.JSONEq(string(be), string(ba))
}

func waitForHealth(ctx context.Context, timeout time.Duration, port string) error {
	client := http.Client{}
	startTime := time.Now().UTC()
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%s/health", port), nil)
		if err != nil {
			return errors.WithStack(err)
		}

		res, err := client.Do(req)
		if err == nil && res.StatusCode == http.StatusOK {
			res.Body.Close()
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if time.Since(startTime) >= timeout {
				return fmt.Errorf("timeout reached while waiting for endpoint")
			}

			time.Sleep(25 * time.Millisecond)
		}
	}
}
