package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const loggerName = "server/http"

type ServerOption func(*Server)

type Server struct {
	server   *http.Server
	listener net.Listener
	network  string
	address  string

	Logger  *logrus.Logger
	timeout time.Duration
	router  *mux.Router
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":8090",
		timeout: time.Second,
		Logger:  defaultLogger(),
		router:  mux.NewRouter(),
	}
	for _, o := range opts {
		o(srv)
	}
	srv.server = &http.Server{Handler: srv}
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
	return
}

func (s *Server) Run() (err error) {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.listener = lis
	s.Logger.Infof("[HTTP] server listening on: %s", lis.Addr().String())
	if err := s.server.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	err = s.server.Shutdown(ctx)
	return
}

func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func Router(router *mux.Router) ServerOption {
	return func(s *Server) {
		s.router = router
	}
}

func Logger(logger *logrus.Logger) ServerOption {
	return func(s *Server) {
		s.Logger = logger
	}
}

func defaultLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
	}
	return logger
}
