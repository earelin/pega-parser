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
    CONSTRAINT FOREIGN KEY (comunidade_autonoma_id) REFERENCES comunidade_autonoma (id)
);

CREATE TABLE concello
(
    id           SMALLINT UNSIGNED PRIMARY KEY,
    provincia_id TINYINT UNSIGNED NOT NULL,
    nome         VARCHAR(128)     NOT NULL,
    CONSTRAINT FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE proceso_electoral
(
    id     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    tipo   TINYINT UNSIGNED NOT NULL,
    ambito TINYINT UNSIGNED,
    data   DATETIME         NOT NULL
);

CREATE TABLE candidatura
(
    proceso_electoral_id INT UNSIGNED NOT NULL,
    id                   INT UNSIGNED NOT NULL,
    siglas               VARCHAR(50)  NOT NULL,
    nome                 VARCHAR(150),
    cabeceira_estatal    INT UNSIGNED,
    cabeceira_autonomica INT UNSIGNED,
    cabeceira_provincial INT UNSIGNED,
    CONSTRAINT PRIMARY KEY (proceso_electoral_id, id),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE
);

CREATE TABLE lista
(
    proceso_electoral_id INT UNSIGNED NOT NULL,
    candidatura_id       INT UNSIGNED NOT NULL,
    ambito               INT UNSIGNED NOT NULL,
    CONSTRAINT PRIMARY KEY (proceso_electoral_id, candidatura_id, ambito),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);

CREATE TABLE candidato
(
    proceso_electoral_id INT UNSIGNED     NOT NULL,
    candidatura_id       INT UNSIGNED     NOT NULL,
    ambito               INT UNSIGNED     NOT NULL,
    posicion             TINYINT UNSIGNED NOT NULL,
    titular              BOOLEAN          NOT NULL,
    nombre               VARCHAR(25)      NOT NULL,
    apelidos             VARCHAR(50)      NOT NULL,
    CONSTRAINT PRIMARY KEY (proceso_electoral_id, candidatura_id, ambito, posicion),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, candidatura_id, ambito) REFERENCES lista (proceso_electoral_id, candidatura_id, ambito) ON DELETE CASCADE
);

CREATE TABLE mesa_electoral
(
    proceso_electoral_id INT UNSIGNED      NOT NULL,
    concello_id          SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    censo                INT UNSIGNED      NOT NULL,
    votos_blanco         INT UNSIGNED      NOT NULL,
    votos_nulos          INT UNSIGNED      NOT NULL,
    votos_candidaturas   INT UNSIGNED      NOT NULL,
    PRIMARY KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (concello_id) REFERENCES concello (id)
);

CREATE TABLE mesa_electoral_votos_candidatura
(
    proceso_electoral_id INT UNSIGNED      NOT NULL,
    concello_id          SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    candidatura_id       INT UNSIGNED      NOT NULL,
    posicion             TINYINT UNSIGNED,
    votos                SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT PRIMARY KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo, candidatura_id),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, concello_id, distrito, seccion, codigo) REFERENCES mesa_electoral (proceso_electoral_id, concello_id, distrito, seccion, codigo) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);

CREATE TABLE circunscripcion_cera
(
    proceso_electoral_id INT UNSIGNED     NOT NULL,
    provincia_id         TINYINT UNSIGNED NOT NULL,
    censo                INT UNSIGNED     NOT NULL,
    votos_blanco         INT UNSIGNED     NOT NULL,
    votos_nulos          INT UNSIGNED     NOT NULL,
    votos_candidaturas   INT UNSIGNED     NOT NULL,
    CONSTRAINT PRIMARY KEY (proceso_electoral_id, provincia_id),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE circunscripcion_cera_votos_candidatura
(
    proceso_electoral_id INT UNSIGNED      NOT NULL,
    provincia_id         TINYINT UNSIGNED  NOT NULL,
    candidatura_id       INT UNSIGNED      NOT NULL,
    posicion             TINYINT UNSIGNED,
    votos                SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, provincia_id) REFERENCES circunscripcion_cera (proceso_electoral_id, provincia_id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id, candidatura_id) REFERENCES candidatura (proceso_electoral_id, id) ON DELETE CASCADE
);
