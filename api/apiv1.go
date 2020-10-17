package api

import (
	"github.com/gorilla/mux"
	"github.com/snehal1112/taxcal/app"
)

type router = map[string]*mux.Router

type API struct {
	App *app.App
	BaseRoutes router
}

func Init(a *app.App, root *mux.Router) *API {
	api := &API{
		App:        a,
		BaseRoutes: make(router),
	}

	api.BaseRoutes["ApiRoot"] = root.PathPrefix("/api/v1").Subrouter()
	baseRoutes := api.BaseRoutes["ApiRoot"]

	api.BaseRoutes["User"] = baseRoutes.PathPrefix("/user").Subrouter()
	api.BaseRoutes["Rule"] = baseRoutes.PathPrefix("/rule").Subrouter()
	api.BaseRoutes["Tax"] = baseRoutes.PathPrefix("/tax").Subrouter()

	api.InitUser()
	api.InitRule()
	api.InitTax()

	return api
}