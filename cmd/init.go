package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new gitter repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		gitterDir := filepath.Join(dir, ".gitter")

		if _, err := os.Stat(gitterDir); err == nil {
			fmt.Println(".gitter already exists")
			return nil
		}

		if err := os.MkdirAll(gitterDir, 0755); err != nil {
			return err
		}

		fmt.Println("Initialized empty gitter repository in", gitterDir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
