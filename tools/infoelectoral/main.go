package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/config"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	start(os.Stdout, os.Args)
}

func start(w io.Writer, args []string) {
	var conf config.Config
	var err error

	conf, err = parseArgs(w, args)
	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	}
	if err != nil {
		fmt.Println("Error arrancando o programa: ", err)
		showUsage()
		os.Exit(1)
	}

	err = validateConfiguration(conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var zipFile *archive_reader.ZipFile
	zipFile, err = archive_reader.NewZipFile(conf.FilePath)
	if err != nil {
		log.Panic("Non se pode abrir o ficheiro: ", err)
	}

	var e = election.NewElection(zipFile)
	fmt.Print(e.String())

}

func parseArgs(w io.Writer, args []string) (config.Config, error) {
	var c config.Config

	fs := flag.NewFlagSet("infoelectoral", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&c.RepositoryConfig.Host, "host", "localhost", "Enderezo da base de datos")
	fs.StringVar(&c.RepositoryConfig.User, "user", "root", "Usuario da base de datos")
	fs.StringVar(&c.RepositoryConfig.Password, "password", "", "Contrasinal da base de datos")
	fs.StringVar(&c.RepositoryConfig.Database, "database", "pega", "Nome da base de datos")

	err := fs.Parse(args[1:])
	if err != nil {
		return c, err
	}

	if fs.NArg() != 1 {
		return c, errors.New("nome de ficheiro non expecificado")
	}

	var filePath = strings.TrimSpace(fs.Arg(0))
	if filePath == "" {
		return c, errors.New("nome de ficherio non expecificado")
	}
	c.FilePath = filePath

	return c, nil
}

func showUsage() {
	fmt.Println(`	Uso:
		infoelectoral [opcions] ficheiro

		ficheiro: Ruta ao ficheiro ZIP

	Opcións:
		-h/--help Show this help info`)
}

func validateConfiguration(conf config.Config) error {
	_, err := os.Stat(conf.FilePath)

	return err
}
