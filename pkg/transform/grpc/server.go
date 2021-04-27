package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ServerOption func(o *Server)

type Server struct {
	*grpc.Server
	app      string
	network  string
	address  string
	timeout  time.Duration
	grpcOpts []grpc.ServerOption

	Logger *logrus.Logger
}

// Network with server network.
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

func Logger(logger *logrus.Logger) ServerOption {
	return func(s *Server) {
		s.Logger = logger
	}
}

func GrpcOpts(opts ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		//s.grpcOpts = append(s.grpcOpts, opts...)
		s.grpcOpts = opts
	}
}

func GrpcHealthCheck(s *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
}

func GrpcKeepAlive(kp keepalive.ServerParameters) ServerOption {
	return func(s *Server) {
		s.grpcOpts = append(s.grpcOpts, grpc.KeepaliveParams(kp))
	}
}

func GrpcUnaryServerInterceptor(interceptors ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.grpcOpts = append(s.grpcOpts, grpc.ChainUnaryInterceptor(interceptors...))
	}
}

func GrpcStreamServerInterceptor(interceptors ...grpc.StreamServerInterceptor) ServerOption {
	return func(s *Server) {
		s.grpcOpts = append(s.grpcOpts, grpc.ChainStreamInterceptor(interceptors...))
	}
}

func GrpcDefaultUnaryServerInterceptor() ServerOption {
	return func(s *Server) {
		s.grpcOpts = append(s.grpcOpts, defaultGrpcOpt(s))
	}
}

func NewServer(app string, opts ...ServerOption) *Server {
	srv := &Server{
		app:      app,
		network:  "tcp",
		address:  ":5080",
		timeout:  time.Second,
		Logger:   defaultLogger(),
		grpcOpts: []grpc.ServerOption{},
	}
	for _, o := range opts {
		o(srv)
	}
	srv.Server = grpc.NewServer(srv.grpcOpts...)
	return srv
}

func (s *Server) Start() error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.Logger.Infof("grpc server listening on: %s", lis.Addr().String())
	return s.Server.Serve(lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop() error {
	s.Server.GracefulStop()
	s.Logger.Info("grpc server stopping")
	return nil
}

func defaultLogger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	return
}

func defaultGrpcOpt(s *Server) (opt grpc.ServerOption) {
	return grpc.ChainUnaryInterceptor(
		recoveryInterceptor(s.Logger),
		timeoutInterceptor(s.Logger),
		logInterceptor(s),
	)
}

func recoveryInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if rerr := recover(); rerr != nil {
				buf := make([]byte, 64<<10)
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				logger.Errorf("recovery: %v: %+v\n%s\n", rerr, req, buf)
				// add err handle
			}
		}()
		return handler(ctx, req)
	}
}

func timeoutInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*60)
		defer cancel()
		return handler(ctx, req)
	}
}

func logInterceptor(s *Server) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var traceId, path, params, method string
		var md metadata.MD
		var ok bool

		md, ok = metadata.FromIncomingContext(ctx)
		if ok {
			if ids, ok := md["traceId"]; ok {
				traceId = ids[0]
			} else {
				traceId = "no-id"
			}
		}
		path = info.FullMethod
		params = req.(fmt.Stringer).String()
		method = "POST"
		logger := s.Logger.WithFields(logrus.Fields{
			"app":     s.app,
			"traceId": traceId,
			"path":    path,
			"method":  method,
			"md":      md,
			"params":  params,
		})
		logger.Infof("receive grpc request")
		ctx = context.WithValue(ctx, "logger", logger)
		resp, err = handler(ctx, req)
		if err != nil {
			logger.Infof("grpc request failled | err: %s", err.Error())
		} else {
			logger.Infof("grpc request success")
		}
		return
	}
}
