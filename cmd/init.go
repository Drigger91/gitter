package cmd

import (
	"fmt"
	"os"

	"github.com/Drigger91/gitter/internal"
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

		repo, err := internal.InitRepo(dir)
		if err != nil {
			return err
		}

		fmt.Println("Initialized empty gitter repository in", repo.GitterDir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
