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
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/config.yml
var configFile []byte

func TestLoad_Success(t *testing.T) {
	// Load valid yml file returns no error
	_, err := LoadConfig(configFile)
	assert.NoError(t, err)
}

func TestLoad_Error(t *testing.T) {
	// Load invalid file returns error
	_, err := LoadConfig([]byte("bad_file"))
	assert.Error(t, err)
}

func getTestConfig() *Config {
	cfg, err := LoadConfig(configFile)
	if err == nil {
		return cfg
	}
	return nil
}

func TestFileConfig_Success(t *testing.T) {
	cfg := getTestConfig()

	// Can obtain the config for a given extension
	fileConfig, err := cfg.FileConfig(".jpeg")
	assert.NoError(t, err)
	assert.Equal(t, ".jpeg", fileConfig.Extension)
	assert.Equal(t, 2, len(fileConfig.DateFields))
	assert.Equal(t, "RandomKey", fileConfig.DateFields[0].Name)
	assert.Equal(t, "CreateDate", fileConfig.DateFields[1].Name)
}

func TestFileConfig_Error(t *testing.T) {
	cfg := getTestConfig()

	_, err := cfg.FileConfig(".docx")
	assert.Error(t, err)
}

func TestFileIsSupported(t *testing.T) {
	cfg := getTestConfig()

	assert.True(t, cfg.FileIsSupported("file.mov"))
	assert.True(t, cfg.FileIsSupported("file.jpeg"))
	assert.False(t, cfg.FileIsSupported("file.docx"))
}
