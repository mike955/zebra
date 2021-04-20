package plugins

import (
	h "github.com/mike955/zebra/pkg/transform/http"
	"github.com/sirupsen/logrus"
)

type auth struct {
	name   string
	logger *logrus.Logger
	ctx    *h.Context
}

func NewAuth(logger *logrus.Logger, ctx *h.Context) *auth {
	return &auth{
		name:   "auth",
		logger: logger,
		ctx:    ctx,
	}
}

func (auth *auth) Handler() (err error) {
	// login
	if auth.ctx.BaseUrl == "/account/login" {
		// call account login

		// add jwt
		return
	}

	// logout
	if auth.ctx.BaseUrl == "/account/logout" {
		// call account login

		// add jwt
		return
	}

	// authentication

	return
}

func (auth *auth) Name() (name string) {
	return auth.name
}
