package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:     "admin",
	Aliases: []string{"u"},
	Short:   "Root of admin command. Subcommands: set, ",
	Run:     admin,
}

func init() {
	RootCmd.AddCommand(adminCmd)
}

func admin(cmd *cobra.Command, args []string) {
	fmt.Println("admin called")
}
