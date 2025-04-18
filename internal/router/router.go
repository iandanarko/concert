package router

import (
	"database/sql"

	"github.com/iandanarko/concert/config"
	"github.com/labstack/echo/v4"
)

type route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

func BuildRoutes(e *echo.Echo, cfg config.Config, db *sql.DB) {
	middlewares := []echo.MiddlewareFunc{}
	routes := []route{}

	for _, r := range routes {
		e.Add(r.method, r.path, r.handler, middlewares...)
	}
}
