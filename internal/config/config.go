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
