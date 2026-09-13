package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/gorkagg10/equity-calculator-api/internal/config"
	"github.com/gorkagg10/equity-calculator-api/internal/http/handler"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/http/yahoofinance"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
	"github.com/gorkagg10/equity-calculator-api/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	conf, err := config.NewConfig()
	if err != nil {
		slog.ErrorContext(ctx, "loading config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = postgres.Migrate(&conf.DatabaseConfig)
	if err != nil {
		slog.ErrorContext(ctx, "running database migrations", slog.String("error", err.Error()))
		os.Exit(1)
	}

	pgClient, err := postgres.NewDatabaseClient(ctx, &conf.DatabaseConfig)
	if err != nil {
		slog.ErrorContext(ctx, "connecting to database client", slog.String("error", err.Error()))
		os.Exit(1)
	}

	portfolioRepository := postgres.NewPortfolioRepository(pgClient)
	portfolioService := service.NewPortfolio(portfolioRepository)
	portfolioHandler := handler.NewPortfolio(portfolioService)

	assetRepository := postgres.NewAssetRepository(pgClient)
	assetDataRepository := yahoofinance.NewAssetDataRepository()
	assetService := service.NewAsset(assetRepository, assetDataRepository)
	assetHandler := handler.NewAsset(assetService)

	router := chi.NewRouter()
	router.Use(chimw.RequestID)
	router.Use(chimw.Recoverer)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", portfolioHandler.Routes())
		r.Mount("/assets", assetHandler.Routes())
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("server starting", "port", "8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	err = shutdownServer(srv)
	if err != nil {
		slog.Error("graceful shutdown failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func shutdownServer(srv *http.Server) error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
