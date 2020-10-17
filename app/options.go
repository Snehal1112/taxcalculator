package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/snehal1112/taxcal/transport"
)

type Option func(a *App)

func WithCtx(ctx context.Context) Option {
	return func(app *App) {
		app.ctx = ctx
	}
}

func WithLogger(logger logrus.FieldLogger) Option {
	return func(app *App) {
		app.Logger = logger
	}
}

func WithTransport(tr *transport.Transport) Option {
	return func(app *App) {
		app.transport = tr
	}
}

