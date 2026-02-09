//go:build integration

package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("bastet_test"),
		postgres.WithUsername("test"),
		postgres.WithPassword("test"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("failed to start postgres container: %v", err)
	}

	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Printf("failed to terminate container: %v", err)
		}
	}()

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get connection string: %v", err)
	}

	// Run migrations
	// Use golang-migrate or embed SQL directly
	if err := runMigrations(connStr); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Create pool
	pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Close()

	// Run tests
	os.Exit(m.Run())
}

func runMigrations(connStr string) error {
	// Option 1: Use golang-migrate programmatically
	// m, err := migrate.New("file://../../../../migrations", connStr)
	// if err != nil { return err }
	// return m.Up()

	// Option 2: Execute SQL directly from embedded files
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return fmt.Errorf("failed to connect for migrations: %w", err)
	}
	defer conn.Close()

	// Read and execute migration files
	// migrationSQL, _ := os.ReadFile("../../../../migrations/000001_create_notifications.up.sql")
	// _, err = conn.Exec(ctx, string(migrationSQL))
	return err
}
