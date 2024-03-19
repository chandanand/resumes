package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func get_resume_data() (Resume, error) {
	data_file, err := os.ReadFile("data.yml")
	if err != nil {
		return Resume{}, err
	}

	var resume_data Resume
	if err := yaml.Unmarshal(data_file, &resume_data); err != nil {
		return Resume{}, err
	}

	return resume_data, nil
}

func main() {
	resume_data, err := get_resume_data()
	if err != nil {
		log.Fatal("error reading data file: ", err)
	}

	fmt.Printf("%+v\n", resume_data)
}
