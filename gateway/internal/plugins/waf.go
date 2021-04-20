package plugins

import (
	h "github.com/mike955/zebra/pkg/transform/http"
	"github.com/sirupsen/logrus"
)

type waf struct {
	name   string
	logger *logrus.Logger
	ctx    *h.Context
}

func NewWaf(logger *logrus.Logger, ctx *h.Context) *waf {
	return &waf{
		name:   "waf",
		logger: logger,
		ctx:    ctx,
	}
}

func (waf *waf) Handler() (err error) {
	return
}

func (waf *waf) Name() (name string) {
	return waf.name
}
