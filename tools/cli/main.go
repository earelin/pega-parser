/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/earelin/pega/tools/cli/pkg/application"
	"github.com/earelin/pega/tools/cli/pkg/config"
	"io"
	"os"
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
		fmt.Fprintln(w, "Erro arrancando o programa: ", err)
		showUsage()
		os.Exit(1)
	}

	application.RunCommand(conf)

	fmt.Fprintln(w, "Configuracion: ", conf)
}

func parseArgs(w io.Writer, args []string) (config.Config, error) {
	var c config.Config

	fs := flag.NewFlagSet("pega-cli", flag.ContinueOnError)
	fs.SetOutput(w)

	err := fs.Parse(args[1:])
	if err != nil {
		return c, err
	}

	if fs.NArg() < 2 {
		return c, errors.New("falta a o comando a executar")
	}

	c.Entidade = fs.Arg(0)
	c.Comando = fs.Arg(1)

	return c, nil
}

func showUsage() {
	fmt.Println(`	Uso:
		pega [opcions] [entidade] [comando]

	Entidades:
		proceso-electoral
		    Comandos:
				- crear
				- lista
				- borrar
				- actualizar
	Opcións:
		-h/--help Amosar esta mensaxe de axuda`)
}
