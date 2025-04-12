package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/nzoschke/codon/pkg/run"
	"github.com/nzoschke/codon/pkg/sql/q"
	"github.com/olekukonko/errors"
	"github.com/stretchr/testify/assert"
)

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

	bs, err := json.Marshal(q.UserCreateParams{
		Email: "foo@bar.com",
		Name:  "user",
	})
	a.NoError(err)

	res, err := http.Post(fmt.Sprintf("http://localhost:%s/api/users", port), "application/json", bytes.NewBuffer(bs))
	a.NoError(err)

	out := q.UserCreateRes{}
	err = json.NewDecoder(res.Body).Decode(&out)
	a.NoError(err)
	defer res.Body.Close()

	a.Equal(time.Now().Format("2006-01-02"), out.CreatedAt.Format("2006-01-02"))

	a.Equal(q.UserCreateRes{
		CreatedAt: out.CreatedAt,
		Email:     "foo@bar.com",
		Id:        1,
		Name:      "user",
	}, out)
}

func waitForHealth(ctx context.Context, timeout time.Duration, port string) error {
	client := http.Client{}
	startTime := time.Now()
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:%s/health", port), nil)
		if err != nil {
			return errors.WithStack(err)
		}

		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
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
