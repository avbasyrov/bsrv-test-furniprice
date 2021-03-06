package dbcon

import (
	"context"
	"fmt"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

type Db struct {
	Sqlx *sqlx.DB
}

func New(ctx context.Context, cfg config.DbConfig) *Db {
	return &Db{
		Sqlx: connect(ctx, cfg),
	}
}

func connect(ctx context.Context, cfg config.DbConfig) *sqlx.DB {
	dsn := "user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=6"
	dsn = fmt.Sprintf(dsn, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Unable to parse Sqlx config:", err)
		os.Exit(1)
	}
	// PgBouncer workaround
	pgxConfig.ConnConfig.RuntimeParams["standard_conforming_strings"] = "on"
	pgxConfig.ConnConfig.PreferSimpleProtocol = true
	//pgxConfig.ConnConfig.CustomCancel = func(_ *pgx.Conn) error { return nil }
	//connConfig.Logger = myLogger
	connStr := stdlib.RegisterConnConfig(pgxConfig.ConnConfig)

	var db *sqlx.DB
	for try := 1; try <= 10; try = try + 1 {
		log.Printf("Try #%d to connect DB...\n", try)

		db, err = sqlx.Connect("pgx", connStr)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Unable to connect Sqlx:", err)
			time.Sleep(3 * time.Second)
		} else {
			log.Println("Successfully connected to DB")
			break
		}
	}

	if db == nil {
		os.Exit(1)
	}

	return db
}
