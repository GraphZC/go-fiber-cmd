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

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Check migration file is exists
		if !migrate.IsMigrationFileIsExists() {
			log.Fatal("Error: migrations/migrations.csv is not exists please run `migrate init`")
		}

		migrations, err := migrate.Read()

		if err != nil {
			log.Fatal(err)
		}

		if fresh, err := cmd.Flags().GetBool("fresh"); err == nil && fresh {
			n, err := migrations.Down()
			if err != nil {
				log.Fatal(err)
			}
	
			if err := migrations.Done(); err != nil {
				log.Fatal(err)
			}
			log.Printf("Migrate down %d migration successfully\n", n)
		}

		n, err := migrations.Up()
		if err != nil {
			log.Fatal(err)
		}

		if err := migrations.Done(); err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			log.Printf("Migrate up %d migration successfully\n", n)
		} else {
			log.Println("Nothing to migrate")
		}

	},
}

func init() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("Could not load the config: ", err)
	}

	if err := configs.ConnectDB(); err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	migrateCmd.AddCommand(migrateUpCmd)

	migrateUpCmd.Flags().BoolP("fresh", "f", false, "Fresh migrations")
}
