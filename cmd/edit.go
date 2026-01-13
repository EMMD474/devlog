package cmd

import (
	"fmt"
	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
	"github.com/emmd474/devlog/internal/tui/edit"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use: "edit",
	Short: "edit a log",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := storage.LoadEntries()
		if err != nil {
			return err
		}

		if len(data) < 1 {
			fmt.Println(ui.EmptyStyle.Render("No log entries found."))
			return nil
		}

		return edit.Run(data)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}