package cmd

import (
	"fmt"
	"time"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's log entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := storage.LoadEntries()
		if err != nil {
			return err
		}

		today := time.Now().Format("2006-01-02")

		for _, e := range entries {
			if e.Date.Format("2006-01-02") == today {
				fmt.Printf("- %s\n", e.Message)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
