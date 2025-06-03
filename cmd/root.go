package cmd

import (
	"github.com/spf13/cobra"
)

var (
	tickets string
	version string
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
}

func init() {
	// Add persistent flags that will be available to all subcommands
	RootCmd.PersistentFlags().StringVar(&tickets, "tickets", "", "Comma-separated list of JIRA tickets to process")
	RootCmd.PersistentFlags().StringVar(&version, "version", "", "Version for processing")

	// Add commands to the root command
	RootCmd.AddCommand(processCmd)
	RootCmd.AddCommand(partitionCmd)
}
