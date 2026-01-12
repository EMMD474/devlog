package cmd

import (
	"fmt"
	"time"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
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
		var output string
		count := 0

		for _, e := range entries {
			if e.Date.Format("2006-01-02") == today {
				line := fmt.Sprintf(
					"%s %s",
					ui.BulletStyle.Render("•"),
					ui.MessageStyle.Render(e.Message),
				)
				output += line + "\n"
				count++
			}
		}

		if count == 0 {
			fmt.Println(ui.EmptyStyle.Render("No entries for today."))
			return nil
		}

		header := ui.HeaderStyle.Render("Today's Entries")
		fmt.Println(header)
		fmt.Print(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
