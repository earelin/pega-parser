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

CREATE VIEW datos_xerais_provincias
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

CREATE VIEW datos_xerais_concellos
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
