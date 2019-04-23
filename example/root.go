package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "app service",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("root")
		return nil
	},
}
