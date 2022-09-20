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

	"os"
	"testing"

	"github.com/lluissm/media-renamer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:embed testdata/config.yml
var configFile []byte

// dirEntryMock mocks fs.DirEntry
type dirEntryMock struct {
	mock.Mock
}

func (d *dirEntryMock) IsDir() bool {
	args := d.Called()
	return args.Get(0).(bool)
}
func (d *dirEntryMock) Name() string               { return "" }
func (d *dirEntryMock) Type() os.FileMode          { return 0 }
func (d *dirEntryMock) Info() (os.FileInfo, error) { return nil, nil }

func getTestConfig() *config.Config {
	cfg, err := config.LoadConfig(configFile)
	if err == nil {
		return cfg
	}
	return nil
}

///////////////////////////////////
//			NewFileName
///////////////////////////////////

func TestNewFileName_JPEG(t *testing.T) {
	DateFormatJPEG := "2006:01:02 15:04:05"

	dateJPEG := "2019:08:05 14:12:13"
	expected := "2019_08_05_14_12_13"

	fnameJPEG, err := newFileName(DateFormatJPEG, dateJPEG)
	assert.NoError(t, err)
	assert.Equal(t, fnameJPEG, expected)
}

func TestNewFileName_MOV(t *testing.T) {
	DateFormatMOV := "2006:01:02 15:04:05-07:00"

	dateMOV := "2015:07:15 13:56:17+02:00"
	expected := "2015_07_15_13_56_17"

	fnameMOV, err := newFileName(DateFormatMOV, dateMOV)
	assert.NoError(t, err)
	assert.Equal(t, fnameMOV, expected)
}

func TestNewFileName_Error(t *testing.T) {
	WrongDateFormat := "2012:12:12"

	dateMOV := "2015:07:15 13:56:17+02:00"

	_, err := newFileName(WrongDateFormat, dateMOV)
	assert.Error(t, err)
}

///////////////////////////////////
//			shouldIgnoreFile
///////////////////////////////////

func TestShouldIgnoreFile(t *testing.T) {
	cfg := getTestConfig()

	dirEntry := &dirEntryMock{}
	hiddenFilePath := ".hidden.jpeg"
	validFilePath := "IMG001.jpeg"
	validFilePathWrongExtension := "document.docx"

	// No folder, not hidden, supported extension
	dirEntry.On("IsDir").Return(false).Once()
	res := shouldIgnoreFile(validFilePath, cfg, dirEntry)
	assert.False(t, res)
	dirEntry.AssertExpectations(t)

	// No folder, not hidden, non-supported extension
	dirEntry.On("IsDir").Return(false).Once()
	res = shouldIgnoreFile(validFilePathWrongExtension, cfg, dirEntry)
	assert.True(t, res)
	dirEntry.AssertExpectations(t)

	// No folder, hidden, supported extension
	dirEntry.On("IsDir").Return(false).Once()
	res = shouldIgnoreFile(hiddenFilePath, cfg, dirEntry)
	assert.True(t, res)
	dirEntry.AssertExpectations(t)

	// A folder
	dirEntry.On("IsDir").Return(true).Once()
	res = shouldIgnoreFile(validFilePath, cfg, dirEntry)
	assert.True(t, res)
	dirEntry.AssertExpectations(t)
}

///////////////////////////////////
//			tryGetDate
///////////////////////////////////

func TestTryGetDate_Success(t *testing.T) {
	path := "IMG_0001.jpeg"
	key := "CreateDate"
	value := "2019:08:05 14:12:13"
	expected := "2019_08_05_14_12_13"

	fileConfig, err := getTestConfig().FileConfig(".jpeg")
	assert.NoError(t, err)

	date, err := tryGetDate(path, fileConfig, key, value)
	assert.NoError(t, err)
	assert.Equal(t, expected, date)
}

func TestTryGetDate_ErrorDateNotExisting(t *testing.T) {
	path := "IMG_0001.jpg"
	key := "WrongKey"
	value := "2019:08:05 14:12:13"

	fileConfig, err := getTestConfig().FileConfig(".jpeg")
	assert.NoError(t, err)

	_, err = tryGetDate(path, fileConfig, key, value)
	assert.Error(t, err)
}

func TestTryGetDate_ErrorCannotParseDate(t *testing.T) {
	path := "IMG_0001.jpg"
	key := "CreateDate"
	value := "2022"

	fileConfig, err := getTestConfig().FileConfig(".jpeg")
	assert.NoError(t, err)

	_, err = tryGetDate(path, fileConfig, key, value)
	assert.Error(t, err)
}
