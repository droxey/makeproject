package filehandler

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// CreateDirIfNotExist ...
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		check(err)
	}
}

// GetAllFilePathsInDirectory recursively returns all file paths in a soecified directory, including sub-directories.
func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	// Get all the .tmpl files in the directory.
	var paths []string
	extension := ".tmpl"

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == extension {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// process applies the data structure 'vars' onto an already
// parsed template 't', and returns the resulting string.
func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	check(err)

	return tmplBytes.String()
}

// ProcessFile ...
func ProcessFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)
	check(err)

	return process(tmpl, vars)
}

// WriteToFile ...
func WriteToFile(filename string, data string) {
	println("Writing file: " + filename)
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	file.WriteString(data)
}
