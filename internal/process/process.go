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

package process

import (
	_ "embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fmt"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/lluissm/media-renamer/internal/config"
)

// Folder process all files in a given path
func Folder(et *exiftool.Exiftool, path string) error {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if shouldIgnoreFile(path, d) {
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

// processFile tries to rename a file according to its date metadata
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

// tryGetDate tries to obtain the date from metadata in a format to be used for the file name
func tryGetDate(path string, fileType *config.FileType, key, value interface{}) (string, error) {
	for _, dateField := range fileType.DateFields {
		if key == dateField.Name {
			dateStr := fmt.Sprintf("%v", value)
			name, err := newFileName(dateField.DateFormat, dateStr)
			if err != nil {
				return "", err
			}
			return name, nil
		}
	}
	return "", fmt.Errorf("creation date not found in %v", key)
}

// tryRename tries to rename a file according to its metadata
func tryRename(path string, fileInfo exiftool.FileMetadata) error {
	ext := filepath.Ext(path)
	fileConfig, err := config.FileConfig(ext)
	if err != nil {
		return err
	}

	for k, v := range fileInfo.Fields {
		dateStr, err := tryGetDate(path, fileConfig, k, v)
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

// newFileName returns the date formated to be used as a file name
func newFileName(dateFormat, date string) (string, error) {
	parseTime, err := time.Parse(dateFormat, date)

	if err != nil {
		fmt.Println("Error parsing date")
		return "", err
	}

	return fmt.Sprintf("%04d_%02d_%02d_%02d_%02d_%02d", parseTime.Year(), parseTime.Month(), parseTime.Day(), parseTime.Hour(), parseTime.Minute(), parseTime.Second()), nil
}

// shouldIgnoreFile returns true if file should be ignored
func shouldIgnoreFile(path string, d fs.DirEntry) bool {
	// Skip if directory
	if d.IsDir() {
		return true
	}

	// Skip if file extension is not supported
	if !config.FileIsSupported(path) {
		return true
	}

	// Skip if hidden file
	_, filename := filepath.Split(path)
	return strings.HasPrefix(filename, ".")
}
