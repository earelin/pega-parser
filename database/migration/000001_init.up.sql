CREATE TABLE comunidades_autonomas
(
    id TINYINT UNSIGNED PRIMARY KEY,
    nome VARCHAR(32) NOT NULL
);

CREATE TABLE provincias
(
    id TINYINT UNSIGNED PRIMARY KEY,
    nome VARCHAR(32) NOT NULL ,
    comunidad_autonoma_id TINYINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (comunidad_autonoma_id) REFERENCES comunidades_autonomas(id)
);

CREATE TABLE municipios
(
    provincia_id TINYINT UNSIGNED NOT NULL,
    municipio_id SMALLINT NOT NULL,
    nome VARCHAR(128) NOT NULL ,
    CONSTRAINT PRIMARY KEY (provincia_id, municipio_id),
    CONSTRAINT FOREIGN KEY (provincia_id) REFERENCES provincias(id)
);

CREATE TABLE procesos_electorais
(
    id         INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    tipo       TINYINT UNSIGNED NOT NULL,
    ambito_ine TINYINT UNSIGNED, -- Codigo INE (si no es estatal)
    data       DATETIME         NOT NULL
);

CREATE TABLE candidaturas
(
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED NOT NULL,
    siglas               VARCHAR(50)  NOT NULL,
    nome                 VARCHAR(150)
);

CREATE TABLE listas
(
    id             INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    candidatura_id INT UNSIGNED NOT NULL,
    ambito_ine     SMALLINT UNSIGNED,
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidaturas (id)
);

CREATE TABLE candidatos
(
    lista_id INT UNSIGNED     NOT NULL,
    posicion TINYINT UNSIGNED NOT NULL,
    titular  BOOLEAN          NOT NULL,
    nombre   VARCHAR(25)      NOT NULL,
    apelidos VARCHAR(50)      NOT NULL,
    CONSTRAINT PRIMARY KEY (lista_id, posicion),
    CONSTRAINT FOREIGN KEY (lista_id) REFERENCES listas (id)
);

CREATE TABLE mesas_electorais
(
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED      NOT NULL,
    provincia_ine        TINYINT UNSIGNED  NOT NULL,
    municipio_ine        SMALLINT UNSIGNED NOT NULL,
    distrito             TINYINT UNSIGNED  NOT NULL,
    seccion              CHAR(4)           NOT NULL,
    codigo               CHAR(1)           NOT NULL,
    censo                INT UNSIGNED      NOT NULL,
    votos_blanco         INT UNSIGNED      NOT NULL,
    votos_nulos          INT UNSIGNED      NOT NULL,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES procesos_electorais (id)
);

CREATE TABLE mesa_electoral_votos_candidaturas
(
    id                INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    mesa_electoral_id INT UNSIGNED      NOT NULL,
    candidatura_id    INT UNSIGNED      NOT NULL,
    orden             TINYINT UNSIGNED,
    votos             SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (mesa_electoral_id) REFERENCES mesas_electorais (id),
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidaturas (id)
);

CREATE TABLE circuscripcions_cera
(
    id                   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    proceso_electoral_id INT UNSIGNED     NOT NULL,
    provincia_ine        TINYINT UNSIGNED NOT NULL,
    censo                INT UNSIGNED     NOT NULL,
    votos_blanco         INT UNSIGNED     NOT NULL,
    votos_nulos          INT UNSIGNED     NOT NULL,
    CONSTRAINT FOREIGN KEY (proceso_electoral_id) REFERENCES procesos_electorais (id)
);

CREATE TABLE circuscripcions_cera_votos_candidaturas
(
    id                     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    circuscripcion_cera_id INT UNSIGNED      NOT NULL,
    candidatura_id         INT UNSIGNED      NOT NULL,
    orden                  TINYINT UNSIGNED,
    votos                  SMALLINT UNSIGNED NOT NULL,
    CONSTRAINT FOREIGN KEY (circuscripcion_cera_id) REFERENCES circuscripcions_cera (id),
    CONSTRAINT FOREIGN KEY (candidatura_id) REFERENCES candidaturas (id)
);
