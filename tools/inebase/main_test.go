/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
					Filename: "database.sqlite",
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
