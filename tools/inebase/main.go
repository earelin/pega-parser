package inebase

import (
	"errors"
	"flag"
	"github.com/earelin/pega/tools/infoelectoral/pkg/config"
	"io"
	"os"
	"strings"
)

func main() {
	start(os.Stdout, os.Args)
}

func start(w io.Writer, args []string) {

}

func parseArgs(w io.Writer, args []string) (config.Config, error) {
	var c config.Config

	fs := flag.NewFlagSet("inebase", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.StringVar(&c.RepositoryConfig.Host, "host", "", "Enderezo da base de datos")
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
