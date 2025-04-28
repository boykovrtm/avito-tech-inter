package repositories

import (
	"avito-tech-inter/domain/user"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Add(ctx context.Context, user *user.User, hash string) error {
	userQuery := "insert into users (username, balance) values ($1, $2);"
	userAuthQuery := "insert into users_auth (username, hash) values ($1, $2);"
	transactionsQuery := "insert into transactions (from, to, amount, moment) values ($1, $2, $3, $4);"

	tx, err := r.pool.Begin(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	if err != nil {
		return err
	}

	_, err = tx.Conn().Exec(ctx, userQuery, user.GetUsername(), user.GetBalance())
	_, err = tx.Conn().Exec(ctx, userAuthQuery, user.GetUsername(), hash)
	for _, transaction := range user.GetTransactions() {
		_, err = tx.Conn().Exec(ctx,
			transactionsQuery,
			transaction.GetFrom(),
			transaction.GetTo(),
			transaction.GetAmount(),
			transaction.GetMoment())
	}
	_, err = tx.Conn().Exec(ctx, transactionsQuery)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return err
	}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	userQuery := "select username, balance from users_auth where username = $1"
	userTransactionQuery := "select id, username_from, username_to, amount, moment from transactions where username_from = $1 or username_to = $1"

	var (
		dbUsername string
		dbBalance  int32
	)

	result := user.User{}

	row := r.pool.QueryRow(ctx, userQuery, username, userTransactionQuery)
	row.Scan(&dbUsername, &dbBalance)
}

func (r *UserRepository) FindHashByUsername(ctx context.Context, username string) (*string, error) {
	query := "select hash from users_auth where username = $1"

	var result string

	row := r.pool.QueryRow(ctx, query, username)
	err := row.Scan(&result)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}
