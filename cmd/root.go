package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitter",
	Short: "gitter is a simple git-like CLI tool",
	Long:  "gitter is a simple git-like CLI tool written in Go.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// place to add global flags, if needed
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}
