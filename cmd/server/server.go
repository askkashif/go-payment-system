package server

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment-system-one/internal/api"
	"payment-system-one/internal/repository"
	"time"
)

// Run injects all dependencies needed to run the app
func Run(db *gorm.DB, port string) {
	newRepo := repository.NewDB(db)

	Handler := api.NewHTTPHandler(newRepo)
	router := SetupRouter(Handler, newRepo)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Listening and serving HTTP on : %v\n", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Receive terminate and shutdown gracefully", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// Params is a data model of the data in our environment variable
type Params struct {
	Port  string
	DbUrl string
}

// InitDBParams gets environment variables needed to run the app
func InitDBParams() Params {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")
	if port == "" {
		port = "8080"
	}

	return Params{
		Port:  port,
		DbUrl: dbURL,
	}
}
