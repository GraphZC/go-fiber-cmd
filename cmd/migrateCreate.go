/*
Copyright © 2023 TANAROEG O-CHAROEN <graph234@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/GraphZC/go-fiber-cmd/modules/migrate"
	"github.com/spf13/cobra"
)

// migrateCreateCmd represents the migrateCreate command
var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new migration",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableName, _ := cmd.Flags().GetString("table")

		// Check if the --table flag is provided
		if tableName == "" {
			log.Fatal("Error: --table flag is required")
		}

		migrationName := args[0]

		// Check migration file is exists
		if !migrate.IsMigrationFileIsExists() {
			log.Fatal("Error: migrations/migrations.csv is not exists please run `migrate init`")
		}

		migrations, err := migrate.Read()

		if err != nil {
			log.Fatal(err)
		}

		lastId := migrations.GetLastId()
		
		migrations.AddAndCreateSQL(migrate.Migration{
			Id: lastId + 1,
			Name: migrationName,
			IsMigrate: false,
		}, tableName)

		if err := migrations.Done(); err != nil {
			log.Fatal(err)
		}
		
		log.Println("Create migration successfully")
	},
}

func init() {
	migrateCmd.AddCommand(migrateCreateCmd)
	// Flag table name
	migrateCreateCmd.Flags().StringP("table", "t", "", "Table name")
	migrateCmd.MarkFlagRequired("table")
}
