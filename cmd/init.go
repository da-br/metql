/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	"codeburg.com/da-br/metql/internal/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initializes new database",
	Long:  `Initializes a new data at the given location`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		location := viper.GetString("location")
		force := viper.GetBool("force")
		slog.Info("initializing new database", slog.String("location", location), slog.String("name", name), slog.Bool("force", force))
		database.Init(location, name, force)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "overwrites a database if it already exists")
}
