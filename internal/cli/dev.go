// internal/cli/dev.go
package cli

import "github.com/urfave/cli/v2"


func DevCommand() *cli.Command {
	return &cli.Command{
		Name:  "dev",
		Usage: "Run the development server",
		Action: func(c *cli.Context) error {
			// Implement development server logic here
			return nil
		},
	}
}