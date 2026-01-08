package cmd

import (
	"fmt"
	"time"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all log entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := storage.LoadEntries()
		if err != nil {
			return err
		}

		for _, e := range entries {
			fmt.Printf(
				"[%s] %s\n",
				e.Date.Format(time.RFC822),
				e.Message,
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
