package cmd

import (
	"fmt"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/tui/delete"
	"github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a log entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := storage.LoadEntries()
		if err != nil {
			return err
		}

		if len(entries) == 0 {
			fmt.Println(ui.EmptyStyle.Render("No log entries found."))
			return nil
		}

		return delete.Run(entries)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
