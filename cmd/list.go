package cmd

import (
	"fmt"
	"time"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
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

		if len(entries) == 0 {
			fmt.Println(ui.EmptyStyle.Render("No log entries found."))
			return nil
		}

		fmt.Println(ui.HeaderStyle.Render("All Entries"))
		for _, e := range entries {
			fmt.Printf(
				"%s %s  %s\n",
				ui.BulletStyle.Render("•"),
				ui.DateStyle.Render(e.Date.Format(time.RFC822)),
				ui.MessageStyle.Render(e.Message),
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
