// Package main is the CLI entry point for the gopostulate framework
package main

import (
	"log"
	"os"

	"github.com/evanmschultz/gopostulate_dev/internal/cli"
)

func main() {
	app := cli.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}