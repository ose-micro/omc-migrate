package cmd

import (
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := newLogger()
		defer logger.Sync()

		db, err := openDB()
		if err != nil {
			logger.Error("DB open failed", zap.Error(err))
			return err
		}
		defer db.Close()

		if err := goose.Down(db, migrationsDir); err != nil {
			logger.Error("Migration down failed", zap.Error(err))
			return err
		}

		logger.Info("Last migration rolled back successfully")
		return nil
	},
}
