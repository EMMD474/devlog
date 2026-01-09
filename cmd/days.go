package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
	"github.com/emmd474/devlog/internal/storage"
)

var duration int

var daysCmd = &cobra.Command{
	Use: "days",
	Short: "Display logs for a specified number of days",
	RunE: func(cmd *cobra.Command, args []string) error {
		if duration < 1 {
			duration = 1
		}
		entries, err := storage.LoadEntriesSince(time.Duration(duration) * 24 * time.Hour)
		if err != nil {
			return nil
		}

		
		if len(entries) == 0 {
			fmt.Println("No logs for this duration")
			return nil
		}

		for _, entry := range entries {
			fmt.Printf("%s  %s\n", entry.Date.Format("2006-01-02"), entry.Message)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(daysCmd)

	daysCmd.Flags().IntVarP(&duration, "duration", "d", 1, "Number of days")
}