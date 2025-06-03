package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process JIRA tickets",
	Long:  `Process JIRA tickets with optional version specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Process command called")

		processTickets := GetTicketList()
		if len(processTickets) > 0 {
			fmt.Printf("Processing tickets: %s\n", strings.Join(processTickets, ","))
		} else {
			fmt.Println("No tickets specified")
		}

		if version != "" {
			fmt.Printf("Fix-version: %s\n", version)
		} else {
			fmt.Println("No fix-version specified")
		}
	},
}

func init() {
	// Removed flag declarations as they're now defined as persistent flags in root.go
}
