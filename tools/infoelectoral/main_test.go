package main

import (
	"bytes"
	"github.com/earelin/pega/tools/infoelectoral/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	var err error
	byteBuf := new(bytes.Buffer)

	_, err = parseArgs(byteBuf, make([]string, 1))
	assert.Error(t, err, "Should return error for too few args")
	byteBuf.Reset()

	_, err = parseArgs(byteBuf, make([]string, 3))
	assert.Error(t, err, "Should return error for too many args")
	byteBuf.Reset()

	_, err = parseArgs(byteBuf, []string{"infoelectoral", ""})
	assert.Error(t, err, "Should return error for empty filepath")
	byteBuf.Reset()

	var conf config.Config
	conf, _ = parseArgs(byteBuf, []string{"infoelectoral", "file"})
	assert.Equal(t, "file", conf.FilePath)
}

func Test_validateConfiguration(t *testing.T) {
	var err error
	var conf config.Config

	conf = config.Config{
		FilePath: "not existing file",
	}
	err = validateConfiguration(conf)
	assert.Error(t, err, "Should return error on not existing file")

	conf = config.Config{
		FilePath: "testdata/02201911_MESA.zip",
	}
	err = validateConfiguration(conf)
	assert.Nil(t, err, "Should not return error on existing file")
}
