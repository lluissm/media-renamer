/* MIT License

Copyright (c) 2022 Lluis Sanchez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package config

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		fileTypes           []FileType
		supportedExtensions []string
	}

	DateField struct {
		Name       string `yaml:"name"`
		DateFormat string `yaml:"dateFormat"`
	}

	FileType struct {
		Extension  string      `yaml:"extension"`
		DateFields []DateField `yaml:"dateFields"`
	}
)

// Load loads the configuration from the provided yaml file
func LoadConfig(bytes []byte) (*Config, error) {
	var fileTypes []FileType
	err := yaml.Unmarshal(bytes, &fileTypes)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling the file: %w", err)
	}

	supportedExtensions := []string{}
	for _, f := range fileTypes {
		supportedExtensions = append(supportedExtensions, f.Extension)
	}

	return &Config{
		fileTypes:           fileTypes,
		supportedExtensions: supportedExtensions,
	}, nil
}

// FileConfig returns the configuration for a given extension, error if not found
func (c *Config) FileConfig(ext string) (*FileType, error) {
	for _, f := range c.fileTypes {
		if f.Extension == ext {
			return &f, nil
		}
	}
	return nil, fmt.Errorf("could not find a configuration for the given extension")
}

// FileIsSupported returns true if the file extension is present in the config
func (c *Config) FileIsSupported(path string) bool {
	fileExtension := filepath.Ext(path)
	for _, ext := range c.supportedExtensions {
		if fileExtension == ext {
			return true
		}
	}
	return false
}
