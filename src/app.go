package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"mails/src/application/config"
	"mails/src/application/user"
	"mails/src/infrastructure/postgres"
	userRepository "mails/src/infrastructure/postgres/repositories/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("starting application")

	cfg := config.GetConfig()
	ctx := context.Background()

	pool, err := initPostgreSQL(ctx, &cfg.Postgres, logger)
	if err != nil {
		logger.Error("failed to connect to Postgres: %v", err)
		return
	}

	client, err := initRedis(ctx, &cfg.Redis, logger)
	if err != nil {
		logger.Error("failed to connect to Redis: %v", err)
		return
	}

	appUserRepository := userRepository.NewUserRepository(pool)
	userService := user.NewUserService(appUserRepository, logger)
	fmt.Println(userService, client)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(c.Handler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	logger.Info("starting server on port %s", cfg.HTTP.Port)
	err = server.ListenAndServe()
	if err != nil {
		logger.Error("failed to start server: %v", err)
		return
	}
}

func initPostgreSQL(ctx context.Context, cfg *config.PostgresConfig, logger *slog.Logger) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(
		context.Background(),
		postgres.NewConnectionString(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database),
	)
	if err != nil {
		logger.Error("failed to connect to Postgres: %v", err)
		return
	}

	err = pool.Ping(ctx)
	if err != nil {
		logger.Error("failed to ping Postgres: %v", err)
		return
	}

	return pool, nil
}

func initRedis(ctx context.Context, cfg *config.RedisConfig, logger *slog.Logger) (client *redis.Client, err error) {
	logger.Info("connecting to Redis")

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	err = client.Ping(ctx).Err()
	if err != nil {
		logger.Error("failed to ping Redis: %v", err)
		return nil, err
	}

	return client, nil
}
