package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/emmd474/devlog/internal/storage"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new entry",
	Long: "Add a new log entry",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		message := args[0]
		if message == "" {
			return errors.New("log message cannot be empty")
		}
		return storage.SaveEntry(message)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}