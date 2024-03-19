package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	_dataFileName    = "data.yml"
	_htmlFileName    = "chand_anand.html"
	_templateDirName = "tmpl"
)

func getResumeData(dataFilepath string) (Resume, error) {
	dataFile, err := os.ReadFile(dataFilepath)
	if err != nil {
		return Resume{}, err
	}

	var resumeData Resume
	if err := yaml.Unmarshal(dataFile, &resumeData); err != nil {
		return Resume{}, err
	}

	return resumeData, nil
}

func getTmplFilepaths(dirName string) ([]string, error) {
	var filepaths []string

	if err := filepath.Walk(dirName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				filepaths = append(filepaths, path)
			}
			return nil
		}); err != nil {
		return nil, err
	}

	return filepaths, nil
}

func generateHTML(paths []string, resumeData Resume) (string, error) {
	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resumeData); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func writeHTML(path string, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

func main() {
	resumeData, err := getResumeData(_dataFileName)
	if err != nil {
		log.Fatal("error reading data file: ", err)
	}

	filepaths, err := getTmplFilepaths(_templateDirName)
	if err != nil {
		log.Fatal("error reading tmpl dir: ", err)
	}

    html, err := generateHTML(filepaths, resumeData)
    if err != nil {
		log.Fatal("error generating html: ", err)
    }

    if err := writeHTML(_htmlFileName, html); err != nil {
		log.Fatal("error writing html: ", err)
    }
}
