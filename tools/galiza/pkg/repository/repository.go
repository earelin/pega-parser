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

package repository

import (
	"context"
	"database/sql"
	"github.com/earelin/pega/tools/galiza/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Repository struct {
	pool *sql.DB
	ctx  context.Context
}

func NewRepository(c config.Config, ctx context.Context) (*Repository, error) {
	var r Repository

	var pool, err = sql.Open("sqlite3", c.FilePath)
	if err != nil {
		return nil, err
	}

	r.pool = pool
	r.ctx = ctx

	return &r, nil
}

func (r *Repository) CheckConnection() error {
	var ctx, cancel = context.WithTimeout(r.ctx, 5*time.Second)
	defer cancel()
	return r.pool.PingContext(ctx)
}

func (r *Repository) CloseConnection() error {
	return r.pool.Close()
}
