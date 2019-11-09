package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"makeproject/filehandler"
	"makeproject/options"

	"github.com/urfave/cli"
)

// Project stores information about a new project.
type Project struct {
	Name        string
	ProjectType options.Alias
}

func main() {
	// Set up properties of the overall CLI application.
	app := cli.NewApp()
	app.Name = "makeproject"
	app.Usage = "Create new projects easily."
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Dani Roxberry",
			Email: "dani@bitoriented.com",
		},
	}

	// Configure the CLI application's flags.
	var _name, _type string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "--name=BEW-1.2-Project-Starter",
			Destination: &_name,
		},
		cli.StringFlag{
			Name:        "type",
			Usage:       "--type=python",
			Destination: &_type,
		},
	}

	app.Action = func(c *cli.Context) error {
		project := Project{Name: _name, ProjectType: strings.ToLower(_type)}
		templatesPath := "templates/" + project.ProjectType
		templateExt := ".tmpl"

		// Create a directory for the project if it doesn't exist.
		outputPath := "./output/" + project.Name
		filehandler.CreateDirIfNotExist(outputPath)

		// Parse the templates for the project.
		paths, _ := filehandler.GetAllFilePathsInDirectory(templatesPath)
		for _, path := range paths {
			// Execute each template.
			templateString := filehandler.ProcessFile(path, project)
			fileName := strings.Replace(path, templatesPath, "", 1)
			finalFilePath, _ := filepath.Abs(outputPath + fileName[0:len(fileName)-len(templateExt)])
			filehandler.WriteToFile(finalFilePath, templateString)
		}
		return nil
	}

	// Run the CLI application.
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
