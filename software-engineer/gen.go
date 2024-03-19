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
	_dataFile    = "data.yml"
	_htmlFile    = "chand_anand.html"
	_templateDir = "tmpl"
)

func get_resume_data() (Resume, error) {
	data_file, err := os.ReadFile(_dataFile)
	if err != nil {
		return Resume{}, err
	}

	var resume_data Resume
	if err := yaml.Unmarshal(data_file, &resume_data); err != nil {
		return Resume{}, err
	}

	return resume_data, nil
}

func get_tmpl_filepaths(dirName string) ([]string, error) {
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

func generate_html(paths []string, resume_data Resume) (string, error) {
	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, resume_data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func write_html(path string, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}

func main() {
	resume_data, err := get_resume_data()
	if err != nil {
		log.Fatal("error reading data file: ", err)
	}

	filepaths, err := get_tmpl_filepaths(_templateDir)
	if err != nil {
		log.Fatal("error reading tmpl dir: ", err)
	}

    html, err := generate_html(filepaths, resume_data)
    if err != nil {
		log.Fatal("error generating html: ", err)
    }

    if err := write_html(_htmlFile, html); err != nil {
		log.Fatal("error writing html: ", err)
    }
}
