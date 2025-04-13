package api

import (
	"log/slog"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nzoschke/codon/build"
	"github.com/olekukonko/errors"
)

func dist(e *echo.Echo, dev bool) error {
	if dev {
		slog.Info("api", "dist", "proxy")

		url, err := url.Parse("http://localhost:3000")
		if err != nil {
			return errors.WithStack(err)
		}

		e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
			Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
				{
					URL: url,
				},
			}),
			Skipper: func(c echo.Context) bool {
				return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
			},
		}))

		e.GET("", echo.WrapHandler(httputil.NewSingleHostReverseProxy(url)))

		return nil
	}

	slog.Info("api", "dist", "embed")

	e.StaticFS("", echo.MustSubFS(build.Dist, "dist"))

	return nil
}
