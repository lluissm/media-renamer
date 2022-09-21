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
	"testing"

	"github.com/stretchr/testify/assert"
)

const filePathArg = "a-file-path"
const versionFlagName = "-version"

func TestErrorWhenMissingArgs(t *testing.T) {
	args := []string{cmdName}
	_, err := Parse(args)
	assert.NotNil(t, err)

	args = []string{cmdName, filePathArg}
	_, err = Parse(args)
	assert.Nil(t, err)
}

func TestVersion(t *testing.T) {
	args := []string{cmdName, versionFlagName, filePathArg}
	options, err := Parse(args)
	assert.Nil(t, err)
	assert.True(t, options.ShowVersion)

	args = []string{cmdName, filePathArg}
	options, err = Parse(args)
	assert.Nil(t, err)
	assert.False(t, options.ShowVersion)
}

func TestPath(t *testing.T) {
	args := []string{cmdName, filePathArg}
	options, err := Parse(args)
	assert.Nil(t, err)
	assert.True(t, options.Path == filePathArg)
}
