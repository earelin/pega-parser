package archive_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var zipFileContent = []string{"FICHEROS.rtf", "FICHEROS.doc", "10021911.DAT", "09021911.DAT",
	"08021911.DAT", "07021911.DAT", "06021911.DAT", "05021911.DAT", "04021911.DAT",
	"03021911.DAT", "02021911.DAT", "columns_test.txt"}

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

func TestZipFile_FileList(t *testing.T) {
	var zipFile *ZipFile

	zipFile, _ = NewZipFile("../../testdata/02201911_MESA.zip")
	fileList := zipFile.FileList()

	assert.ElementsMatch(t, fileList, zipFileContent, "Should return list of files in ZIP archive")
}
