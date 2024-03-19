package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	_dataFile    = "data.yml"
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

func main() {
	resume_data, err := get_resume_data()
	if err != nil {
		log.Fatal("error reading data file: ", err)
	}

	fmt.Printf("%+v\n", resume_data)

	filepaths, err := get_tmpl_filepaths(_templateDir)
	if err != nil {
		log.Fatal("error reading tmpl dir: ", err)
	}

	fmt.Println(filepaths)
}
