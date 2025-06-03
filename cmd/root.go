package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	tickets    string
	version    string
	ticketList []string // Shared variable to store the parsed ticket list
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cobra-cmd",
	Short: "A CLI application for processing and partitioning tickets",
	Long: `A CLI application that provides functionality for processing 
and partitioning JIRA tickets through two commands: process and partition.`,
	Run: func(cmd *cobra.Command, args []string) {
		processCmd.Run(processCmd, args)
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip validation for help and completion commands
		if cmd.Name() == "help" || cmd.Name() == "completion" {
			return nil
		}

		// Validate tickets flag if it's not empty
		if tickets != "" {
			ticketList = strings.Split(tickets, ",")
			for _, ticket := range ticketList {
				if !strings.HasPrefix(ticket, "ABC-") {
					return fmt.Errorf("invalid ticket format: %s, must start with ABC- prefix", ticket)
				}
			}
		}
		return nil
	},
}

func init() {
	// Add persistent flags that will be available to all subcommands
	RootCmd.PersistentFlags().StringVar(&tickets, "tickets", "", "Comma-separated list of JIRA tickets to process (must start with ABC- prefix)")
	RootCmd.PersistentFlags().StringVar(&version, "fix-version", "", "Version for processing")

	// Add commands to the root command
	RootCmd.AddCommand(processCmd)
	RootCmd.AddCommand(partitionCmd)
}

// GetTicketList returns the parsed ticket list for use in any command
func GetTicketList() []string {
	return ticketList
}
