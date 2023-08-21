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

func (z *ZipFile) FindFileWithPrefixAndExtension(prefix string, extension string) (fs.File, error) {
	for _, f := range z.file.File {
		if strings.HasPrefix(f.Name, prefix) && strings.HasSuffix(f.Name, extension) {
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
