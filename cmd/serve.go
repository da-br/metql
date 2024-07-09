package cmd

import (
	"log"

	"codeburg.com/da-br/metql/internal/database"
	"codeburg.com/da-br/metql/internal/tcp"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := database.LoadConfig()

		db := database.NewDatabase(config.DatabaseFile)
		err := db.Start()
		if err != nil {
			log.Fatalf("Error loading database: %v", err)
		}
		defer func() {
			if err := db.Exit(); err != nil {
				log.Fatalf("Error saving database: %v", err)
			}
		}()

		server := tcp.NewServer(db)
		err = server.ListenAndServe(config.ServerAddress)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Define flags and configuration settings.
}
