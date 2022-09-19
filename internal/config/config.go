package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed config.yml
var configFile []byte
var fileTypes []FileType
var supportedExtensions []string

type (
	DateField struct {
		Name       string `yaml:"name"`
		DateFormat string `yaml:"dateFormat"`
	}

	FileType struct {
		Extension  string      `yaml:"extension"`
		DateFields []DateField `yaml:"dateFields"`
	}
)

func Unmarshal() error {
	err := yaml.Unmarshal(configFile, &fileTypes)
	if err != nil {
		fmt.Println("error")

		return fmt.Errorf("error unmarshaling the file: %w", err)
	}

	supportedExtensions = []string{}
	for _, f := range fileTypes {
		supportedExtensions = append(supportedExtensions, f.Extension)
	}

	return nil
}

func SupportedExtensions() []string {
	return supportedExtensions
}

func FileConfig(ext string) *FileType {
	for _, f := range fileTypes {
		if f.Extension == ext {
			return &f
		}
	}
	return nil
}
