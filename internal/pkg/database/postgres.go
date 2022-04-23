package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/paul-ss/pgram-backend/internal/pkg/config"
	log "github.com/sirupsen/logrus"
)

type PgxConn interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{},
		f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

var c *pgxpool.Pool

func newConn() *pgxpool.Pool {
	conf := config.C().Postgres

	conn, err := pgxpool.Connect(
		context.Background(),
		fmt.Sprintf("user=%s password=%s host=localhost port=%s database=%s sslmode=disable pool_max_conns=%d",
			conf.User, conf.Password, conf.Port, conf.Database, conf.MaxConns),
	)

	if err != nil {
		panic(err.Error())
	}

	if err := conn.Ping(context.Background()); err != nil {
		panic(err.Error())
	}

	return conn
}

// GetConn return pgx conn, can panic
func GetConn() PgxConn {
	if c == nil {
		c = newConn()
	}

	return c
}

// Init can panic, returns teardown function
func Init() func() {
	if c == nil {
		c = newConn()
	}

	return func() {
		if c == nil {
			log.Info("pgx conn is nil")
			return
		}

		c.Close()
		log.Info("pgx conn is closed")
	}
}
