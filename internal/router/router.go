package router

import (
	"database/sql"
	"net/http"

	"github.com/iandanarko/concert/config"
	_ "github.com/iandanarko/concert/docs"
	"github.com/iandanarko/concert/internal/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type route struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

func BuildRoutes(e *echo.Echo, cfg config.Config, db *sql.DB) {
	middlewares := []echo.MiddlewareFunc{
		middleware.Auth(cfg),
	}
	routes := []route{}
	routes = append(routes, concertRouters(db)...)
	routes = append(routes, bookingRouter(db)...)
	e.Add(http.MethodGet, "/swagger/*", echoSwagger.WrapHandler)

	for _, r := range routes {
		e.Add(r.method, r.path, r.handler, middlewares...)
	}
}
