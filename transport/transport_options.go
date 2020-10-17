package transport

import "github.com/sirupsen/logrus"

type Options func(s *Transport)

func WithDBName(name string) Options {
	return func(tr *Transport) {
		tr.dbName = name
	}
}

func WithDBUser(user string) Options {
	return func(tr *Transport) {
		tr.dbUser = user
	}
}

func WithDBPass(pass string) Options {
	return func(tr *Transport) {
		tr.dbPass = pass
	}
}

func WithDBURI(uri string) Options {
	return func(tr *Transport) {
		tr.dbURI = uri
	}
}

func WithLogger(logger logrus.FieldLogger) Options {
	return func(app *Transport) {
		app.logger = logger
	}
}
