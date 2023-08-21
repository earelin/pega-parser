package election

import (
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
)

type Election struct {
	zipFile *archive_reader.ZipFile
}

func (e Election) String() string {
	return "Election file for: "
}

func NewElection(zipFile *archive_reader.ZipFile) Election {
	return Election{
		zipFile: zipFile,
	}
}
