package transport

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/snehal1112/transport/client"
)

var dbSession interface{}

type Transport struct {
	dbUser string
	dbPass string
	dbName string
	dbURI string

	logger logrus.FieldLogger
}

func NewTransport(options ...Options) *Transport {
	t := &Transport{}
	for _, option := range options{
		option(t)
	}
	t.logger.WithField("data_URI", t.dbURI).Infoln("database url is configured.")
	return t
}

func (t *Transport) GetSession() *client.Connect {
	if dbSession != nil {
		return dbSession.(*client.Connect)
	}

	dbSession = client.NewConnection(
		client.WithURL(t.dbURI),
		client.WithCtx(context.Background()),
		client.WithDatabase(t.dbName))
	return dbSession.(*client.Connect)
}