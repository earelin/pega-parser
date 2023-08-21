package main

import (
	"errors"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	"log"
	"os"
)

type config struct {
	showHelp bool
	filePath string
}

func main() {
	var conf config
	var err error

	conf, err = parseArgs(os.Args)
	if err != nil {
		fmt.Println("Error executing command: ", err)
		showUsage()
		os.Exit(1)
	}

	if conf.showHelp {
		showHelp()
		os.Exit(0)
	}

	err = validateConfiguration(conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var zipFile *archive_reader.ZipFile
	zipFile, err = archive_reader.NewZipFile(conf.filePath)
	if err != nil {
		log.Panic("Cannot open archive: ", err)
	}

	var e = election.NewElection(zipFile)
	fmt.Print(e.String())
}

func parseArgs(arguments []string) (config, error) {
	if len(arguments) < 2 {
		return config{}, errors.New("missing argument")
	}
	if len(arguments) > 2 {
		return config{}, errors.New("too many arguments")
	}
	argument := arguments[1]

	if argument == "" {
		return config{}, errors.New("empty filepath")
	}

	if argument == "-h" || argument == "--help" {
		return config{
			showHelp: true,
		}, nil
	}

	return config{
		filePath: argument,
	}, nil
}

func showHelp() {
	fmt.Println("Extracts polling information from infoelectoral.interior.gob.es ZIP files")
	showUsage()
}

func showUsage() {
	fmt.Println(`	Usage:
		infoelectoral [options] filepath

		filepath: Path to the infoelectoral ZIP file

	Options:
		-h/--help Show this help info`)
}

func validateConfiguration(conf config) error {
	var err error

	if !conf.showHelp {
		_, err = os.Stat(conf.filePath)
	}

	return err
}
