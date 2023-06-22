/*
Copyright © 2023 TANAROEG O-CHAROEN <graph234@gmail.com>
*/

package main

import (
	"log"

	"github.com/GraphZC/go-fiber-cmd/cmd"
	"github.com/GraphZC/go-fiber-cmd/configs"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Could not load the config: " + err.Error())
	}
	cmd.Execute()
}
