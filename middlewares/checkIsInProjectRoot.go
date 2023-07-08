package middlewares

import (
	"log"
	"os"
)

func CheckIsInProjectRoot() {
	var CURRENT_DIR = os.Getenv("PWD")

	// if go.mod is not found, then we are not in the project root
	if _, err := os.Stat(CURRENT_DIR + "/go.mod"); err != nil {
		log.Fatal("You must run this command in the go project root (now in " + CURRENT_DIR + ")")
	}
}