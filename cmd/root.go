package cmd

import (
	"fmt"
	"os"

	"github.com/emmd474/devlog/internal/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devlog",
	Short: "Developer activity logger",
	Long:  "Devlog helps you track what you worked on each day.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println(
			ui.BoxStyle.Render(
				ui.HeaderStyle.Render("devlog") + "\n" +
					"Track your daily development progress",
			),
		)
		fmt.Println()

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
}
