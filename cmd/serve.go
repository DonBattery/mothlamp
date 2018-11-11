package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/donbattery/mothlamp/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Web server",
	Run:   serve,
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 8080, "Mothlamp webserver's PORT")
	if err := viper.BindPFlag("MOTHLAMP_PORT", serveCmd.Flags().Lookup("port")); err != nil {
		fmt.Printf("Cannot bind flag port %v\n", err)
	}
}

func serve(cmd *cobra.Command, args []string) {
	server.Run()
}
