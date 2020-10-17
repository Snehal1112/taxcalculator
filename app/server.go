package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Router     *mux.Router
	Server     *http.Server
	listenAddr string
}

func (a *App) StartServer(options ...ServerOption) {
	for _,option := range options{
		option(a.Srv)
	}
	startServer(a)
}

func startServer(a *App) {
	cc := cors.AllowAll()
	a.Srv.Server = &http.Server{
		Handler: cc.Handler(a.Srv.Router),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	listener, err := net.Listen("tcp", a.Srv.listenAddr)
	if err != nil {
		a.Logger.WithError(err)
		return
	}

	a.Logger.WithField("listenAddr", a.Srv.listenAddr).Infoln("starting http listener")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := a.Srv.Server.Serve(listener); err != nil {
			time.Sleep(time.Second)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	a.Srv.Server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Server shutting down")
	os.Exit(0)
}
