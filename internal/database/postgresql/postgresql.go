package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type StorageConfig struct {
	username, password, host, port, database string
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, sc StorageConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", sc.username, sc.password, sc.host, sc.database)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Println("failed to connect to postgresql: ", err)
		return nil, err
	}
	return pool, nil
}
