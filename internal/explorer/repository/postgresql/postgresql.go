package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"remy_explorer/internal/config"
	"remy_explorer/internal/utils"
	"time"
)

// Client is a subset of the pgx.Conn interface.
// It provides methods for executing SQL queries and transactions.
type Client interface {
	// Exec executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	// Query sends a query to the database and returns the rows.
	// The args are for any placeholder parameters in the query.
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// QueryRow sends a query to the database and returns a single row.
	// The args are for any placeholder parameters in the query.
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	// Begin starts a new transaction.
	Begin(ctx context.Context) (pgx.Tx, error)
}

// New creates a new Client from a pgx.Conn.
// It connects to the database using the provided connection details and returns a pool of connections.
// The function will attempt to connect to the database maxAttempts times before failing.
func New(ctx context.Context, conn config.StorageConfig, maxAttempts int) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.Host, conn.User, conn.Password, conn.Database)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return pool, nil
}

//CREATE TABLE IF NOT EXISTS public.folder
//(
// id         BIGSERIAL PRIMARY KEY,
// owner_id   BIGINT                  NOT NULL,
// name       VARCHAR(255)            NOT NULL,
// parent_id  BIGINT,
// created_at TIMESTAMP DEFAULT NOW() NOT NULL,
// updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
// CONSTRAINT fk_folder_parent_id FOREIGN KEY (parent_id) REFERENCES public.folder (id)
//);
//
//CREATE TABLE IF NOT EXISTS public.file
//(
// id          BIGSERIAL PRIMARY KEY,
// owner_id    BIGINT                  NOT NULL,
// name        VARCHAR(255)            NOT NULL,
// folder_id   BIGINT                  NOT NULL,
// object_path VARCHAR(255)            NOT NULL,
// size        INT                     NOT NULL,
// type        VARCHAR(255)            NOT NULL,
// created_at  TIMESTAMP DEFAULT NOW() NOT NULL,
// updated_at  TIMESTAMP DEFAULT NOW() NOT NULL,
// tags        JSONB,
// CONSTRAINT fk_file_folder_id FOREIGN KEY (folder_id) REFERENCES public.folder (id)
//);
