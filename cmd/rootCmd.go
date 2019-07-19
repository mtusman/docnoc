package cmd

import (
	"fmt"
	"os"

	"docnoc_priv/pkg"

	"github.com/spf13/cobra"
)

var (
	flags   = pkg.NewFlags()
	rootCmd = &cobra.Command{
		Use:   "docnoc",
		Short: "Simple notifier for Docker events",
		Long:  "Docnoc scans your containers and notifies you of potential issues",
		Run:   runDocNoc,
	}
)

func init() {
	rootCmd.Flags().StringVarP(
		flags.ConfigFile,
		"file",
		"f",
		"",
		"Location of the docnoc config file",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Sprintf("Error: Cannot start docnoc: %s", err)
		os.Exit(1)
	}
}

func runDocNoc(*cobra.Command, []string) {
	fmt.Println("Success: docnoc is up and running")
	fmt.Println(*(flags.ConfigFile))
	os.Exit(0)
}
