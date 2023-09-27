package archive_reader

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
)

func TestNewZipFile(t *testing.T) {
	var zipFile *ZipFile
	var err error

	_, err = NewZipFile("no existing file")
	assert.Error(t, err, "Should return error if file does not exists")

	zipFile, err = NewZipFile("../../testdata/02201911_MESA.zip")
	assert.Nil(t, err)
	if zipFile.Close() != nil {
		t.Error("Could not close opened zip file")
	}
}

func TestZipFileFindFileWithPrefixAndExtension(t *testing.T) {
	zipFile, err := NewZipFile("../../testdata/02201911_MESA.zip")
	if err != nil {
		t.Error(err)
	}

	var file fs.File
	var fileInfo fs.FileInfo
	file, err = zipFile.FindFileWithPrefix("01")
	assert.Nil(t, err)
	fileInfo, _ = file.Stat()
	assert.Equal(t, "01021911.DAT", fileInfo.Name())

	_, err = zipFile.FindFileWithPrefix("KK")

	assert.Equal(t, FileNotFound, err)
}

func TestZipFileFileList(t *testing.T) {
	zipFile, err := NewZipFile("../../testdata/02201911_MESA.zip")
	if err != nil {
		t.Error(err)
	}

	var fileList = zipFile.FileList()
	assert.Equal(t, []string{"01021911.DAT", "02021911.DAT", "03021911.DAT", "04021911.DAT", "05021911.DAT",
		"06021911.DAT", "07021911.DAT", "08021911.DAT", "09021911.DAT", "10021911.DAT", "FICHEROS.doc",
		"FICHEROS.rtf"}, fileList)
}
