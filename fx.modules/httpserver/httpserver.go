package httpserver

import (
	"context"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newHttpServer(lc fx.Lifecycle, mux *http.ServeMux, logger *zap.Logger) *http.Server {
	server := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error { return onStartHook(ctx, server, logger) },
			OnStop:  func(ctx context.Context) error { return onStopHook(ctx, server, logger) },
		},
	)

	return server
}

func onStartHook(ctx context.Context, server *http.Server, logger *zap.Logger) error {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	logger.Info("Starting HTTP server", zap.String("addr", server.Addr))
	go server.Serve(ln)

	return nil
}

func onStopHook(ctx context.Context, server *http.Server, logger *zap.Logger) error {
	logger.Info("Stopping HTTP service")
	return server.Shutdown(ctx)
}

var Module = fx.Options(
	fx.Provide(newHttpServer),
)
