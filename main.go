package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/donbattery/mothlamp/cmd"
	"github.com/donbattery/mothlamp/utils"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("Error during execution : %v\n", err)
	}
}

func init() {
	// Load config from mothlamp.yaml and ENV
	if err := utils.ReadConfig("./", "mothlamp", nil); err != nil {
		fmt.Printf("Cannot set configuration %v\n", err)
	}
	// Setup global flags
	cmd.RootCmd.PersistentFlags().StringP("dir", "d", "./drive", "Directory to be served by Mothlamp")
	cmd.RootCmd.PersistentFlags().StringP("token", "t", "", "Mothlamp webserver API token")
	cmd.RootCmd.PersistentFlags().StringP("url", "u", "", "Mothlamp webserver URL")
	if err := viper.BindPFlag("MOTHLAMP_DIR", cmd.RootCmd.PersistentFlags().Lookup("dir")); err != nil {
		fmt.Printf("Cannot bind flag MOTHLAMP_DIR %v\n", err)
	}
	if err := viper.BindPFlag("MOTHLAMP_API_TOKEN", cmd.RootCmd.PersistentFlags().Lookup("token")); err != nil {
		fmt.Printf("Cannot bind flag MOTHLAMP_API_TOKEN %v\n", err)
	}
	if err := viper.BindPFlag("MOTHLAMP_API_URL", cmd.RootCmd.PersistentFlags().Lookup("url")); err != nil {
		fmt.Printf("Cannot bind flag MOTHLAMP_API_URL %v\n", err)
	}
}
