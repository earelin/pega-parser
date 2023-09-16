CREATE VIEW concello_datos
AS
SELECT proceso_electoral_id,
       concello_id,
       SUM(censo)        AS censo,
       SUM(votos_blanco) AS votos_blanco,
       SUM(votos_nulos)  AS votos_nulos
FROM mesa_electoral
GROUP BY proceso_electoral_id, concello_id;

CREATE VIEW concello_votos_candidatura
AS
SELECT proceso_electoral_id, concello_id, candidatura_id, SUM(votos) AS votos
FROM mesa_electoral_votos_candidatura mevc
LEFT JOIN mesa_electoral me ON mevc.mesa_electoral_id = me.id
GROUP BY proceso_electoral_id, concello_id, candidatura_id;
