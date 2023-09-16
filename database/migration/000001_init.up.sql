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
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED NOT NULL,
    siglas               VARCHAR(50)  NOT NULL,
    nome                 VARCHAR(150),
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE
);

CREATE TABLE lista
(
    id             INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    candidatura_id INT UNSIGNED NOT NULL,
    ambito         INT UNSIGNED,
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidatura (id) ON DELETE CASCADE
);

CREATE TABLE candidato
(
    lista_id INT UNSIGNED     NOT NULL,
    posicion TINYINT UNSIGNED NOT NULL,
    titular  BOOLEAN          NOT NULL,
    nombre   VARCHAR(25)      NOT NULL,
    apelidos VARCHAR(50)      NOT NULL,
    CONSTRAINT PRIMARY KEY (lista_id, posicion),
    CONSTRAINT FOREIGN KEY (lista_id) REFERENCES lista (id) ON DELETE CASCADE
);

CREATE TABLE mesa_electoral
(
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED      NOT NULL,
    concello_id          SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    censo                INT UNSIGNED      NOT NULL,
    votos_blanco         INT UNSIGNED      NOT NULL,
    votos_nulos          INT UNSIGNED      NOT NULL,
    votos_candidaturas   INT UNSIGNED      NOT NULL,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (concello_id) REFERENCES concello (id)
);

CREATE TABLE mesa_electoral_votos_candidatura
(
    id                INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    mesa_electoral_id INT UNSIGNED      NOT NULL,
    candidatura_id    INT UNSIGNED      NOT NULL,
    posicion          TINYINT UNSIGNED,
    votos             SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (mesa_electoral_id) REFERENCES mesa_electoral (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidatura (id) ON DELETE CASCADE
);

CREATE TABLE circunscripcion_cera
(
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED     NOT NULL,
    provincia_id         TINYINT UNSIGNED NOT NULL,
    censo                INT UNSIGNED     NOT NULL,
    votos_blanco         INT UNSIGNED     NOT NULL,
    votos_nulos          INT UNSIGNED     NOT NULL,
    votos_candidaturas   INT UNSIGNED     NOT NULL,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES proceso_electoral (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (provincia_id) REFERENCES provincia (id)
);

CREATE TABLE circunscripcion_cera_votos_candidatura
(
    id                     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    circuscripcion_cera_id INT UNSIGNED      NOT NULL,
    candidatura_id         INT UNSIGNED      NOT NULL,
    posicion               TINYINT UNSIGNED,
    votos                  SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (circuscripcion_cera_id) REFERENCES circunscripcion_cera (id) ON DELETE CASCADE,
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidatura (id) ON DELETE CASCADE
);
