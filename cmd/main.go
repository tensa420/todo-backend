package main

import (
	"context"
	"log"
	"time"
	"todo-backend/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = srv.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
