package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of all migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger()
		defer logger.Sync()

		db, err := openDB()
		if err != nil {
			logger.Error("DB open failed", zap.Error(err))
			return err
		}
		defer db.Close()

		if err := goose.Status(db, migrationsDir); err != nil {
			logger.Error("Migration status failed", zap.Error(err))
			return err
		}

		return nil
	},
}
