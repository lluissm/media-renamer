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
	err := Load(configFile)
	assert.NoError(t, err)
}

func TestLoad_Error(t *testing.T) {
	// Load invalid file returns error
	err := Load([]byte("bad_file"))
	assert.Error(t, err)
}

func TestSupportedExtensions(t *testing.T) {
	Load(configFile)

	// Supported extensions are properly parsed
	extensions := SupportedExtensions()
	assert.Equal(t, 2, len(extensions))
	assert.Equal(t, ".mov", extensions[0])
	assert.Equal(t, ".jpeg", extensions[1])
}

func TestFileConfig_Success(t *testing.T) {
	Load(configFile)

	// Can obtain the config for a given extension
	fileConfig, err := FileConfig(".jpeg")
	assert.NoError(t, err)
	assert.Equal(t, ".jpeg", fileConfig.Extension)
	assert.Equal(t, 2, len(fileConfig.DateFields))
	assert.Equal(t, "RandomKey", fileConfig.DateFields[0].Name)
	assert.Equal(t, "CreateDate", fileConfig.DateFields[1].Name)
}

func TestFileConfig_Error(t *testing.T) {
	Load(configFile)

	_, err := FileConfig(".docx")
	assert.Error(t, err)
}
