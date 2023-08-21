package processor

import (
	"errors"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"io/fs"
	"log"
	"strings"
)

func ProcessDataFile(filePath string) {
	var zipFile *archive_reader.ZipFile
	var err error

	zipFile, err = archive_reader.NewZipFile(filePath)
	if err != nil {
		log.Panic("Cannot open archive: ", err)
	}

	f, err := findControlFile(zipFile)
	if err != nil {
		log.Panic("Cannot find control file", err)
	}

}

func findControlFile(zipFile *archive_reader.ZipFile) (fs.File, error) {
	fileList := zipFile.FileList()
	for _, filename := range fileList {
		if strings.HasPrefix(filename, "01") &&
			strings.HasSuffix(filename, ".DAT") {
			return zipFile.Open(filename)
		}
	}
	return nil, errors.New("control file not found")
}
