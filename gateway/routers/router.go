package routers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mike955/zebra/gateway/internal/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
	h "github.com/mike955/zebra/pkg/transform/http"
)

func Route(router *mux.Router, logger *logrus.Logger) {
	healthRouter(router)

	pprofRouter(router)

	serviceRouter(router, logger)
}

func serviceRouter(router *mux.Router, logger *logrus.Logger)  {
	router.HandleFunc("*", func(w http.ResponseWriter, req *http.Request) {
		ctx := newServerContext(req, w, logger)
		gateway := new(service.Gateway)
		gateway.SetContext(ctx)
		gateway.BeforeAction()
		gateway.Action()
		gateway.AfterAction()
		return
	})
}

func newServerContext(req *http.Request, w http.ResponseWriter,logger *logrus.Logger) *h.Context {
	traceId := req.Header.Get("traceId")
	if traceId == "" {
		traceId = fmt.Sprintf("%s", uuid.New())
	}
	
	method := req.Method
	ctx := &h.Context{
		Request: req,
		Response: h.ResponseWriter{
			ResponseWriter: w,
		},
		
		ClientId: req.RemoteAddr,
		TraceId:  traceId,
		
		Method:  method,
		BaseUrl: req.URL.Path,
		Url:     req.URL.String(),
	}
	
	ctx.Logger = logger.WithFields(logrus.Fields{
		"traceId":    ctx.TraceId,
		"remoteAddr": ctx.Request.RemoteAddr,
		"method":     method,
		"url":        ctx.Url,
	})
	
	if method == "post" {
		body, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
		
		}
		ctx.Body = body
	}
	return ctx
}

func healthRouter(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func pprofRouter(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	router.HandleFunc("/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	router.HandleFunc("/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	router.HandleFunc("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	router.HandleFunc("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.HandleFunc("/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
	router.HandleFunc("/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
}


