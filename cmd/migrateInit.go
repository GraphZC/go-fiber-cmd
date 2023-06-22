/*
Copyright © 2023 TANAROEG O-CHAROEN <graph234@gmail.com>
*/
package cmd

import (

	"log"

	"github.com/GraphZC/go-fiber-cmd/modules/migrate"
	"github.com/spf13/cobra"
)

var migrateInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize migrations file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrate.Init(); err != nil {
			log.Fatal(err.Error())
		}
		log.Println("migrations/migrations.csv created")
	},
}

func init() {
	migrateCmd.AddCommand(migrateInitCmd)
}