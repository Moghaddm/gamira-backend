package main

import (
	"fmt"
	"gamira/common"
	"gamira/configs"
	"gamira/internal/auth"
	"gamira/internal/divar"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// logger
	logger := configs.NewLogger()
	slog.SetDefault(logger)

	// clients
	redis := configs.NewRedis()
	mongo := configs.NewMongoClient()

	// db
	dbName := os.Getenv("MONGO_DB_NAME")
	db := mongo.Database(dbName)

	// repositories
	userRepo := auth.NewRepository(db.Collection("user"))

	// handlers
	divarHandler := divar.NewHandler(redis)
	authHandler := auth.NewHandler(redis, userRepo)

	// end-points
	http.Handle("POST /divar/start-flow", common.Handler(divarHandler.StartFlow))
	http.Handle("POST /auth/init", common.Handler(authHandler.Init))
	http.Handle("POST /auth/callback", common.Handler(authHandler.Callback))

	// http
	slog.Info("Application starting...")
	port := os.Getenv("APP_PORT")
	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, http.DefaultServeMux)
	if err != nil {
		slog.Error("Error starting server!", err)
	}
}
