package cmd

import (
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all up migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger()
		defer logger.Sync()

		db, err := openDB()
		if err != nil {
			logger.Error("DB open failed", zap.Error(err))
			return err
		}
		defer db.Close()

		if err := goose.Up(db, migrationsDir); err != nil {
			logger.Error("Migration up failed", zap.Error(err))
			return err
		}

		logger.Info("Migrations applied successfully")
		return nil
	},
}
