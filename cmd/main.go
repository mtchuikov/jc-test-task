package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	v1handlers "github.com/mtchuikov/jc-test-task/internal/handlers/v1"
	"github.com/mtchuikov/jc-test-task/internal/repo/postgres"
	"github.com/mtchuikov/jc-test-task/internal/services"
	"github.com/mtchuikov/jc-test-task/pkg/closer"
	"github.com/mtchuikov/jc-test-task/pkg/middlewares"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const pgConnURL = "postgres://user:password@127.0.0.1:5432/postgres?sslmode=disable"

func main() {
	rootCtx := context.Background()

	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}
	stopCtx, stop := signal.NotifyContext(rootCtx, signals...)
	defer stop()

	closer.InitGlobal()
	defer func() {
		timeoutCtx, cancel := context.WithTimeout(rootCtx, 3*time.Second)
		closer.Global.Close(timeoutCtx)
		cancel()
	}()

	zerolog.LevelFieldName = "lvl"
	zerolog.ErrorFieldName = "err"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = time.RFC1123

	log.Logger = log.Logger.
		Level(zerolog.InfoLevel).With().
		Timestamp().Logger()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	compress := middleware.Compress(5)
	router.Use(middlewares.Decompress, compress)

	apiV1Router := chi.NewRouter()
	router.Mount("/api/v1", apiV1Router)

	pgConn, err := pgx.Connect(rootCtx, pgConnURL)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	closer.Global.Add(closer.Task{
		Sync: false,
		Fn: func(ctx context.Context) {
			log.Info().Msg("shutting down pg conn...")
			pgConn.Close(ctx)
			log.Info().Msg("pg conn shut down")
		},
	})

	txsRepo := postgres.NewTransactions(pgConn)
	transactor := services.NewTransactor(log.Logger, txsRepo)
	v1handlers.RegisterTransact(apiV1Router, transactor)

	balancesRepo := postgres.NewBalances(pgConn)
	balanceGetter := services.NewBalanceGetter(log.Logger, balancesRepo)
	v1handlers.RegisterGetBalance(apiV1Router, balanceGetter)

	server := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	closer.Global.AddWithPriority(0, closer.Task{
		Sync: true,
		Fn: func(ctx context.Context) {
			log.Info().Msg("shutting down http server...")
			server.Shutdown(ctx)
			log.Info().Msg("http server shut down")
		},
	})

	log.Info().Msg("listening http server...")
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen server")
		}
	}()

	<-stopCtx.Done()
}
