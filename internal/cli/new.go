package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// NewApp creates and returns a new cli.App instance configured for the gopostulate framework.
// This function sets up the entire command-line interface for the framework, including all
// available commands, global flags, and application metadata.
//
// The returned cli.App is the entry point for all CLI interactions with the gopostulate framework.
// It includes commands for creating new projects, running the development server, and building
// projects for production.
//
// Usage:
//
//	app := cli.NewApp()
//	err := app.Run(os.Args)
//
// The app can be further customized before running, if necessary.
func NewApp() *cli.App {
	return &cli.App{
		Name:  "gopostulate",
		Usage: "A modern fullstack Go framework",
		Commands: []*cli.Command{
			NewCommand(),
			DevCommand(),
			BuildCommand(),
		},
		// You can add global flags here if needed
		// Flags: []cli.Flag{ ... },
	}
}

// NewCommand creates and returns a new CLI command for creating a new gopostulate project.
// The command accepts a single flag, --name (or -n), which specifies the name of the new project.
// When the command is executed, it creates a new directory with the specified project name and
// copies the project template files into the new directory.
func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Create a new gopostulate project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Name of the new project",
			},
		},
		Action: func(c *cli.Context) error {
			projectName := c.String("name")
			if projectName == "" {
				return fmt.Errorf("project name is required")
			}

			// Create project directory
			err := os.Mkdir(projectName, 0755)
			if err != nil {
				return fmt.Errorf("failed to create project directory: %w", err)
			}

			// Copy template files
			err = copyTemplateFiles(projectName)
			if err != nil {
				return fmt.Errorf("failed to copy template files: %w", err)
			}

			fmt.Printf("Created new gopostulate project: %s\n", projectName)
			return nil
		},
	}
}

func copyTemplateFiles(projectName string) error {
	templateDir := "templates/project"
	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(projectName, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(destPath, data, info.Mode())
	})
}
