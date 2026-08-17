package main

import (
	"context"
	"gamira/configs"
	"gamira/internal/game"
	"log/slog"
	"os"
)

func main() {
	// logger
	logger := configs.NewLogger()
	slog.SetDefault(logger)

	slog.Info("Seeder starting...")

	// db
	mongo := configs.NewMongoClient()
	dbName := os.Getenv("MONGO_DB_NAME")
	db := mongo.Database(dbName)

	// minio
	minio := configs.NewMinioClient()

	// game seed
	gameRepo := game.NewRepository(db.Collection("game"))
	gameService := game.NewService(gameRepo, minio)
	err := gameService.SeedAll(context.Background())
	if err != nil {
		panic(err)
	}

	slog.Info("Seed completed successfully.")
}
