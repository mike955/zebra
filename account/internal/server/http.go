package server

import (
	configs "github.com/mike955/zebra/account/configs"
	"github.com/mike955/zebra/account/internal/service"
	pb "github.com/mike955/zebra/api/account"
	h "github.com/mike955/zrpc/transform/http"
)

func NewHTTPServer() (server *h.Server) {
	config := configs.GlobalConfig.Server
	var opts = []h.ServerOption{
		h.Address(config.HttpAddr),
		h.ReadTimeout(config.Timeout),
	}

	server = h.NewServer(config.AppName, opts...)
	s := service.NewAccountService(server.Logger)
	server.SetHandler(pb.NewAccountServiceHTTPServer(s, server.Logger))
	return server
}

func RunHTTPServer(server *h.Server) (err error) {
	err = server.Run()
	if err != nil {
		server.Logger.Errorf("server start error: %s", err.Error())
	}
	return
}
