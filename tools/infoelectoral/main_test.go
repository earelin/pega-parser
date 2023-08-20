package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	var err error

	_, err = parseArgs(make([]string, 1))
	assert.Error(t, err, "Should return error for too few args")

	_, err = parseArgs(make([]string, 3))
	assert.Error(t, err, "Should return error for too many args")

	_, err = parseArgs([]string{"infoelectoral", ""})
	assert.Error(t, err, "Should return error for empty filepath")

	var conf config
	conf, _ = parseArgs([]string{"infoelectoral", "-h"})
	assert.True(t, conf.showHelp, "Should set showHelp")

	conf, _ = parseArgs([]string{"infoelectoral", "--help"})
	assert.True(t, conf.showHelp, "Should set showHelp")

	conf, _ = parseArgs([]string{"infoelectoral", "file"})
	assert.False(t, conf.showHelp, "Should not set showHelp")
	assert.Equal(t, conf.filePath, "file")
}

func Test_validateConfiguration(t *testing.T) {
	var err error
	var conf config

	conf = config{
		showHelp: true,
		filePath: "not existing file",
	}
	err = validateConfiguration(conf)
	assert.Nil(t, err, "Should not return error on showHelp")

	conf = config{
		showHelp: false,
		filePath: "not existing file",
	}
	err = validateConfiguration(conf)
	assert.Error(t, err, "Should return error on not existing file")

	conf = config{
		showHelp: false,
		filePath: "testdata/02201911_MESA.zip",
	}
	err = validateConfiguration(conf)
	assert.Nil(t, err, "Should not return error on existing file")
}
