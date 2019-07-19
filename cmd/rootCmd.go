package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docnoc",
	Short: "Simple notifier for Docker events",
	Long:  "Docnoc scans your containers and notifies you of potential issues",
	Run:   runDocNoc,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Sprintf("Error: Cannot start docnoc: %s", err)
		os.Exit(1)
	}
}

func runDocNoc(*cobra.Command, []string) {
	fmt.Println("Success: docnoc is up and running")
	os.Exit(0)
}
