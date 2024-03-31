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

INSERT INTO concello (id, nome, provincia_id)
VALUES (1, 'A Coruña', 1),
       (2, 'Arteixo', 1),
       (3, 'Cambre', 1),
       (4, 'Culleredo', 1),
       (5, 'Ferrol', 1),
       (6, 'Narón', 1),
       (7, 'Oleiros', 1),
       (8, 'Sada', 1),
       (9, 'Ames', 1),
       (19, 'Santiago de Compostela', 1),
       (20, 'Teo', 1),
       (21, 'Val do Dubra', 1),
       (23, 'Vedra', 1);

INSERT INTO organizacion_politica(id, siglas, nome, logo, pai_id)
VALUES (1, 'PP', 'Partido Popular', 'pp.png', NULL),
       (2, 'PSOE', 'Partido Socialista Obrero Español', 'psoe.png', NULL),
       (3, 'BNG', 'Bloque Nacionalista Galego', 'bng.png', NULL),
       (4, 'Cs', 'Ciudadanos', 'cs.png', NULL),
       (5, 'VOX', 'VOX', 'vox.png', NULL),
       (6, 'UP', 'Unidas Podemos', 'up.png', NULL),
       (7, 'Marea Atlántica', 'Marea Atlántica', 'marea.png', NULL),
       (8, 'Alternativa dos Veciños', 'Alternativa dos Veciños', 'av.png', NULL),
       (9, 'Compromiso por Galicia', 'Compromiso por Galicia', 'cpg.png', NULL),
       (10, 'Partido Popular de Galicia', 'Partido Popular de Galicia', 'ppg.png', 1),
       (11, 'Partido Socialista de Galicia', 'Partido Socialista de Galicia', 'psg.png', 2),
       (12, 'Bloque Nacionalista Galego', 'Bloque Nacionalista Galego', 'bng.png', 3),
       (13, 'Ciudadanos Galicia', 'Ciudadanos Galicia', 'csg.png', 4),
       (14, 'VOX Galicia', 'VOX Galicia', 'voxg.png', 5),
       (15, 'Unidas Podemos Galicia', 'Unidas Podemos Galicia', 'upg.png', 6),
       (16, 'Marea Atlántica', 'Marea Atlántica', 'marea.png', 7),
       (17, 'Alternativa dos Veciños', 'Alternativa dos Veciños', 'av.png', 8),
       (18, 'Compromiso por Galicia', 'Compromiso por Galicia', 'cpg.png', 9);

INSERT INTO proceso_electoral(id, tipo, ambito, data)
VALUES (1, 12, 12, '2020-07-12');
