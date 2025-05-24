package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create new empty migration files (up/down)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		logger := newLogger()
		defer logger.Sync()

		absDir, err := filepath.Abs(migrationsDir)
		if err != nil {
			logger.Error("Failed to get absolute migrations dir", zap.Error(err))
			return err
		}

		if err := goose.Create(nil, absDir, name, "sql"); err != nil {
			logger.Error("Failed to create migration", zap.Error(err))
			return err
		}

		logger.Info("Migration files created", zap.String("name", name), zap.String("dir", absDir))
		fmt.Printf("Created new migration files in %s:\n - %s.up.sql\n - %s.down.sql\n", absDir, name, name)

		return nil
	},
}
