package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-uts/internal/config"
	"go-uts/internal/db"
)

const migrationsDir = "migrations"

func main() {
	command := flag.String("command", "up", "command: up|fresh")
	flag.Parse()

	cfg := config.Load()
	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatalf("gagal konek database: %v", err)
	}
	defer pool.Close()

	switch *command {
	case "up":
		if err := migrateUp(ctx, pool); err != nil {
			log.Fatalf("migrasi gagal: %v", err)
		}
		log.Println("migrasi selesai")
	case "fresh":
		if err := dropAll(ctx, pool); err != nil {
			log.Fatalf("fresh gagal: %v", err)
		}
		if err := migrateUp(ctx, pool); err != nil {
			log.Fatalf("migrasi gagal: %v", err)
		}
		log.Println("fresh selesai")
	default:
		log.Fatalf("command tidak dikenal: %s", *command)
	}
}

func migrateUp(ctx context.Context, pool *pgxpool.Pool) error {
	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return err
	}

	files, err := migrationFiles()
	if err != nil {
		return err
	}

	applied, err := appliedMigrations(ctx, pool)
	if err != nil {
		return err
	}

	for _, file := range files {
		if applied[file] {
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsDir, file))
		if err != nil {
			return err
		}

		if err := execStatements(ctx, pool, string(content)); err != nil {
			return fmt.Errorf("%s: %w", file, err)
		}

		_, err = pool.Exec(ctx, "INSERT INTO schema_migrations(filename, applied_at) VALUES ($1, $2)", file, time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationsTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			filename TEXT PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL
		)
	`)
	return err
}

func migrationFiles() ([]string, error) {
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}

	sort.Strings(files)
	return files, nil
}

func appliedMigrations(ctx context.Context, pool *pgxpool.Pool) (map[string]bool, error) {
	rows, err := pool.Query(ctx, "SELECT filename FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			return nil, err
		}
		applied[filename] = true
	}

	return applied, rows.Err()
}

func execStatements(ctx context.Context, pool *pgxpool.Pool, sqlText string) error {
	statements := strings.Split(sqlText, ";")
	for _, statement := range statements {
		trimmed := strings.TrimSpace(statement)
		if trimmed == "" {
			continue
		}
		if _, err := pool.Exec(ctx, trimmed); err != nil {
			return err
		}
	}
	return nil
}

func dropAll(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
		DROP TABLE IF EXISTS pembayaran;
		DROP TABLE IF EXISTS peminjaman;
		DROP TABLE IF EXISTS anggota;
		DROP TABLE IF EXISTS schema_migrations;
	`)
	return err
}
