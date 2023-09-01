CREATE TABLE organizacion
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    siglas VARCHAR(128),
    nome VARCHAR(256)
);

CREATE TABLE proceso_electoral
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    tipo TINYINT,
    data DATETIME,
    ambito TINYINT
);
