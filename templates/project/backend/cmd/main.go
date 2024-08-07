// This is the main entry point to your GoPostulate project.

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
