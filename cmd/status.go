package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the gitter repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		// placeholder: later you can read .gitter metadata, etc.
		fmt.Println("On branch main")
		fmt.Println("Nothing to commit, working tree clean")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
