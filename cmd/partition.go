package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	filename string
)

// partitionCmd represents the partition command
var partitionCmd = &cobra.Command{
	Use:   "partition",
	Short: "Partition JIRA tickets",
	Long:  `Partition JIRA tickets into a specified file or default file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Partition command called")

		partitionTickets := GetTicketList()
		if len(partitionTickets) > 0 {
			fmt.Printf("Partitioning tickets: %s\n", strings.Join(partitionTickets, ","))
		} else {
			fmt.Println("No tickets specified")
		}

		fmt.Printf("Using filename: %s\n", filename)
	},
}

func init() {
	partitionCmd.Flags().StringVar(&filename, "filename", "default_partition.txt", "Output filename for partitioned tickets")
	// Remove the local tickets flag as we're using the global one now
}
