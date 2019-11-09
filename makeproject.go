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
	Name         string
	ProjectType  options.Alias
	OutputDir    string
	TemplatesDir string
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
	var _name, _tmpl string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "--name=BEW-1.2-Project-Starter",
			Destination: &_name,
		},
		cli.StringFlag{
			Name:        "template",
			Usage:       "--template=django",
			Destination: &_tmpl,
		},
	}

	app.Action = func(c *cli.Context) error {
		// Store the data about the new project.
		project := Project{Name: _name, ProjectType: strings.ToLower(_tmpl)}
		project.TemplatesDir = "./templates/" + project.ProjectType

		// Create a directory for the project if it doesn't exist.
		project.OutputDir = "./output/" + project.Name
		filehandler.CreateDirIfNotExist(project.OutputDir)

		// Parse the templates for the project.
		templateFiles, _ := filehandler.GetAllFilePathsInDirectory(project.TemplatesDir)
		commonFiles, _ := filehandler.GetAllFilePathsInDirectory("./templates/common")
		templateFiles = append(templateFiles, commonFiles...)
		templateExt := ".tmpl"
		for _, path := range templateFiles {
			// Execute each template.
			templateString := filehandler.ProcessFile(path, project)
			fileName := filepath.Base(path)
			finalFilePath, _ := filepath.Abs(project.OutputDir + "/" + fileName[0:len(fileName)-len(templateExt)])
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
