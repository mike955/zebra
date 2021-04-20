package plugins

import (
	h "github.com/mike955/zebra/pkg/transform/http"
	"github.com/sirupsen/logrus"
)

type limit struct {
	name   string
	logger *logrus.Logger
	ctx    *h.Context
}

func NewLimit(logger *logrus.Logger, ctx *h.Context) *limit {
	return &limit{
		name:   "limit",
		logger: logger,
		ctx:    ctx,
	}
}

func (limit *limit) Handler() (err error) {
	return
}

func (limit *limit) Name() (name string) {
	return limit.name
}
