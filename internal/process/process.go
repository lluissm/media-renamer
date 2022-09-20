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
	"path/filepath"
	"strings"

	"fmt"
	"time"

	"github.com/lluissm/media-renamer/internal/config"
)

// NewFileName returns the date formated to be used as a file name
func NewFileName(dateFormat, date string) (string, error) {
	parseTime, err := time.Parse(dateFormat, date)

	if err != nil {
		fmt.Println("Error parsing date")
		return "", err
	}

	return fmt.Sprintf("%04d_%02d_%02d_%02d_%02d_%02d", parseTime.Year(), parseTime.Month(), parseTime.Day(), parseTime.Hour(), parseTime.Minute(), parseTime.Second()), nil
}

// ShouldIgnoreFile returns true if file should be ignored
func ShouldIgnoreFile(path string, d fs.DirEntry) bool {
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
