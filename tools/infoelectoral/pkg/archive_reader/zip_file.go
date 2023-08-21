package archive_reader

import (
	"archive/zip"
)

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
