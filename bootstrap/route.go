package bootstrap

import (
	"github.com/attitudefx7/goblog/pkg/route"
	"github.com/attitudefx7/goblog/routes"
	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
