CREATE TABLE tipo_proceso_electoral
(
  id TINYINT UNSIGNED PRIMARY KEY,
  nome VARCHAR(100) NOT NULL
);

INSERT INTO tipo_proceso_electoral(id, nome)
VALUES (1, 'Referéndum'),
       (2, 'Congreso'),
       (3, 'Senado'),
       (4, 'Municipais'),
       (5, 'Autonómicas'),
       (6, 'Cabildos Insulares'),
       (7, 'Parlamento Europeu'),
       (15, 'Xuntas Xerais');
