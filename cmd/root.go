package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mothlamp",
	Short: "Mothlamp is a tiny Golang webserver",
	Long: `
It serves http://mothlamp.com
`,
}
