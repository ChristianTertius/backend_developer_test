package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChristianTertius/backend_developer_test/internal/config"
	"github.com/ChristianTertius/backend_developer_test/internal/database"
	nationalityHTTP "github.com/ChristianTertius/backend_developer_test/internal/nationality/delivery/http"
	nationalityRepo "github.com/ChristianTertius/backend_developer_test/internal/nationality/repository"
	nationalityUC "github.com/ChristianTertius/backend_developer_test/internal/nationality/usecase"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("koneksi database gagal: %v", err)
	}
	defer db.Close()

	const timeout = 10 * time.Second

	nRepo := nationalityRepo.NewPostgresNationalityRepository(db)
	nUC := nationalityUC.NewNationalityUseCase(nRepo, timeout)

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	nationalityHTTP.NewNationalityHandler(api, nUC)

	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Printf("server running at http://localhost:%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server shutdown ....")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
}
