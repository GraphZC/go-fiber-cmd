/*
Copyright © 2023 TANAROEG O-CHAROEN <graph234@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/GraphZC/go-fiber-cmd/configs"
	"github.com/GraphZC/go-fiber-cmd/modules/migrate"
	"github.com/spf13/cobra"
)

// migrateDownCmd represents the migrateDown command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Down migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Check migration file is exists
		if !migrate.IsMigrationFileIsExists() {
			log.Fatal("Error: migrations/migrations.csv is not exists please run `migrate init`")
		}

		migrations, err := migrate.Read()

		if err != nil {
			log.Fatal(err)
		}

		n, err := migrations.Down()
		if err != nil {
			log.Fatal(err)
		}

		if err := migrations.Done(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Migrate down %d migration successfully\n", n)
	},
}

func init() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("Could not load the config: ", err)
	}

	if err := configs.ConnectDB(); err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	migrateCmd.AddCommand(migrateDownCmd)
}
