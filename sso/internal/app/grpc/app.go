package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authrpc "sso/internal/grpc/auth"
	"time"

	"google.golang.org/grpc"
)

type App struct {
	log         *slog.Logger
	gRPCServer  *grpc.Server
	port        int
	storagePath string
	TokenTTL    time.Duration
}

func New(log *slog.Logger, port int, storage string, token time.Duration) *App {
	gRPCServer := grpc.NewServer()
	authrpc.RegisterServerAPI(gRPCServer)

	return &App{
		log:         log,
		gRPCServer:  gRPCServer,
		port:        port,
		storagePath: storage,
		TokenTTL:    token,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op),
		slog.Int("port", a.port))
	log.Info("starting grpc server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op),
		slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
