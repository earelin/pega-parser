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

package db

import (
	"database/sql"
	"log"
)

type Config struct {
	Filename string
}

func (rc Config) BuildPool() *sql.DB {
	var pool, err = sql.Open("sqlite3", rc.Filename)
	if err != nil {
		log.Panicf("Error connecting to database: %s", err)
	}

	return pool
}
