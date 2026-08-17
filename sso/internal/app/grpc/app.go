package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authrpc "sso/internal/grpc/auth" // Проверь, чтобы путь совпадал с твоим

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New создает новое gRPC приложение.
// Обрати внимание: мы принимаем authrpc.Auth — это наша бизнес-логика!
func New(log *slog.Logger, authService authrpc.Auth, port int) *App {
	gRPCServer := grpc.NewServer()

	// Регистрируем хендлеры, передавая им сервис бизнес-логики
	// Убедись, что функция в файле server.go называется именно Register!
	authrpc.RegisterServerAPI(gRPCServer, authService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
