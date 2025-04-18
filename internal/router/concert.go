package router

import (
	"database/sql"
	"net/http"

	"github.com/iandanarko/concert/internal/handler/concert"
	concertrepo "github.com/iandanarko/concert/internal/repository/db/concert"
	concertsvc "github.com/iandanarko/concert/internal/service/concert"
)

func concertRouters(db *sql.DB) []route {
	getAvailableRepo := concertrepo.New(db)
	getAvailableSvc := concertsvc.NewGetAvailable(getAvailableRepo)
	getAvailableHandler := concert.NewGetAvailableHandler(getAvailableSvc)

	return []route{
		{
			method:  http.MethodGet,
			path:    "/concerts",
			handler: getAvailableHandler.Handle,
		},
	}
}
