package main

import (
	"avito-tech-inter/api"
	"avito-tech-inter/facade"
	"avito-tech-inter/infra/repositories"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("PG_CONN_STR"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	authRepo := repositories.NewAuthRepository(pool)
	controller := facade.NewController(authRepo)
	}

	srv, err := api.NewServer(&controller)
}
