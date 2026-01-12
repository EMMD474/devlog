package cmd

import (
	"fmt"
	"time"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"
)

var duration int

var daysCmd = &cobra.Command{
	Use:   "days",
	Short: "Display logs for a specified number of days",
	RunE: func(cmd *cobra.Command, args []string) error {
		if duration < 1 {
			duration = 1
		}
		entries, err := storage.LoadEntriesSince(time.Duration(duration) * 24 * time.Hour)
		if err != nil {
			return err
		}

		if len(entries) == 0 {
			fmt.Println(ui.EmptyStyle.Render(fmt.Sprintf("No logs for the last %d day(s).", duration)))
			return nil
		}

		header := fmt.Sprintf("Entries from the last %d day(s)", duration)
		fmt.Println(ui.HeaderStyle.Render(header))
		for _, entry := range entries {
			fmt.Printf(
				"%s %s  %s\n",
				ui.BulletStyle.Render("•"),
				ui.DateStyle.Render(entry.Date.Format("2006-01-02")),
				ui.MessageStyle.Render(entry.Message),
			)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(daysCmd)

	daysCmd.Flags().IntVarP(&duration, "duration", "d", 1, "Number of days")
}