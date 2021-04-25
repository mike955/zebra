package service

import (
	"github.com/mike955/zebra/gateway/internal/plugins"
	h "github.com/mike955/zebra/pkg/transform/http"
)

type Gateway struct {
	ctx     *h.Context
	plugins []plugins.Plugin
}

func (gate *Gateway) SetContext(ctx *h.Context) {
	gate.ctx = ctx
}

func (gate *Gateway) BeforeAction() {
	// register plugins
	gate.plugins = []plugins.Plugin{
		plugins.NewWaf(gate.ctx.Logger.Logger, gate.ctx),
		plugins.NewLimit(gate.ctx.Logger.Logger, gate.ctx),
		plugins.NewAuth(gate.ctx.Logger.Logger, gate.ctx),
	}
}

func (gate *Gateway) Action() {
	for i := 0; i < len(gate.plugins); i++ {
		plugin := gate.plugins[i]
		err := plugin.Handler()
		if err != nil {
			gate.ctx.Response.Write([]byte(err.Error()))
			break
		}
	}
}

func (gate *Gateway) AfterAction() {
}
