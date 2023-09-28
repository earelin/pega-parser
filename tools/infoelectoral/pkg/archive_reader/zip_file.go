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

package archive_reader

import (
	"archive/zip"
	"errors"
	"io/fs"
	"strings"
)

var FileNotFound = errors.New("file not found")

type ZipFile struct {
	file *zip.ReadCloser
}

func (z *ZipFile) FileList() []string {
	var filenames []string
	for _, f := range z.file.File {
		filenames = append(filenames, f.Name)
	}
	return filenames
}

func (z *ZipFile) FindFileWithPrefix(prefix string) (fs.File, error) {
	for _, f := range z.file.File {
		if strings.HasPrefix(f.Name, prefix) && strings.HasSuffix(f.Name, ".DAT") {
			return z.Open(f.Name)
		}
	}
	return nil, FileNotFound
}

func (z *ZipFile) Open(filename string) (fs.File, error) {
	return z.file.Open(filename)
}

func (z *ZipFile) Close() error {
	return z.file.Close()
}

func NewZipFile(path string) (*ZipFile, error) {
	var err error
	zipFile := &ZipFile{}

	zipFile.file, err = zip.OpenReader(path)
	if err != nil {
		return nil, err
	}

	return zipFile, nil
}
