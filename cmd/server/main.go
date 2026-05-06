package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-uts/internal/config"
	"go-uts/internal/db"
	"go-uts/internal/handler"
	"go-uts/internal/repository"
	"go-uts/internal/route"
	"go-uts/internal/usecase"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatalf("gagal konek database: %v", err)
	}
	defer pool.Close()

	peminjamanRepo := repository.NewPeminjamanPGX(pool)
	pembayaranRepo := repository.NewPembayaranPGX(pool)

	peminjamanUC := usecase.NewPeminjamanUsecase(peminjamanRepo)
	pembayaranUC := usecase.NewPembayaranUsecase(pembayaranRepo, peminjamanRepo)

	peminjamanHandler := handler.NewPeminjamanHandler(peminjamanUC)
	pembayaranHandler := handler.NewPembayaranHandler(pembayaranUC)

	mux := route.NewRouter(peminjamanHandler, pembayaranHandler)

	server := &http.Server{
		Addr:              cfg.ServerAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("server berjalan di %s", cfg.ServerAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Printf("gagal shutdown server: %v", err)
	}
}
