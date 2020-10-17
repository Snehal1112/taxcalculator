package app

type ServerOption func(s *Server)

func WithListen(listen string) ServerOption {
	return func(srv *Server) {
		srv.listenAddr = listen
	}
}