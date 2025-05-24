package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"go.uber.org/zap"
)

var (
	driver        string
	dsn           string
	migrationsDir string
	verbose       bool
)

var rootCmd = &cobra.Command{
	Use:   "ose-migrate",
	Short: "ose-migrate - a simple DB migration CLI",
	Long:  "A CLI tool for applying and managing database migrations using Goose internally.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		if driver == "" || dsn == "" || migrationsDir == "" {
			return errors.New("flags --driver, --dsn and --migrations-dir are required")
		}
		return nil
	},
}

func Execute() error {
	// Add commands
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(createCmd)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&driver, "driver", "", "Database driver (required), e.g. postgres")
	rootCmd.PersistentFlags().StringVar(&dsn, "dsn", "", "Database DSN (required)")
	rootCmd.PersistentFlags().StringVar(&migrationsDir, "migrations-dir", "./migrations", "Directory with migration files")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose logging")

	return rootCmd.Execute()
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Try pinging with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}

func newLogger() *zap.Logger {
	if verbose {
		logger, _ := zap.NewDevelopment()
		return logger
	}
	logger, _ := zap.NewProduction()
	return logger
}
