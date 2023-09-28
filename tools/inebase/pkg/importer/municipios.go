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

package importer

import (
	"fmt"
	"github.com/earelin/pega/tools/inebase/pkg/config"
	"github.com/earelin/pega/tools/inebase/pkg/model"
	"github.com/earelin/pega/tools/inebase/pkg/repository"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func ImportarConcellos(c config.Config, r *repository.Repository) error {
	f, err := excelize.OpenFile(c.FilePath)
	if err != nil {
		return fmt.Errorf("non se puido abrir o ficheiro %s: %w", c.FilePath, err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var sheets = f.GetSheetList()
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return fmt.Errorf("non se puido extraer as filas da folla %s: %w", sheets[0], err)
	}

	var concellos = make([]model.Concello, 0, len(rows)-2)
	for i := 2; i < len(rows); i++ {
		var ierr error
		var concelloId, provinciaId int
		concelloId, ierr = strconv.Atoi(rows[i][2])
		if ierr != nil {
			return fmt.Errorf("codigo de concello non é un número: %s", rows[i][2])
		}
		provinciaId, ierr = strconv.Atoi(rows[i][1])
		if ierr != nil {
			return fmt.Errorf("codigo de provincia non é un número: %s", rows[i][1])
		}

		concellos = append(concellos, model.Concello{
			CodigoProvincia: provinciaId,
			CodigoConcello:  concelloId,
			Nome:            rows[i][4],
		})
	}

	err = r.GardarConcellos(concellos)

	return err
}
