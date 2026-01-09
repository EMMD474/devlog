package cmd

import (
	"fmt"
	"time"
	"github.com/emmd474/devlog/internal/storage"
	"github.com/spf13/cobra"

)

var weekCmd = &cobra.Command{
	Use: "week",
	Short: "Show logs for the last 7 days",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := storage.LoadEntriesSince(7 * 24 * time.Hour)
		if err != nil {
			return err
		}
		
		if len(entries) == 0 {
			fmt.Println("No entries")
			return nil
		}

		for _, entry := range entries {
			fmt.Printf(
				"%s  %s\n",
				entry.Date.Format("2006-01-02"),
				entry.Message,
			)
		}
		
		return nil
	},

}

func init() {
	rootCmd.AddCommand(weekCmd)
}