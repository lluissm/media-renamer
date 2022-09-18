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

	"github.com/barasher/go-exiftool"
	"github.com/lluissm/media-renamer/internal/rename"
)

func main() {
	argsWithoutProg := os.Args[1:]
	path := argsWithoutProg[0]
	fmt.Printf("Path: %s\n", path)
	iterate(path)
}

func processFile(et *exiftool.Exiftool, path string) {
	fileInfos := et.ExtractMetadata(path)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			if filepath.Ext(path) == ".mov" {

				if k == "CreationDate" {
					dir, file := filepath.Split(path)
					ext := filepath.Ext(path)

					fmt.Println("\n\nPath:", path)
					fmt.Println("Dir:", dir)
					fmt.Println("File:", file)
					fmt.Println("Extension:", ext)

					fmt.Printf("%v: %v\n", k, v)
					dateStr := fmt.Sprintf("%v", v)
					dateFmt := "2006:01:02 15:04:05-07:00"

					name, err := rename.NewFileName(dateFmt, dateStr)
					if err != nil {
						continue
					}

					newPath := fmt.Sprintf("%s%s%s", dir, name, ext)
					fmt.Printf("New path: %s\n\n", newPath)

					os.Rename(path, newPath)

				}

			} else if filepath.Ext(path) == ".jpeg" {
				if k == "CreateDate" {
					dir, file := filepath.Split(path)
					ext := filepath.Ext(path)

					fmt.Println("\n\nPath:", path)
					fmt.Println("Dir:", dir)
					fmt.Println("File:", file)
					fmt.Println("Extension:", ext)

					fmt.Printf("%v: %v\n", k, v)

					dateStr := fmt.Sprintf("%v", v)
					dateFmt := "2006:01:02 15:04:05"

					name, err := rename.NewFileName(dateFmt, dateStr)
					if err != nil {
						continue
					}

					newPath := fmt.Sprintf("%s%s%s", dir, name, ext)
					fmt.Printf("New path: %s\n\n", newPath)

					os.Rename(path, newPath)

				}
			}
		}
	}
}

func iterate(path string) {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error intializing exiftool: %v\n", err)
		return
	}
	defer et.Close()

	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}

		if d.IsDir() {
			return nil
		}

		processFile(et, path)
		return nil
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}
