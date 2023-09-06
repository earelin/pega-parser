package main

import (
	"bytes"
	"github.com/earelin/pega/tools/inebase/pkg/config"
	"github.com/earelin/pega/tools/inebase/pkg/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	var err error
	byteBuf := new(bytes.Buffer)

	_, err = parseArgs(byteBuf, make([]string, 2))
	assert.Error(t, err, "Should return error for too few args")
	byteBuf.Reset()

	_, err = parseArgs(byteBuf, make([]string, 4))
	assert.Error(t, err, "Should return error for too many args")
	byteBuf.Reset()

	_, err = parseArgs(byteBuf, []string{"inebase", ""})
	assert.Error(t, err, "Should return error for empty command")
	byteBuf.Reset()

	_, err = parseArgs(byteBuf, []string{"inebase", "municipios", ""})
	assert.Error(t, err, "Should return error for empty filepath")
	byteBuf.Reset()
}

func Test_parseArgs_loadConfig(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want config.Config
	}{
		{
			name: "Only file and dataset",
			args: []string{"inebase", "dataset", "file"},
			want: config.Config{
				FilePath: "file",
				DataSet:  "dataset",
				RepositoryConfig: repository.Config{
					Host:     "",
					Database: "pega",
					User:     "root",
					Password: "",
				},
			},
		},
		{
			name: "host",
			args: []string{"inebase", "-host", "test", "dataset", "file"},
			want: config.Config{
				FilePath: "file",
				DataSet:  "dataset",
				RepositoryConfig: repository.Config{
					Host:     "test",
					Database: "pega",
					User:     "root",
					Password: "",
				},
			},
		},
		{
			name: "database name",
			args: []string{"inebase", "-database", "test", "dataset", "file"},
			want: config.Config{
				FilePath: "file",
				DataSet:  "dataset",
				RepositoryConfig: repository.Config{
					Host:     "",
					Database: "test",
					User:     "root",
					Password: "",
				},
			},
		},
		{
			name: "user",
			args: []string{"inebase", "-user", "test", "dataset", "file"},
			want: config.Config{
				FilePath: "file",
				DataSet:  "dataset",
				RepositoryConfig: repository.Config{
					Host:     "",
					Database: "pega",
					User:     "test",
					Password: "",
				},
			},
		},
		{
			name: "password",
			args: []string{"inebase", "-password", "test", "dataset", "file"},
			want: config.Config{
				FilePath: "file",
				DataSet:  "dataset",
				RepositoryConfig: repository.Config{
					Host:     "",
					Database: "pega",
					User:     "root",
					Password: "test",
				},
			},
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tt := range tests {
		var conf, _ = parseArgs(byteBuf, tt.args)
		assert.Equal(t, tt.want, conf)
	}
}
