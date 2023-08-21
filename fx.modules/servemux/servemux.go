package servemux

import (
	"net/http"

	"go.uber.org/fx"
)

const SERVE_MUX_GROUP_NAME = `group:"routes"`

type Route interface {
	http.Handler
	Pattern() string
}

func NewServerMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}

	return mux
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewServerMux,
			fx.ParamTags(SERVE_MUX_GROUP_NAME),
		),
	),
)
