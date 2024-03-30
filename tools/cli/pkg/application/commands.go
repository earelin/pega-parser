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

package application

import "github.com/earelin/pega/tools/cli/pkg/config"

var commands = map[string]map[string]func(){
	config.ProcesoElectoral: {
		config.Crear:      func() {},
		config.Listado:    func() {},
		config.Borrar:     func() {},
		config.Actualizar: func() {},
	},
}

func RunCommand(conf config.Config) {
	if commands[conf.Entidade] != nil {
		if commands[conf.Entidade][conf.Comando] != nil {
			commands[conf.Entidade][conf.Comando]()
		}
	}
}
