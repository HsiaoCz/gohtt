package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Get("/foo", handleFoo)
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	srv := http.Server{
		Handler:      router,
		Addr:         listenAddr,
		WriteTimeout: time.Millisecond * 1500,
		ReadTimeout:  time.Millisecond * 1500,
	}
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("http server run error", "err", err)
		return
	}
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
