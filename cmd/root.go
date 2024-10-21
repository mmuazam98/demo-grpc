package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A gRPC server and client in one binary",
	Long:  "This application can run either a gRPC server or client based on the command.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add server and client commands
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
}
