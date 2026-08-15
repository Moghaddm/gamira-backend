package main

import (
	"fmt"
	"gamira/cache"
	"gamira/common"
	"gamira/internal/auth"
	"gamira/internal/divar"
	"gamira/log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := log.NewLogger()
	slog.SetDefault(logger)

	redis := cache.NewRedis()

	divarHandler := divar.NewHandler(redis)
	authHandler := auth.NewHandler(redis)

	// divar
	http.Handle("/divar/start-flow", common.Handler(divarHandler.StartFlow))
	http.Handle("/auth/init", common.Handler(authHandler.Init))

	slog.Info("Application starting...")
	port := os.Getenv("APP_PORT")
	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, http.DefaultServeMux)
	if err != nil {
		slog.Error("Error starting server!", err)
	}
}
