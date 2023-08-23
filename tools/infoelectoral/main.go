package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	"github.com/earelin/pega/tools/infoelectoral/pkg/writers"
	"io"
	"log"
	"os"
	"strings"
)

type config struct {
	showHelp bool
	filePath string
}

func main() {
	var conf config
	var err error

	conf, err = parseArgs(os.Stdout, os.Args)
	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	}
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

	err = e.ExportToFiles(writers.CreateJsonFile)
	if err != nil {
		log.Panic("Error exporting files to JSON", err)
	}
}

func parseArgs(w io.Writer, args []string) (config, error) {
	var c config

	fs := flag.NewFlagSet("infoelectoral", flag.ContinueOnError)
	fs.SetOutput(w)
	err := fs.Parse(args[1:])
	if err != nil {
		return c, err
	}

	if fs.NArg() != 1 {
		return c, errors.New("one positional argument required")
	}

	var filePath = strings.TrimSpace(fs.Arg(0))
	if filePath == "" {
		return c, errors.New("invalid file name")
	}
	c.filePath = filePath

	return c, nil
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
