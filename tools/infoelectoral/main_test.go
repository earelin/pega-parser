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
	"github.com/earelin/pega/tools/infoelectoral/pkg/config"
	"github.com/earelin/pega/tools/infoelectoral/pkg/repository"
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

	tests := []struct {
		name string
		args []string
		want config.Config
	}{
		{
			name: "Only file",
			args: []string{"infoelectoral", "file"},
			want: config.Config{
				FilePath: "file",
				RepositoryConfig: repository.Config{
					Filename: "database.sqlite",
				},
			},
		},
	}

	for _, tt := range tests {
		var conf, _ = parseArgs(byteBuf, tt.args)
		assert.Equal(t, tt.want, conf)
	}
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
