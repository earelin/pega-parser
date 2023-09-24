CREATE VIEW datos_xerais
AS
SELECT p.id AS id, censo_ine.censo AS censo_ine, censo_cera.censo AS censo_cera
FROM proceso_electoral p
         JOIN (SELECT proceso_electoral_id AS id, SUM(censo) AS censo
               FROM mesa_electoral
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id) AS censo_ine ON p.id = censo_ine.id
         JOIN (SELECT proceso_electoral_id AS id, SUM(censo) AS censo
               FROM circunscripcion_cera
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id) AS censo_cera ON p.id = censo_cera.id;

CREATE VIEW datos_xerais_autonomicos
AS
SELECT censo_ine.id AS id, ca.id AS comunidade_autonoma_id, censo_ine.censo AS censo_ine, censo_cera.censo AS censo_cera
FROM comunidade_autonoma ca
         JOIN (SELECT proceso_electoral_id     AS id,
                      p.comunidade_autonoma_id AS comunidade_autonoma_id,
                      SUM(censo)               AS censo
               FROM mesa_electoral me
                        LEFT JOIN concello c ON c.id = me.concello_id
                        LEFT JOIN provincia p ON p.id = c.provincia_id
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id, p.comunidade_autonoma_id) AS censo_ine ON ca.id = censo_ine.comunidade_autonoma_id
         JOIN (SELECT proceso_electoral_id     AS id,
                      p.comunidade_autonoma_id AS comunidade_autonoma_id,
                      SUM(censo)               AS censo
               FROM circunscripcion_cera cc
                        LEFT JOIN provincia p ON p.id = cc.provincia_id
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id, p.comunidade_autonoma_id) AS censo_cera ON ca.id = censo_cera.comunidade_autonoma_id;

CREATE VIEW datos_xerais_provincias
AS
SELECT censo_ine.id AS id, p.id AS provincia_id, censo_ine.censo AS censo_ine, censo_cera.censo AS censo_cera
FROM provincia p
         JOIN (SELECT proceso_electoral_id AS id, c.provincia_id AS provincia_id, SUM(censo) AS censo
               FROM mesa_electoral me
                        LEFT JOIN concello c ON me.concello_id = c.id
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id, provincia_id) AS censo_ine ON p.id = censo_ine.provincia_id
         JOIN (SELECT proceso_electoral_id AS id, provincia_id, SUM(censo) AS censo
               FROM circunscripcion_cera
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id, provincia_id) AS censo_cera ON p.id = censo_cera.provincia_id;

CREATE VIEW datos_xerais_concellos
AS
SELECT p.id AS id, censo_ine.concello_id, censo_ine.censo AS censo_ine
FROM proceso_electoral p
         JOIN (SELECT proceso_electoral_id AS id, concello_id, SUM(censo) AS censo
               FROM mesa_electoral
               WHERE proceso_electoral_id = 1
               GROUP BY proceso_electoral_id, concello_id) AS censo_ine ON p.id = censo_ine.id
