package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Using a different variable for tickets to avoid conflicts
	partitionTickets string
	filename         string
)

// partitionCmd represents the partition command
var partitionCmd = &cobra.Command{
	Use:   "partition",
	Short: "Partition JIRA tickets",
	Long:  `Partition JIRA tickets into a specified file or default file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Partition command called")

		if partitionTickets != "" {
			fmt.Printf("Partitioning tickets: %s\n", partitionTickets)
		} else {
			fmt.Println("No tickets specified")
		}

		fmt.Printf("Using filename: %s\n", filename)
	},
}

func init() {
	partitionCmd.Flags().StringVar(&partitionTickets, "tickets", "", "Comma-separated list of JIRA tickets to partition")
	partitionCmd.Flags().StringVar(&filename, "filename", "default_partition.txt", "Output filename for partitioned tickets")
}
