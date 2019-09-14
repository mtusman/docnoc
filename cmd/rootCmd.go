package cmd

import (
	"fmt"
	"os"

	"github.com/mtusman/docnoc/pkg"

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

// Execute is used to run the default cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: Cannot start docnoc: ", err)
		os.Exit(1)
	}
}

// runDocNoc starts the process of scanning and outputting the report
func runDocNoc(*cobra.Command, []string) {
	dN := pkg.NewDocNoc(flags)
	dN.StartScrubbingDefault()
	dN.ProcessReport()
}
