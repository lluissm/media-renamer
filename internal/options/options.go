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

package options

import (
	"errors"
	"flag"
	"fmt"
)

const cmdName = "media-renamer"

// Options are the process.Options parsed from command line flags/args
type Options struct {
	ShowVersion      bool
	Verbose          bool
	Path             string
	CustomConfigPath string
}

// Parse returns the parsed Options from command line flags/args
func Parse(osArgs []string) (*Options, error) {

	flagSet := flag.NewFlagSet("mrn", flag.ExitOnError)
	flagSet.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\033[1;4mSYNOPSIS\033[0m\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "%s ~/Desktop/my-trip\n\n", cmdName)
		fmt.Fprintf(flag.CommandLine.Output(), "\033[1;4mOPTIONS\033[0m\n\n")
		flagSet.PrintDefaults()
	}

	showVersionFlag := flagSet.Bool("version", false, "Display version number")
	verboseFlag := flagSet.Bool("v", false, "Diplay detailed information of the processing during execution")
	configFileFlag := flagSet.String("c", "", "Path to custom configuration file (optional)")

	if err := flagSet.Parse(osArgs[1:]); err != nil {
		return nil, err
	}

	args := flagSet.Args()

	if *showVersionFlag {
		return &Options{
			true,
			true,
			"",
			"",
		}, nil
	}

	if len(args) < 1 {
		return nil, errors.New("Missing arguments, please see documentation")
	}

	path := args[0]

	return &Options{
		*showVersionFlag,
		*verboseFlag,
		path,
		*configFileFlag,
	}, nil
}
