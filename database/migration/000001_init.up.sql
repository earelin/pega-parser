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

CREATE TABLE comunidade_autonoma
(
    id   INTEGER UNSIGNED PRIMARY KEY,
    nome VARCHAR(32) NOT NULL
);

CREATE TABLE provincia
(
    id                     INTEGER UNSIGNED PRIMARY KEY,
    nome                   VARCHAR(32)      NOT NULL,
    comunidade_autonoma_id TINYINT UNSIGNED NOT NULL,
    FOREIGN KEY (comunidade_autonoma_id) REFERENCES comunidade_autonoma (id)
);

CREATE TABLE concello
(
    id           INTEGER UNSIGNED PRIMARY KEY,
    provincia_id TINYINT UNSIGNED NOT NULL,
    nome         VARCHAR(128)     NOT NULL,
    FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE organizacion_politica
(
    id     INTEGER UNSIGNED PRIMARY KEY,
    siglas VARCHAR(50)  NOT NULL,
    nome   VARCHAR(150) NOT NULL
);

CREATE TABLE proceso_electoral
(
    id     INTEGER UNSIGNED PRIMARY KEY,
    tipo   TINYINT UNSIGNED NOT NULL,
    ambito TINYINT UNSIGNED,
    data   DATETIME         NOT NULL
);

CREATE TABLE candidatura
(
    proceso_electoral_id INTEGER UNSIGNED NOT NULL,
    id                   INTEGER UNSIGNED NOT NULL,
    siglas               VARCHAR(50)      NOT NULL,
    nome                 VARCHAR(150),
    PRIMARY KEY (proceso_electoral_id, id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE
);

CREATE TABLE candidatura_organizacion_politica
(
    candidatura_id           INTEGER UNSIGNED NOT NULL,
    organizacion_politica_id INTEGER UNSIGNED NOT NULL,
    PRIMARY KEY (candidatura_id, organizacion_politica_id),
    FOREIGN KEY (candidatura_id) REFERENCES candidatura (id) ON DELETE CASCADE,
    FOREIGN KEY (organizacion_politica_id) REFERENCES organizacion_politica (id)
);

CREATE TABLE concello_resultados
(
    proceso_electoral_id INTEGER UNSIGNED  NOT NULL,
    concello_id          INTEGER UNSIGNED NOT NULL,
    censo                INTEGER UNSIGNED  NOT NULL,
    votos_blanco         INTEGER UNSIGNED  NOT NULL,
    votos_nulos          INTEGER UNSIGNED  NOT NULL,
    votos_candidaturas   INTEGER UNSIGNED  NOT NULL,
    PRIMARY KEY (proceso_electoral_id, concello_id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    FOREIGN KEY (concello_id) REFERENCES concello (id)
);

CREATE TABLE concello_votos_candidaturas
(
    proceso_electoral_id INTEGER UNSIGNED  NOT NULL,
    concello_id          INTEGER UNSIGNED NOT NULL,
    candidatura_id       INTEGER UNSIGNED  NOT NULL,
    votos                INTEGER UNSIGNED  NOT NULL,
    PRIMARY KEY (proceso_electoral_id, concello_id, candidatura_id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    FOREIGN KEY (concello_id) REFERENCES concello (id),
    FOREIGN KEY (candidatura_id) REFERENCES candidatura (id)
);

CREATE TABLE cera_resultados
(
    proceso_electoral_id INTEGER UNSIGNED NOT NULL,
    provincia_id         INTEGER UNSIGNED NOT NULL,
    censo                INTEGER UNSIGNED NOT NULL,
    votos_blanco         INTEGER UNSIGNED NOT NULL,
    votos_nulos          INTEGER UNSIGNED NOT NULL,
    votos_candidaturas   INTEGER UNSIGNED NOT NULL,
    PRIMARY KEY (proceso_electoral_id, provincia_id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE cera_votos_candidaturas
(
    proceso_electoral_id INTEGER UNSIGNED  NOT NULL,
    provincia_id         INTEGER UNSIGNED  NOT NULL,
    candidatura_id       INTEGER UNSIGNED  NOT NULL,
    votos                SMALLINT UNSIGNED NOT NULL,
    FOREIGN KEY (proceso_electoral_id, provincia_id) REFERENCES cera_resultados (proceso_electoral_id, provincia_id) ON DELETE CASCADE,
    FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);
