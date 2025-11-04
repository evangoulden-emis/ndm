package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the `add` subcommand of `deviceCmd`
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new device to the system",
	Long: `Add a new managed device to the system.

Examples:
  appname deviceCmd add --name MyDevice --ip_address 192.168.1.1
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flag values
		name, _ := cmd.Flags().GetString("name")
		ipAddress, _ := cmd.Flags().GetString("ip_address")

		if name == "" || ipAddress == "" {
			fmt.Println("Both --name and --ip_address are required.")
			return
		}

		// Add deviceCmd logic here
		fmt.Printf("Device added: Name=%s, IP Address=%s\n", name, ipAddress)
	},
}

func init() {
	// Add "add" as a subcommand under "deviceCmd"
	deviceCmd.AddCommand(addCmd)

	// Define flags for `add` subcommand
	addCmd.Flags().String("name", "", "Name of the deviceCmd (required)")
	addCmd.Flags().String("ip_address", "", "IP address of the deviceCmd (required)")

	// Make flags required
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("ip_address")
}
