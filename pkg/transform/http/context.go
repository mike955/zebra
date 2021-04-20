package http

import (
	"net/http"
	"net/url"
	
	"github.com/sirupsen/logrus"
)

type Context struct {
	Request  *http.Request
	Response ResponseWriter

	ClientId string
	TraceId  string
	Uid      string
	err  		bool

	Method  string
	Body    []byte
	Level   uint64
	BaseUrl string
	Url     string
	Values url.Values

	Logger *logrus.Entry
}
