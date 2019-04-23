package main

import (
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/gocmds"
	"github.com/spf13/cobra"
	"os"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
	},
}

func main() {
	rootCmd := RootCmd
	rootCmd.AddCommand(
		VersionCmd,
	)
	assert.MustNotError(gocmds.PrepareBaseCmd(rootCmd, "cmd",
		os.ExpandEnv("$PWD/kdata")).Execute())
}
