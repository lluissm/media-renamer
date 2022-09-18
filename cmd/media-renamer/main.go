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

package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/lluissm/media-renamer/internal/rename"
)

var allowedExtensions = [2]string{".jpeg", ".mov"}

func main() {
	path := os.Args[1]
	if err := iterate(path); err != nil {
		log.Fatalf("Error intializing exiftool: %v\n", err)
	}
}

func fileNotSupported(path string) bool {
	fileExtension := filepath.Ext(path)
	for _, ext := range allowedExtensions {
		if fileExtension == ext {
			return false
		}
	}
	return true
}

func iterate(path string) error {
	et, err := exiftool.NewExiftool()
	if err != nil {
		return err
	}
	defer et.Close()

	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if d.IsDir() {
			return nil
		}

		if fileNotSupported(path) {
			return nil
		}

		_, filename := filepath.Split(path)
		if strings.HasPrefix(filename, ".") {
			return nil
		}

		processFile(et, path)
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func getDateFromExifKey(path string, key, value interface{}) (string, error) {
	if filepath.Ext(path) == ".mov" {
		if key == "CreationDate" {
			dateStr := fmt.Sprintf("%v", value)

			name, err := rename.NewFileName(rename.DateFormatMOV, dateStr)
			if err != nil {
				return "", err
			}

			return name, nil
		}

	} else if filepath.Ext(path) == ".jpeg" {
		if key == "CreateDate" {
			dateStr := fmt.Sprintf("%v", value)

			name, err := rename.NewFileName(rename.DateFormatJPEG, dateStr)
			if err != nil {
				return "", err
			}

			return name, nil
		}
	}
	return "", fmt.Errorf("creation date not found in %v", key)
}

func tryRename(path string, fileInfo exiftool.FileMetadata) error {
	for k, v := range fileInfo.Fields {
		dateStr, err := getDateFromExifKey(path, k, v)
		if err == nil {
			dir, _ := filepath.Split(path)
			ext := filepath.Ext(path)
			newPath := fmt.Sprintf("%s%s%s", dir, dateStr, ext)

			if err = os.Rename(path, newPath); err != nil {
				return fmt.Errorf("Could not rename file %s to %s. %w", path, newPath, err)
			} else {
				log.Printf("Renamed %s to %s", path, newPath)
				return nil
			}
		}
	}
	return fmt.Errorf("could not find information in metadata for file %s", path)
}

func processFile(et *exiftool.Exiftool, path string) {
	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			log.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		if err := tryRename(path, fileInfo); err != nil {
			log.Printf("ERROR: %s", err.Error())
		}
	}
}
