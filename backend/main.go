package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"auto-parts-catalog/backend/api"
	"auto-parts-catalog/backend/catalogservice/grpcimpl"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// interceptorLogger adapts go-kit logger to interceptor logger.
func interceptorLogger(l kitlog.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = level.Debug(l).Log(largs...)
		case logging.LevelInfo:
			_ = level.Info(l).Log(largs...)
		case logging.LevelWarn:
			_ = level.Warn(l).Log(largs...)
		case logging.LevelError:
			_ = level.Error(l).Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {

	// Set up a connection to the catatalog service
	logger := kitlog.NewLogfmtLogger(os.Stdout)
	conn, err := grpc.NewClient(
		os.Getenv("CATALOG_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(interceptorLogger(logger))),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(interceptorLogger(logger))),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	catalogService := grpcimpl.NewCatalogService(conn)

	// Set up server
	server := api.NewServer(catalogService)
	err = server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
