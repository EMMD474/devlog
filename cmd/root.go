package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "devlog",
	Short: "Developer activity logger",
	Long: "Devlog helps you track what you worked on each day.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Println(err)
		os.Exit(1)
	}
}