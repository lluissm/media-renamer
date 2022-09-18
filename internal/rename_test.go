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

package rename

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileName(t *testing.T) {

	dateJPEG := "2019:08:05 14:12:13"
	expectedJPEG := "2019_08_05_14_12_13"
	dateMOV := "2015:07:15 13:56:17+02:00"
	expectedMOV := "2015_07_15_13_56_17"

	fnameJPEG, err := newFileName(formatJPEG, dateJPEG)
	assert.NoError(t, err)
	assert.Equal(t, fnameJPEG, expectedJPEG)

	fnameMOV, err := newFileName(formatMOV, dateMOV)
	assert.NoError(t, err)
	assert.Equal(t, fnameMOV, expectedMOV)
}
