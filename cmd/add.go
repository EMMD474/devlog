package cmd

import (
	"errors"
	"fmt"

	"github.com/emmd474/devlog/internal/storage"
	"github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
	Long:  "Add a new log entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		message := args[0]
		if message == "" {
			return errors.New("log message cannot be empty")
		}
		if err := storage.SaveEntry(message); err != nil {
			return err
		}
		fmt.Println(ui.SuccessStyle.Render("Entry added successfully."))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}