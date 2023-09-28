CREATE TABLE comunidade_autonoma
(
    id   TINYINT UNSIGNED PRIMARY KEY,
    nome VARCHAR(32) NOT NULL
);

CREATE TABLE provincia
(
    id                     TINYINT UNSIGNED PRIMARY KEY,
    nome                   VARCHAR(32)      NOT NULL,
    comunidade_autonoma_id TINYINT UNSIGNED NOT NULL,
    FOREIGN KEY (comunidade_autonoma_id) REFERENCES comunidade_autonoma (id)
);

CREATE TABLE concello
(
    id           SMALLINT UNSIGNED PRIMARY KEY,
    provincia_id TINYINT UNSIGNED NOT NULL,
    nome         VARCHAR(128)     NOT NULL,
    FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE proceso_electoral
(
    id     INTEGER PRIMARY KEY AUTOINCREMENT,
    tipo   TINYINT UNSIGNED NOT NULL,
    ambito TINYINT UNSIGNED,
    data   DATETIME         NOT NULL
);

CREATE TABLE candidatura
(
    proceso_electoral_id INTEGER UNSIGNED NOT NULL,
    id                   INTEGER UNSIGNED NOT NULL,
    siglas               VARCHAR(50)  NOT NULL,
    nome                 VARCHAR(150),
    cabeceira_estatal    INTEGER UNSIGNED,
    cabeceira_autonomica INTEGER UNSIGNED,
    cabeceira_provincial INTEGER UNSIGNED,
    PRIMARY KEY (proceso_electoral_id, id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE
);

CREATE TABLE candidato
(
    proceso_electoral_id INTEGER UNSIGNED     NOT NULL,
    candidatura_id       INTEGER UNSIGNED     NOT NULL,
    ambito               INTEGER UNSIGNED     NOT NULL,
    posicion             TINYINT UNSIGNED NOT NULL,
    titular              BOOLEAN          NOT NULL,
    nombre               VARCHAR(25)      NOT NULL,
    apelidos             VARCHAR(50)      NOT NULL,
    eleito              BOOLEAN          NOT NULL,
    PRIMARY KEY (proceso_electoral_id, candidatura_id, ambito, posicion),
    FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);

CREATE TABLE mesa_electoral
(
    proceso_electoral_id INTEGER UNSIGNED      NOT NULL,
    concello_id          SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    censo                INTEGER UNSIGNED      NOT NULL,
    votos_blanco         INTEGER UNSIGNED      NOT NULL,
    votos_nulos          INTEGER UNSIGNED      NOT NULL,
    votos_candidaturas   INTEGER UNSIGNED      NOT NULL,
    PRIMARY KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    FOREIGN KEY (concello_id) REFERENCES concello (id)
);

CREATE TABLE mesa_electoral_votos_candidatura
(
    proceso_electoral_id INTEGER UNSIGNED      NOT NULL,
    concello_id          SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    candidatura_id       INTEGER UNSIGNED      NOT NULL,
    posicion             TINYINT UNSIGNED,
    votos                SMALLINT UNSIGNED NOT NULL,
    PRIMARY KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo, candidatura_id),
    FOREIGN KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo) REFERENCES mesa_electoral (proceso_electoral_id, concello_id, distrito, seccion, codigo) ON DELETE CASCADE,
    FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);

CREATE TABLE circunscripcion_cera
(
    proceso_electoral_id INTEGER UNSIGNED     NOT NULL,
    provincia_id         TINYINT UNSIGNED NOT NULL,
    censo                INTEGER UNSIGNED     NOT NULL,
    votos_blanco         INTEGER UNSIGNED     NOT NULL,
    votos_nulos          INTEGER UNSIGNED     NOT NULL,
    votos_candidaturas   INTEGER UNSIGNED     NOT NULL,
    PRIMARY KEY (proceso_electoral_id, provincia_id),
    FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE circunscripcion_cera_votos_candidatura
(
    proceso_electoral_id INTEGER UNSIGNED      NOT NULL,
    provincia_id         TINYINT UNSIGNED  NOT NULL,
    candidatura_id       INTEGER UNSIGNED      NOT NULL,
    posicion             TINYINT UNSIGNED,
    votos                SMALLINT UNSIGNED NOT NULL,
    FOREIGN KEY (proceso_electoral_id, provincia_id) REFERENCES circunscripcion_cera (proceso_electoral_id, provincia_id) ON DELETE CASCADE,
    FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);
