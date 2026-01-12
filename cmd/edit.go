package cmd

import (
	"fmt"
	// "github.com/emmd474/devlog/internal/storage"
	// "github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use: "edit",
	Short: "edit a log",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Edit a log!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}