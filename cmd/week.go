package cmd

import (
	"fmt"
	"time"
	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"

)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show logs for the last 7 days",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := storage.LoadEntriesSince(7 * 24 * time.Hour)
		if err != nil {
			return err
		}

		if len(entries) == 0 {
			fmt.Println(ui.HeaderStyle.Render("No entries in the last 7 days."))
			return nil
		}

		var output string
		for _, entry := range entries {
			line := fmt.Sprintf(
				"%s  %s",
				ui.DateStyle.Render(entry.Date.Format("2006-01-02")),
				ui.MessageStyle.Render(entry.Message),
			)
			output += line + "\n"
		}

		fmt.Println(ui.BoxStyle.Render(output))
		return nil
	},
}


func init() {
	rootCmd.AddCommand(weekCmd)
}