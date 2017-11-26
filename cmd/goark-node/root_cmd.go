package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCommand will setup and return the root command
func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "example",
		Run: run,
	}

	// this is where we will configure everything!
	rootCmd.Flags().IntP("port", "p", 0, "the port to do things on")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("--- here ---")
}
