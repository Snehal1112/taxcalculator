package api

import (
	"github.com/snehal1112/taxcal/app"
	"net/http"
)

type handler struct {
	app            *app.App
	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	requireSession bool
	trustRequester bool
	requireMfa     bool
}

type Context struct {
	App           *app.App
	RequestId     string
	IpAddress     string
	Path          string
}

func (api *API) ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		app:            api.App,
		handleFunc:     h,
		requireSession: false,
		trustRequester: false,
		requireMfa:     false,
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{}
	c.App = h.app
	c.IpAddress = r.RemoteAddr

	// handle session and authentication stuff
	w.Header().Set("Content-Type", "application/json")
	c.Path = r.URL.Path

	c.App.Logger.Infoln(c.Path)

	h.handleFunc(c, w, r)

	// handle error
}