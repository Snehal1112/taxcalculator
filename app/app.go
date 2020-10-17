package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/snehal1112/taxcal/transport"
	"github.com/snehal1112/transport/client"
)

type App struct {
	Srv   *Server
	Logger logrus.FieldLogger

	transport *transport.Transport
	ctx context.Context
}

func NewApp(options ...Option) *App {
	app := &App{
		Srv: &Server{
			Router: mux.NewRouter(),
		},
	}

	for _,option := range options{
		option(app)
	}

	return app
}

func getSession(a *App) *client.Connect {
	return a.transport.GetSession()
}