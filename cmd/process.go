package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	tickets string
	version string
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process JIRA tickets",
	Long:  `Process JIRA tickets with optional version specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Process command called")

		if tickets != "" {
			fmt.Printf("Processing tickets: %s\n", tickets)
		} else {
			fmt.Println("No tickets specified")
		}

		if version != "" {
			fmt.Printf("Version: %s\n", version)
		} else {
			fmt.Println("No version specified")
		}
	},
}

func init() {
	processCmd.Flags().StringVar(&tickets, "tickets", "", "Comma-separated list of JIRA tickets to process")
	processCmd.Flags().StringVar(&version, "version", "", "Version for processing")
}
